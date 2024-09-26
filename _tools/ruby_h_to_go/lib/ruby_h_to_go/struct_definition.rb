# frozen_string_literal: true

module RubyHToGo
  # Proxy class for generating go type
  class StructDefinition
    # @!attribute [r] header_dir
    #   @return [String]
    attr_reader :header_dir

    extend Forwardable

    def_delegators :@definition, :==, :name, :name=, :filepath, :filepath=

    include GeneratorHelper

    # @param definition [RubyHeaderParser::StructDefinition]
    # @param header_dir [String]
    def initialize(definition:, header_dir:)
      @definition = definition
      @header_dir = header_dir
    end

    # Write definition as go file
    # @param [String] dist_dir
    def write_go_file(dist_dir)
      go_file_path = File.join(dist_dir, "struct_generated.go")

      generate_initial_go_file(go_file_path)

      File.open(go_file_path, "a") do |f|
        f.write(generate_go_content)
      end
    end

    # @return [String]
    def generate_go_content
      go_type_name = snake_to_camel(name)

      <<~GO
        // #{go_type_name} is a type for passing `C.#{name}` in and out of package
        type #{go_type_name} C.#{name}

      GO
    end
  end
end
