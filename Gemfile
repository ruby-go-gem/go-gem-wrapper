# frozen_string_literal: true

source "https://rubygems.org"

group :development do
  gem "rake"
  gem "rdoc"
  gem "rubocop", require: false
  gem "rubocop_auto_corrector", require: false
  gem "rubocop-on-rbs", require: false
  gem "ruby_header_parser", ">= 0.4.2"
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
  gem "uri", ">= 1.0.3"
end

gemspec path: "./_gem/"
