# frozen_string_literal: true

# @return [String]
def repo_root
  __dir__
end

Dir["#{__dir__}/tasks/*.rake"].each { |f| load f }

desc "Release package"
task :release do
  Dir.chdir(File.join(__dir__, "gem")) do
    sh "rake release"
  end
end

task build_all: %w[ruby:build_all go:build_all go_gem:test ruby_h_to_go:test patch_for_go_gem:test]

task default: :build_all
