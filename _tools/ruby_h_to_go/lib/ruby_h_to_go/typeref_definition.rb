# frozen_string_literal: true

module RubyHToGo
  # Proxy class for generating typeref in go function
  class TyperefDefinition
    # @!attribute [r] header_dir
    #   @return [String]
    attr_reader :header_dir

    extend Forwardable

    def_delegators :@definition, :==, :type, :type=, :pointer, :pointer=, :pointer?

    include GeneratorHelper

    # @param definition [RubyHeaderParser::TyperefDefinition]
    def initialize(definition:, header_dir:)
      @definition = definition
      @header_dir = header_dir
    end

    # @return [String]
    def go_function_typeref
      return "" if type == "void" && !pointer?

      ruby_c_type_to_go_type(type, type: :return, pointer:)
    end
  end
end
