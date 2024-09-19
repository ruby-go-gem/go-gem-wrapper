# frozen_string_literal: true

module RubyHeaderParser
  # Manager for `data.yml`
  class Data
    # @!attribute [r] data
    #   @return [Hash]
    attr_reader :data

    def initialize
      yaml = File.read(File.join(__dir__, "data.yml"))
      @data = YAML.safe_load(yaml, permitted_classes: [Regexp])
    end

    # @param function_name [String]
    # @param index [Integer] arg position (0 origin)
    # @return [Symbol] :ref, :array
    def function_arg_pointer_hint(function_name:, index:)
      pointer_hint = data["function"]["pointer_hint"].dig(function_name, index)
      return pointer_hint.to_sym if pointer_hint

      :ref
    end

    # Whether generate C function to go
    # @param function_name [String]
    # @return [Boolean]
    def should_generate_function?(function_name)
      function_name = function_name.downcase

      return false if data["function"]["deny_name"].any? { |format| format === function_name } # rubocop:disable Style/CaseEquality

      data["function"]["allow_name"].any? { |format| format === function_name } # rubocop:disable Style/CaseEquality
    end

    # Whether generate C struct to go
    # @param struct_name [String]
    # @return [Boolean]
    def should_generate_struct?(struct_name)
      struct_name = struct_name.downcase

      data["struct"]["allow_name"].any? { |format| format === struct_name } # rubocop:disable Style/CaseEquality
    end

    # Whether generate C type to go
    # @param type_name [String]
    # @return [Boolean]
    def should_generate_type?(type_name)
      type_name = type_name.downcase

      data["type"]["allow_name"].any? { |format| format === type_name } # rubocop:disable Style/CaseEquality
    end
  end
end