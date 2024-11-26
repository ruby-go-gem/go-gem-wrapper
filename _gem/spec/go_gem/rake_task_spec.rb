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
      it { should be_task_defined("go:testrace") }
      it { should be_task_defined("go:fmt") }
      it { should be_task_defined("go:build_envs") }

      describe "Add additional tasks" do
        include Rake::DSL

        subject do
          t = GoGem::RakeTask.new(gem_name)

          namespace :go do
            task :test2 do
              t.within_target_dir do
                sh "go test"
              end
            end
          end

          Rake::Task
        end

        it { should be_task_defined("go:test2") }
      end
    end

    context "with params" do
      let(:gem_name) { "my_gem" }

      subject do
        GoGem::RakeTask.new(gem_name) do |config|
          config.task_namespace = :go5
        end
        Rake::Task
      end

      it { should be_task_defined("go5:test") }
      it { should be_task_defined("go5:testrace") }
      it { should be_task_defined("go5:fmt") }
    end
  end
end
