package libavfilter

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/dwdcth/ffmpeg-go/libavutil"
	"github.com/ebitengine/purego"
)

/*
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

//#ifndef AVFILTER_BUFFERSINK_H
//#define AVFILTER_BUFFERSINK_H

/**
 * @file
 * @ingroup lavfi_buffersink
 * memory buffer sink API for audio and video
 */

//#include "avfilter.h"

/**
 * @defgroup lavfi_buffersink Buffer sink API
 * @ingroup lavfi
 * @{
 *
 * The buffersink and abuffersink filters are there to connect filter graphs
 * to applications. They have a single input, connected to the graph, and no
 * output. Frames must be extracted using av_buffersink_get_frame() or
 * av_buffersink_get_samples().
 *
 * The format negotiated by the graph during configuration can be obtained
 * using the accessor functions:
 * - av_buffersink_get_time_base(),
 * - av_buffersink_get_format(),
 * - av_buffersink_get_frame_rate(),
 * - av_buffersink_get_w(),
 * - av_buffersink_get_h(),
 * - av_buffersink_get_sample_aspect_ratio(),
 * - av_buffersink_get_channels(),
 * - av_buffersink_get_channel_layout(),
 * - av_buffersink_get_sample_rate().
 *
 * The format can be constrained by setting options, using av_opt_set() and
 * related functions with the AV_OPT_SEARCH_CHILDREN flag.
 *  - pix_fmts (int list),
 *  - sample_fmts (int list),
 *  - sample_rates (int list),
 *  - channel_layouts (int64_t),
 *  - channel_counts (int list),
 *  - all_channel_counts (bool).
 * Most of these options are of type binary, and should be set using
 * av_opt_set_int_list() or av_opt_set_bin(). If they are not set, all
 * corresponding formats are accepted.
 *
 * As a special case, if neither channel_layouts nor channel_counts is set,
 * all valid channel layouts are accepted, but channel counts without a
 * layout are not, unless all_channel_counts is set.
 * Also, channel_layouts must not contain a channel layout already accepted
 * by a value in channel_counts; for example, if channel_counts contains 2,
 * then channel_layouts must not contain stereo.
 */

/**
 * Get a frame with filtered data from sink and put it in frame.
 *
 * @param ctx    pointer to a buffersink or abuffersink filter context.
 * @param frame  pointer to an allocated frame that will be filled with data.
 *               The data must be freed using av_frame_unref() / av_frame_free()
 * @param flags  a combination of AV_BUFFERSINK_FLAG_* flags
 *
 * @return  >= 0 in for success, a negative AVERROR code for failure.
 */
//int av_buffersink_get_frame_flags(AVFilterContext *ctx, AVFrame *frame, int flags);
func (ctx *AVFilterContext) AvBuffersinkGetFrameFlags(frame *AVFrame, flags ffcommon.FInt) ffcommon.FInt {
	var avBuffersinkGetFrameFlags func(*AVFilterContext, *AVFrame, ffcommon.FInt) ffcommon.FInt
	var avBuffersinkGetFrameFlagsOnce sync.Once

	avBuffersinkGetFrameFlagsOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetFrameFlags, ffcommon.GetAvfilterDll(), "av_buffersink_get_frame_flags")
	})

	return avBuffersinkGetFrameFlags(ctx, frame, flags)
}

/**
 * Tell av_buffersink_get_buffer_ref() to read video/samples buffer
 * reference, but not remove it from the buffer. This is useful if you
 * need only to read a video/samples buffer, without to fetch it.
 */
const AV_BUFFERSINK_FLAG_PEEK = 1

/**
 * Tell av_buffersink_get_buffer_ref() not to request a frame from its input.
 * If a frame is already buffered, it is read (and removed from the buffer),
 * but if no frame is present, return AVERROR(EAGAIN).
 */
const AV_BUFFERSINK_FLAG_NO_REQUEST = 2

//#if FF_API_BUFFERSINK_ALLOC
/**
 * Deprecated and unused struct to use for initializing a buffersink context.
 */
type AVPixelFormat = libavutil.AVPixelFormat
type AVBufferSinkParams struct {
	PixelFmts *AVPixelFormat ///< list of allowed pixel formats, terminated by AV_PIX_FMT_NONE
}

/**
 * Create an AVBufferSinkParams structure.
 *
 * Must be freed with av_free().
 */
//attribute_deprecated
//AVBufferSinkParams *av_buffersink_params_alloc(void);
//todo
var avBuffersinkParamsAlloc func() uintptr
var avBuffersinkParamsAllocOnce sync.Once

func AvBuffersinkParamsAlloc() string {
	avBuffersinkParamsAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkParamsAlloc, ffcommon.GetAvfilterDll(), "av_buffersink_params_alloc")
	})

	t := avBuffersinkParamsAlloc()
	return ffcommon.StringFromPtr(t)
}

/**
 * Deprecated and unused struct to use for initializing an abuffersink context.
 */
