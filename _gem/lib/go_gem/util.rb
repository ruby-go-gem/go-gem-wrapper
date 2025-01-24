# frozen_string_literal: true

module GoGem
  # Common utility methods for {GoGem::Mkmf} and {GoGem::RakeTask}
  module Util
    # Return ruby version build tag for `go build` and `go test`
    #
    # @param ruby_version [String]
    # @return [String]
    #
    # @example
    #   GoGem::Util.ruby_minor_version_build_tag("3.4.1")
    #   #=> "ruby_3_4"
    def self.ruby_minor_version_build_tag(ruby_version = RUBY_VERSION)
      "ruby_#{ruby_version.to_f.to_s.gsub(".", "_")}"
    end

    # @return [String]
    def self.generate_ldflags
      ldflags = "-L#{RbConfig::CONFIG["libdir"]} -l#{RbConfig::CONFIG["RUBY_SO_NAME"]}"

      case `#{RbConfig::CONFIG["CC"]} --version` # rubocop:disable Lint/LiteralAsCondition
      when /Free Software Foundation/
        ldflags << " -Wl,--unresolved-symbols=ignore-all"
      when /clang/
        ldflags << " -undefined dynamic_lookup"
      end

      # FIXME: Workaround for Ubuntu (GitHub Actions)
      ldflags.gsub!("-Wl,--unresolved-symbols=ignore-all", "") if RUBY_PLATFORM =~ /linux/i

      ldflags.strip
    end

    # @return [String]
    def self.generate_cflags
      cflags =
        [
          RbConfig::CONFIG["CFLAGS"],
          "-I#{RbConfig::CONFIG["rubyarchhdrdir"]}",
          "-I#{RbConfig::CONFIG["rubyhdrdir"]}",
        ].join(" ")

      # FIXME: Workaround for Ubuntu (GitHub Actions)
      if RUBY_PLATFORM =~ /linux/i
        cflags.gsub!("-Wno-self-assign", "")
        cflags.gsub!("-Wno-parentheses-equality", "")
        cflags.gsub!("-Wno-constant-logical-operand", "")
        cflags.gsub!("-Wsuggest-attribute=format", "")
        cflags.gsub!("-Wold-style-definition", "")
        cflags.gsub!("-Wsuggest-attribute=noreturn", "")
      end

      # FIXME: Workaround for Alpine
      cflags.gsub!("-Wpointer-arith", "") if RUBY_PLATFORM =~ /linux-musl/i

      cflags.strip
    end

    # @return [String]
    def self.generate_goflags
      "-tags=#{ruby_minor_version_build_tag}"
    end
  end
end
