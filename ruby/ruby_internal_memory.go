package ruby

/*
#include "ruby.h"
*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/memory.h

// RB_GC_GUARD calls `RB_GC_GUARD` in C
//
// ref.
//   - https://docs.ruby-lang.org/capi/en/master/dc/d18/memory_8h.html
//   - https://docs.ruby-lang.org/en/master/extension_rdoc.html
func RB_GC_GUARD(v VALUE) {
	C.RB_GC_GUARD(C.VALUE(v))
}
