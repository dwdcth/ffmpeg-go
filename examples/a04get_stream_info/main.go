package main

import (
	"flag"
	"fmt"
	"os/exec"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v6/libavformat"
)

//go run main.go  -file

func main() {

	fileName := flag.String("file", "", "video file to open")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("usage: -file 视频文件")
		return
	}

	// ffcommon.SetAvformatPath(*libPath)
	err := ffcommon.AutoSetAvLib("")
	if err != nil {
		fmt.Println("AutoSetAvLib err = ", err)
		return
	}
	fmt_ctx := libavformat.AvformatAllocContext() //创建对象并初始化
	ret := int32(0)

	for {
		//打开文件
		ret = libavformat.AvformatOpenInput(&fmt_ctx, *fileName, nil, nil)
		if ret < 0 {
			fmt.Printf("Cannot open video file\n")
			break //Cannot open video file
		}

		//查找流信息（音频流和视频流）
		ret = fmt_ctx.AvformatFindStreamInfo(nil)
		if ret < 0 {
			fmt.Printf("Cannot find stream information\n")
			break
		}

		fmt_ctx.AvDumpFormat(0, *fileName, 0) //输出视频信息
		break
	}

	libavformat.AvformatCloseInput(&fmt_ctx) //关闭文件

	fmt.Println("\n---------------------------------\n")
	cmd := exec.Command("ffprobe", *fileName)
	data, err2 := cmd.CombinedOutput()
	if err2 != nil {
		fmt.Println("ffprobe err = ", err2)
		return
	}
	fmt.Println(string(data))
}
