package main

/*
#include "example.h"

void  rb_example_tests_nop_rb_define_method_id(VALUE self);
void  rb_example_tests_nop_rb_define_private_method(VALUE self);
void  rb_example_tests_nop_rb_define_protected_method(VALUE self);
VALUE rb_example_tests_rb_ivar_get(VALUE self);
void  rb_example_tests_rb_ivar_set(VALUE self, VALUE value);
VALUE rb_example_tests_rb_yield(VALUE self, VALUE arg);
VALUE rb_example_tests_rb_block_proc(VALUE self, VALUE arg);
VALUE rb_example_tests_rb_funcall2(VALUE self, VALUE num, VALUE ndigits);
VALUE rb_example_tests_rb_funcall3(VALUE self, VALUE num, VALUE ndigits);
void  rb_example_tests_rb_alias(VALUE self, VALUE dst, VALUE src);
VALUE rb_example_tests_rb_class2name(VALUE self);
void  rb_example_tests_rb_attr(VALUE self, VALUE name, VALUE needReader, VALUE needWriter, VALUE honourVisibility);
VALUE rb_example_tests_rb_const_get(VALUE self, VALUE name);
VALUE rb_example_tests_rb_const_get_at(VALUE self, VALUE name);
void  rb_example_tests_rb_const_set(VALUE self, VALUE name, VALUE val);
VALUE rb_example_tests_rb_const_defined(VALUE self, VALUE name);
VALUE rb_example_tests_rb_const_defined_at(VALUE self, VALUE name);
VALUE rb_example_tests_rb_eval_string(VALUE self, VALUE str);
VALUE rb_example_tests_rb_eval_string_protect(VALUE self, VALUE str);
VALUE rb_example_tests_rb_eval_string_wrap(VALUE self, VALUE str);
VALUE rb_example_tests_rb_ary_new(VALUE self);
VALUE rb_example_tests_rb_ary_new_capa(VALUE self, VALUE capa);
VALUE rb_example_tests_rb_ary_push(VALUE self, VALUE ary, VALUE elem);
VALUE rb_example_tests_rb_ary_pop(VALUE self, VALUE ary);
VALUE rb_example_tests_rb_ary_shift(VALUE self, VALUE ary);
VALUE rb_example_tests_rb_ary_unshift(VALUE self, VALUE ary, VALUE elem);
void  rb_example_tests_rb_define_variable(VALUE self, VALUE name, VALUE v);
void  rb_example_tests_rb_define_const(VALUE self, VALUE name, VALUE val);
*/
import "C"

import (
	"github.com/ruby-go-gem/go-gem-wrapper/ruby"
)

//export rb_example_tests_nop_rb_define_method_id
func rb_example_tests_nop_rb_define_method_id(_ C.VALUE) {
}

//export rb_example_tests_nop_rb_define_private_method
func rb_example_tests_nop_rb_define_private_method(_ C.VALUE) {
}

//export rb_example_tests_nop_rb_define_protected_method
func rb_example_tests_nop_rb_define_protected_method(_ C.VALUE) {
}

//export rb_example_tests_rb_ivar_get
func rb_example_tests_rb_ivar_get(self C.VALUE) C.VALUE {
	ivarID := ruby.RbIntern("@ivar")
	ivarValue := ruby.RbIvarGet(ruby.VALUE(self), ivarID)

	return C.VALUE(ivarValue)
}

//export rb_example_tests_rb_ivar_set
func rb_example_tests_rb_ivar_set(self C.VALUE, value C.VALUE) {
	ivarID := ruby.RbIntern("@ivar")
	ruby.RbIvarSet(ruby.VALUE(self), ivarID, ruby.VALUE(value))
}

//export rb_example_tests_rb_yield
func rb_example_tests_rb_yield(_ C.VALUE, arg C.VALUE) C.VALUE {
	if ruby.RbBlockGivenP() == 0 {
		ruby.RbRaise(ruby.VALUE(C.rb_eArgError), "Block not given")
	}

	blockResult := ruby.RbYield(ruby.VALUE(arg))
	return C.VALUE(blockResult)
}

//export rb_example_tests_rb_block_proc
func rb_example_tests_rb_block_proc(_ C.VALUE, arg C.VALUE) C.VALUE {
	if ruby.RbBlockGivenP() == 0 {
		ruby.RbRaise(ruby.VALUE(C.rb_eArgError), "Block not given")
	}

	block := ruby.RbBlockProc()

	// Call Proc#call
	blockResult := ruby.RbFuncall2(ruby.VALUE(block), ruby.RbIntern("call"), 1, []ruby.VALUE{ruby.VALUE(arg)})

	return C.VALUE(blockResult)
}

//export rb_example_tests_rb_funcall2
func rb_example_tests_rb_funcall2(_ C.VALUE, num C.VALUE, ndigits C.VALUE) C.VALUE {
	// Call Integer#round
	result := ruby.RbFuncall2(ruby.VALUE(num), ruby.RbIntern("round"), 1, []ruby.VALUE{ruby.VALUE(ndigits)})

	return C.VALUE(result)
}

