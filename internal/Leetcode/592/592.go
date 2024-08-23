package _592

/*
#cgo LDFLAGS: -L. -l592
#include <stdlib.h>

extern const char* fractionAddition(const char* expression);
*/
import "C"
import (
	"unsafe"
)

func fractionAddition(expression string) string {
	cExpression := C.CString(expression)
	defer C.free(unsafe.Pointer(cExpression))
	result := C.fractionAddition(cExpression)
	goResult := C.GoString(result)
	return goResult
}
