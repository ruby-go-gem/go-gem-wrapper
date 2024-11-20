# frozen_string_literal: true

require "rake"
require "rake/tasklib"

module GoGem
  # Provides rake tasks for `go test` with CRuby
  #
  # @example Without config
  #   # Rakefile
  #   require "go_gem/rake_task"
  #
  #   GoGem::RakeTask.new("gem_name")
  #
  # @example With config
  #   # Rakefile
  #   require "go_gem/rake_task"
  #
  #   GoGem::RakeTask.new("gem_name") do |t|
  #     t.task_namespace = "go5"
  #     t.go_bin_path = "/path/to/go"
  #     t.go_test_args = "-mod=readonly"
  #     t.target_dir = "/dir/to/go-mod/"
  #   end
  class RakeTask < ::Rake::TaskLib
    DEFAULT_TASK_NAMESPACE = :go

    DEFAULT_GO_BIN_PATH = "go"

    DEFAULT_GO_TEST_ARGS = "-mod=readonly -count=1"

    # @!attribute [r] gem_name
    # @return [String]
    attr_reader :gem_name

    # @!attribute task_namespace
    #   @return [Symbol] task namespace (default: `:go`)
    attr_accessor :task_namespace

    # @!attribute go_bin_path
    #   @return [String] path to go binary (default: `"go"`)
    attr_accessor :go_bin_path

    # @!attribute go_test_args
    #   @return [String] argument passed to `go test` (default: `"-mod=readonly -count=1"`)
    attr_accessor :go_test_args

    # @!attribute cwd
    #   @return [String] directory when executing go commands. (default: `"ext/#{gem_name}"`)
    attr_accessor :target_dir

    # @param gem_name [String]
    # @yield configuration of {RakeTask}
    # @yieldparam t [RakeTask]
    def initialize(gem_name)
      super()

      @gem_name = gem_name

      @task_namespace = DEFAULT_TASK_NAMESPACE
      @go_bin_path = DEFAULT_GO_BIN_PATH
      @go_test_args = DEFAULT_GO_TEST_ARGS
      @target_dir = ext_dir

      yield(self) if block_given?

      namespace(task_namespace) do
        define_go_test_task
        define_go_testrace_task
        define_go_fmt_task
      end
    end

    # Generate environment variables to build go programs in the Go gem
    #
    # @return [Hash<String, String>]
    def self.build_env_vars
      ldflags = "-L#{RbConfig::CONFIG["libdir"]} -l#{RbConfig::CONFIG["RUBY_SO_NAME"]}"

      case `#{RbConfig::CONFIG["CC"]} --version` # rubocop:disable Lint/LiteralAsCondition
      when /Free Software Foundation/
        ldflags << " -Wl,--unresolved-symbols=ignore-all"
      when /clang/
        ldflags << " -undefined dynamic_lookup"
      end

      cflags = [
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
        ldflags.gsub!("-Wl,--unresolved-symbols=ignore-all", "")
      end

      ld_library_path = RbConfig::CONFIG["libdir"]

      {
        "CGO_CFLAGS"      => cflags,
        "CGO_LDFLAGS"     => ldflags,
        "LD_LIBRARY_PATH" => ld_library_path,
      }
    end

    private

    def define_go_test_task
      desc "Run #{go_bin_path} test"
      task(:test) do
        within_target_dir do
          sh RakeTask.build_env_vars, "#{go_bin_path} test #{go_test_args} ./..."
        end
      end
    end

    def define_go_testrace_task
      desc "Run #{go_bin_path} test -race"
      task(:testrace) do
        within_target_dir do
          sh RakeTask.build_env_vars, "#{go_bin_path} test #{go_test_args} -race ./..."
        end
      end
    end

    def define_go_fmt_task
      desc "Run #{go_bin_path} fmt"
      task(:fmt) do
        within_target_dir do
          sh "#{go_bin_path} fmt ./..."
        end
      end
    end

    # @yield
    def within_target_dir
      Dir.chdir(target_dir) do # rubocop:disable Style/ExplicitBlockArgument
        yield
      end
    end

    # @return [String]
    def ext_dir
      File.join("ext", gem_name)
    end
  end
end
