package examples

import (
	"flag"
	"fmt"
	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"os"
)

func Setup() *string {
	err := ffcommon.AutoSetAvLib("")
	if err != nil {
		fmt.Println("AutoSetAvLib err = ", err)
		os.Exit(-1)
		return nil
	}
	fileName := flag.String("file", "", "video file to open")
	flag.Parse()
	if *fileName == "" {
		fmt.Println("usage: -file 视频文件")
		os.Exit(-1)
		return nil
	}
	return fileName
}
