package darkside

//import (
//	"fmt"
//
//	"github.com/alangpierce/go-forceexport"
//	"github.com/bouk/monkey"
//)
//
//type funcval struct {
//	fn uintptr
//	// variable-size, fn-specific data here
//}
//
//func SetupGoroutineLocal() {
//	var newproc func(siz int32, fn *funcval)
//	var getcallerpc func() uintptr
//	var systemstack func(func())
//	var newproc1 func(fn *funcval, argp *uint8, narg int32, callerpc uintptr)
//	forceexport.GetFunc(&newproc, "runtime.newproc")
//	forceexport.GetFunc(&getcallerpc, "runtime.getcallerpc")
//	forceexport.GetFunc(&systemstack, "runtime.systemstack")
//	forceexport.GetFunc(&newproc1, "runtime.newproc1")
//	monkey.Patch(newproc, func(siz int32, fn *funcval) {
//		fmt.Println("New goroutine here")
//
//	})
//}
