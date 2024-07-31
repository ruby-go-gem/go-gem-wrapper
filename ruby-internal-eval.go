package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/eval.h

// RbFuncallv calls `rb_funcallv` in C
//
// Original definition is following
//
//	VALUE rb_funcallv(VALUE recv, ID mid, int argc, const VALUE *argv)
func RbFuncallv(recv VALUE, mid ID, argc int, argv []VALUE) VALUE {
	return VALUE(C.rb_funcallv(C.VALUE(recv), C.ID(mid), C.int(argc), toCValueArray(argv)))
}

// RbFuncall2 is alias to [RbFuncallv]
func RbFuncall2(recv VALUE, mid ID, argc int, argv []VALUE) VALUE {
	return RbFuncallv(recv, mid, argc, argv)
}