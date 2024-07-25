package main

import (
	// 导入 purego 包
	"fmt"
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"

	// "github.com/dwdcth/ffmpeg-go/v7/libavcodec"
	"github.com/ebitengine/purego"
)

// func AvcodecGetName(id int32) (res ffcommon.FCharP) {
// 	var avcodec_get_name func(int32) ffcommon.FCharP
// 	purego.RegisterLibFunc(&avcodec_get_name, ffcommon.GetAvcodecLib(), "av_version_info")

// 	res = avcodec_get_name(id)
// 	return
// }

/*
var av_version_info func() string
var av_version_info_once sync.Once
func AvVersionInfo() (res ffcommon.FCharP){
	av_version_info_once.Do(func() {
		purego.RegisterLibFunc(&av_version_info, ffcommon.GetAvutilDll(), "av_version_info")
	})
	res = av_version_info()
	return
}
*/

var av_version_info func() string
var av_version_info_once sync.Once

func AvVersionInfo() (res ffcommon.FCharP) {
	av_version_info_once.Do(func() {
		purego.RegisterLibFunc(&av_version_info, ffcommon.GetAvutilDll(), "av_version_info")
	})
	res = av_version_info()
	return
}
func main() {
	// // 加载动态库
	// lib, err := purego.Dlopen("/lib/x86_64-linux-gnu/libavutil.so.56", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	// if err != nil {
	// 	panic(err)
	// }
	// defer purego.Dlclose(lib)

	// var av_version_info func() string
	// purego.RegisterLibFunc(&av_version_info, lib, "av_version_info")

	// res := av_version_info()
	// fmt.Println(res)

	// lib2, err := purego.Dlopen("/lib/x86_64-linux-gnu/libavcodec.so.58", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	// if err != nil {
	// 	panic(err)
	// }
	// defer purego.Dlclose(lib2)

	// var avcodec_get_name func(int32) string
	// purego.RegisterLibFunc(&avcodec_get_name, lib2, "avcodec_get_name")

	// res1 := avcodec_get_name(27)
	// fmt.Println(res1)

	res := AvVersionInfo()
	res = AvVersionInfo()
	res = AvVersionInfo()

	fmt.Println(res)
}

/*

//dylib
func AvVersionInfo() (res ffcommon.FCharP) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_version_info").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}
//purego
func AvVersionInfo() (res ffcommon.FCharP){

	var av_version_info func() string
	purego.RegisterLibFunc(&av_version_info, ffcommon.GetAvutilDll(), "av_version_info")
	res = av_version_info()
	return
}
*/
