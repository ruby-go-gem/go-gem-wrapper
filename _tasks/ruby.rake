# frozen_string_literal: true

require "rubocop/rake_task"

RuboCop::RakeTask.new("ruby:rubocop") do |task|
  task.plugins << "rubocop-on-rbs"
end

namespace :ruby do
  namespace :example do
    desc "Build ruby/testdata/example/"
    task :build do
      Dir.chdir(File.join(repo_root, "ruby", "testdata", "example")) do
        sh "bundle exec rake all"
      end
    end
  end

  namespace :rbs do
    desc "`rbs collection install` and `git commit`"
    task :install do
      sh "rbs collection install"
      sh "git add rbs_collection.lock.yaml"
      sh "git commit -m 'rbs collection install' || true"
    end
  end

  desc "Check rbs"
  task :rbs do
    sh "rbs validate"
    sh "steep check"
  end

  desc "Run all build tasks in ruby"
  task build_all: %w[example:build rubocop rbs]
end
