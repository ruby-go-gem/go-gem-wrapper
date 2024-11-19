# frozen_string_literal: true

namespace :go do
  desc "Run go test"
  task :test do
    sh GoGem::RakeTask.build_env_vars, "go test -mod=readonly -count=1 #{ENV["GO_TEST_ARGS"]} ./..."
  end

  desc "Run go test -race"
  task :testrace do
    sh GoGem::RakeTask.build_env_vars, "go test -mod=readonly -count=1 #{ENV["GO_TEST_ARGS"]} -race  ./..."
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
    sh GoGem::RakeTask.build_env_vars, "golangci-lint run"
  end

  desc "Run all build tasks in go"
  task build_all: %i[test fmt lint]
end
