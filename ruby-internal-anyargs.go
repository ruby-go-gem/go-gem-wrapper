package ruby

/*
#include "ruby.h"
*/
import "C"

import (
	"unsafe"
)

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/anyargs.h

// RbDefineSingletonMethod calls `rb_define_singleton_method` in C
func RbDefineSingletonMethod(klass VALUE, name string, fun unsafe.Pointer, args int) {
	cname := string2Char(name)
	C.rb_define_singleton_method(C.VALUE(klass), cname, (*[0]byte)(fun), C.int(args))
}