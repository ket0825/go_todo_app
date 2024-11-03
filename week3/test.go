// package main

// // #include <stdio.h>
// // void printHelloWorld() {
// // 	printf("Hello, World!\n");
// // }
// import "C"

// func main() {
// 	C.printHelloWorld()
// }

// case 2
package main

// #include <stdio.h>
// #include <stdlib.h>
// void printString(char* str) {
// 	printf("%s\n", str);
// }
import "C"
import "unsafe"

func main() {
	a := C.CString("This is from Golang!")
	C.printString(a)
	C.free(unsafe.Pointer(a))
}
