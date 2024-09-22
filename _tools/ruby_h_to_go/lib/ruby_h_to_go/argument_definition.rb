# frozen_string_literal: true

module RubyHToGo
  # Proxy class for generating argument in go function
  class ArgumentDefinition
    # @!attribute [r] header_dir
    #   @return [String]
    attr_reader :header_dir

    extend Forwardable

    def_delegators :@definition, :==, :type, :type=, :name, :name=, :pointer, :pointer=, :pointer?, :length, :length=

    include GeneratorHelper

    # @param definition [RubyHeaderParser::ArgumentDefinition]
    # @param header_dir [String]
    def initialize(definition:, header_dir:)
      @definition = definition
      @header_dir = header_dir
    end

    C_NAME_TO_GO_NAME = {
      # These are reserved in Go
      "var"   => "v",
      "func"  => "fun",
      "range" => "r",
      "type"  => "r",

      # Can't use "_" as a value
      "_"     => "arg",
    }.freeze

    # @return [String] Variable name available in Go
    def go_name
      return C_NAME_TO_GO_NAME[name] if C_NAME_TO_GO_NAME[name]

      name
    end

    # @return [String]
    def go_function_arg
      "#{go_name} #{ruby_c_type_to_go_type(type, pointer:, type: :arg)}"
    end

    # @return [String]
    def cast_to_cgo
      return "toCArray[#{ruby_c_type_to_go_type(type)}, #{cast_to_cgo_type(type)}](#{go_name})" if pointer == :array

      "#{cast_to_cgo_type(type)}(#{go_name})"
    end
  end
end
