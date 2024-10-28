# frozen_string_literal: true
namespace :patch_for_go_gem do
  desc "Run _tools/patch_for_go_gem test"
  task :test do
    Dir.chdir(File.join(repo_root, "_tools", "patch_for_go_gem")) do
      sh "rspec"
    end
  end
end
