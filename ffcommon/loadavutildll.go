package ffcommon

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"
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

func SetAvutilPath(path0 string) {
	avutilPath = path0
}

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
		avUtilLib, err = openLibrary(avutilPath)
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
		avcodecLib, err = openLibrary(avcodecPath)
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
		avdeviceLib, err = openLibrary(avdevicePath)
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
		avfilterLib, err = openLibrary(avfilterPath)
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
		avformatLib, err = openLibrary(avformatPath)
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
		avpostprocLib, err = openLibrary(avpostprocPath)
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
		avswresampleLib, err = openLibrary(avswresamplePath)
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
		avswscaleLib, err = openLibrary(avswscalePath)
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
			closeLibrary(libs[i])
		}
	}
}

var fnMap = map[string]func(path string){
	"avutil":     SetAvutilPath,
	"avcodec":    SetAvcodecPath,
	"avdevice":   SetAvdevicePath,
	"avfilter":   SetAvfilterPath,
	"avformat":   SetAvformatPath,
	"postproc":   SetAvpostprocPath,
	"swresample": SetAvswresamplePath,
	"swscale":    SetAvswscalePath,
}

func AutoSetAvLib(libpath string) error {
	var libloaded = make(map[string]bool)
	searchDirs := []string{".", "/usr/local/lib", "/usr/lib/x86_64-linux-gnu", "/usr/lib"}
	if libpath != "" {
		searchDirs = append([]string{libpath}, searchDirs...)
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("SetAvLib %s\n", err)
		}
		for k, v := range libloaded {
			if !v {
				fmt.Printf("lib not found %s\n", k)
			}
		}
	}()
	for _, dir := range searchDirs {
		_, err := os.Stat(dir)
		if os.IsNotExist(err) {
			continue
		}
		for k, f := range fnMap {
			if libloaded[k] {
				continue
			}
			switch runtime.GOOS {
			case "darwin":
				lib := fmt.Sprintf("%s/lib%s.dylib", dir, k)
				if _, err := os.Stat(lib); err == nil {
					fmt.Println("load lib", lib)
					f(lib)
					libloaded[k] = true
				}
			case "windows":
				r := regexp.MustCompile(k + "-\\d+\\.dll")
				filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if libloaded[k] {
						return nil
					}
					if err != nil {
						fmt.Println(err)
						return nil
					}
					if !info.IsDir() && r.MatchString(info.Name()) {
						fmt.Println("load lib", path)
						f(path)
						libloaded[k] = true
					}
					return nil
				})
			default:
				// lib := fmt.Sprintf("%s/lib%s.so", dir, k)
				// if _, err := os.Stat(lib); err == nil {
				// 	plugin.Info("load lib", zap.String("path", lib))
				// 	f(lib)
				// 	libloaded[k] = true
				// }
				r := regexp.MustCompile(k + ".so")
				filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
					if libloaded[k] {
						return nil
					}
					if err != nil {
						fmt.Println(err)
						return nil
					}
					if !info.IsDir() && r.MatchString(info.Name()) {
						fmt.Println("load lib", path)
						f(path)
						libloaded[k] = true
					}
					return nil
				})
			}
		}
	}

	//libavutil.AvLogSetLevel(libavutil.AV_LOG_VERBOSE)
	//av_log_set_level(AV_LOG_VERBOSE);
	return nil
}
