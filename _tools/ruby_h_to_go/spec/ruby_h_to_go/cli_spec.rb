# frozen_string_literal: true

RSpec.describe RubyHToGo::Cli do
  include_context "uses temp dir"

  let(:cli) do
    RubyHToGo::Cli.new(
      header_dir: RbConfig::CONFIG["rubyhdrdir"],
      dist_dir:   temp_dir,
    )
  end

  describe "#perform" do
    subject { cli.perform }

    it { expect { subject }.not_to raise_error }
  end
end
