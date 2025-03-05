package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/arithmetic/long.h

// NUM2LONG is alias to [RbNum2Long]
func NUM2LONG(num VALUE) Long {
	return RbNum2Long(num)
}

// LONG2NUM is alias to [RbLong2NumInline]
func LONG2NUM(v Long) VALUE {
	return RbLong2NumInline(v)
}
