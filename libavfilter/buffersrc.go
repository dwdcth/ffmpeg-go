package libavfilter

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v7/libavutil"
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

//#ifndef AVFILTER_BUFFERSRC_H
//#define AVFILTER_BUFFERSRC_H
//
///**
// * @file
// * @ingroup lavfi_buffersrc
// * Memory buffer source API.
// */
//
//#include "avfilter.h"

/**
 * @defgroup lavfi_buffersrc Buffer source API
 * @ingroup lavfi
 * @{
 */

const (

	/**
	 * Do not check for format changes.
	 */
	AV_BUFFERSRC_FLAG_NO_CHECK_FORMAT = 1

	/**
	 * Immediately push the frame to the output.
	 */
	AV_BUFFERSRC_FLAG_PUSH = 4

	/**
	 * Keep a reference to the frame.
	 * If the frame if reference-counted, create a new reference; otherwise
	 * copy the frame data.
	 */
	AV_BUFFERSRC_FLAG_KEEP_REF = 8
)

/**
 * Get the number of failed requests.
 *
 * A failed request is when the request_frame method is called while no
 * frame is present in the buffer.
 * The number is reset when a frame is added.
 */
//unsigned av_buffersrc_get_nb_failed_requests(AVFilterContext *buffer_src);
var avBuffersrcGetNbFailedRequests func(bufferSrc *AVFilterContext) ffcommon.FUnsigned
var avBuffersrcGetNbFailedRequestsOnce sync.Once

func (bufferSrc *AVFilterContext) AvBuffersrcGetNbFailedRequests() ffcommon.FUnsigned {
	avBuffersrcGetNbFailedRequestsOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcGetNbFailedRequests, ffcommon.GetAvfilterDll(), "av_buffersrc_get_nb_failed_requests")
	})
	return avBuffersrcGetNbFailedRequests(bufferSrc)
}

/**
 * This structure contains the parameters describing the frames that will be
 * passed to this filter.
 *
 * It should be allocated with av_buffersrc_parameters_alloc() and freed with
 * av_free(). All the allocated fields in it remain owned by the caller.
 */
type AVRational = libavutil.AVRational
type AVBufferRef = libavutil.AVBufferRef
type AVBufferSrcParameters struct {

	/**
	 * video: the pixel format, value corresponds to enum AVPixelFormat
	 * audio: the sample format, value corresponds to enum AVSampleFormat
	 */
	Format ffcommon.FInt
	/**
	 * The timebase to be used for the timestamps on the input frames.
	 */
	TimeBase AVRational

	/**
	 * Video only, the display dimensions of the input frames.
	 */
	Width, Height ffcommon.FInt

	/**
	 * Video only, the sample (pixel) aspect ratio.
	 */
	SampleAspectRatio AVRational

	/**
	 * Video only, the frame rate of the input video. This field must only be
	 * set to a non-zero value if input stream has a known constant framerate
	 * and should be left at its initial value if the framerate is variable or
	 * unknown.
	 */
	FrameRate AVRational

	/**
	 * Video with a hwaccel pixel format only. This should be a reference to an
	 * AVHWFramesContext instance describing the input frames.
	 */
	HwFramesCtx *AVBufferRef

	/**
	 * Audio only, the audio sampling rate in samples per second.
	 */
	SampleRate ffcommon.FInt

	/**
	 * Audio only, the audio channel layout
	 */
	ChannelLayout ffcommon.FUint64T
}

/**
 * Allocate a new AVBufferSrcParameters instance. It should be freed by the
 * caller with av_free().
 */
//AVBufferSrcParameters *av_buffersrc_parameters_alloc(void);
var avBuffersrcParametersAlloc func() *AVBufferSrcParameters
var avBuffersrcParametersAllocOnce sync.Once

func AvBuffersrcParametersAlloc() *AVBufferSrcParameters {
	avBuffersrcParametersAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcParametersAlloc, ffcommon.GetAvfilterDll(), "av_buffersrc_parameters_alloc")
	})
	return avBuffersrcParametersAlloc()
}

/**
 * Initialize the buffersrc or abuffersrc filter with the provided parameters.
 * This function may be called multiple times, the later calls override the
 * previous ones. Some of the parameters may also be set through AVOptions, then
 * whatever method is used last takes precedence.
 *
 * @param ctx an instance of the buffersrc or abuffersrc filter
 * @param param the stream parameters. The frames later passed to this filter
 *              must conform to those parameters. All the allocated fields in
 *              param remain owned by the caller, libavfilter will make internal
 *              copies or references when necessary.
 * @return 0 on success, a negative AVERROR code on failure.
 */
