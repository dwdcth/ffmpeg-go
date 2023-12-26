package ffcommon

import (
	"sync"

	"github.com/ebitengine/purego"
)

var (
	avutilPath       = "/lib/x86_64-linux-gnu/libavutil.so.56"  //"avutil-56.dll"
	avcodecPath      = "/lib/x86_64-linux-gnu/libavcodec.so.58" //"avcodec-56.dll"
	avdevicePath     = "avdevice-56.dll"
	avfilterPath     = "avfilter-56.dll"
	avformatPath     = "avformat-58.dll"
	avpostprocPath   = "postproc-55.dll"
	avswresamplePath = "swresample-3.dll"
	avswscalePath    = "swscale-5.dll"
)

func SetAvcodecPath(path0 string) {
	avcodecPath = path0
}

func SetAvdevicePath(path0 string) {
	avdevicePath = path0
}

func SetAvfilterPath(path0 string) {
	avfilterPath = path0
}

func SetAvformatPath(path0 string) {
	avformatPath = path0
}

func SetAvpostprocPath(path0 string) {
	avpostprocPath = path0
}

func SetAvswresamplePath(path0 string) {
	avswresamplePath = path0
}

func SetAvswscalePath(path0 string) {
	avswscalePath = path0
}

var avUtilLib uintptr
var avUtilLibOnce sync.Once

func GetAvutilDll() uintptr {
	avUtilLibOnce.Do(func() {
		var err error
		avUtilLib, err = purego.Dlopen(avutilPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Hanpuregoe error appropriately
		}
	})
	return avUtilLib
}

var avcodecLib uintptr
var avcodecLibOnce sync.Once

func GetAvcodecDll() uintptr {
	avcodecLibOnce.Do(func() {
		var err error
		avcodecLib, err = purego.Dlopen(avcodecPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Hanpuregoe error appropriately
		}
	})
	return avcodecLib
}

var avdeviceLib uintptr
var avdeviceLibOnce sync.Once

func GetAvdeviceDll() uintptr {
	avdeviceLibOnce.Do(func() {
		var err error
		avdeviceLib, err = purego.Dlopen(avdevicePath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Hanpuregoe error appropriately
		}
	})
	return avdeviceLib
}

var avfilterLib uintptr
var avfilterLibOnce sync.Once

func GetAvfilterDll() uintptr {
	avfilterLibOnce.Do(func() {
		var err error
		avfilterLib, err = purego.Dlopen(avfilterPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Handle error appropriately
		}
	})
	return avfilterLib
}

var avformatLib uintptr
var avformatLibOnce sync.Once

func GetAvformatDll() uintptr {
	avformatLibOnce.Do(func() {
		var err error
		avformatLib, err = purego.Dlopen(avformatPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Handle error appropriately
		}
	})
	return avformatLib
}

var avpostprocLib uintptr
var avpostprocLibOnce sync.Once

func GetAvpostprocDll() uintptr {
	avpostprocLibOnce.Do(func() {
		var err error
		avpostprocLib, err = purego.Dlopen(avpostprocPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Handle error appropriately
		}
	})
	return avpostprocLib
}

var avswresampleLib uintptr
var avswresampleLibOnce sync.Once

func GetAvswresampleDll() uintptr {
	avswresampleLibOnce.Do(func() {
		var err error
		avswresampleLib, err = purego.Dlopen(avswresamplePath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Handle error appropriately
		}
	})
	return avswresampleLib
}

var avswscaleLib uintptr
var avswscaleLibOnce sync.Once

func GetAvswscaleDll() uintptr {
	avswscaleLibOnce.Do(func() {
		var err error
		avswscaleLib, err = purego.Dlopen(avswscalePath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			panic(err) // Handle error appropriately
		}
	})
	return avswscaleLib
}

func CloseDll() {
	libs := []uintptr{
		avUtilLib, avcodecLib, avdeviceLib,
		avfilterLib, avformatLib, avpostprocLib,
		avswresampleLib, avswscaleLib}
	for i := 0; i < len(libs); i++ {
		if libs[i] != 0 {
			purego.Dlclose(libs[i])
		}
	}
}
