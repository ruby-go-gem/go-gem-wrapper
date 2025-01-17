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

      ldflags
    end
  end
end
