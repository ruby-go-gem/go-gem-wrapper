# frozen_string_literal: true

RSpec.describe RubyHToGo::FunctionDefinition do
  describe "#generate_go_content" do
    subject { RubyHToGo::FunctionDefinition.new(definition:).generate_go_content }

    context "rb_define_method" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_define_method",
          definition: "void rb_define_method(VALUE klass, const char *mid, VALUE (*func)(ANYARGS), int arity)",
          typeref:    typeref(type: "void"),
          args:       [
            argument(type: "VALUE", name: "klass"),
            argument(type: "char", name: "mid", pointer: :ref),
            argument(type: "void", name: "arg3", pointer: :function),
            argument(type: "int", name: "arity"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbDefineMethod calls `rb_define_method` in C
          //
          // Original definition is following
          //
          //	void rb_define_method(VALUE klass, const char *mid, VALUE (*func)(ANYARGS), int arity)
          func RbDefineMethod(klass VALUE, mid string, arg3 unsafe.Pointer, arity int)  {
          char, clean := string2Char(mid)
          defer clean()

          C.rb_define_method(C.VALUE(klass), char, toCFunctionPointer(arg3), C.int(arity))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_block_call" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_block_call",
          definition: "VALUE rb_block_call(VALUE obj, ID mid, int argc, const VALUE *argv, rb_block_call_func_t proc, VALUE data2)", # rubocop:disable Layout/LineLength
          typeref:    typeref(type: "VALUE"),
          args:       [
            argument(type: "VALUE", name: "obj"),
            argument(type: "ID", name: "mid"),
            argument(type: "int", name: "argc"),
            argument(type: "VALUE", name: "argv", pointer: :ref),
            argument(type: "rb_block_call_func_t", name: "proc"),
            argument(type: "VALUE", name: "data2"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbBlockCall calls `rb_block_call` in C
          //
          // Original definition is following
          //
          //	VALUE rb_block_call(VALUE obj, ID mid, int argc, const VALUE *argv, rb_block_call_func_t proc, VALUE data2)
          func RbBlockCall(obj VALUE, mid ID, argc int, argv *VALUE, proc RbBlockCallFuncT, data2 VALUE) VALUE {
          var cArgv C.VALUE
          ret := VALUE(C.rb_block_call(C.VALUE(obj), C.ID(mid), C.int(argc), &cArgv, C.rb_block_call_func_t(proc), C.VALUE(data2)))
          *argv = VALUE(cArgv)
          return ret
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_funcallv" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_funcallv",
          definition: "VALUE rb_funcallv(VALUE recv, ID mid, int argc, const VALUE *argv)",
          typeref:    typeref(type: "VALUE"),
          args:       [
            argument(type: "VALUE", name: "recv"),
            argument(type: "ID", name: "mid"),
            argument(type: "int", name: "argc"),
            argument(type: "VALUE", name: "argv", pointer: :array),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbFuncallv calls `rb_funcallv` in C
          //
          // Original definition is following
          //
          //	VALUE rb_funcallv(VALUE recv, ID mid, int argc, const VALUE *argv)
          func RbFuncallv(recv VALUE, mid ID, argc int, argv []VALUE) VALUE {
          return VALUE(C.rb_funcallv(C.VALUE(recv), C.ID(mid), C.int(argc), toCArray[VALUE, C.VALUE](argv)))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_thread_call_with_gvl" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_thread_call_with_gvl",
          definition: "void *rb_thread_call_with_gvl(void *(*func)(void *), void *data1)",
          typeref:    typeref(type: "void", pointer: :ref),
          args:       [
            argument(type: "void", name: "arg1", pointer: :function),
            argument(type: "void", name: "data1", pointer: :ref),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbThreadCallWithGvl calls `rb_thread_call_with_gvl` in C
          //
          // Original definition is following
          //
          //	void *rb_thread_call_with_gvl(void *(*func)(void *), void *data1)
          func RbThreadCallWithGvl(arg1 unsafe.Pointer, data1 unsafe.Pointer) unsafe.Pointer {
          return unsafe.Pointer(C.rb_thread_call_with_gvl(toCFunctionPointer(arg1), data1))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_uv_to_utf8" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_uv_to_utf8",
          definition: "int rb_uv_to_utf8(char buf[6], unsigned long uv)",
          typeref:    typeref(type: "int"),
          args:       [
            argument(type: "char", name: "buf", pointer: :array, length: 6),
            argument(type: "unsigned long", name: "uv"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbUvToUtf8 calls `rb_uv_to_utf8` in C
          //
          // Original definition is following
          //
          //	int rb_uv_to_utf8(char buf[6], unsigned long uv)
          func RbUvToUtf8(buf []Char, uv uint) int {
          return int(C.rb_uv_to_utf8(toCArray[Char, C.char](buf), C.ulong(uv)))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_errno_ptr" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_errno_ptr",
          definition: "int *rb_errno_ptr(void)",
          typeref:    typeref(type: "int", pointer: :ref),
          args:       [],
        )
      end

      let(:go_content) do
        <<~GO
          // RbErrnoPtr calls `rb_errno_ptr` in C
          //
          // Original definition is following
          //
          //	int *rb_errno_ptr(void)
          func RbErrnoPtr() *Int {
          return (*Int)(C.rb_errno_ptr())
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_big2ll" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_big2ll",
          definition: "rb_big2ll(VALUE)",
          typeref:    typeref(type: "long long"),
          args:       [
            argument(type: "VALUE", name: "arg1"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbBig2Ll calls `rb_big2ll` in C
          //
          // Original definition is following
          //
          //	rb_big2ll(VALUE)
          func RbBig2Ll(arg1 VALUE) Longlong {
          return Longlong(C.rb_big2ll(C.VALUE(arg1)))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_big2ull" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_big2ull",
          definition: "rb_big2ull(VALUE)",
          typeref:    typeref(type: "unsigned long long"),
          args:       [
            argument(type: "VALUE", name: "arg1"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbBig2Ull calls `rb_big2ull` in C
          //
          // Original definition is following
          //
          //	rb_big2ull(VALUE)
          func RbBig2Ull(arg1 VALUE) Ulonglong {
          return Ulonglong(C.rb_big2ull(C.VALUE(arg1)))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_scan_args_set" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_scan_args_set",
          definition: "rb_scan_args_set(int kw_flag, int argc, const VALUE *argv,",
          typeref:    typeref(type: "int"),
          args:       [
            argument(type: "int", name: "kw_flag"),
            argument(type: "int", name: "argc"),
            argument(type: "VALUE", name: "argv", pointer: :ref),
            argument(type: "int", name: "n_lead"),
            argument(type: "int", name: "n_opt"),
            argument(type: "int", name: "n_trail"),
            argument(type: "_Bool", name: "f_var"),
            argument(type: "_Bool", name: "f_hash"),
            argument(type: "_Bool", name: "f_block"),
            argument(type: "VALUE", name: "vars", pointer: :ref_array),
            argument(type: "char", name: "fmt", pointer: :ref),
            argument(type: "int", name: "varc"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbScanArgsSet calls `rb_scan_args_set` in C
          //
          // Original definition is following
          //
          //	rb_scan_args_set(int kw_flag, int argc, const VALUE *argv,
          func RbScanArgsSet(kw_flag int, argc int, argv *VALUE, n_lead int, n_opt int, n_trail int, f_var Bool, f_hash Bool, f_block Bool, vars []*VALUE, fmt string, varc int) int {
          var cArgv C.VALUE
          char, clean := string2Char(fmt)
          defer clean()

          ret := int(C.rb_scan_args_set(C.int(kw_flag), C.int(argc), &cArgv, C.int(n_lead), C.int(n_opt), C.int(n_trail), C._Bool(f_var), C._Bool(f_hash), C._Bool(f_block), toCArray[*VALUE, *C.VALUE](vars), char, C.int(varc)))
          *argv = VALUE(cArgv)
          return ret
          }

        GO
      end

      it { should eq go_content }
    end

    context "RSTRING_END" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "RSTRING_END",
          definition: "RSTRING_END(VALUE str)",
          typeref:    typeref(type: "char", pointer: :raw),
          args:       [
            argument(type: "VALUE", name: "str"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RSTRING_END calls `RSTRING_END` in C
          //
          // Original definition is following
          //
          //	RSTRING_END(VALUE str)
          func RSTRING_END(str VALUE) *Char {
          return (*Char)(C.RSTRING_END(C.VALUE(str)))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_const_list" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_const_list",
          definition: "VALUE rb_const_list(void*)",
          typeref:    typeref(type: "VALUE"),
          args:       [
            argument(type: "void", name: "arg1", pointer: :ref),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbConstList calls `rb_const_list` in C
          //
          // Original definition is following
          //
          //	VALUE rb_const_list(void*)
          func RbConstList(arg1 unsafe.Pointer) VALUE {
          return VALUE(C.rb_const_list(arg1))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_feature_provided" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_feature_provided",
          definition: "int rb_feature_provided(const char *feature, const char **loading)",
          typeref:    typeref(type: "int"),
          args:       [
            argument(type: "char", name: "feature", pointer: :ref),
            argument(type: "char", name: "loading", pointer: :sref, length: 2),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbFeatureProvided calls `rb_feature_provided` in C
          //
          // Original definition is following
          //
          //	int rb_feature_provided(const char *feature, const char **loading)
          func RbFeatureProvided(feature string, loading **Char) int {
          char, clean := string2Char(feature)
          defer clean()

          return int(C.rb_feature_provided(char, (**C.char)(unsafe.Pointer(loading))))
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_find_file_ext" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_find_file_ext",
          definition: "int rb_find_file_ext(VALUE *feature, const char *const *exts)",
          typeref:    typeref(type: "int"),
          args:       [
            argument(type: "VALUE", name: "feature", pointer: :ref),
            argument(type: "char", name: "exts", pointer: :str_array),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbFindFileExt calls `rb_find_file_ext` in C
          //
          // Original definition is following
          //
          //	int rb_find_file_ext(VALUE *feature, const char *const *exts)
          func RbFindFileExt(feature *VALUE, exts []string) int {
          var cFeature C.VALUE
          chars, cleanChars := strings2Chars(exts)
          defer cleanChars()

          ret := int(C.rb_find_file_ext(&cFeature, chars))
          *feature = VALUE(cFeature)
          return ret
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_data_typed_object_make" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_data_typed_object_make",
          definition: "rb_data_typed_object_make(VALUE klass, const rb_data_type_t *type, void **datap, size_t size)",
          typeref:    typeref(type: "VALUE"),
          args:       [
            argument(type: "VALUE", name: "klass"),
            argument(type: "rb_data_type_t", name: "type", pointer: :ref),
            argument(type: "void", name: "datap", pointer: :sref, length: 2),
            argument(type: "size_t", name: "size"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbDataTypedObjectMake calls `rb_data_typed_object_make` in C
          //
          // Original definition is following
          //
          //	rb_data_typed_object_make(VALUE klass, const rb_data_type_t *type, void **datap, size_t size)
          func RbDataTypedObjectMake(klass VALUE, t *RbDataTypeT, datap *unsafe.Pointer, size SizeT) VALUE {
          var cT C.rb_data_type_t
          ret := VALUE(C.rb_data_typed_object_make(C.VALUE(klass), &cT, datap, C.size_t(size)))
          *t = RbDataTypeT(cT)
          return ret
          }

        GO
      end

      it { should eq go_content }
    end

    context "rb_define_variable" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "rb_define_variable",
          definition: "void rb_define_variable(const char *name, VALUE *var)",
          typeref:    typeref(type: "void"),
          args:       [
            argument(type: "char", name: "name", pointer: :ref),
            argument(type: "VALUE", name: "var", pointer: :in_ref),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RbDefineVariable calls `rb_define_variable` in C
          //
          // Original definition is following
          //
          //	void rb_define_variable(const char *name, VALUE *var)
          func RbDefineVariable(name string, v *VALUE)  {
          char, clean := string2Char(name)
          defer clean()

          C.rb_define_variable(char, (*C.VALUE)(v))
          }

        GO
      end

      it { should eq go_content }
    end

    context "RSTRING_PTR" do
      let(:definition) do
        RubyHeaderParser::FunctionDefinition.new(
          name:       "RSTRING_PTR",
          definition: "RSTRING_PTR(VALUE str)",
          typeref:    typeref(type: "char", pointer: :raw),
          args:       [
            argument(type: "VALUE", name: "str"),
          ],
        )
      end

      let(:go_content) do
        <<~GO
          // RSTRING_PTR calls `RSTRING_PTR` in C
          //
          // Original definition is following
          //
          //	RSTRING_PTR(VALUE str)
          func RSTRING_PTR(str VALUE) *Char {
          return (*Char)(C.RSTRING_PTR(C.VALUE(str)))
          }

        GO
      end

      it { should eq go_content }
    end
  end

  describe "#go_function_name" do
    subject { RubyHToGo::FunctionDefinition.new(definition:).go_function_name }

    let(:definition) do
      RubyHeaderParser::FunctionDefinition.new(
        name:,
        definition: "",
        typeref:    typeref(type: "void"),
        args:       [],
      )
    end

    using RSpec::Parameterized::TableSyntax

    where(:name, :expected) do
      "RB_FIX2INT" | "RB_FIX2INT"
      "rb_fix2int" | "RbFix2Int"
    end

    with_them do
      it { should eq expected }
    end
  end
end