//export rb_example_tests_rb_funcall3
func rb_example_tests_rb_funcall3(_ C.VALUE, num C.VALUE, ndigits C.VALUE) C.VALUE {
	// Call Integer#round
	result := ruby.RbFuncall3(ruby.VALUE(num), ruby.RbIntern("round"), 1, []ruby.VALUE{ruby.VALUE(ndigits)})

	return C.VALUE(result)
}

//export rb_example_tests_rb_alias
func rb_example_tests_rb_alias(klass C.VALUE, dst C.VALUE, src C.VALUE) {
	dstName := ruby.Value2String(ruby.VALUE(dst))
	dstID := ruby.RbIntern(dstName)

	srcName := ruby.Value2String(ruby.VALUE(src))
	srcID := ruby.RbIntern(srcName)

	ruby.RbAlias(ruby.VALUE(klass), dstID, srcID)
}

//export rb_example_tests_rb_class2name
func rb_example_tests_rb_class2name(klass C.VALUE) C.VALUE {
	str := ruby.RbClass2Name(ruby.VALUE(klass))
	value := ruby.String2Value(str)
	return C.VALUE(value)
}

//export rb_example_tests_rb_attr
func rb_example_tests_rb_attr(klass C.VALUE, name C.VALUE, needReader C.VALUE, needWriter C.VALUE, honourVisibility C.VALUE) {
	ivarName := ruby.Value2String(ruby.VALUE(name))
	intNeedReader := ruby.NUM2INT(ruby.VALUE(needReader))
	intNeedWriter := ruby.NUM2INT(ruby.VALUE(needWriter))
	intHonourVisibility := ruby.NUM2INT(ruby.VALUE(honourVisibility))

	ruby.RbAttr(ruby.VALUE(klass), ruby.RbIntern(ivarName), intNeedReader, intNeedWriter, intHonourVisibility)
}

//export rb_example_tests_rb_const_get
func rb_example_tests_rb_const_get(klass C.VALUE, name C.VALUE) C.VALUE {
	constName := ruby.Value2String(ruby.VALUE(name))
	constID := ruby.RbIntern(constName)
	return C.VALUE(ruby.RbConstGet(ruby.VALUE(klass), constID))
}

//export rb_example_tests_rb_const_get_at
func rb_example_tests_rb_const_get_at(klass C.VALUE, name C.VALUE) C.VALUE {
	constName := ruby.Value2String(ruby.VALUE(name))
	constID := ruby.RbIntern(constName)
	return C.VALUE(ruby.RbConstGetAt(ruby.VALUE(klass), constID))
}

//export rb_example_tests_rb_const_set
func rb_example_tests_rb_const_set(klass C.VALUE, name C.VALUE, val C.VALUE) {
	constName := ruby.Value2String(ruby.VALUE(name))
	constID := ruby.RbIntern(constName)

	ruby.RbConstSet(ruby.VALUE(klass), constID, ruby.VALUE(val))
}

//export rb_example_tests_rb_const_defined
func rb_example_tests_rb_const_defined(klass C.VALUE, name C.VALUE) C.VALUE {
	constName := ruby.Value2String(ruby.VALUE(name))
	constID := ruby.RbIntern(constName)

	defined := ruby.RbConstDefined(ruby.VALUE(klass), constID)
	if defined != 0 {
		return C.VALUE(ruby.Qtrue())
	}

	return C.VALUE(ruby.Qfalse())
}

//export rb_example_tests_rb_const_defined_at
func rb_example_tests_rb_const_defined_at(klass C.VALUE, name C.VALUE) C.VALUE {
	constName := ruby.Value2String(ruby.VALUE(name))
	constID := ruby.RbIntern(constName)

	defined := ruby.RbConstDefinedAt(ruby.VALUE(klass), constID)
	if defined != 0 {
		return C.VALUE(ruby.Qtrue())
	}

	return C.VALUE(ruby.Qfalse())
}

//export rb_example_tests_rb_eval_string
func rb_example_tests_rb_eval_string(_ C.VALUE, str C.VALUE) C.VALUE {
	goStr := ruby.Value2String(ruby.VALUE(str))
	ret := ruby.RbEvalString(goStr)

	return C.VALUE(ret)
}

//export rb_example_tests_rb_eval_string_protect
func rb_example_tests_rb_eval_string_protect(_ C.VALUE, str C.VALUE) C.VALUE {
	goStr := ruby.Value2String(ruby.VALUE(str))

	var state int
	ret := ruby.RbEvalStringProtect(goStr, &state)

	slice := []ruby.VALUE{ret, ruby.INT2NUM(state)}
	ary := ruby.Slice2rbAry(slice)

	return C.VALUE(ary)
}

//export rb_example_tests_rb_eval_string_wrap
func rb_example_tests_rb_eval_string_wrap(_ C.VALUE, str C.VALUE) C.VALUE {
	goStr := ruby.Value2String(ruby.VALUE(str))

	var state int
	ret := ruby.RbEvalStringWrap(goStr, &state)

	slice := []ruby.VALUE{ret, ruby.INT2NUM(state)}
	ary := ruby.Slice2rbAry(slice)

	return C.VALUE(ary)
}

