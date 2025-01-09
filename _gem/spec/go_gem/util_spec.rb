# frozen_string_literal: true

RSpec.describe GoGem::Util do
  describe ".ruby_minor_version_build_tag" do
    subject { GoGem::Util.ruby_minor_version_build_tag(ruby_version) }

    let(:ruby_version) { "3.4.1" }

    it { should eq "ruby_3_4" }
  end
end
