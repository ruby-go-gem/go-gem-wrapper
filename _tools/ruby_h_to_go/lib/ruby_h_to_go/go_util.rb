# frozen_string_literal: true

module RubyHToGo
  # helper methods for generating go code
  module GoUtil # rubocop:disable Metrics/ModuleLength
    # @param str [String]
    # @return [String]
    def self.snake_to_camel(str)
      return str if %w[VALUE ID].include?(str)

      str.split("_").map(&:capitalize).join.gsub(/(?<=\d)([a-z])/) { _1.upcase } # rubocop:disable Style/SymbolProc
    end

    # Generate initial go file whether not exists
    # @param go_file_path [String]
    def self.generate_initial_go_file(go_file_path)
      return if File.exist?(go_file_path)

      header = +<<~GO
        // THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

        // WARNING: This file has automatically been generated
        // Code generated by ruby_h_to_go. DO NOT EDIT.

      GO

      ruby_build_tag = GoGem::Util.ruby_minor_version_build_tag

      header <<
        if ruby_build_tag == RubyHToGo.config.default_tag
          other_tags = RubyHToGo.config.available_tags - [RubyHToGo.config.default_tag]
          condition = other_tags.map { |tag| "!#{tag}" }.join(" && ")

          <<~GO
            // FIXME: https://pkg.go.dev/ doesn't support custom build tag.
            //        Therefore, if no build tag is passed, treat it as the default tag
            //go:build #{ruby_build_tag} || (#{condition})

          GO
        else
          <<~GO
            //go:build #{ruby_build_tag}

          GO
        end

      header << <<~GO
        package ruby

        /*
        #include "ruby.h"
        */
        import "C"

        import (
          "unsafe"
        )

      GO

      File.binwrite(go_file_path, header)
    end

    C_TYPE_TO_GO_TYPE = {
      "RUBY_DATA_FUNC"     => "unsafe.Pointer",
      "long long"          => "Longlong",
      "rb_alloc_func_t"    => "unsafe.Pointer",
      "unsigned char"      => "Uchar",
      "unsigned int"       => "uint",
      "unsigned long"      => "uint",
      "unsigned long long" => "Ulonglong",
      "unsigned short"     => "Ushort",
    }.freeze

    # Convert C type to Go type. (used in wrapper function args and return type etc)
    # @param typename [String]
    # @param pos [Symbol,nil] :arg, :typeref, :return
    # @param pointer [Symbol,nil] pointer hint (:ref, :array, :ref_array, :function, :sref, :str_array, :in_ref, :raw)
    # @param pointer_length [Integer]
    # @return [String]
    def self.ruby_c_type_to_go_type(typename, pos: nil, pointer: nil, pointer_length: 0)
      return ruby_pointer_c_type_to_go_type(typename, pos:, pointer:, pointer_length:) if pointer

      return C_TYPE_TO_GO_TYPE[typename] if C_TYPE_TO_GO_TYPE[typename]

      case typename
      when /^[A-Z]+$/, "int"
        # e.g. VALUE
        return typename
      when "void"
        return "unsafe.Pointer" if pointer == :ref && type == :typeref
      end

      snake_to_camel(typename)
    end

    C_TYPE_TO_CGO_TYPE = {
      "RUBY_DATA_FUNC"       => "toCFunctionPointer",
      "long long"            => "C.longlong",
      "rb_io_wait_readwrite" => "C.enum_rb_io_wait_readwrite",
      "ruby_value_type"      => "C.enum_ruby_value_type",
      "unsigned char"        => "C.uchar",
      "unsigned int"         => "C.uint",
      "unsigned long"        => "C.ulong",
      "unsigned long long"   => "C.ulonglong",
      "unsigned short"       => "C.ushort",
      "st_hash_type"         => "C.struct_st_hash_type",
      "timespec"             => "C.struct_timespec",
      "timeval"              => "C.struct_timeval",
    }.freeze

    # Cast C type to cgo type. (Used in wrapper function)
    # @param typename [String]
    # @return [String]
    def self.cast_to_cgo_type(typename)
      return C_TYPE_TO_CGO_TYPE[typename] if C_TYPE_TO_CGO_TYPE[typename]

      "C.#{typename}"
    end

    # Convert pointer C type to Go type. (used in wrapper function args and return type etc)
    # @param typename [String]
    # @param pos [Symbol,nil] :arg, :typeref, :return
    # @param pointer [Symbol,nil] pointer hint (:ref, :array, :ref_array, :function, :sref, :str_array, :in_ref, :raw)
    # @param pointer_length [Integer]
    # @return [String]
    def self.ruby_pointer_c_type_to_go_type(typename, pos:, pointer:, pointer_length:)
      go_type_name =
        if typename == "int" && %i[return typeref].include?(pos)
          "Int"
        else
          ruby_c_type_to_go_type(typename, pos:, pointer: nil)
        end

      case pointer
      when :sref
        return "*unsafe.Pointer" if typename == "void" && pointer_length == 2

        return "#{"*" * pointer_length}#{go_type_name}"

      when :str_array
        return "[]string"

      when :array
        return "[]#{go_type_name}"

      when :ref_array
        return "[]*#{go_type_name}"

      when :ref
        if typename == "char"
          case pos
          when :arg, :typeref
            return "string"
          else
            return "char2String"
          end
        end
      end

      return "unsafe.Pointer" if typename == "void"

      "*#{go_type_name}"
    end

    private_class_method :ruby_pointer_c_type_to_go_type
  end
end
