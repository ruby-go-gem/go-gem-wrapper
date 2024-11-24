package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/core/rarray.h

// RARRAY_LEN is alias to [RbArrayLen]
func RARRAY_LEN(a VALUE) Long {
	return RbArrayLen(a)
}

// RARRAY_CONST_PTR is alias to [RbArrayConstPtr]
func RARRAY_CONST_PTR(a VALUE) *VALUE {
	return RbArrayConstPtr(a)
}