type AVSampleFormat = libavutil.AVSampleFormat
type AVABufferSinkParams struct {
	SampleFmts       *AVSampleFormat   ///< list of allowed sample formats, terminated by AV_SAMPLE_FMT_NONE
	ChannelLayouts   *ffcommon.FInt64T ///< list of allowed channel layouts, terminated by -1
	ChannelCounts    *ffcommon.FInt    ///< list of allowed channel counts, terminated by -1
	AllChannelCounts ffcommon.FInt     ///< if not 0, accept any channel count or layout
	SampleRates      *ffcommon.FInt    ///< list of allowed sample rates, terminated by -1
}

/**
 * Create an AVABufferSinkParams structure.
 *
 * Must be freed with av_free().
 */
//attribute_deprecated
//AVABufferSinkParams *av_abuffersink_params_alloc(void);

var avAbuffersinkParamsAlloc func() *AVABufferSinkParams
var avAbuffersinkParamsAllocOnce sync.Once

func AvAbuffersinkParamsAlloc() *AVABufferSinkParams {
	avAbuffersinkParamsAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avAbuffersinkParamsAlloc, ffcommon.GetAvfilterDll(), "av_abuffersink_params_alloc")
	})
	return avAbuffersinkParamsAlloc()
}

//#endif

/**
 * Set the frame size for an audio buffer sink.
 *
 * All calls to av_buffersink_get_buffer_ref will return a buffer with
 * exactly the specified number of samples, or AVERROR(EAGAIN) if there is
 * not enough. The last buffer at EOF will be padded with 0.
 */
//void av_buffersink_set_frame_size(AVFilterContext *ctx, unsigned frame_size);
var avBuffersinkSetFrameSize func(ctx *AVFilterContext, frame_size ffcommon.FUnsigned)
var avBuffersinkSetFrameSizeOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkSetFrameSize(frame_size ffcommon.FUnsigned) {
	avBuffersinkSetFrameSizeOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkSetFrameSize, ffcommon.GetAvfilterDll(), "av_buffersink_set_frame_size")
	})
	avBuffersinkSetFrameSize(ctx, frame_size)
}

/**
 * @defgroup lavfi_buffersink_accessors Buffer sink accessors
 * Get the properties of the stream
 * @{
 */

// enum AVMediaType av_buffersink_get_type                (const AVFilterContext *ctx);
var avBuffersinkGetType func(ctx *AVFilterContext) AVMediaType
var avBuffersinkGetTypeOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetType() AVMediaType {
	avBuffersinkGetTypeOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetType, ffcommon.GetAvfilterDll(), "av_buffersink_get_type")
	})
	return avBuffersinkGetType(ctx)
}

// AVRational       av_buffersink_get_time_base           (const AVFilterContext *ctx);
var avBuffersinkGetTimeBase func(ctx *AVFilterContext) AVRational
var avBuffersinkGetTimeBaseOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetTimeBase() AVRational {
	avBuffersinkGetTimeBaseOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetTimeBase, ffcommon.GetAvfilterDll(), "av_buffersink_get_time_base")
	})
	return avBuffersinkGetTimeBase(ctx)
}

// int              av_buffersink_get_format              (const AVFilterContext *ctx);
var avBuffersinkGetFormat func(ctx *AVFilterContext) ffcommon.FInt
var avBuffersinkGetFormatOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetFormat() ffcommon.FInt {
	avBuffersinkGetFormatOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetFormat, ffcommon.GetAvfilterDll(), "av_buffersink_get_format")
	})
	return avBuffersinkGetFormat(ctx)
}

// AVRational       av_buffersink_get_frame_rate          (const AVFilterContext *ctx);
var avBuffersinkGetFrameRate func(ctx *AVFilterContext) AVRational
var avBuffersinkGetFrameRateOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetFrameRate() AVRational {
	avBuffersinkGetFrameRateOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetFrameRate, ffcommon.GetAvfilterDll(), "av_buffersink_get_frame_rate")
	})
	return avBuffersinkGetFrameRate(ctx)
}

// int              av_buffersink_get_w                   (const AVFilterContext *ctx);
var avBuffersinkGetW func(ctx *AVFilterContext) ffcommon.FInt
var avBuffersinkGetWOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetW() ffcommon.FInt {
	avBuffersinkGetWOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetW, ffcommon.GetAvfilterDll(), "av_buffersink_get_w")
	})
	return avBuffersinkGetW(ctx)
}

// int              av_buffersink_get_h                   (const AVFilterContext *ctx);
var avBuffersinkGetH func(ctx *AVFilterContext) ffcommon.FInt
var avBuffersinkGetHOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetH() ffcommon.FInt {
	avBuffersinkGetHOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetH, ffcommon.GetAvfilterDll(), "av_buffersink_get_h")
	})
	return avBuffersinkGetH(ctx)
}

