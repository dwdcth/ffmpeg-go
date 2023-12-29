package main

import (
	"fmt"
	"os/exec"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/dwdcth/ffmpeg-go/libavcodec"
	"github.com/dwdcth/ffmpeg-go/libavutil"
)

// go run main.go
func main() {
	//flags
	// utilPath := flag.String("util", "", "avutilpath")
	// codecPath := flag.String("codec", "", "avcodecpth")
	// flag.Parse()
	// if *utilPath == "" || *codecPath == "" {
	// 	fmt.Println("usage: -util avutil 路径 -codec avcodec 路径")
	// 	return
	// }
	// ffcommon.SetAvutilPath(*utilPath)
	// ffcommon.SetAvcodecPath(*codecPath)

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
