package main

import (
	"fmt"
	"os"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v7/libavformat"
	"github.com/dwdcth/ffmpeg-go/v7/libavutil"
)

func main() {

	os.Setenv("Path", os.Getenv("Path")+";./lib")
	ffcommon.SetAvutilPath("avutil-56.dll")
	ffcommon.SetAvcodecPath("avcodec-58.dll")
	ffcommon.SetAvdevicePath("avdevice-58.dll")
	ffcommon.SetAvfilterPath("avfilter-56.dll")
	ffcommon.SetAvformatPath("avformat-58.dll")
	ffcommon.SetAvpostprocPath("postproc-55.dll")
	ffcommon.SetAvswresamplePath("swresample-3.dll")
	ffcommon.SetAvswscalePath("swscale-5.dll")

	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}
	main0()
}

func main0() (ret ffcommon.FInt) {
	var fmt_ctx *libavformat.AVFormatContext
	var tag *libavutil.AVDictionaryEntry

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <input_file>\nexample program to demonstrate the use of the libavformat metadata API.\n\n", os.Args[0])
		return 1
	}
	ret = libavformat.AvformatOpenInput(&fmt_ctx, os.Args[1], nil, nil)
	if ret != 0 {
		return ret
	}

	ret = fmt_ctx.AvformatFindStreamInfo(nil)
	if ret < 0 {
		libavutil.AvLog(uintptr(0), libavutil.AV_LOG_ERROR, "Cannot find stream information\n")
		return ret
	}

	tag = fmt_ctx.Metadata.AvDictGet("", tag, libavutil.AV_DICT_IGNORE_SUFFIX)
	for tag != nil {
		fmt.Printf("%s=%s\n", ffcommon.StringFromPtr(tag.Key), ffcommon.StringFromPtr(tag.Value))
		tag = fmt_ctx.Metadata.AvDictGet("", tag, libavutil.AV_DICT_IGNORE_SUFFIX)
	}

	libavformat.AvformatCloseInput(&fmt_ctx)
	return 0
}
