package main

import (
	"fmt"
	// "os"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/dwdcth/ffmpeg-go/libavutil"
)

func main() {
	// os.Setenv("Path", os.Getenv("Path")+";/lib/x86_64-linux-gnu/")
	ffcommon.SetAvutilPath("/lib/x86_64-linux-gnu/libavutil.so.56")
	ffcommon.SetAvcodecPath("/lib/x86_64-linux-gnu/libavcodec.so.58")
	ffcommon.SetAvdevicePath("/lib/x86_64-linux-gnu/libavdevice.so.58")
	ffcommon.SetAvfilterPath("/lib/x86_64-linux-gnu/libavfilter.so.7")
	ffcommon.SetAvformatPath("/lib/x86_64-linux-gnu/libavformat.so.58")
	ffcommon.SetAvpostprocPath("/lib/x86_64-linux-gnu/libpostproc.so.55")
	ffcommon.SetAvswresamplePath("/lib/x86_64-linux-gnu/libswresample.so.3")
	ffcommon.SetAvswscalePath("/lib/x86_64-linux-gnu/libswscale.so.5")
	if true {
		ret := libavutil.AvFrameAlloc()
		fmt.Println(ret)
		fmt.Println(libavutil.AV_NUM_DATA_POINTERS)
	}
	// if true {
	// 	libavutil.AvLog(0, 0, "a", "b")
	// }
	// if true {
	// 	ret := libavcodec.AvcodecVersion()
	// 	fmt.Println(ret)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecLicense())
	// }
	// if true {
	// 	fmt.Println(libavcodec.AvcodecConfiguration())
	// }
	// if true {
	// 	ans := libavutil.AvAdler32Update(111, nil, 0)
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	ans := libavutil.AvAesAlloc()
	// 	fmt.Println(ans)
	// }
	// if true {
	// 	fmt.Println(libavutil.AV_MATRIX_ENCODING_DOLBYHEADPHONE)
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilVersion())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvVersionInfo())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilConfiguration())
	// }
	// if true {
	// 	fmt.Println(libavutil.AvutilLicense())
	// }
	// if true {
	// 	fmt.Println(libavutil.AVMEDIA_TYPE_VIDEO)
	// 	fmt.Println(libavutil.AvGetMediaTypeString(libavutil.AVMEDIA_TYPE_AUDIO))
	// }
	// if true {
	// 	fmt.Println(libavutil.AvGetTimeBaseQ())
	// 	libavutil.AvGetTimeBaseQ()
	// }
}
