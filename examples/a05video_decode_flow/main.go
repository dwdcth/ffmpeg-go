package main

import (
	"fmt"
	"github.com/dwdcth/ffmpeg-go/v7/examples"

	"github.com/dwdcth/ffmpeg-go/v7/libavcodec"
	"github.com/dwdcth/ffmpeg-go/v7/libavformat"
	"github.com/dwdcth/ffmpeg-go/v7/libavutil"
)

//go run main.go  -file

func main() {
	fileName := examples.Setup()

	videoStreamIndex := -1 //视频流所在流序列中的索引
	ret := int32(0)        //默认返回值

	//需要的变量名并初始化
	var fmtCtx *libavformat.AVFormatContext
	var pkt *libavcodec.AVPacket
	var codecCtx *libavcodec.AVCodecContext
	var avCodecPara *libavcodec.AVCodecParameters
	var codec *libavcodec.AVCodec
	libavformat.AvformatNetworkInit()

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

		//=========================== 分配AVPacket结构体 ===============================//
		i := 0                                            //用于帧计数
		pkt = libavcodec.AvPacketAlloc()                  //分配一个packet
		pkt.AvNewPacket(codecCtx.Width * codecCtx.Height) //调整packet的数据

		//===========================  读取视频信息 ===============================//
		for fmtCtx.AvReadFrame(pkt) >= 0 { //读取的是一帧视频  数据存入一个AVPacket的结构中
			if pkt.StreamIndex == uint32(videoStreamIndex) {
				i++ //只计算视频帧
			}
			pkt.AvPacketUnref() //重置pkt的内容
		}
		fmt.Printf("There are %d frames int total.\n", i)
		break
	}
	//===========================释放所有指针===============================//
	libavcodec.AvPacketFree(&pkt)
	codecCtx.AvcodecClose()
	libavformat.AvformatCloseInput(&fmtCtx)
	fmtCtx.AvformatFreeContext()
}
