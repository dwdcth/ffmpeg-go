// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2023 The Ebitengine Authors

package ffcommon

import "golang.org/x/sys/windows"

func openLibrary(name string) (uintptr, error) {
	handle, err := windows.LoadLibrary(name)
	return uintptr(handle), err
}

func closeLibrary(lib uintptr) {
	windows.FreeLibrary(windows.Handle(lib))
}
