# frozen_string_literal: true

require_relative "util"

module GoGem
  # Helper module for creating Go Makefiles
  module Mkmf
    # Create Makefile for go-gem
    #
    # @param target [String]
    # @param srcprefix [String,nil]
    # @param go_build_args [String,nil] Arguments passed to `go build`
    #
    # @example
    #   require "mkmf"
    #   require "go_gem/mkmf" # Append this
    #
    #   # Use create_go_makefile instead of create_makefile
    #   # create_makefile("example/example")
    #   create_go_makefile("example/example")
    #
    # @example Pass debug flags to `go build`
    #   create_go_makefile("example/example", go_build_args: "-gcflags='all=-N -l'")
    def create_go_makefile(target, srcprefix: nil, go_build_args: nil)
      find_executable("go")

      # rubocop:disable Style/GlobalVars
      $objs = []
      # @private
      def $objs.empty?; false; end
      # rubocop:enable Style/GlobalVars

      create_makefile(target, srcprefix)

      ldflags = GoGem::Util.generate_ldflags
      current_dir = File.expand_path(".")

      goflags = "-tags=#{GoGem::Util.ruby_minor_version_build_tag}"

      File.open("Makefile", "a") do |f|
        f.write <<~MAKEFILE.gsub(/^ {8}/, "\t")
          $(DLLIB): Makefile $(srcdir)/*.go
                  cd $(srcdir); \
                  CGO_CFLAGS='$(INCFLAGS)' CGO_LDFLAGS='#{ldflags}' GOFLAGS='#{goflags}' \
                    go build -p 4 -buildmode=c-shared -o #{current_dir}/$(DLLIB) #{go_build_args}
        MAKEFILE
      end
    end
  end
end

include GoGem::Mkmf # rubocop:disable Style/MixinUsage
