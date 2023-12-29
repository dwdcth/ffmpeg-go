package ffcommon

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sync"

	"github.com/ying32/dylib"
)

var avUtilDll *dylib.LazyDLL
var avUtilDllOnce sync.Once

func GetAvutilDll() (ans *dylib.LazyDLL) {
	avUtilDllOnce.Do(func() {
		avUtilDll = dylib.NewLazyDLL(avutilPath)
	})
	ans = avUtilDll
	return
}

var avutilPath = "avutil-56.dll"

func SetAvutilPath(path0 string) {
	avutilPath = path0
}

var avcodecDll *dylib.LazyDLL
var avcodecDllOnce sync.Once

func GetAvcodecDll() (ans *dylib.LazyDLL) {
	avcodecDllOnce.Do(func() {
		avcodecDll = dylib.NewLazyDLL(avcodecPath)
	})
	ans = avcodecDll
	return
}

var avcodecPath = "avcodec-56.dll"

func SetAvcodecPath(path0 string) {
	avcodecPath = path0
}

var avdeviceDll *dylib.LazyDLL
var avdeviceDllOnce sync.Once

func GetAvdeviceDll() (ans *dylib.LazyDLL) {
	avdeviceDllOnce.Do(func() {
		avdeviceDll = dylib.NewLazyDLL(avdevicePath)
	})
	ans = avdeviceDll
	return
}

var avdevicePath = "avdevice-56.dll"

func SetAvdevicePath(path0 string) {
	avdevicePath = path0
}

var avfilterDll *dylib.LazyDLL
var avfilterDllOnce sync.Once

func GetAvfilterDll() (ans *dylib.LazyDLL) {
	avfilterDllOnce.Do(func() {
		avfilterDll = dylib.NewLazyDLL(avfilterPath)
	})
	ans = avfilterDll
	return
}

var avfilterPath = "avfilter-56.dll"

func SetAvfilterPath(path0 string) {
	avfilterPath = path0
}

var avformatDll *dylib.LazyDLL
var avformatDllOnce sync.Once

func GetAvformatDll() (ans *dylib.LazyDLL) {
	avformatDllOnce.Do(func() {
		avformatDll = dylib.NewLazyDLL(avformatPath)
	})
	ans = avformatDll
	return
}

var avformatPath = "avformat-58.dll"

func SetAvformatPath(path0 string) {
	avformatPath = path0
}

var avpostprocDll *dylib.LazyDLL
var avpostprocDllOnce sync.Once

func GetAvpostprocDll() (ans *dylib.LazyDLL) {
	avpostprocDllOnce.Do(func() {
		avpostprocDll = dylib.NewLazyDLL(avpostprocPath)
	})
	ans = avpostprocDll
	return
}

var avpostprocPath = "postproc-55.dll"

func SetAvpostprocPath(path0 string) {
	avpostprocPath = path0
}

var avswresampleDll *dylib.LazyDLL
var avswresampleDllOnce sync.Once

func GetAvswresampleDll() (ans *dylib.LazyDLL) {
	avswresampleDllOnce.Do(func() {
		avswresampleDll = dylib.NewLazyDLL(avswresamplePath)
	})
	ans = avswresampleDll
	return
}

var avswresamplePath = "swresample-3.dll"

func SetAvswresamplePath(path0 string) {
	avswresamplePath = path0
}

var avswscaleDll *dylib.LazyDLL
var avswscaleDllOnce sync.Once

func GetAvswscaleDll() (ans *dylib.LazyDLL) {
	avswscaleDllOnce.Do(func() {
		avswscaleDll = dylib.NewLazyDLL(avswscalePath)
	})
	ans = avswscaleDll
	return
}

var avswscalePath = "swscale-5.dll"

func SetAvswscalePath(path0 string) {
	avswscalePath = path0
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
