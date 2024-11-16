# frozen_string_literal: true

namespace :ruby_h_to_go do
  desc "Run _tools/ruby_h_to_go test"
  task :test do
    Dir.chdir(File.join(repo_root, "_tools", "ruby_h_to_go")) do
      sh "rspec"
    end
  end
end

desc "Run _tools/ruby_h_to_go"
task :ruby_h_to_go do
  sh "./_tools/ruby_h_to_go/exe/ruby_h_to_go"
end
