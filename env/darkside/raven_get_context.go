package darkside

import (
	"unsafe"

	"github.com/alangpierce/go-forceexport"
)

func GetRavenClientInterfaces(client unsafe.Pointer) {
	var interfaces func(interface{}, string) error
	forceexport.GetFunc(&interfaces, "github.com/getsentry/raven-go.*context.interfaces")
}
