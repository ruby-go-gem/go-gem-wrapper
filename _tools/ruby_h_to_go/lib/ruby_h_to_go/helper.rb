# frozen_string_literal: true

module RubyHToGo
  # helper methods for generating go coce
  module Helper
    # @param header_dir [String]
    # @param ruby_header_file [String]
    # @return [String]
    def go_file_name(header_dir:, ruby_header_file:)
      ruby_header_file.delete_prefix(header_dir + File::SEPARATOR).gsub(File::SEPARATOR, "-").gsub(/\.h$/, ".go")
    end

    # @param str [String]
    # @return [String]
    def snake_to_camel(str)
      return str if %w[VALUE ID].include?(str)

      str.split("_").map(&:capitalize).join.gsub(/(?<=\d)([a-z])/) { _1.upcase } # rubocop:disable Style/SymbolProc
    end

    # Generate initial go file whether not exists
    # @param go_file_path [String]
    def generate_initial_go_file(go_file_path)
      return if File.exist?(go_file_path)

      File.binwrite(go_file_path, <<~GO)
        package ruby

        /*
        #include "ruby.h"
        */
        import "C"

        import (
          "unsafe"
        )

      GO
    end

    # Convert C type to Go type. (used in wrapper function args and return type etc)
    # @param typename [String]
    # @param type [Symbol,nil] :arg, :return
    # @param pointer [Symbol,nil] Whether pointer hint
    # @return [String]
    def ruby_c_type_to_go_type(typename, type: nil, pointer: nil)
      typename = typename.delete_prefix("struct ").delete_prefix("volatile ")

      if pointer
        case typename
        when "char", "const char"
          case type
          when :arg, :return
            return "string"
          else
            return "char2String"
          end
        when "void"
          return "unsafe.Pointer"
        end

        go_type_name = ruby_c_type_to_go_type(typename, type:, pointer: nil)

        return "[]#{go_type_name}" if pointer == :array

        return "*#{go_type_name}"
      end

      case typename
      when "unsigned int", "unsigned long"
        return "uint"
      when "unsigned short"
        return "Ushort"
      when "unsigned char"
        return "Uchar"
      when "LONG_LONG"
        return "Longlong"
      when "unsigned LONG_LONG"
        return "Ulonglong"
      when /^VALUE\s*\(\*func\)\s*\(ANYARGS\)$/
        return "unsafe.Pointer"
      when /^[A-Z]+$/, "int"
        # e.g. VALUE
        return typename
      when "void"
        return "unsafe.Pointer" if pointer == :ref && type == :return
      end

      snake_to_camel(typename)
    end

    # Cast C type to cgo type. (Used in wrapper function)
    # @param typename [String]
    # @return [String]
    def cast_to_cgo_type(typename)
      case typename
      when "unsigned long"
        return "C.ulong"
      when "unsigned int"
        return "C.uint"
      when "unsigned char"
        return "C.uchar"
      when "unsigned short"
        return "C.ushort"
      when "LONG_LONG"
        return "C.Longlong"
      when "unsigned LONG_LONG"
        return "C.Ulonglong"
      when "VALUE*"
        return "toCValueArray"
      when /^VALUE\s*\(\*func\)\s*\(ANYARGS\)$/
        return "toCPointer"
      end

      "C.#{typename}"
    end
  end
end
