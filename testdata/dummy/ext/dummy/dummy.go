package main

/*
#include "dummy.h"

VALUE rb_dummy_sum(VALUE self, VALUE a, VALUE b);
VALUE rb_dummy_with_block(VALUE self, VALUE arg);
*/
import "C"

import (
	"github.com/sue445/go-gem-wrapper"
)

//export rb_dummy_sum
func rb_dummy_sum(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	aLong := ruby.RbNum2long(ruby.VALUE(a))
	bLong := ruby.RbNum2long(ruby.VALUE(b))

	sum := aLong + bLong

	return C.VALUE(ruby.RbLong2numInline(sum))
}

//export rb_dummy_with_block
func rb_dummy_with_block(_ C.VALUE, arg C.VALUE) C.VALUE {
	if !ruby.RbBlockGivenP() {
		ruby.RbRaise(ruby.VALUE(C.rb_eArgError), "Block not given")
	}

	blockResult := ruby.RbYield(ruby.VALUE(arg))
	return C.VALUE(blockResult)
}

var rb_mDummy ruby.VALUE

//export Init_dummy
func Init_dummy() {
	rb_mDummy = ruby.RbDefineModule("Dummy")
	ruby.RbDefineSingletonMethod(rb_mDummy, "sum", C.rb_dummy_sum, 2)
	ruby.RbDefineSingletonMethod(rb_mDummy, "with_block", C.rb_dummy_with_block, 1)
}

func main() {
}
