package ruby

/*
#include "ruby.h"
*/
import "C"

import (
	"unsafe"
)

func string2Char(str string) *C.char {
	bytes := append([]byte(str), '\000')

	return (*C.char)(unsafe.Pointer(&bytes[0]))
}

// String2Char convert from Go string to `C.char`
func String2Char(str string) *Char {
	return (*Char)(string2Char(str))
}

// Value2String convert from `VALUE` to Go string
func Value2String(str VALUE) string {
	return value2String(C.VALUE(str))
}

// value2String convert from `C.VALUE` to Go string
func value2String(str C.VALUE) string {
	return C.GoStringN(rstringPtr(str), rstringLenint(str))
}

// StringLen returns string length as `Long`
func StringLen(str string) Long {
	return Long(stringLen(str))
}

// stringLen returns string length as `C.long`
func stringLen(str string) C.long {
	return C.long(len(str))
}
