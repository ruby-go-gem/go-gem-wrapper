# frozen_string_literal: true
# @return [Hash<String, String>]
def env_vars
  ldflags = "-L#{RbConfig::CONFIG["libdir"]} -l#{RbConfig::CONFIG["RUBY_SO_NAME"]}"

  case `#{RbConfig::CONFIG["CC"]} --version` # rubocop:disable Lint/LiteralAsCondition
  when /Free Software Foundation/
    ldflags << " -Wl,--unresolved-symbols=ignore-all"
  when /clang/
    ldflags << " -undefined dynamic_lookup"
  end

  cflags = "#{RbConfig::CONFIG["CFLAGS"]} -I#{RbConfig::CONFIG["rubyarchhdrdir"]} -I#{RbConfig::CONFIG["rubyhdrdir"]}"

  # FIXME: Workaround for GitHub Actions
  if ENV["GITHUB_ACTIONS"]
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

namespace :go do
  desc "Run go test"
  task :test do
    sh env_vars, "go test -mod=readonly -count=1 #{ENV["GO_TEST_ARGS"]} ./..."
  end

  desc "Run go test -race"
  task :testrace do
    sh env_vars, "go test -mod=readonly -count=1 #{ENV["GO_TEST_ARGS"]} -race  ./..."
  end

  desc "Run go fmt"
  task :fmt do
    sh "go fmt ./..."
  end

  desc "Run golangci-lint"
  task :lint do
    sh "which golangci-lint" do |ok, _|
      raise "golangci-lint isn't installed. See. https://golangci-lint.run/welcome/install/" unless ok
    end
    sh env_vars, "golangci-lint run"
  end

  desc "Run all build tasks in go"
  task build_all: %i[test fmt lint]
end
