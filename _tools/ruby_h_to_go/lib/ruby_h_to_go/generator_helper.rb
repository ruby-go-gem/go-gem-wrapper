# frozen_string_literal: true

module RubyHToGo
  # helper methods for generating go code
  module GeneratorHelper
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
        // THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

        // WARNING: This file has automatically been generated
        // Code generated by ruby_h_to_go. DO NOT EDIT.

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
      return ruby_pointer_c_type_to_go_type(typename, type:, pointer:) if pointer

      case typename
      when "unsigned int", "unsigned long"
        return "uint"
      when "unsigned short"
        return "Ushort"
      when "unsigned char"
        return "Uchar"
      when "long long"
        return "Longlong"
      when "unsigned long long"
        return "Ulonglong"
      when /^VALUE\s*\(\*func\)\s*\(ANYARGS\)$/, "RUBY_DATA_FUNC"
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
      when "long long"
        return "C.longlong"
      when "unsigned long long"
        return "C.ulonglong"
      when "timeval"
        return "C.struct_timeval"
      when "timespec"
        return "C.struct_timespec"
      when "st_hash_type"
        return "C.struct_st_hash_type"
      when "ruby_value_type"
        return "C.enum_ruby_value_type"
      when "rb_io_wait_readwrite"
        return "C.enum_rb_io_wait_readwrite"
      when /^VALUE\s*\(\*func\)\s*\(ANYARGS\)$/, "RUBY_DATA_FUNC"
        return "toCPointer"
      end

      "C.#{typename}"
    end

    private

    # Convert pointer C type to Go type. (used in wrapper function args and return type etc)
    # @param typename [String]
    # @param type [Symbol,nil] :arg, :return
    # @param pointer [Symbol,nil] Whether pointer hint
    # @return [String]
    def ruby_pointer_c_type_to_go_type(typename, type:, pointer:)
      case typename
      when "char", "const char"
        if pointer == :ref
          case type
          when :arg, :return
            return "string"
          else
            return "char2String"
          end
        end
      when "void"
        return "unsafe.Pointer"
      end

      go_type_name = ruby_c_type_to_go_type(typename, type:, pointer: nil)

      case pointer
      when :array
        return "[]#{go_type_name}"
      when :ref_array
        return "[]*#{go_type_name}"
      end

      "*#{go_type_name}"
    end
  end
end
