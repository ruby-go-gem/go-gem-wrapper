# frozen_string_literal: true

require_relative "_gem/lib/go_gem/rake_task"

# @return [String]
def repo_root
  __dir__
end

Dir["#{__dir__}/_tasks/*.rake"].each { |f| load f }

desc "Generate changelog entry"
task :changelog, [:before, :after] do |_, params|
  args = []
  args << "--before #{params.before}" if params.before
  args << "--after #{params.after}" if params.after

  sh "ruby _tools/changelog_generator/changelog_generator.rb #{args.join(" ")}"
end

task build_all: %w[ruby:build_all go:build_all go_gem:test ruby_h_to_go:test]

task default: :build_all
