package ruby

/*
#include "ruby.h"

// FIXME: static inline volatile functions can't be called from cgo.
static inline volatile VALUE *
__rb_gc_guarded_ptr(volatile VALUE *ptr)
{
    return ptr;
}

*/
import "C"

// c.f. https://github.com/ruby/ruby/blob/master/include/ruby/internal/memory.h

// rbGcGuardedPtr is wrapper for `__rb_gc_guarded_ptr`
func rbGcGuardedPtr(v *VALUE) *C.VALUE {
	return C.__rb_gc_guarded_ptr((*C.VALUE)(v))
}

// RB_GC_GUARD calls `RB_GC_GUARD` in C
//
// ref.
//   - https://docs.ruby-lang.org/capi/en/master/dc/d18/memory_8h.html
//   - https://docs.ruby-lang.org/en/master/extension_rdoc.html
func RB_GC_GUARD(v VALUE) *VALUE {
	return (*VALUE)(rbGcGuardedPtr(&v))
}
