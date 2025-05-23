# frozen_string_literal: true

require "forwardable"
require "ruby_header_parser"
require "go_gem/util"
require "yaml"

require_relative "ruby_h_to_go/type_helper"

require_relative "ruby_h_to_go/argument_definition"
require_relative "ruby_h_to_go/cli"
require_relative "ruby_h_to_go/config"
require_relative "ruby_h_to_go/go_util"
require_relative "ruby_h_to_go/enum_definition"
require_relative "ruby_h_to_go/function_definition"
require_relative "ruby_h_to_go/struct_definition"
require_relative "ruby_h_to_go/type_definition"
require_relative "ruby_h_to_go/typeref_definition"

# Generate Go binding from ruby.h
module RubyHToGo
  # @return [RubyHToGo::Config]
  def self.config
    @config ||= RubyHToGo::Config.new
  end
end
