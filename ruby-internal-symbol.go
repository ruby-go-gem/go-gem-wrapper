package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/symbol.h

// RbIntern calls `rb_intern` in C
//
// Original definition is following
//
//	ID rb_intern(const char *name)
func RbIntern(name string) ID {
	char, clean := string2Char(name)
	defer clean()

	return ID(C.rb_intern(char))
}
