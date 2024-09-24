// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated
// Code generated by ruby_h_to_go. DO NOT EDIT.

package ruby

/*
#include "ruby.h"
*/
import "C"

// RbArray calls `rb_Array` in C
//
// Original definition is following
//
//	VALUE rb_Array(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbArray(val VALUE) VALUE {
	ret := VALUE(C.rb_Array(C.VALUE(val)))
	return ret
}

// RbFloat calls `rb_Float` in C
//
// Original definition is following
//
//	VALUE rb_Float(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbFloat(val VALUE) VALUE {
	ret := VALUE(C.rb_Float(C.VALUE(val)))
	return ret
}

// RbInteger calls `rb_Integer` in C
//
// Original definition is following
//
//	VALUE rb_Integer(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbInteger(val VALUE) VALUE {
	ret := VALUE(C.rb_Integer(C.VALUE(val)))
	return ret
}

// RbString calls `rb_String` in C
//
// Original definition is following
//
//	VALUE rb_String(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbString(val VALUE) VALUE {
	ret := VALUE(C.rb_String(C.VALUE(val)))
	return ret
}

// RbAnyToS calls `rb_any_to_s` in C
//
// Original definition is following
//
//	VALUE rb_any_to_s(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbAnyToS(obj VALUE) VALUE {
	ret := VALUE(C.rb_any_to_s(C.VALUE(obj)))
	return ret
}

// RbCheckConvertType calls `rb_check_convert_type` in C
//
// Original definition is following
//
//	VALUE rb_check_convert_type(VALUE val, int type, const char *name, const char *mid)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbCheckConvertType(val VALUE, r int, name string, mid string) VALUE {
	charName, cleanCharname := string2Char(name)
	defer cleanCharname()

	charMid, cleanCharmid := string2Char(mid)
	defer cleanCharmid()

	ret := VALUE(C.rb_check_convert_type(C.VALUE(val), C.int(r), charName, charMid))
	return ret
}

// RbCheckToFloat calls `rb_check_to_float` in C
//
// Original definition is following
//
//	VALUE rb_check_to_float(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbCheckToFloat(val VALUE) VALUE {
	ret := VALUE(C.rb_check_to_float(C.VALUE(val)))
	return ret
}

// RbCheckToInt calls `rb_check_to_int` in C
//
// Original definition is following
//
//	VALUE rb_check_to_int(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbCheckToInt(val VALUE) VALUE {
	ret := VALUE(C.rb_check_to_int(C.VALUE(val)))
	return ret
}

// RbCheckToInteger calls `rb_check_to_integer` in C
//
// Original definition is following
//
//	VALUE rb_check_to_integer(VALUE val, const char *mid)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbCheckToInteger(val VALUE, mid string) VALUE {
	char, clean := string2Char(mid)
	defer clean()

	ret := VALUE(C.rb_check_to_integer(C.VALUE(val), char))
	return ret
}

// RbClassNewInstance calls `rb_class_new_instance` in C
//
// Original definition is following
//
//	VALUE rb_class_new_instance(int argc, const VALUE *argv, VALUE klass)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbClassNewInstance(argc int, argv *VALUE, klass VALUE) VALUE {
	var cArgv C.VALUE
	ret := VALUE(C.rb_class_new_instance(C.int(argc), &cArgv, C.VALUE(klass)))
	*argv = VALUE(cArgv)
	return ret
}

// RbClassNewInstanceKw calls `rb_class_new_instance_kw` in C
//
// Original definition is following
//
//	VALUE rb_class_new_instance_kw(int argc, const VALUE *argv, VALUE klass, int kw_splat)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbClassNewInstanceKw(argc int, argv *VALUE, klass VALUE, kw_splat int) VALUE {
	var cArgv C.VALUE
	ret := VALUE(C.rb_class_new_instance_kw(C.int(argc), &cArgv, C.VALUE(klass), C.int(kw_splat)))
	*argv = VALUE(cArgv)
	return ret
}

// RbConvertType calls `rb_convert_type` in C
//
// Original definition is following
//
//	VALUE rb_convert_type(VALUE val, int type, const char *name, const char *mid)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbConvertType(val VALUE, r int, name string, mid string) VALUE {
	charName, cleanCharname := string2Char(name)
	defer cleanCharname()

	charMid, cleanCharmid := string2Char(mid)
	defer cleanCharmid()

	ret := VALUE(C.rb_convert_type(C.VALUE(val), C.int(r), charName, charMid))
	return ret
}

