package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/dwdcth/ffmpeg-go/libavutil"
	"github.com/ebitengine/purego"
)

/*
 * The Video Decode and Presentation API for UNIX (VDPAU) is used for
 * hardware-accelerated decoding of MPEG-1/2, H.264 and VC-1.
 *
 * Copyright (C) 2008 NVIDIA
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
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with FFmpeg; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVCODEC_VDPAU_H
//#define AVCODEC_VDPAU_H

/**
 * @file
 * @ingroup lavc_codec_hwaccel_vdpau
 * Public libavcodec VDPAU header.
 */

/**
 * @defgroup lavc_codec_hwaccel_vdpau VDPAU Decoder and Renderer
 * @ingroup lavc_codec_hwaccel
 *
 * VDPAU hardware acceleration has two modules
 * - VDPAU decoding
 * - VDPAU presentation
 *
 * The VDPAU decoding module parses all headers using FFmpeg
 * parsing mechanisms and uses VDPAU for the actual decoding.
 *
 * As per the current implementation, the actual decoding
 * and rendering (API calls) are done as part of the VDPAU
 * presentation (vo_vdpau.c) module.
 *
 * @{
 */

//#include <vdpau/vdpau.h>
//
//#include "../libavutil/avconfig.h"
//#include "../libavutil/attributes.h"
//
//#include "avcodec.h"
//#include "version.h"

// struct AVCodecContext;
// struct AVFrame;
type AVFrame = libavutil.AVFrame

//typedef int (*AVVDPAU_Render2)(struct AVCodecContext *, struct AVFrame *,
//const VdpPictureInfo *, uint32_t,
//const VdpBitstreamBuffer *);

/**
 * This structure is used to share data between the libavcodec library and
 * the client video application.
 * The user shall allocate the structure via the av_alloc_vdpau_hwaccel
 * function and make it available as
 * AVCodecContext.hwaccel_context. Members can be set by the user once
 * during initialization or through each AVCodecContext.get_buffer()
 * function call. In any case, they must be valid prior to calling
 * decoding functions.
 *
 * The size of this structure is not a part of the public ABI and must not
 * be used outside of libavcodec. Use av_vdpau_alloc_context() to allocate an
 * AVVDPAUContext.
 */
type AVVDPAUContext struct {

	/**
	 * VDPAU decoder handle
	 *
	 * Set by user.
	 */
	//decoder VdpDecoder
	Decoder uintptr

	/**
	 * VDPAU decoder render callback
	 *
	 * Set by the user.
	 */
	//render *VdpDecoderRender
	Render uintptr

	//render2 AVVDPAU_Render2
	Render2 uintptr
}

/**
 * @brief allocation function for AVVDPAUContext
 *
 * Allows extending the struct without breaking API/ABI
 */
//AVVDPAUContext *av_alloc_vdpaucontext(void);
var avAllocVdpaucontext func() *AVVDPAUContext
var avAllocVdpaucontextOnce sync.Once

func AvAllocVdpaucontext() *AVVDPAUContext {
	avAllocVdpaucontextOnce.Do(func() {
		purego.RegisterLibFunc(&avAllocVdpaucontext, ffcommon.GetAvcodecDll(), "av_alloc_vdpaucontext")
	})
	return avAllocVdpaucontext()
}

// AVVDPAU_Render2 av_vdpau_hwaccel_get_render2(const AVVDPAUContext *);
var avVdpauHwaccelGetRender2 func(c *AVVDPAUContext) uintptr
var avVdpauHwaccelGetRender2Once sync.Once

func (c *AVVDPAUContext) AvVdpauHwaccelGetRender2() uintptr {
	avVdpauHwaccelGetRender2Once.Do(func() {
		purego.RegisterLibFunc(&avVdpauHwaccelGetRender2, ffcommon.GetAvcodecDll(), "av_vdpau_hwaccel_get_render2")
	})
	return avVdpauHwaccelGetRender2(c)
}

// void av_vdpau_hwaccel_set_render2(AVVDPAUContext *, AVVDPAU_Render2);
var avVdpauHwaccelSetRender2 func(c *AVVDPAUContext, r2 uintptr)
var avVdpauHwaccelSetRender2Once sync.Once

func (c *AVVDPAUContext) AvVdpauHwaccelSetRender2(r2 uintptr) {
	avVdpauHwaccelSetRender2Once.Do(func() {
		purego.RegisterLibFunc(&avVdpauHwaccelSetRender2, ffcommon.GetAvcodecDll(), "av_vdpau_hwaccel_set_render2")
	})
	avVdpauHwaccelSetRender2(c, r2)
}

