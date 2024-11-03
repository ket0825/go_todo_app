// File: concatStr.go
//go:build concatStr
// +build concatStr

package main

/*
#cgo CFLAGS: -O0
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char* concatStr(const char* str1, const char* str2) {
    char* result = (char*)malloc(strlen(str1) + strlen(str2) + 1);

    strcpy(result, str1);
    strcat(result, str2);

    return result;
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str1 := C.CString("Hello ")
	str2 := C.CString("World!")
	defer C.free(unsafe.Pointer(str1))
	defer C.free(unsafe.Pointer(str2))

	result := C.concatStr(str1, str2)
	defer C.free(unsafe.Pointer(result))

	goResult := C.GoString(result)
	fmt.Println(goResult)
}
