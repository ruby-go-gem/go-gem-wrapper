# frozen_string_literal: true

namespace :go_gem do
  desc "Run go_gem test"
  task :test do
    Dir.chdir(File.join(repo_root, "gem")) do
      sh "rspec"
    end
  end
end
