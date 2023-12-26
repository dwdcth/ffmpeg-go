package ffcommon

import (
	"github.com/ebitengine/purego"
)

func NewCallback(fn interface{}) uintptr {
	if fn == nil {
		return uintptr(0)
	} else {
		return purego.NewCallback(fn)
	}
}
