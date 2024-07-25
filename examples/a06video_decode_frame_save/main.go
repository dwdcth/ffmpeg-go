package main

import (
	"fmt"
	"github.com/dwdcth/ffmpeg-go/v7/examples"
	"os"
	"os/exec"
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/v7/libavcodec"
	"github.com/dwdcth/ffmpeg-go/v7/libavformat"
	"github.com/dwdcth/ffmpeg-go/v7/libavutil"
	"github.com/dwdcth/ffmpeg-go/v7/libswscale"
)

func main() {

	fileName := examples.Setup()

	videoStreamIndex := -1 //视频流所在流序列中的索引
	ret := int32(0)        //默认返回值

	//需要的变量名并初始化
	var fmtCtx *libavformat.AVFormatContext
	var pkt *libavformat.AVPacket
	var codecCtx *libavcodec.AVCodecContext
	var avCodecPara *libavcodec.AVCodecParameters
	var codec *libavcodec.AVCodec
	yuvFrame := libavutil.AvFrameAlloc()
	rgbFrame := libavutil.AvFrameAlloc()

	for {
		//=========================== 创建AVFormatContext结构体 ===============================//
		//分配一个AVFormatContext，FFMPEG所有的操作都要通过这个AVFormatContext来进行
		fmtCtx = libavformat.AvformatAllocContext()
		//==================================== 打开文件 ======================================//
		ret = libavformat.AvformatOpenInput(&fmtCtx, *fileName, nil, nil)
		if ret != 0 {
			fmt.Printf("cannot open video file\n")
			break
		}

		//=================================== 获取视频流信息 ===================================//
		ret = fmtCtx.AvformatFindStreamInfo(nil)
		if ret < 0 {
			fmt.Printf("cannot retrive video info\n")
			break
		}

		//循环查找视频中包含的流信息，直到找到视频类型的流
		//便将其记录下来 保存到videoStreamIndex变量中
		for i := uint32(0); i < fmtCtx.NbStreams; i++ {
			if fmtCtx.GetStream(i).Codecpar.CodecType == libavutil.AVMEDIA_TYPE_VIDEO {
				videoStreamIndex = int(i)
				break //找到视频流就退出
			}
		}

		//如果videoStream为-1 说明没有找到视频流
		if videoStreamIndex == -1 {
			fmt.Printf("cannot find video stream\n")
			break
		}

		//打印输入和输出信息：长度 比特率 流格式等
		fmtCtx.AvDumpFormat(0, *fileName, 0)

		//=================================  查找解码器 ===================================//
		avCodecPara = fmtCtx.GetStream(uint32(videoStreamIndex)).Codecpar
		codec = libavcodec.AvcodecFindDecoder(avCodecPara.CodecId)
		if codec == nil {
			fmt.Printf("cannot find decoder\n")
			break
		}
		//根据解码器参数来创建解码器内容
		codecCtx = codec.AvcodecAllocContext3()
		codecCtx.AvcodecParametersToContext(avCodecPara)
		if codecCtx == nil {
			fmt.Printf("Cannot alloc context.")
			break
		}

		//================================  打开解码器 ===================================//
		ret = codecCtx.AvcodecOpen2(codec, nil)
		if ret < 0 { // 具体采用什么解码器ffmpeg经过封装 我们无须知道
			fmt.Printf("cannot open decoder\n")
			break
		}

		// //================================ 设置数据转换参数 ================================//
		img_ctx := libswscale.SwsGetContext(codecCtx.Width, codecCtx.Height, codecCtx.PixFmt, //源地址长宽以及数据格式
			codecCtx.Width, codecCtx.Height, libavutil.AV_PIX_FMT_RGB32, //目的地址长宽以及数据格式
			libswscale.SWS_BICUBIC, nil, nil, nil) //算法类型  AV_PIX_FMT_YUVJ420P   AV_PIX_FMT_BGR24

		// //==================================== 分配空间 ==================================//
		// //一帧图像数据大小
		numBytes := libavutil.AvImageGetBufferSize(libavutil.AV_PIX_FMT_RGB32, codecCtx.Width, codecCtx.Height, 1)
		out_buffer := libavutil.AvMalloc(uint64(numBytes))

		//=========================== 分配AVPacket结构体 ===============================//
		i := 0                                            //用于帧计数
		pkt = libavcodec.AvPacketAlloc()                  //分配一个packet
		pkt.AvNewPacket(codecCtx.Width * codecCtx.Height) //调整packet的数据

		//会将pFrameRGB的数据按RGB格式自动"关联"到buffer  即pFrameRGB中的数据改变了
		//out_buffer中的数据也会相应的改变
		libavutil.AvImageFillArrays((*[4]*byte)(unsafe.Pointer(&rgbFrame.Data)), (*[4]int32)(unsafe.Pointer(&rgbFrame.Linesize)), (*byte)(unsafe.Pointer(out_buffer)), libavutil.AV_PIX_FMT_RGB32,
			codecCtx.Width, codecCtx.Height, 1)

		//===========================  读取视频信息 ===============================//
		for fmtCtx.AvReadFrame(pkt) >= 0 { //读取的是一帧视频  数据存入一个AVPacket的结构中
			if pkt.StreamIndex == uint32(videoStreamIndex) {
				if codecCtx.AvcodecSendPacket(pkt) == 0 {
					for codecCtx.AvcodecReceiveFrame(yuvFrame) == 0 {
						i++
						if i <= 100 && i >= 0 {
							img_ctx.SwsScale((**byte)(unsafe.Pointer(&yuvFrame.Data)),
								(*int32)(unsafe.Pointer(&yuvFrame.Linesize)),
								0,
								uint32(codecCtx.Height),
								(**byte)(unsafe.Pointer(&rgbFrame.Data)),
								(*int32)(unsafe.Pointer(&rgbFrame.Linesize)))

							saveFrame(rgbFrame, int(codecCtx.Width), int(codecCtx.Height), i)
						}
					}
				}
			}
			pkt.AvPacketUnref() //重置pkt的内容
		}
		fmt.Printf("There are %d frames int total.\n", i)
		break
	}
	// //===========================释放所有指针===============================//
	libavcodec.AvPacketFree(&pkt)
	codecCtx.AvcodecClose()
	libavformat.AvformatCloseInput(&fmtCtx)
	fmtCtx.AvformatFreeContext()
	libavutil.AvFrameFree(&yuvFrame)
	libavutil.AvFrameFree(&rgbFrame)

	fmt.Println("---------------------------------")
	exec.Command("ffmpeg", "-ss", "0", "-i", *fileName, "-f", "image2", "./out/capture.jpg", "-y").Output()
	go exec.Command("ffplay", "-top", "100", "./out/capture.jpg").Output()
	exec.Command("ffplay", "-top", "500", "./out/frame1.ppm").Output()
}

// 将FFmpeg解码后的数据保存到本地文件
func saveFrame(pFrame *libavutil.AVFrame, width, height, iFrame int) {
	genDir := "./out"
	_, err := os.Stat(genDir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(genDir, 0777) //  Everyone can read write and execute
		}
	}
	file, err := os.OpenFile(fmt.Sprintf("./out/frame%d.ppm", iFrame), os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	// 关闭文件
	defer file.Close()
	// 写入文件头
	file.WriteString(fmt.Sprintf("P6\n%d %d\n255\n", width, height))

	// 写入像素数据
	for y := 0; y < height; y++ {
		ptr := uintptr(unsafe.Pointer(pFrame.Data[0])) + uintptr(y*int(pFrame.Linesize[0]))
		bytes := []byte{}
		for z := 0; z < width*3; z++ {
			bytes = append(bytes, *(*byte)(unsafe.Pointer(ptr)))
			ptr++
		}
		file.Write(bytes)
	}
}
