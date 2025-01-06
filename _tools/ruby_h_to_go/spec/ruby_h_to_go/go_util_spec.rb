# frozen_string_literal: true

RSpec.describe RubyHToGo::GoUtil do
  describe ".ruby_minor_version_build_tag" do
    subject { RubyHToGo::GoUtil.ruby_minor_version_build_tag(ruby_version) }

    let(:ruby_version) { "3.4.1" }

    it { should eq "ruby_3_4" }
  end
end