//int av_buffersrc_parameters_set(AVFilterContext *ctx, AVBufferSrcParameters *param);
func (ctx *AVFilterContext) AvBuffersrcParametersSet(param *AVBufferSrcParameters) ffcommon.FInt {
	var avBuffersrcParametersSet func(*AVFilterContext, *AVBufferSrcParameters) ffcommon.FInt
	var avBuffersrcParametersSetOnce sync.Once

	avBuffersrcParametersSetOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcParametersSet, ffcommon.GetAvfilterDll(), "av_buffersrc_parameters_set")
	})

	return avBuffersrcParametersSet(ctx, param)
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will make a new reference to it. Otherwise the frame data will be
 * copied.
 *
 * @return 0 on success, a negative AVERROR on error
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() with the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
//av_warn_unused_result
//int av_buffersrc_write_frame(AVFilterContext *ctx, const AVFrame *frame);
func (ctx *AVFilterContext) AvBuffersrcWriteFrame(frame *AVFrame) ffcommon.FInt {
	var avBuffersrcWriteFrame func(*AVFilterContext, *AVFrame) ffcommon.FInt
	var avBuffersrcWriteFrameOnce sync.Once

	avBuffersrcWriteFrameOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcWriteFrame, ffcommon.GetAvfilterDll(), "av_buffersrc_write_frame")
	})

	return avBuffersrcWriteFrame(ctx, frame)
}

/**
 * Add a frame to the buffer source.
 *
 * @param ctx   an instance of the buffersrc filter
 * @param frame frame to be added. If the frame is reference counted, this
 * function will take ownership of the reference(s) and reset the frame.
 * Otherwise the frame data will be copied. If this function returns an error,
 * the input frame is not touched.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @note the difference between this function and av_buffersrc_write_frame() is
 * that av_buffersrc_write_frame() creates a new reference to the input frame,
 * while this function takes ownership of the reference passed to it.
 *
 * This function is equivalent to av_buffersrc_add_frame_flags() without the
 * AV_BUFFERSRC_FLAG_KEEP_REF flag.
 */
//av_warn_unused_result
//int av_buffersrc_add_frame(AVFilterContext *ctx, AVFrame *frame);
func (ctx *AVFilterContext) AvBuffersrcAddFrame(frame *AVFrame) ffcommon.FInt {
	var avBuffersrcAddFrame func(*AVFilterContext, *AVFrame) ffcommon.FInt
	var avBuffersrcAddFrameOnce sync.Once

	avBuffersrcAddFrameOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcAddFrame, ffcommon.GetAvfilterDll(), "av_buffersrc_add_frame")
	})

	return avBuffersrcAddFrame(ctx, frame)
}

/**
 * Add a frame to the buffer source.
 *
 * By default, if the frame is reference-counted, this function will take
 * ownership of the reference(s) and reset the frame. This can be controlled
 * using the flags.
 *
 * If this function returns an error, the input frame is not touched.
 *
 * @param buffer_src  pointer to a buffer source context
 * @param frame       a frame, or NULL to mark EOF
 * @param flags       a combination of AV_BUFFERSRC_FLAG_*
 * @return            >= 0 in case of success, a negative AVERROR code
 *                    in case of failure
 */
//av_warn_unused_result
//int av_buffersrc_add_frame_flags(AVFilterContext *buffer_src,
//AVFrame *frame, int flags);
func (ctx *AVFilterContext) AvBuffersrcAddFrameFlags(frame *AVFrame, flags ffcommon.FInt) ffcommon.FInt {
	var avBuffersrcAddFrameFlags func(*AVFilterContext, *AVFrame, ffcommon.FInt) ffcommon.FInt
	var avBuffersrcAddFrameFlagsOnce sync.Once

	avBuffersrcAddFrameFlagsOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcAddFrameFlags, ffcommon.GetAvfilterDll(), "av_buffersrc_add_frame_flags")
	})

	return avBuffersrcAddFrameFlags(ctx, frame, flags)
}

/**
 * Close the buffer source after EOF.
 *
 * This is similar to passing NULL to av_buffersrc_add_frame_flags()
 * except it takes the timestamp of the EOF, i.e. the timestamp of the end
 * of the last frame.
 */
//int av_buffersrc_close(AVFilterContext *ctx, int64_t pts, unsigned flags);

func (ctx *AVFilterContext) AvBuffersrcClose(pts ffcommon.FInt, flags ffcommon.FUnsigned) ffcommon.FInt {
	var avBuffersrcClose func(*AVFilterContext, ffcommon.FInt, ffcommon.FUnsigned) ffcommon.FInt
	var avBuffersrcCloseOnce sync.Once

	avBuffersrcCloseOnce.Do(func() {
		purego.RegisterLibFunc(&avBuffersrcClose, ffcommon.GetAvfilterDll(), "av_buffersrc_close")
	})

	return avBuffersrcClose(ctx, pts, flags)
}

/**
 * @}
 */

//#endif /* AVFILTER_BUFFERSRC_H */
