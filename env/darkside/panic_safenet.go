package darkside

import (
	"fmt"
	"unsafe"

	"github.com/alangpierce/go-forceexport"
	"github.com/bouk/monkey"
)

type Panic struct {
	Argp      unsafe.Pointer // pointer to arguments of deferred call run during panic; cannot move - known to liblink
	Arg       interface{}    // argument to panic
	Link      *Panic         // link to earlier panic
	Recovered bool           // whether this panic is over
	Aborted   bool           // the panic was aborted
}

//noinspection SpellCheckingInspection,GoSnakeCaseUsage
func SetUnrecoveredPanicHandler(hanler func(p *Panic)) {
	var throw func(string)
	var startpanic func()
	var systemstack func(func())
	var startpanic_m func()
	var preprintpanics func(p *Panic)
	forceexport.GetFunc(&throw, "runtime.throw")
	forceexport.GetFunc(&startpanic, "runtime.startpanic")
	forceexport.GetFunc(&systemstack, "runtime.systemstack")
	forceexport.GetFunc(&startpanic_m, "runtime.startpanic_m")
	forceexport.GetFunc(&preprintpanics, "runtime.preprintpanics")

	monkey.Patch(startpanic, func() {
		defer func() {
			if recover() != nil {
				throw("panic while printing panic value")
			}
		}()

		p := getg()._panic
		hanler(p)

		for p != nil {
			switch v := p.Arg.(type) {
			case error:
				p.Arg = v.Error()
			case fmt.Stringer:
				p.Arg = v.String()
			}
			p = p.Link
		}

		systemstack(startpanic_m)
	})
	monkey.Patch(preprintpanics, func(p *Panic) {
		// noop
	})
}
