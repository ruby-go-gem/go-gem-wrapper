package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/arithmetic/int.h

// NUM2INT is alias to [RbNum2IntInline]
func NUM2INT(x VALUE) int {
	return RbNum2IntInline(x)
}

// INT2NUM is alias to [RbInt2NumInline]
func INT2NUM(v int) VALUE {
	return RbInt2NumInline(v)
}