/**
 * Associate a VDPAU device with a codec context for hardware acceleration.
 * This function is meant to be called from the get_format() codec callback,
 * or earlier. It can also be called after avcodec_flush_buffers() to change
 * the underlying VDPAU device mid-stream (e.g. to recover from non-transparent
 * display preemption).
 *
 * @note get_format() must return AV_PIX_FMT_VDPAU if this function completes
 * successfully.
 *
 * @param avctx decoding context whose get_format() callback is invoked
 * @param device VDPAU device handle to use for hardware acceleration
 * @param get_proc_address VDPAU device driver
 * @param flags zero of more OR'd AV_HWACCEL_FLAG_* flags
 *
 * @return 0 on success, an AVERROR code on failure.
 */
//int av_vdpau_bind_context(AVCodecContext *avctx, VdpDevice device,
//VdpGetProcAddress *get_proc_address, unsigned flags);
//todo
var avVdpauBindContext func() ffcommon.FCharP
var avVdpauBindContextOnce sync.Once

func AvVdpauBindContext() ffcommon.FCharP {
	avVdpauBindContextOnce.Do(func() {
		purego.RegisterLibFunc(&avVdpauBindContext, ffcommon.GetAvcodecDll(), "av_vdpau_bind_context")
	})
	return avVdpauBindContext()
}

/**
 * Gets the parameters to create an adequate VDPAU video surface for the codec
 * context using VDPAU hardware decoding acceleration.
 *
 * @note Behavior is undefined if the context was not successfully bound to a
 * VDPAU device using av_vdpau_bind_context().
 *
 * @param avctx the codec context being used for decoding the stream
 * @param type storage space for the VDPAU video surface chroma type
 *              (or NULL to ignore)
 * @param width storage space for the VDPAU video surface pixel width
 *              (or NULL to ignore)
 * @param height storage space for the VDPAU video surface pixel height
 *              (or NULL to ignore)
 *
 * @return 0 on success, a negative AVERROR code on failure.
 */
//int av_vdpau_get_surface_parameters(AVCodecContext *avctx, VdpChromaType *type,
//uint32_t *width, uint32_t *height);
//todo
var avVdpauGetSurfaceParameters func() ffcommon.FCharP
var avVdpauGetSurfaceParametersOnce sync.Once

func AvVdpauGetSurfaceParameters() ffcommon.FCharP {
	avVdpauGetSurfaceParametersOnce.Do(func() {
		purego.RegisterLibFunc(&avVdpauGetSurfaceParameters, ffcommon.GetAvcodecDll(), "av_vdpau_get_surface_parameters")
	})
	return avVdpauGetSurfaceParameters()
}

/**
 * Allocate an AVVDPAUContext.
 *
 * @return Newly-allocated AVVDPAUContext or NULL on failure.
 */
//AVVDPAUContext *av_vdpau_alloc_context(void);
//todo
var avVdpauAllocContext func() ffcommon.FCharP
var avVdpauAllocContextOnce sync.Once

func AvVdpauAllocContext() ffcommon.FCharP {
	avVdpauAllocContextOnce.Do(func() {
		purego.RegisterLibFunc(&avVdpauAllocContext, ffcommon.GetAvcodecDll(), "av_vdpau_alloc_context")
	})
	return avVdpauAllocContext()
}

//#if FF_API_VDPAU_PROFILE
/**
 * Get a decoder profile that should be used for initializing a VDPAU decoder.
 * Should be called from the AVCodecContext.get_format() callback.
 *
 * @deprecated Use av_vdpau_bind_context() instead.
 *
 * @param avctx the codec context being used for decoding the stream
 * @param profile a pointer into which the result will be written on success.
 *                The contents of profile are undefined if this function returns
 *                an error.
 *
 * @return 0 on success (non-negative), a negative AVERROR on failure.
 */
//attribute_deprecated
//int av_vdpau_get_profile(AVCodecContext *avctx, VdpDecoderProfile *profile);
//todo
var avVdpauGetProfile func(avctx *AVCodecContext) ffcommon.FCharP
var avVdpauGetProfileOnce sync.Once

func (avctx *AVCodecContext) AvVdpauGetProfile() ffcommon.FCharP {
	avVdpauGetProfileOnce.Do(func() {
		purego.RegisterLibFunc(&avVdpauGetProfile, ffcommon.GetAvcodecDll(), "av_vdpau_get_profile")
	})
	return avVdpauGetProfile(avctx)
}

//#endif

/* @}*/

//#endif /* AVCODEC_VDPAU_H */
