# frozen_string_literal: true

module RubyHToGo
  class Config
    def initialize
      @config = YAML.load_file(File.expand_path("../../config.yml", __dir__))
    end

    # @return [String]
    def default_tag
      @config["default_tag"]
    end

    # @return [Array<String>]
    def available_tags
      @config["available_tags"]
    end
  end
end
