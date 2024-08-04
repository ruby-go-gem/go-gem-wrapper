package main

/*
#include "dummy.h"
*/
import "C"

import (
	"github.com/sue445/go-gem-wrapper"
)

//export rb_dummy_sum
func rb_dummy_sum(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	aLong := ruby.NUM2LONG(ruby.VALUE(a))
	bLong := ruby.NUM2LONG(ruby.VALUE(b))

	sum := aLong + bLong

	return C.VALUE(ruby.LONG2NUM(sum))
}

//export rb_dummy_with_block2
func rb_dummy_with_block2(_ C.VALUE, arg C.VALUE) C.VALUE {
	if !ruby.RbBlockGivenP() {
		ruby.RbRaise(ruby.VALUE(C.rb_eArgError), "Block not given")
	}

	block := ruby.RbBlockProc()

	// Call Proc#call
	blockResult := ruby.RbFuncall2(ruby.VALUE(block), ruby.RbIntern("call"), 1, []ruby.VALUE{ruby.VALUE(arg)})

	return C.VALUE(blockResult)
}

//export rb_dummy_hello
func rb_dummy_hello(_ C.VALUE, name C.VALUE) C.VALUE {
	nameString := ruby.Value2String(ruby.VALUE(name))
	result := "Hello, " + nameString
	return C.VALUE(ruby.String2Value(result))
}

//export rb_dummy_round_num2
func rb_dummy_round_num2(_ C.VALUE, num C.VALUE, ndigits C.VALUE) C.VALUE {
	// Call Integer#round
	result := ruby.RbFuncall2(ruby.VALUE(num), ruby.RbIntern("round"), 1, []ruby.VALUE{ruby.VALUE(ndigits)})

	return C.VALUE(result)
}

//export rb_dummy_round_num3
func rb_dummy_round_num3(_ C.VALUE, num C.VALUE, ndigits C.VALUE) C.VALUE {
	// Call Integer#round
	result := ruby.RbFuncall3(ruby.VALUE(num), ruby.RbIntern("round"), 1, []ruby.VALUE{ruby.VALUE(ndigits)})

	return C.VALUE(result)
}

//export rb_dummy_to_string
func rb_dummy_to_string(_ C.VALUE, source C.VALUE) C.VALUE {
	// Call Object#to_s
	result := ruby.CallFunction(ruby.VALUE(source), "to_s")

	return C.VALUE(result)
}

//export rb_dummy_max
func rb_dummy_max(_ C.VALUE, a C.VALUE, b C.VALUE) C.VALUE {
	aLong := ruby.NUM2INT(ruby.VALUE(a))
	bLong := ruby.NUM2INT(ruby.VALUE(b))

	if aLong > bLong {
		return C.VALUE(ruby.INT2NUM(aLong))
	}

	return C.VALUE(ruby.INT2NUM(bLong))
}

var rb_mDummy ruby.VALUE

//export Init_dummy
func Init_dummy() {
	rb_mDummy = ruby.RbDefineModule("Dummy")
	ruby.RbDefineSingletonMethod(rb_mDummy, "sum", C.rb_dummy_sum, 2)
	ruby.RbDefineSingletonMethod(rb_mDummy, "with_block2", C.rb_dummy_with_block2, 1)
	ruby.RbDefineSingletonMethod(rb_mDummy, "hello", C.rb_dummy_hello, 1)
	ruby.RbDefineSingletonMethod(rb_mDummy, "round_num2", C.rb_dummy_round_num2, 2)
	ruby.RbDefineSingletonMethod(rb_mDummy, "round_num3", C.rb_dummy_round_num3, 2)
	ruby.RbDefineSingletonMethod(rb_mDummy, "to_string", C.rb_dummy_to_string, 1)
	ruby.RbDefineModuleFunction(rb_mDummy, "max", C.rb_dummy_max, 2)

	// Create Dummy::InnerClass class
	ruby.RbDefineClassUnder(rb_mDummy, "InnerClass", ruby.VALUE(C.rb_cObject))

	// Create Dummy::InnerModule module
	ruby.RbDefineModuleUnder(rb_mDummy, "InnerModule")

	// Create OuterClass class
	ruby.RbDefineClass("OuterClass", ruby.VALUE(C.rb_cObject))

	defineMethodsToDummyTests(rb_mDummy)
}

func main() {
}