//export rb_example_tests_rb_ary_new
func rb_example_tests_rb_ary_new(_ C.VALUE) C.VALUE {
	ret := ruby.RbAryNew()

	return C.VALUE(ret)
}

//export rb_example_tests_rb_ary_new_capa
func rb_example_tests_rb_ary_new_capa(_ C.VALUE, capa C.VALUE) C.VALUE {
	longCapa := ruby.NUM2LONG(ruby.VALUE(capa))
	ret := ruby.RbAryNewCapa(ruby.Long(longCapa))

	return C.VALUE(ret)
}

//export rb_example_tests_rb_ary_push
func rb_example_tests_rb_ary_push(_ C.VALUE, ary C.VALUE, elem C.VALUE) C.VALUE {
	ret := ruby.RbAryPush(ruby.VALUE(ary), ruby.VALUE(elem))

	return C.VALUE(ret)
}

//export rb_example_tests_rb_ary_pop
func rb_example_tests_rb_ary_pop(_ C.VALUE, ary C.VALUE) C.VALUE {
	ret := ruby.RbAryPop(ruby.VALUE(ary))

	return C.VALUE(ret)
}

//export rb_example_tests_rb_ary_shift
func rb_example_tests_rb_ary_shift(_ C.VALUE, ary C.VALUE) C.VALUE {
	ret := ruby.RbAryShift(ruby.VALUE(ary))

	return C.VALUE(ret)
}

//export rb_example_tests_rb_ary_unshift
func rb_example_tests_rb_ary_unshift(_ C.VALUE, ary C.VALUE, elem C.VALUE) C.VALUE {
	ret := ruby.RbAryUnshift(ruby.VALUE(ary), ruby.VALUE(elem))

	return C.VALUE(ret)
}

//export rb_example_tests_rb_define_variable
func rb_example_tests_rb_define_variable(_ C.VALUE, name C.VALUE, v C.VALUE) {
	strName := ruby.Value2String(ruby.VALUE(name))
	ruby.RbDefineVariable(strName, (*ruby.VALUE)(&v))
}

//export rb_example_tests_rb_define_const
func rb_example_tests_rb_define_const(self C.VALUE, name C.VALUE, val C.VALUE) {
	strName := ruby.Value2String(ruby.VALUE(name))
	ruby.RbDefineConst(ruby.VALUE(self), strName, ruby.VALUE(val))
}

// defineMethodsToExampleTests define methods in Example::Tests
func defineMethodsToExampleTests(rb_mExample ruby.VALUE) {
	rb_cTests := ruby.RbDefineClassUnder(rb_mExample, "Tests", ruby.VALUE(C.rb_cObject))

	ruby.RbDefineMethod(rb_cTests, "rb_ivar_get", C.rb_example_tests_rb_ivar_get, 0)
	ruby.RbDefineMethod(rb_cTests, "rb_ivar_set", C.rb_example_tests_rb_ivar_set, 1)

	ruby.RbDefineMethodId(rb_cTests, ruby.RbIntern("nop_rb_define_method_id"), C.rb_example_tests_nop_rb_define_method_id, 0)
	ruby.RbDefinePrivateMethod(rb_cTests, "nop_rb_define_private_method", C.rb_example_tests_nop_rb_define_private_method, 0)
	ruby.RbDefineProtectedMethod(rb_cTests, "nop_rb_define_protected_method", C.rb_example_tests_nop_rb_define_protected_method, 0)

	ruby.RbDefineSingletonMethod(rb_cTests, "rb_yield", C.rb_example_tests_rb_yield, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_block_proc", C.rb_example_tests_rb_block_proc, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_funcall2", C.rb_example_tests_rb_funcall2, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_funcall3", C.rb_example_tests_rb_funcall3, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_alias", C.rb_example_tests_rb_alias, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_class2name", C.rb_example_tests_rb_class2name, 0)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_attr", C.rb_example_tests_rb_attr, 4)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_const_get", C.rb_example_tests_rb_const_get, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_const_get_at", C.rb_example_tests_rb_const_get_at, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_const_set", C.rb_example_tests_rb_const_set, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_const_defined", C.rb_example_tests_rb_const_defined, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_const_defined_at", C.rb_example_tests_rb_const_defined_at, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_eval_string", C.rb_example_tests_rb_eval_string, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_eval_string_protect", C.rb_example_tests_rb_eval_string_protect, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_eval_string_wrap", C.rb_example_tests_rb_eval_string_wrap, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_new", C.rb_example_tests_rb_ary_new, 0)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_new_capa", C.rb_example_tests_rb_ary_new_capa, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_push", C.rb_example_tests_rb_ary_push, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_pop", C.rb_example_tests_rb_ary_pop, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_shift", C.rb_example_tests_rb_ary_shift, 1)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_ary_unshift", C.rb_example_tests_rb_ary_unshift, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_define_variable", C.rb_example_tests_rb_define_variable, 2)
	ruby.RbDefineSingletonMethod(rb_cTests, "rb_define_const", C.rb_example_tests_rb_define_const, 2)
}