// AVRational       av_buffersink_get_sample_aspect_ratio (const AVFilterContext *ctx);
var avBuffersinkGetSampleAspectRatio func(ctx *AVFilterContext) AVRational
var avBuffersinkGetSampleAspectRatioOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetSampleAspectRatio() AVRational {
	avBuffersinkGetSampleAspectRatioOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetSampleAspectRatio, ffcommon.GetAvfilterDll(), "av_buffersink_get_sample_aspect_ratio")
	})
	return avBuffersinkGetSampleAspectRatio(ctx)
}

// int              av_buffersink_get_channels            (const AVFilterContext *ctx);
var avBuffersinkGetChannels func(ctx *AVFilterContext) ffcommon.FInt
var avBuffersinkGetChannelsOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetChannels() ffcommon.FInt {
	avBuffersinkGetChannelsOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetChannels, ffcommon.GetAvfilterDll(), "av_buffersink_get_channels")
	})
	return avBuffersinkGetChannels(ctx)
}

// uint64_t         av_buffersink_get_channel_layout      (const AVFilterContext *ctx);
var avBuffersinkGetChannelLayout func(ctx *AVFilterContext) ffcommon.FUint64T
var avBuffersinkGetChannelLayoutOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetChannelLayout() ffcommon.FUint64T {
	avBuffersinkGetChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetChannelLayout, ffcommon.GetAvfilterDll(), "av_buffersink_get_channel_layout")
	})
	return avBuffersinkGetChannelLayout(ctx)
}

// int              av_buffersink_get_sample_rate         (const AVFilterContext *ctx);
var avBuffersinkGetSampleRate func(ctx *AVFilterContext) ffcommon.FInt
var avBuffersinkGetSampleRateOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetSampleRate() ffcommon.FInt {
	avBuffersinkGetSampleRateOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetSampleRate, ffcommon.GetAvfilterDll(), "av_buffersink_get_sample_rate")
	})
	return avBuffersinkGetSampleRate(ctx)
}

// AVBufferRef *    av_buffersink_get_hw_frames_ctx       (const AVFilterContext *ctx);
var avBuffersinkGetHwFramesCtx func(ctx *AVFilterContext) *AVBufferRef
var avBuffersinkGetHwFramesCtxOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetHwFramesCtx() *AVBufferRef {
	avBuffersinkGetHwFramesCtxOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetHwFramesCtx, ffcommon.GetAvfilterDll(), "av_buffersink_get_hw_frames_ctx")
	})
	return avBuffersinkGetHwFramesCtx(ctx)
}

/** @} */

/**
 * Get a frame with filtered data from sink and put it in frame.
 *
 * @param ctx pointer to a context of a buffersink or abuffersink AVFilter.
 * @param frame pointer to an allocated frame that will be filled with data.
 *              The data must be freed using av_frame_unref() / av_frame_free()
 *
 * @return
 *         - >= 0 if a frame was successfully returned.
 *         - AVERROR(EAGAIN) if no frames are available at this point; more
 *           input frames must be added to the filtergraph to get more output.
 *         - AVERROR_EOF if there will be no more output frames on this sink.
 *         - A different negative AVERROR code in other failure cases.
 */
//int av_buffersink_get_frame(AVFilterContext *ctx, AVFrame *frame);
var avBuffersinkGetFrame func(ctx *AVFilterContext, frame *AVFrame) ffcommon.FInt
var avBuffersinkGetFrameOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetFrame(frame *AVFrame) ffcommon.FInt {
	avBuffersinkGetFrameOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetFrame, ffcommon.GetAvfilterDll(), "av_buffersink_get_frame")
	})
	return avBuffersinkGetFrame(ctx, frame)
}

/**
 * Same as av_buffersink_get_frame(), but with the ability to specify the number
 * of samples read. This function is less efficient than
 * av_buffersink_get_frame(), because it copies the data around.
 *
 * @param ctx pointer to a context of the abuffersink AVFilter.
 * @param frame pointer to an allocated frame that will be filled with data.
 *              The data must be freed using av_frame_unref() / av_frame_free()
 *              frame will contain exactly nb_samples audio samples, except at
 *              the end of stream, when it can contain less than nb_samples.
 *
 * @return The return codes have the same meaning as for
 *         av_buffersink_get_frame().
 *
 * @warning do not mix this function with av_buffersink_get_frame(). Use only one or
 * the other with a single sink, not both.
 */
//int av_buffersink_get_samples(AVFilterContext *ctx, AVFrame *frame, int nb_samples);
var avBuffersinkGetSamples func(ctx *AVFilterContext, frame *AVFrame, nb_samples ffcommon.FInt) ffcommon.FInt
var avBuffersinkGetSamplesOnce sync.Once

func (ctx *AVFilterContext) AvBuffersinkGetSamples(frame *AVFrame, nb_samples ffcommon.FInt) ffcommon.FInt {
	avBuffersinkGetSamplesOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersinkGetSamples, ffcommon.GetAvfilterDll(), "av_buffersink_get_samples")
	})
	return avBuffersinkGetSamples(ctx, frame, nb_samples)
}

/**
 * @}
 */

//#endif /* AVFILTER_BUFFERSINK_H */
