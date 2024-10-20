# frozen_string_literal: true

require "optparse"

gemspec_file = nil
dry_run = false

opt = OptionParser.new
opt.on("-f", "--file=GEMSPEC_FILE") { |v| gemspec_file = v }
opt.on("--dry-run") { |v| dry_run = v }

opt.parse!(ARGV)

raise "--file is required" unless gemspec_file
raise "#{gemspec_file} isn't gemspec" unless File.extname(gemspec_file) == ".gemspec"
raise "#{gemspec_file} isn't found" unless File.exist?(gemspec_file)

# Patch to make a gem into a Go gem right after `bundle gem`
class GemPatcher # rubocop:disable Metrics/ClassLength
  attr_reader :gemspec_file

  # @param gemspec_file [String]
  # @param dry_run [Boolean]
  def initialize(gemspec_file:, dry_run:)
    @gemspec_file = gemspec_file
    @dry_run = dry_run
  end

  def perform
    create_gem_name_go
    create_go_mod
    update_gem_name_c
    update_extconf_rb
    update_gemspec
  end

  private

  # @return [Boolean]
  def dry_run?
    @dry_run
  end

  # @return [String] path to ext dir. (e.g. /path/to/gem_name/ext/gem_name)
  def ext_dir
    File.join(File.absolute_path(File.dirname(gemspec_file)), "ext", gem_name)
  end

  # @return [String]
  def gem_name
    File.basename(gemspec_file, ".gemspec")
  end

  # @return [String]
  def module_name
    snake_to_camel(gem_name)
  end

  # @param str [String]
  # @return [String]
  def snake_to_camel(str)
    str.split("_").map(&:capitalize).join.gsub(/(?<=\d)([a-z])/) { _1.upcase } # rubocop:disable Style/SymbolProc
  end

  # Create <gem_name>.go
  def create_gem_name_go
    gem_name_go_path = File.join(ext_dir, "#{gem_name}.go")

    return if File.exist?(gem_name_go_path)

    content = <<~GO
      package main

      /*
      #include "#{gem_name}.h"
      */
      import "C"

      import (
      \t"github.com/ruby-go-gem/go-gem-wrapper/ruby"
      )

      //export Init_#{gem_name}
      func Init_#{gem_name}() {
      \trb_m#{module_name} := ruby.RbDefineModule("#{module_name}")
      }

      func main() {
      }
    GO

    save_file(file_path: gem_name_go_path, content:)
  end

  def create_go_mod
    go_mod_path = File.join(ext_dir, "go.mod")

    return if File.exist?(go_mod_path)

    `go version` =~ /go version go([.\d]+)/
    go_version = ::Regexp.last_match(1)

    raise "go isn't found in PATH" unless go_version

    content = <<~GO
      module github.com/username/#{gem_name}

      go #{go_version}
    GO

    save_file(file_path: go_mod_path, content:)
  end

  def update_gem_name_c
    gem_name_c_path = File.join(ext_dir, "#{gem_name}.c")

    content = File.read(gem_name_c_path)

    return if content.include?('#include "_cgo_export.h"')

    content = <<~C
      #include "#{gem_name}.h"
      #include "_cgo_export.h"
    C

    save_file(file_path: gem_name_c_path, content:)
  end

  def update_extconf_rb
    extconf_rb_path = File.join(ext_dir, "extconf.rb")

    content = File.read(extconf_rb_path)

    unless content.include?(%(require "go_gem/mkmf"))
      content.gsub!(<<~RUBY, <<~RUBY)
        require "mkmf"
      RUBY
        require "mkmf"
        require "go_gem/mkmf"
      RUBY
    end

    unless content.include?(%(create_go_makefile("#{gem_name}/#{gem_name}")))
      content.gsub!(<<~RUBY, <<~RUBY)
        create_makefile("#{gem_name}/#{gem_name}")
      RUBY
        create_go_makefile("#{gem_name}/#{gem_name}")
      RUBY
    end

    save_file(file_path: extconf_rb_path, content:)
  end

  def update_gemspec
    content = File.read(gemspec_file)

    return if content.include?(%(.add_dependency "go_gem")) || content.include?(%(.add_runtime_dependency "go_gem"))

    content =~ /Gem::Specification\.new\s+do\s+\|(.+)\|/
    spec_var_name = ::Regexp.last_match(1)

    content.gsub!(/^end\n/, <<~RUBY)
        #{spec_var_name}.add_dependency "go_gem"
      end
    RUBY

    save_file(file_path: gemspec_file, content:)
  end

  # @param file_path [String]
  # @param content [String]
  def save_file(file_path:, content:)
    is_updated = File.exist?(file_path)
    if is_updated
      before_content = File.read(file_path)
      return if content == before_content
    end

    if dry_run?
      if is_updated
        puts "[INFO] #{file_path} will be updated (dry-run)"
      else
        puts "[INFO] #{file_path} will be created (dry-run)"
      end

      return
    end

    File.binwrite(file_path, content)

    if is_updated
      puts "[INFO] #{file_path} is updated"
    else
      puts "[INFO] #{file_path} is created"
    end
  end
end

GemPatcher.new(gemspec_file:, dry_run:).perform
