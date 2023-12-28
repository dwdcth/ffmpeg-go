package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Videotoolbox hardware acceleration
 *
 * copyright (c) 2012 Sebastien Zwickert
 *
 * This file is part of FFmpeg.
 *
 * FFmpeg is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVCODEC_VIDEOTOOLBOX_H
//#define AVCODEC_VIDEOTOOLBOX_H
//
///**
// * @file
// * @ingroup lavc_codec_hwaccel_videotoolbox
// * Public libavcodec Videotoolbox header.
// */
//
//#include <stdint.h>
//
//#define Picture QuickdrawPicture
//#include <VideoToolbox/VideoToolbox.h>
//#undef Picture
//
//#include "../libavcodec/avcodec.h"

/**
 * This struct holds all the information that needs to be passed
 * between the caller and libavcodec for initializing Videotoolbox decoding.
 * Its size is not a part of the public ABI, it must be allocated with
 * av_videotoolbox_alloc_context() and freed with av_free().
 */
type AVVideotoolboxContext struct {

	/**
	 * Videotoolbox decompression session object.
	 * Created and freed the caller.
	 */
	//VTDecompressionSessionRef session;
	Session uintptr
	/**
	 * The output callback that must be passed to the session.
	 * Set by av_videottoolbox_default_init()
	 */
	//VTDecompressionOutputCallback output_callback;
	OutputCallback uintptr
	/**
	 * CVPixelBuffer Format Type that Videotoolbox will use for decoded frames.
	 * set by the caller. If this is set to 0, then no specific format is
	 * requested from the decoder, and its native format is output.
	 */
	//OSType cv_pix_fmt_type;
	CvPixFmtType uintptr
	/**
	 * CoreMedia Format Description that Videotoolbox will use to create the decompression session.
	 * Set by the caller.
	 */
	//CMVideoFormatDescriptionRef cm_fmt_desc;
	CmFmtDesc uintptr
	/**
	 * CoreMedia codec type that Videotoolbox will use to create the decompression session.
	 * Set by the caller.
	 */
	CmCodecType ffcommon.FInt
}

/**
 * Allocate and initialize a Videotoolbox context.
 *
 * This function should be called from the get_format() callback when the caller
 * selects the AV_PIX_FMT_VIDETOOLBOX format. The caller must then create
 * the decoder object (using the output callback provided by libavcodec) that
 * will be used for Videotoolbox-accelerated decoding.
 *
 * When decoding with Videotoolbox is finished, the caller must destroy the decoder
 * object and free the Videotoolbox context using av_free().
 *
 * @return the newly allocated context or NULL on failure
 */
//AVVideotoolboxContext *av_videotoolbox_alloc_context(void);
var avVideotoolboxAllocContext func() *AVVideotoolboxContext
var avVideotoolboxAllocContextOnce sync.Once

func AvVideotoolboxAllocContext() *AVVideotoolboxContext {
	avVideotoolboxAllocContextOnce.Do(func() {
		purego.RegisterLibFunc(&avVideotoolboxAllocContext, ffcommon.GetAvcodecDll(), "av_videotoolbox_alloc_context")
	})
	return avVideotoolboxAllocContext()
}

/**
 * This is a convenience function that creates and sets up the Videotoolbox context using
 * an internal implementation.
 *
 * @param avctx the corresponding codec context
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
//int av_videotoolbox_default_init(AVCodecContext *avctx);
var avVideotoolboxDefaultInit func(avctx *AVCodecContext) ffcommon.FInt
var avVideotoolboxDefaultInitOnce sync.Once

func (avctx *AVCodecContext) AvVideotoolboxDefaultInit() ffcommon.FInt {
	avVideotoolboxDefaultInitOnce.Do(func() {
		purego.RegisterLibFunc(&avVideotoolboxDefaultInit, ffcommon.GetAvcodecDll(), "av_videotoolbox_default_init")
	})
	return avVideotoolboxDefaultInit(avctx)
}

/**
 * This is a convenience function that creates and sets up the Videotoolbox context using
 * an internal implementation.
 *
 * @param avctx the corresponding codec context
 * @param vtctx the Videotoolbox context to use
 *
 * @return >= 0 on success, a negative AVERROR code on failure
 */
//int av_videotoolbox_default_init2(AVCodecContext *avctx, AVVideotoolboxContext *vtctx);
var avVideotoolboxDefaultInit2 func(avctx *AVCodecContext, vtctx *AVVideotoolboxContext) ffcommon.FInt
var avVideotoolboxDefaultInit2Once sync.Once

func (avctx *AVCodecContext) AvVideotoolboxDefaultInit2(vtctx *AVVideotoolboxContext) ffcommon.FInt {
	avVideotoolboxDefaultInit2Once.Do(func() {
		purego.RegisterLibFunc(&avVideotoolboxDefaultInit2, ffcommon.GetAvcodecDll(), "av_videotoolbox_default_init2")
	})
	return avVideotoolboxDefaultInit2(avctx, vtctx)
}

/**
 * This function must be called to free the Videotoolbox context initialized with
 * av_videotoolbox_default_init().
 *
 * @param avctx the corresponding codec context
 */
//void av_videotoolbox_default_free(AVCodecContext *avctx);
var avVideotoolboxDefaultFree func(avctx *AVCodecContext)
var avVideotoolboxDefaultFreeOnce sync.Once

func (avctx *AVCodecContext) AvVideotoolboxDefaultFree() {
	avVideotoolboxDefaultFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avVideotoolboxDefaultFree, ffcommon.GetAvcodecDll(), "av_videotoolbox_default_free")
	})
	avVideotoolboxDefaultFree(avctx)
}

/**
 * @}
 */

//#endif /* AVCODEC_VIDEOTOOLBOX_H */