// RbCstrToDbl calls `rb_cstr_to_dbl` in C
//
// Original definition is following
//
//	double rb_cstr_to_dbl(const char *str, int mode)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbCstrToDbl(str string, mode int) Double {
	char, clean := string2Char(str)
	defer clean()

	ret := Double(C.rb_cstr_to_dbl(char, C.int(mode)))
	return ret
}

// RbEql calls `rb_eql` in C
//
// Original definition is following
//
//	int rb_eql(VALUE lhs, VALUE rhs)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbEql(lhs VALUE, rhs VALUE) int {
	ret := int(C.rb_eql(C.VALUE(lhs), C.VALUE(rhs)))
	return ret
}

// RbInspect calls `rb_inspect` in C
//
// Original definition is following
//
//	VALUE rb_inspect(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbInspect(obj VALUE) VALUE {
	ret := VALUE(C.rb_inspect(C.VALUE(obj)))
	return ret
}

// RbObjAlloc calls `rb_obj_alloc` in C
//
// Original definition is following
//
//	VALUE rb_obj_alloc(VALUE klass)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjAlloc(klass VALUE) VALUE {
	ret := VALUE(C.rb_obj_alloc(C.VALUE(klass)))
	return ret
}

// RbObjClone calls `rb_obj_clone` in C
//
// Original definition is following
//
//	VALUE rb_obj_clone(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjClone(obj VALUE) VALUE {
	ret := VALUE(C.rb_obj_clone(C.VALUE(obj)))
	return ret
}

// RbObjDup calls `rb_obj_dup` in C
//
// Original definition is following
//
//	VALUE rb_obj_dup(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjDup(obj VALUE) VALUE {
	ret := VALUE(C.rb_obj_dup(C.VALUE(obj)))
	return ret
}

// RbObjFreeze calls `rb_obj_freeze` in C
//
// Original definition is following
//
//	VALUE rb_obj_freeze(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjFreeze(obj VALUE) VALUE {
	ret := VALUE(C.rb_obj_freeze(C.VALUE(obj)))
	return ret
}

// RbObjId calls `rb_obj_id` in C
//
// Original definition is following
//
//	VALUE rb_obj_id(VALUE obj)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjId(obj VALUE) VALUE {
	ret := VALUE(C.rb_obj_id(C.VALUE(obj)))
	return ret
}

// RbObjInitCopy calls `rb_obj_init_copy` in C
//
// Original definition is following
//
//	VALUE rb_obj_init_copy(VALUE src, VALUE dst)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjInitCopy(src VALUE, dst VALUE) VALUE {
	ret := VALUE(C.rb_obj_init_copy(C.VALUE(src), C.VALUE(dst)))
	return ret
}

// RbObjIsInstanceOf calls `rb_obj_is_instance_of` in C
//
// Original definition is following
//
//	VALUE rb_obj_is_instance_of(VALUE obj, VALUE klass)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjIsInstanceOf(obj VALUE, klass VALUE) VALUE {
	ret := VALUE(C.rb_obj_is_instance_of(C.VALUE(obj), C.VALUE(klass)))
	return ret
}

// RbObjIsKindOf calls `rb_obj_is_kind_of` in C
//
// Original definition is following
//
//	VALUE rb_obj_is_kind_of(VALUE obj, VALUE klass)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbObjIsKindOf(obj VALUE, klass VALUE) VALUE {
	ret := VALUE(C.rb_obj_is_kind_of(C.VALUE(obj), C.VALUE(klass)))
	return ret
}

// RbStrToDbl calls `rb_str_to_dbl` in C
//
// Original definition is following
//
//	double rb_str_to_dbl(VALUE str, int mode)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbStrToDbl(str VALUE, mode int) Double {
	ret := Double(C.rb_str_to_dbl(C.VALUE(str), C.int(mode)))
	return ret
}

// RbToFloat calls `rb_to_float` in C
//
// Original definition is following
//
//	VALUE rb_to_float(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbToFloat(val VALUE) VALUE {
	ret := VALUE(C.rb_to_float(C.VALUE(val)))
	return ret
}

// RbToInt calls `rb_to_int` in C
//
// Original definition is following
//
//	VALUE rb_to_int(VALUE val)
//
// ref. https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/object.h
func RbToInt(val VALUE) VALUE {
	ret := VALUE(C.rb_to_int(C.VALUE(val)))
	return ret
}