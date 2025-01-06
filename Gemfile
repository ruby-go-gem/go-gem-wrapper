# frozen_string_literal: true

source "https://rubygems.org"

group :development do
  gem "rake"
  gem "rubocop", require: false
  gem "rubocop_auto_corrector", require: false
  # gem "ruby_header_parser", ">= 0.4.0"
  gem "ruby_header_parser", github: "ruby-go-gem/ruby_header_parser", branch: "fix_github_actions"
  gem "yard"
end

group :test do
  gem "rspec"
  gem "rspec-its"
  gem "rspec-parameterized"
  gem "rspec-temp_dir"
  gem "serverspec"

  # for ruby/testdata/example/
  gem "rake-compiler"
  gem "steep"
  gem "test-unit"
end

gemspec path: "./_gem/"
