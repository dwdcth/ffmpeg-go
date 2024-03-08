package main

import (
	"fmt"
	"os/exec"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v6/libavcodec"
	"github.com/dwdcth/ffmpeg-go/v6/libavutil"
)

// go run main.go
func main() {
	err := ffcommon.AutoSetAvLib("")
	if err != nil {
		fmt.Println("AutoSetAvLib err = ", err)
		return
	}

	codecVer := libavcodec.AvcodecVersion()
	ver_major := (codecVer >> 16) & 0xff
	ver_minor := (codecVer >> 8) & 0xff
	ver_micro := (codecVer) & 0xff
	fmt.Printf("FFmpeg version is: %s .\navcodec version is: %d=%d.%d.%d.\n", libavutil.AvVersionInfo(), codecVer, ver_major, ver_minor, ver_micro)

	fmt.Println("\n---------------------------------\n")
	data, err := exec.Command("ffmpeg", "-version").Output()
	if err != nil {
		fmt.Println("ffmpeg err = ", err)
	}
	fmt.Println(string(data))
}
