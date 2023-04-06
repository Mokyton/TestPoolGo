package main

/*
 #cgo CFLAGS: -x objective-c -DGL_SILENCE_DEPRECATION
 #cgo LDFLAGS: -framework Cocoa
 #import <Cocoa/Cocoa.h>
 #include "application.h"
 #include "window.h"
*/
import "C"

func main() {
	C.InitApplication()
	w := C.Window_Create(150, 150, 300, 200, C.CString("School 21"))
	C.Window_MakeKeyAndOrderFront(w)
	C.RunApplication()
}
