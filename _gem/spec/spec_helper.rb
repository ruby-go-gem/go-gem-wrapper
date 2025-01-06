# frozen_string_literal: true

require "go_gem"
require "go_gem/mkmf"
require "go_gem/rake_task"
require "go_gem/util"

require "tmpdir"
require "serverspec"
require "mkmf"

set :backend, :exec

RSpec.configure do |config|
  # Enable flags like --only-failures and --next-failure
  config.example_status_persistence_file_path = ".rspec_status"

  # Disable RSpec exposing methods globally on `Module` and `main`
  config.disable_monkey_patching!

  config.expect_with :rspec do |c|
    c.syntax = :expect
  end

  config.filter_run_when_matching :focus

  config.define_derived_metadata do |meta|
    meta[:aggregate_failures] = true
  end
end
