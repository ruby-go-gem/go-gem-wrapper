# frozen_string_literal: true

RSpec.describe GoGem::RakeTask do
  before { Rake::Task.clear }

  after { Rake::Task.clear }

  describe "defining tasks" do
    context "with default params" do
      let(:gem_name) { "my_gem" }

      subject do
        GoGem::RakeTask.new(gem_name)
        Rake::Task
      end

      it { should be_task_defined("go:test") }
    end
  end
end
