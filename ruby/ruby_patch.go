package ruby

/*
#include "ruby.h"
#include <stdlib.h>
#include "cgo_helpers.h"

// Go's variable-length arguments couldn't be passed directly to C, so they are passed through another function to avoid this
void __rb_raise(VALUE exception, const char *str) {
    rb_raise(exception, "%s", str);
}
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// RbRaise function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/error.h
func RbRaise(exc VALUE, format string, a ...interface{}) {
	str := fmt.Sprintf(format, a...)
	cexc, cexcAllocMap := (C.VALUE)(exc), cgoAllocsUnknown
	cstr, cfmtAllocMap := unpackPCharString(str)
	C.__rb_raise(cexc, cstr)
	runtime.KeepAlive(str)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cfmtAllocMap)
	runtime.KeepAlive(cexcAllocMap)
}

// FIXME: Monkey patched from C `VALUE(*func)(ANYARGS)` argument to Go `unsafe.Pointer` argument

// RbDefineSingletonMethod function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/class.h
func RbDefineSingletonMethod(obj VALUE, mid string, _func unsafe.Pointer, arity int32) {
	cobj, cobjAllocMap := (C.VALUE)(obj), cgoAllocsUnknown
	mid = safeString(mid)
	cmid, cmidAllocMap := unpackPCharString(mid)
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_singleton_method(cobj, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cobjAllocMap)
}

// RbDefineModuleFunction function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/method.h
func RbDefineModuleFunction(klass VALUE, mid string, _func unsafe.Pointer, arity int32) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	mid = safeString(mid)
	cmid, cmidAllocMap := unpackPCharString(mid)
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_module_function(cklass, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cklassAllocMap)
}

// RbDefineMethod function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/method.h
func RbDefineMethod(klass VALUE, mid string, _func unsafe.Pointer, arity int32) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	mid = safeString(mid)
	cmid, cmidAllocMap := unpackPCharString(mid)
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_method(cklass, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cklassAllocMap)
}

// RbDefinePrivateMethod function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/class.h
func RbDefinePrivateMethod(klass VALUE, mid string, _func unsafe.Pointer, arity int32) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	mid = safeString(mid)
	cmid, cmidAllocMap := unpackPCharString(mid)
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_private_method(cklass, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cklassAllocMap)
}

// RbDefineProtectedMethod function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/class.h
func RbDefineProtectedMethod(klass VALUE, mid string, _func unsafe.Pointer, arity int32) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	mid = safeString(mid)
	cmid, cmidAllocMap := unpackPCharString(mid)
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_protected_method(cklass, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cklassAllocMap)
}

// RbDefineMethodId function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/class.h
func RbDefineMethodId(klass VALUE, mid ID, _func unsafe.Pointer, arity int32) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	cmid, cmidAllocMap := (C.ID)(mid), cgoAllocsUnknown
	carity, carityAllocMap := (C.int)(arity), cgoAllocsUnknown
	C.rb_define_method_id(cklass, cmid, toFunctionPointer(_func), carity)
	runtime.KeepAlive(carityAllocMap)
	runtime.KeepAlive(cmidAllocMap)
	runtime.KeepAlive(cklassAllocMap)
}

// RbDefineAllocFunc function as declared in https://github.com/ruby/ruby/blob/master/include/ruby/internal/intern/vm.h
func RbDefineAllocFunc(klass VALUE, _func unsafe.Pointer) {
	cklass, cklassAllocMap := (C.VALUE)(klass), cgoAllocsUnknown
	C.rb_define_alloc_func(cklass, toFunctionPointer(_func))
	runtime.KeepAlive(cklassAllocMap)
}
