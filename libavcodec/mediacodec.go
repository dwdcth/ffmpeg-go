package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Android MediaCodec public API
 *
 * Copyright (c) 2016 Matthieu Bouron <matthieu.bouron stupeflix.com>
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

//#ifndef AVCODEC_MEDIACODEC_H
//#define AVCODEC_MEDIACODEC_H
//
//#include "../libavcodec/avcodec.h"

/**
 * This structure holds a reference to a android/view/Surface object that will
 * be used as output by the decoder.
 *
 */
type AVMediaCodecContext struct {

	/**
	 * android/view/Surface object reference.
	 */
	Surface ffcommon.FVoidP
}

/**
 * Allocate and initialize a MediaCodec context.
 *
 * When decoding with MediaCodec is finished, the caller must free the
 * MediaCodec context with av_mediacodec_default_free.
 *
 * @return a pointer to a newly allocated AVMediaCodecContext on success, NULL otherwise
 */
//AVMediaCodecContext *av_mediacodec_alloc_context(void);
var avMediacodecAllocContext func() *AVMediaCodecContext
var avMediacodecAllocContextOnce sync.Once

func AvMediacodecAllocContext() *AVMediaCodecContext {
	avMediacodecAllocContextOnce.Do(func() {
		purego.RegisterLibFunc(&avMediacodecAllocContext, ffcommon.GetAvcodecDll(), "av_mediacodec_alloc_context")
	})
	return avMediacodecAllocContext()
}

/**
 * Convenience function that sets up the MediaCodec context.
 *
 * @param avctx codec context
 * @param ctx MediaCodec context to initialize
 * @param surface reference to an android/view/Surface
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_default_init(AVCodecContext *avctx, AVMediaCodecContext *ctx, void *surface);
var avMediacodecDefaultInit func(avctx *AVCodecContext, ctx *AVMediaCodecContext, surface ffcommon.FVoidP) ffcommon.FInt
var avMediacodecDefaultInitOnce sync.Once

func (avctx *AVCodecContext) AvMediacodecDefaultInit(ctx *AVMediaCodecContext, surface ffcommon.FVoidP) ffcommon.FInt {
	avMediacodecDefaultInitOnce.Do(func() {
		purego.RegisterLibFunc(&avMediacodecDefaultInit, ffcommon.GetAvcodecDll(), "av_mediacodec_default_init")
	})
	return avMediacodecDefaultInit(avctx, ctx, surface)
}

/**
 * This function must be called to free the MediaCodec context initialized with
 * av_mediacodec_default_init().
 *
 * @param avctx codec context
 */
//void av_mediacodec_default_free(AVCodecContext *avctx);
var avMediacodecDefaultFree func(avctx *AVCodecContext)
var avMediacodecDefaultFreeOnce sync.Once

func (avctx *AVCodecContext) AvMediacodecDefaultFree() {
	avMediacodecDefaultFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avMediacodecDefaultFree, ffcommon.GetAvcodecDll(), "av_mediacodec_default_free")
	})
	avMediacodecDefaultFree(avctx)
}

/**
 * Opaque structure representing a MediaCodec buffer to render.
 */
//typedef struct MediaCodecBuffer AVMediaCodecBuffer;
type AVMediaCodecBuffer struct {
}

/**
 * Release a MediaCodec buffer and render it to the surface that is associated
 * with the decoder. This function should only be called once on a given
 * buffer, once released the underlying buffer returns to the codec, thus
 * subsequent calls to this function will have no effect.
 *
 * @param buffer the buffer to render
 * @param render 1 to release and render the buffer to the surface or 0 to
 * discard the buffer
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_release_buffer(AVMediaCodecBuffer *buffer, int render);
var avMediacodecReleaseBuffer func(buffer *AVMediaCodecBuffer, render ffcommon.FInt) ffcommon.FInt
var avMediacodecReleaseBufferOnce sync.Once

func (buffer *AVMediaCodecBuffer) AvMediacodecReleaseBuffer(render ffcommon.FInt) ffcommon.FInt {
	avMediacodecReleaseBufferOnce.Do(func() {
		purego.RegisterLibFunc(&avMediacodecReleaseBuffer, ffcommon.GetAvcodecDll(), "av_mediacodec_release_buffer")
	})
	return avMediacodecReleaseBuffer(buffer, render)
}

/**
 * Release a MediaCodec buffer and render it at the given time to the surface
 * that is associated with the decoder. The timestamp must be within one second
 * of the current java/lang/System#nanoTime() (which is implemented using
 * CLOCK_MONOTONIC on Android). See the Android MediaCodec documentation
 * of android/media/MediaCodec#releaseOutputBuffer(int,long) for more details.
 *
 * @param buffer the buffer to render
 * @param time timestamp in nanoseconds of when to render the buffer
 * @return 0 on success, < 0 otherwise
 */
//int av_mediacodec_render_buffer_at_time(AVMediaCodecBuffer *buffer, int64_t time);
var avMediacodecRenderBufferAtTime func(buffer *AVMediaCodecBuffer, time ffcommon.FInt64T) ffcommon.FInt
var avMediacodecRenderBufferAtTimeOnce sync.Once

func (buffer *AVMediaCodecBuffer) AvMediacodecRenderBufferAtTime(time ffcommon.FInt64T) ffcommon.FInt {
	avMediacodecRenderBufferAtTimeOnce.Do(func() {
		purego.RegisterLibFunc(&avMediacodecRenderBufferAtTime, ffcommon.GetAvcodecDll(), "av_mediacodec_render_buffer_at_time")
	})
	return avMediacodecRenderBufferAtTime(buffer, time)
}

//#endif /* AVCODEC_MEDIACODEC_H */
