package wchart

import (
	"syscall/js"

	"github.com/soypat/gwasm"
)

// Chart.js global library handle
var chart js.Value

func Init() {
	chart = js.Global().Get("Chart")
	if !chart.Truthy() {
		panic("unable to find Chart in global namespace. Have you added the script?")
	}
}

// objectify converts a struct with `js` field tags to
// a javascript Object type with the non-zero, non-nil
// fields set to the struct's values.
func objectify(Struct interface{}) js.Value {
	return gwasm.ValueFromStruct(Struct, true)
}
