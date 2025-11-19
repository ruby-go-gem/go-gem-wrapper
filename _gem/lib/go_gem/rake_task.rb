# frozen_string_literal: true

require "rake"
require "rake/tasklib"

require_relative "util"

module GoGem
  # rubocop:disable Layout/LineLength

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
  #
  # @example additional tasks
  #   # Rakefile
  #   require "go_gem/rake_task"
  #
  #   t = GoGem::RakeTask.new("gem_name")
  #
  #   namespace :go do
  #     desc "Run golangci-lint"
  #     task :lint do
  #       t.within_target_dir do
  #         sh "which golangci-lint" do |ok, _|
  #           raise "golangci-lint isn't installed. See. https://golangci-lint.run/welcome/install/" unless ok
  #         end
  #
  #         build_tag = GoGem::Util.ruby_minor_version_build_tag
  #         sh GoGem::RakeTask.build_env_vars, "golangci-lint run --build-tags #{build_tag} --modules-download-mode=readonly"
  #       end
  #     end
  #   end
  class RakeTask < ::Rake::TaskLib
    # rubocop:enable Layout/LineLength

    DEFAULT_TASK_NAMESPACE = :go

    DEFAULT_GO_BIN_PATH = "go"

    DEFAULT_GO_TEST_ARGS = "-mod=readonly -count=1"

    # @!attribute [r] gem_name
    # @return [String]
    attr_reader :gem_name

    # @!attribute task_namespace
    #   @return [Symbol,String] task namespace (default: `:go`)
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
        define_go_build_envs_task
        define_go_build_tag_task
        define_go_mod_tidy_task
      end
    end

    # Generate environment variables to build go programs in the Go gem
    #
    # @return [Hash<String, String>]
    def self.build_env_vars
      ldflags = GoGem::Util.generate_ldflags
      cflags = GoGem::Util.generate_cflags

      ld_library_path = RbConfig::CONFIG["libdir"].to_s

      {
        "GOFLAGS"         => GoGem::Util.generate_goflags,
        "CGO_CFLAGS"      => cflags,
        "CGO_LDFLAGS"     => ldflags,
        "LD_LIBRARY_PATH" => ld_library_path,
      }
    end

    # Change current working dir inside `ext/<GEM_NAME>/`
    #
    # @yield
    def within_target_dir
      Dir.chdir(target_dir) do # rubocop:disable Style/ExplicitBlockArgument
        yield
      end
    end

    # @return [String] ext dir name (e.g.`ext/<GEM_NAME>/`)
    def ext_dir
      File.join("ext", gem_name)
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

    def define_go_build_envs_task
      desc "Print build envs for `go build`"
      task(:build_envs, [:env_name]) do |_, args|
        if args[:env_name]
          value = RakeTask.build_env_vars[args[:env_name]]
          puts "#{args[:env_name]}=#{value}"
        else
          RakeTask.build_env_vars.each do |name, value|
            puts "#{name}=#{value}"
          end
        end
      end
    end

    def define_go_build_tag_task
      desc "Print build tag"
      task(:build_tag) do
        puts GoGem::Util.ruby_minor_version_build_tag
      end
    end

    def define_go_mod_tidy_task
      desc "Run #{go_bin_path} mod tidy"
      task(:mod_tidy) do
        within_target_dir do
          sh "#{go_bin_path} mod tidy"
        end
      end
    end
  end
end
