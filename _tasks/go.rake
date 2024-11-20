# frozen_string_literal: true

GoGem::RakeTask.new("") do |t|
  t.target_dir = repo_root
  t.go_test_args = "#{GoGem::RakeTask::DEFAULT_GO_TEST_ARGS} #{ENV["GO_TEST_ARGS"]}"
end

namespace :go do
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
