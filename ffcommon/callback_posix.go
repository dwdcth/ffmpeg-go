//go:build !windows
// +build !windows

package ffcommon

import "github.com/ebitengine/purego"

func NewCallback(fn interface{}) uintptr {
	u := purego.NewCallback(fn)
	return u
}
