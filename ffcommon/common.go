package ffcommon

import (
	"syscall"
	"unsafe"
)

func BytePtrFromString(str string) (res *byte) {
	res, _ = syscall.BytePtrFromString(str)
	return
}

func UintPtrFromString(str string) uintptr {
	if str == "" {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(BytePtrFromString(str)))
}

func UintPtrFromContainsEmptyString(str string) uintptr {
	return uintptr(unsafe.Pointer(BytePtrFromString(str)))
}

//func BoolFromUintptr(ptr uintptr) bool {
//	if ptr == 0 {
//		return false
//	}
//	return true
//}

func StringFromPtr(sptr uintptr) (res string) {
	if sptr <= 0 {
		return
	}
	buf := make([]byte, 0)
	for i := 0; *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != 0; i++ {
		buf = append(buf, *(*byte)(unsafe.Pointer(sptr + uintptr(i))))
	}
	res = string(buf)
	return
}

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func CBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

func ByteSliceFromByteP(data *byte, len0 int) []byte {
	if data == nil {
		return nil
	}
	if len0 == 0 {
		return []byte{}
	}

	var sliceHeader sliceHeader
	sliceHeader.Data = data
	sliceHeader.Len = len0
	sliceHeader.Cap = len0
	return *(*[]byte)(unsafe.Pointer(&sliceHeader))
}

type sliceHeader struct {
	Data *byte
	Len  int
	Cap  int
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// CString converts a go string to *byte that can be passed to C code.
func CString(name string) *byte {
	if hasSuffix(name, "\x00") {
		return &(*(*[]byte)(unsafe.Pointer(&name)))[0]
	}
	b := make([]byte, len(name)+1)
	copy(b, name)
	return &b[0]
}

// GoString copies a null-terminated char* to a Go string.
func GoString(c uintptr) string {
	// We take the address and then dereference it to trick go vet from creating a possible misuse of unsafe.Pointer
	ptr := *(*unsafe.Pointer)(unsafe.Pointer(&c))
	if ptr == nil {
		return ""
	}
	var length int
	for {
		if *(*byte)(unsafe.Add(ptr, uintptr(length))) == '\x00' {
			break
		}
		length++
	}
	return string(unsafe.Slice((*byte)(ptr), length))
}

func GoStringFromBytePtr(p *byte) string {
	if p == nil {
		return ""
	}

	// 找到字符串的结束位置（null 终止符）
	end := unsafe.Pointer(p)
	for *(*byte)(end) != 0 {
		end = unsafe.Pointer(uintptr(end) + 1)
	}

	// 计算字符串长度
	length := uintptr(end) - uintptr(unsafe.Pointer(p))

	// 使用 unsafe.Slice 创建一个字节切片
	slice := unsafe.Slice(p, length)

	// 将字节切片转换为字符串
	return string(slice)
}
