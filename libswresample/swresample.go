package libswresample

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v6/libavutil"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (C) 2011-2013 Michael Niedermayer (michaelni@gmx.at)
 *
 * This file is part of libswresample
 *
 * libswresample is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * libswresample is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with libswresample; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef SWRESAMPLE_SWRESAMPLE_H
//#define SWRESAMPLE_SWRESAMPLE_H

/**
 * @file
 * @ingroup lswr
 * libswresample public header
 */

/**
 * @defgroup lswr libswresample
 * @{
 *
 * Audio resampling, sample format conversion and mixing library.
 *
 * Interaction with lswr is done through SwrContext, which is
 * allocated with swr_alloc() or swr_alloc_set_opts(). It is opaque, so all parameters
 * must be set with the @ref avoptions API.
 *
 * The first thing you will need to do in order to use lswr is to allocate
 * SwrContext. This can be done with swr_alloc() or swr_alloc_set_opts(). If you
 * are using the former, you must set options through the @ref avoptions API.
 * The latter function provides the same feature, but it allows you to set some
 * common options in the same statement.
 *
 * For example the following code will setup conversion from planar float sample
 * format to interleaved signed 16-bit integer, downsampling from 48kHz to
 * 44.1kHz and downmixing from 5.1 channels to stereo (using the default mixing
 * matrix). This is using the swr_alloc() function.
 * @code
 * SwrContext *swr = swr_alloc();
 * av_opt_set_channel_layout(swr, "in_channel_layout",  AV_CH_LAYOUT_5POINT1, 0);
 * av_opt_set_channel_layout(swr, "out_channel_layout", AV_CH_LAYOUT_STEREO,  0);
 * av_opt_set_int(swr, "in_sample_rate",     48000,                0);
 * av_opt_set_int(swr, "out_sample_rate",    44100,                0);
 * av_opt_set_sample_fmt(swr, "in_sample_fmt",  AV_SAMPLE_FMT_FLTP, 0);
 * av_opt_set_sample_fmt(swr, "out_sample_fmt", AV_SAMPLE_FMT_S16,  0);
 * @endcode
 *
 * The same job can be done using swr_alloc_set_opts() as well:
 * @code
 * SwrContext *swr = swr_alloc_set_opts(NULL,  // we're allocating a new context
 *                       AV_CH_LAYOUT_STEREO,  // out_ch_layout
 *                       AV_SAMPLE_FMT_S16,    // out_sample_fmt
 *                       44100,                // out_sample_rate
 *                       AV_CH_LAYOUT_5POINT1, // in_ch_layout
 *                       AV_SAMPLE_FMT_FLTP,   // in_sample_fmt
 *                       48000,                // in_sample_rate
 *                       0,                    // log_offset
 *                       NULL);                // log_ctx
 * @endcode
 *
 * Once all values have been set, it must be initialized with swr_init(). If
 * you need to change the conversion parameters, you can change the parameters
 * using @ref AVOptions, as described above in the first example; or by using
 * swr_alloc_set_opts(), but with the first argument the allocated context.
 * You must then call swr_init() again.
 *
 * The conversion itself is done by repeatedly calling swr_convert().
 * Note that the samples may get buffered in swr if you provide insufficient
 * output space or if sample rate conversion is done, which requires "future"
 * samples. Samples that do not require future input can be retrieved at any
 * time by using swr_convert() (in_count can be set to 0).
 * At the end of conversion the resampling buffer can be flushed by calling
 * swr_convert() with NULL in and 0 in_count.
 *
 * The samples used in the conversion process can be managed with the libavutil
 * @ref lavu_sampmanip "samples manipulation" API, including av_samples_alloc()
 * function used in the following example.
 *
 * The delay between input and output, can at any time be found by using
 * swr_get_delay().
 *
 * The following code demonstrates the conversion loop assuming the parameters
 * from above and caller-defined functions get_input() and handle_output():
 * @code
 * uint8_t **input;
 * int in_samples;
 *
 * while (get_input(&input, &in_samples)) {
 *     uint8_t *output;
 *     int out_samples = av_rescale_rnd(swr_get_delay(swr, 48000) +
 *                                      in_samples, 44100, 48000, AV_ROUND_UP);
 *     av_samples_alloc(&output, NULL, 2, out_samples,
 *                      AV_SAMPLE_FMT_S16, 0);
 *     out_samples = swr_convert(swr, &output, out_samples,
 *                                      input, in_samples);
 *     handle_output(output, out_samples);
 *     av_freep(&output);
 * }
 * @endcode
 *
 * When the conversion is finished, the conversion
 * context and everything associated with it must be freed with swr_free().
 * A swr_close() function is also available, but it exists mainly for
 * compatibility with libavresample, and is not required to be called.
 *
 * There will be no memory leak if the data is not completely flushed before
 * swr_free().
 */

//#include <stdint.h>
//#include "../libavutil/channel_layout.h"
//#include "../libavutil/frame.h"
//#include "../libavutil/samplefmt.h"
//
//#include "../libswresample/version.h"

/**
 * @name Option constants
 * These constants are used for the @ref avoptions interface for lswr.
 * @{
 *
 */

const SWR_FLAG_RESAMPLE = 1 ///< Force resampling even if equal sample rate
//TODO use int resample ?
//long term TODO can we enable this dynamically?

/** Dithering algorithms */
type SwrDitherType int32

const (
	SWR_DITHER_NONE = iota
	SWR_DITHER_RECTANGULAR
	SWR_DITHER_TRIANGULAR
	SWR_DITHER_TRIANGULAR_HIGHPASS
)
const (
	SWR_DITHER_NS = 64 + iota ///< not part of API/ABI
	SWR_DITHER_NS_LIPSHITZ
	SWR_DITHER_NS_F_WEIGHTED
	SWR_DITHER_NS_MODIFIED_E_WEIGHTED
	SWR_DITHER_NS_IMPROVED_E_WEIGHTED
	SWR_DITHER_NS_SHIBATA
	SWR_DITHER_NS_LOW_SHIBATA
	SWR_DITHER_NS_HIGH_SHIBATA
	SWR_DITHER_NB ///< not part of API/ABI
)

/** Resampling Engines */
type SwrEngine int32

const (
	SWR_ENGINE_SWR  = iota /**< SW Resampler */
	SWR_ENGINE_SOXR        /**< SoX Resampler */
	SWR_ENGINE_NB          ///< not part of API/ABI
)

/** Resampling Filter Types */
type SwrFilterType int32

const (
	SWR_FILTER_TYPE_CUBIC            = iota /**< Cubic */
	SWR_FILTER_TYPE_BLACKMAN_NUTTALL        /**< Blackman Nuttall windowed sinc */
	SWR_FILTER_TYPE_KAISER                  /**< Kaiser windowed sinc */
)

/**
 * @}
 */

/**
 * The libswresample context. Unlike libavcodec and libavformat, this structure
 * is opaque. This means that if you would like to set options, you must use
 * the @ref avoptions API and cannot directly set values to members of the
 * structure.
 */
//typedef struct SwrContext SwrContext;
type SwrContext struct {
}

type AVClass = libavutil.AVClass

/**
 * Get the AVClass for SwrContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 * @return the AVClass of SwrContext
 */
//const AVClass *swr_get_class(void);
var swrGetClass func() *AVClass
var swrGetClassOnce sync.Once

// SwrGetClass is a purego function to get the AVClass for SwrContext.
func SwrGetClass() *AVClass {
	swrGetClassOnce.Do(func() {
		purego.RegisterLibFunc(&swrGetClass, ffcommon.GetAvswresampleDll(), "swr_get_class")
	})
	return swrGetClass()
}

/**
 * @name SwrContext constructor functions
 * @{
 */

/**
 * Allocate SwrContext.
 *
 * If you use this function you will need to set the parameters (manually or
 * with swr_alloc_set_opts()) before calling swr_init().
 *
 * @see swr_alloc_set_opts(), swr_init(), swr_free()
 * @return NULL on error, allocated context otherwise
 */
//struct SwrContext *swr_alloc(void);
var swrAlloc func() *SwrContext
var swrAllocOnce sync.Once

// SwrAlloc is a purego function to allocate a SwrContext.
func SwrAlloc() *SwrContext {
	swrAllocOnce.Do(func() {
		purego.RegisterLibFunc(&swrAlloc, ffcommon.GetAvswresampleDll(), "swr_alloc")
	})
	return swrAlloc()
}

/**
 * Initialize context after user parameters have been set.
 * @note The context must be configured using the AVOption API.
 *
 * @see av_opt_set_int()
 * @see av_opt_set_dict()
 *
 * @param[in,out]   s Swr context to initialize
 * @return AVERROR error code in case of failure.
 */
//int swr_init(struct SwrContext *s);
var swrInit func(s *SwrContext) ffcommon.FInt
var swrInitOnce sync.Once

// SwrInit is a purego method to initialize a SwrContext.
func (s *SwrContext) SwrInit() ffcommon.FInt {
	swrInitOnce.Do(func() {
		purego.RegisterLibFunc(&swrInit, ffcommon.GetAvswresampleDll(), "swr_init")
	})
	return swrInit(s)
}

/**
 * Check whether an swr context has been initialized or not.
 *
 * @param[in]       s Swr context to check
 * @see swr_init()
 * @return positive if it has been initialized, 0 if not initialized
 */
//int swr_is_initialized(struct SwrContext *s);
var swrIsInitialized func(s *SwrContext) ffcommon.FInt
var swrIsInitializedOnce sync.Once

// SwrIsInitialized is a purego method to check if a SwrContext is initialized.
func (s *SwrContext) SwrIsInitialized() ffcommon.FInt {
	swrIsInitializedOnce.Do(func() {
		purego.RegisterLibFunc(&swrIsInitialized, ffcommon.GetAvswresampleDll(), "swr_is_initialized")
	})
	return swrIsInitialized(s)
}

/**
 * Allocate SwrContext if needed and set/reset common parameters.
 *
 * This function does not require s to be allocated with swr_alloc(). On the
 * other hand, swr_alloc() can use swr_alloc_set_opts() to set the parameters
 * on the allocated context.
 *
 * @param s               existing Swr context if available, or NULL if not
 * @param out_ch_layout   output channel layout (AV_CH_LAYOUT_*)
 * @param out_sample_fmt  output sample format (AV_SAMPLE_FMT_*).
 * @param out_sample_rate output sample rate (frequency in Hz)
 * @param in_ch_layout    input channel layout (AV_CH_LAYOUT_*)
 * @param in_sample_fmt   input sample format (AV_SAMPLE_FMT_*).
 * @param in_sample_rate  input sample rate (frequency in Hz)
 * @param log_offset      logging level offset
 * @param log_ctx         parent logging context, can be NULL
 *
 * @see swr_init(), swr_free()
 * @return NULL on error, allocated context otherwise
 */
//struct SwrContext *swr_alloc_set_opts(struct SwrContext *s,
//int64_t out_ch_layout, enum AVSampleFormat out_sample_fmt, int out_sample_rate,
//int64_t  in_ch_layout, enum AVSampleFormat  in_sample_fmt, int  in_sample_rate,
//int log_offset, void *log_ctx);
type AVSampleFormat = libavutil.AVSampleFormat

var swrAllocSetOpts func(s *SwrContext, outChLayout ffcommon.FInt64T, outSampleFmt AVSampleFormat, outSampleRate ffcommon.FInt, inChLayout ffcommon.FInt64T, inSampleFmt AVSampleFormat, inSampleRate, logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) *SwrContext
var swrAllocSetOptsOnce sync.Once

// SwrAllocSetOpts is a purego method to allocate a SwrContext with options.
func (s *SwrContext) SwrAllocSetOpts(outChLayout ffcommon.FInt64T, outSampleFmt AVSampleFormat, outSampleRate ffcommon.FInt, inChLayout ffcommon.FInt64T, inSampleFmt AVSampleFormat, inSampleRate, logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) *SwrContext {
	swrAllocSetOptsOnce.Do(func() {
		purego.RegisterLibFunc(&swrAllocSetOpts, ffcommon.GetAvswresampleDll(), "swr_alloc_set_opts")
	})
	return swrAllocSetOpts(s, outChLayout, outSampleFmt, outSampleRate, inChLayout, inSampleFmt, inSampleRate, logOffset, logCtx)
}

/**
 * @}
 *
 * @name SwrContext destructor functions
 * @{
 */

/**
 * Free the given SwrContext and set the pointer to NULL.
 *
 * @param[in] s a pointer to a pointer to Swr context
 */
//void swr_free(struct SwrContext **s);
var swrFree func(s **SwrContext)
var swrFreeOnce sync.Once

// SwrFree is a purego function to free a SwrContext.
func SwrFree(s **SwrContext) {
	swrFreeOnce.Do(func() {
		purego.RegisterLibFunc(&swrFree, ffcommon.GetAvswresampleDll(), "swr_free")
	})
	swrFree(s)
}

/**
 * Closes the context so that swr_is_initialized() returns 0.
 *
 * The context can be brought back to life by running swr_init(),
 * swr_init() can also be used without swr_close().
 * This function is mainly provided for simplifying the usecase
 * where one tries to support libavresample and libswresample.
 *
 * @param[in,out] s Swr context to be closed
 */
//void swr_close(struct SwrContext *s);
var swrClose func(s *SwrContext)
var swrCloseOnce sync.Once

// SwrClose is a purego method to close a SwrContext.
func (s *SwrContext) SwrClose() {
	swrCloseOnce.Do(func() {
		purego.RegisterLibFunc(&swrClose, ffcommon.GetAvswresampleDll(), "swr_close")
	})
	swrClose(s)
}

/**
 * @}
 *
 * @name Core conversion functions
 * @{
 */

/** Convert audio.
 *
 * in and in_count can be set to 0 to flush the last few samples out at the
 * end.
 *
 * If more input is provided than output space, then the input will be buffered.
 * You can avoid this buffering by using swr_get_out_samples() to retrieve an
 * upper bound on the required number of output samples for the given number of
 * input samples. Conversion will run directly without copying whenever possible.
 *
 * @param s         allocated Swr context, with parameters set
 * @param out       output buffers, only the first one need be set in case of packed audio
 * @param out_count amount of space available for output in samples per channel
 * @param in        input buffers, only the first one need to be set in case of packed audio
 * @param in_count  number of input samples available in one channel
 *
 * @return number of samples output per channel, negative value on error
 */
//int swr_convert(struct SwrContext *s, uint8_t **out, int out_count,
//const uint8_t **in , int in_count);
var swrConvert func(s *SwrContext, out **ffcommon.FUint8T, outCount ffcommon.FInt, in **ffcommon.FUint8T, inCount ffcommon.FInt) ffcommon.FInt
var swrConvertOnce sync.Once

// SwrConvert is a purego method to convert audio samples using SwrContext.
func (s *SwrContext) SwrConvert(out **ffcommon.FUint8T, outCount ffcommon.FInt, in **ffcommon.FUint8T, inCount ffcommon.FInt) ffcommon.FInt {
	swrConvertOnce.Do(func() {
		purego.RegisterLibFunc(&swrConvert, ffcommon.GetAvswresampleDll(), "swr_convert")
	})
	return swrConvert(s, out, outCount, in, inCount)
}

/**
 * Convert the next timestamp from input to output
 * timestamps are in 1/(in_sample_rate * out_sample_rate) units.
 *
 * @note There are 2 slightly differently behaving modes.
 *       @li When automatic timestamp compensation is not used, (min_compensation >= FLT_MAX)
 *              in this case timestamps will be passed through with delays compensated
 *       @li When automatic timestamp compensation is used, (min_compensation < FLT_MAX)
 *              in this case the output timestamps will match output sample numbers.
 *              See ffmpeg-resampler(1) for the two modes of compensation.
 *
 * @param s[in]     initialized Swr context
 * @param pts[in]   timestamp for the next input sample, INT64_MIN if unknown
 * @see swr_set_compensation(), swr_drop_output(), and swr_inject_silence() are
 *      function used internally for timestamp compensation.
 * @return the output timestamp for the next output sample
 */
//int64_t swr_next_pts(struct SwrContext *s, int64_t pts);
var swrNextPts func(s *SwrContext, pts ffcommon.FInt64T) ffcommon.FInt64T
var swrNextPtsOnce sync.Once

// SwrNextPts is a purego method to get the next PTS (presentation timestamp) from SwrContext.
func (s *SwrContext) SwrNextPts(pts ffcommon.FInt64T) ffcommon.FInt64T {
	swrNextPtsOnce.Do(func() {
		purego.RegisterLibFunc(&swrNextPts, ffcommon.GetAvswresampleDll(), "swr_next_pts")
	})
	return swrNextPts(s, pts)
}

/**
 * @}
 *
 * @name Low-level option setting functions
 * These functons provide a means to set low-level options that is not possible
 * with the AVOption API.
 * @{
 */

/**
 * Activate resampling compensation ("soft" compensation). This function is
 * internally called when needed in swr_next_pts().
 *
 * @param[in,out] s             allocated Swr context. If it is not initialized,
 *                              or SWR_FLAG_RESAMPLE is not set, swr_init() is
 *                              called with the flag set.
 * @param[in]     sample_delta  delta in PTS per sample
 * @param[in]     compensation_distance number of samples to compensate for
 * @return    >= 0 on success, AVERROR error codes if:
 *            @li @c s is NULL,
 *            @li @c compensation_distance is less than 0,
 *            @li @c compensation_distance is 0 but sample_delta is not,
 *            @li compensation unsupported by resampler, or
 *            @li swr_init() fails when called.
 */
//int swr_set_compensation(struct SwrContext *s, int sample_delta, int compensation_distance);
var swrSetCompensation func(s *SwrContext, sampleDelta, compensationDistance ffcommon.FInt) ffcommon.FInt
var swrSetCompensationOnce sync.Once

// SwrSetCompensation is a purego method to set compensation parameters for SwrContext.
func (s *SwrContext) SwrSetCompensation(sampleDelta, compensationDistance ffcommon.FInt) ffcommon.FInt {
	swrSetCompensationOnce.Do(func() {
		purego.RegisterLibFunc(&swrSetCompensation, ffcommon.GetAvswresampleDll(), "swr_set_compensation")
	})
	return swrSetCompensation(s, sampleDelta, compensationDistance)
}

/**
 * Set a customized input channel mapping.
 *
 * @param[in,out] s           allocated Swr context, not yet initialized
 * @param[in]     channel_map customized input channel mapping (array of channel
 *                            indexes, -1 for a muted channel)
 * @return >= 0 on success, or AVERROR error code in case of failure.
 */
//int swr_set_channel_mapping(struct SwrContext *s, const int *channel_map);
var swrSetChannelMapping func(s *SwrContext, channelMap *ffcommon.FInt) ffcommon.FInt
var swrSetChannelMappingOnce sync.Once

// SwrSetChannelMapping is a purego method to set channel mapping for SwrContext.
func (s *SwrContext) SwrSetChannelMapping(channelMap *ffcommon.FInt) ffcommon.FInt {
	swrSetChannelMappingOnce.Do(func() {
		purego.RegisterLibFunc(&swrSetChannelMapping, ffcommon.GetAvswresampleDll(), "swr_set_channel_mapping")
	})
	return swrSetChannelMapping(s, channelMap)
}

/**
 * Generate a channel mixing matrix.
 *
 * This function is the one used internally by libswresample for building the
 * default mixing matrix. It is made public just as a utility function for
 * building custom matrices.
 *
 * @param in_layout           input channel layout
 * @param out_layout          output channel layout
 * @param center_mix_level    mix level for the center channel
 * @param surround_mix_level  mix level for the surround channel(s)
 * @param lfe_mix_level       mix level for the low-frequency effects channel
 * @param rematrix_maxval     if 1.0, coefficients will be normalized to prevent
 *                            overflow. if INT_MAX, coefficients will not be
 *                            normalized.
 * @param[out] matrix         mixing coefficients; matrix[i + stride * o] is
 *                            the weight of input channel i in output channel o.
 * @param stride              distance between adjacent input channels in the
 *                            matrix array
 * @param matrix_encoding     matrixed stereo downmix mode (e.g. dplii)
 * @param log_ctx             parent logging context, can be NULL
 * @return                    0 on success, negative AVERROR code on failure
 */
//int swr_build_matrix(uint64_t in_layout, uint64_t out_layout,
//double center_mix_level, double surround_mix_level,
//double lfe_mix_level, double rematrix_maxval,
//double rematrix_volume, double *matrix,
//int stride, enum AVMatrixEncoding matrix_encoding,
//void *log_ctx);
type AVMatrixEncoding = libavutil.AVMatrixEncoding

var swrBuildMatrix func(inLayout, outLayout ffcommon.FUint64T, centerMixLevel, surroundMixLevel, lfeMixLevel, rematrixMaxval, rematrixVolume *ffcommon.FDouble, matrix *ffcommon.FDouble, stride ffcommon.FInt, matrixEncoding AVMatrixEncoding, logCtx ffcommon.FVoidP) ffcommon.FInt
var swrBuildMatrixOnce sync.Once

// SwrBuildMatrix is a purego function to build a rematrixing matrix.
func SwrBuildMatrix(inLayout, outLayout ffcommon.FUint64T, centerMixLevel, surroundMixLevel, lfeMixLevel, rematrixMaxval, rematrixVolume *ffcommon.FDouble, matrix *ffcommon.FDouble, stride ffcommon.FInt, matrixEncoding AVMatrixEncoding, logCtx ffcommon.FVoidP) ffcommon.FInt {
	swrBuildMatrixOnce.Do(func() {
		purego.RegisterLibFunc(&swrBuildMatrix, ffcommon.GetAvswresampleDll(), "swr_build_matrix")
	})
	return swrBuildMatrix(inLayout, outLayout, centerMixLevel, surroundMixLevel, lfeMixLevel, rematrixMaxval, rematrixVolume, matrix, stride, matrixEncoding, logCtx)
}

/**
 * Set a customized remix matrix.
 *
 * @param s       allocated Swr context, not yet initialized
 * @param matrix  remix coefficients; matrix[i + stride * o] is
 *                the weight of input channel i in output channel o
 * @param stride  offset between lines of the matrix
 * @return  >= 0 on success, or AVERROR error code in case of failure.
 */
//int swr_set_matrix(struct SwrContext *s, const double *matrix, int stride);
var swrSetMatrix func(s *SwrContext, matrix *ffcommon.FDouble, stride ffcommon.FInt) ffcommon.FInt
var swrSetMatrixOnce sync.Once

// SwrSetMatrix is a purego method to set a rematrixing matrix for SwrContext.
func (s *SwrContext) SwrSetMatrix(matrix *ffcommon.FDouble, stride ffcommon.FInt) ffcommon.FInt {
	swrSetMatrixOnce.Do(func() {
		purego.RegisterLibFunc(&swrSetMatrix, ffcommon.GetAvswresampleDll(), "swr_set_matrix")
	})
	return swrSetMatrix(s, matrix, stride)
}

/**
 * @}
 *
 * @name Sample handling functions
 * @{
 */

/**
 * Drops the specified number of output samples.
 *
 * This function, along with swr_inject_silence(), is called by swr_next_pts()
 * if needed for "hard" compensation.
 *
 * @param s     allocated Swr context
 * @param count number of samples to be dropped
 *
 * @return >= 0 on success, or a negative AVERROR code on failure
 */
//int swr_drop_output(struct SwrContext *s, int count);
var swrDropOutput func(s *SwrContext, count ffcommon.FInt) ffcommon.FInt
var swrDropOutputOnce sync.Once

// SwrDropOutput is a purego method to drop audio output samples from SwrContext.
func (s *SwrContext) SwrDropOutput(count ffcommon.FInt) ffcommon.FInt {
	swrDropOutputOnce.Do(func() {
		purego.RegisterLibFunc(&swrDropOutput, ffcommon.GetAvswresampleDll(), "swr_drop_output")
	})
	return swrDropOutput(s, count)
}

/**
 * Injects the specified number of silence samples.
 *
 * This function, along with swr_drop_output(), is called by swr_next_pts()
 * if needed for "hard" compensation.
 *
 * @param s     allocated Swr context
 * @param count number of samples to be dropped
 *
 * @return >= 0 on success, or a negative AVERROR code on failure
 */
//int swr_inject_silence(struct SwrContext *s, int count);
var swrInjectSilence func(s *SwrContext, count ffcommon.FInt) ffcommon.FInt
var swrInjectSilenceOnce sync.Once

// SwrInjectSilence is a purego method to inject silence samples into SwrContext.
func (s *SwrContext) SwrInjectSilence(count ffcommon.FInt) ffcommon.FInt {
	swrInjectSilenceOnce.Do(func() {
		purego.RegisterLibFunc(&swrInjectSilence, ffcommon.GetAvswresampleDll(), "swr_inject_silence")
	})
	return swrInjectSilence(s, count)
}

/**
 * Gets the delay the next input sample will experience relative to the next output sample.
 *
 * Swresample can buffer data if more input has been provided than available
 * output space, also converting between sample rates needs a delay.
 * This function returns the sum of all such delays.
 * The exact delay is not necessarily an integer value in either input or
 * output sample rate. Especially when downsampling by a large value, the
 * output sample rate may be a poor choice to represent the delay, similarly
 * for upsampling and the input sample rate.
 *
 * @param s     swr context
 * @param base  timebase in which the returned delay will be:
 *              @li if it's set to 1 the returned delay is in seconds
 *              @li if it's set to 1000 the returned delay is in milliseconds
 *              @li if it's set to the input sample rate then the returned
 *                  delay is in input samples
 *              @li if it's set to the output sample rate then the returned
 *                  delay is in output samples
 *              @li if it's the least common multiple of in_sample_rate and
 *                  out_sample_rate then an exact rounding-free delay will be
 *                  returned
 * @returns     the delay in 1 / @c base units.
 */
//int64_t swr_get_delay(struct SwrContext *s, int64_t base);
var swrGetDelay func(s *SwrContext, base ffcommon.FInt64T) ffcommon.FInt64T
var swrGetDelayOnce sync.Once

// SwrGetDelay is a purego method to get the delay in samples for SwrContext.
func (s *SwrContext) SwrGetDelay(base ffcommon.FInt64T) ffcommon.FInt64T {
	swrGetDelayOnce.Do(func() {
		purego.RegisterLibFunc(&swrGetDelay, ffcommon.GetAvswresampleDll(), "swr_get_delay")
	})
	return swrGetDelay(s, base)
}

/**
 * Find an upper bound on the number of samples that the next swr_convert
 * call will output, if called with in_samples of input samples. This
 * depends on the internal state, and anything changing the internal state
 * (like further swr_convert() calls) will may change the number of samples
 * swr_get_out_samples() returns for the same number of input samples.
 *
 * @param in_samples    number of input samples.
 * @note any call to swr_inject_silence(), swr_convert(), swr_next_pts()
 *       or swr_set_compensation() invalidates this limit
 * @note it is recommended to pass the correct available buffer size
 *       to all functions like swr_convert() even if swr_get_out_samples()
 *       indicates that less would be used.
 * @returns an upper bound on the number of samples that the next swr_convert
 *          will output or a negative value to indicate an error
 */
//int swr_get_out_samples(struct SwrContext *s, int in_samples);
var swrGetOutSamples func(s *SwrContext, inSamples ffcommon.FInt) ffcommon.FInt
var swrGetOutSamplesOnce sync.Once

// SwrGetOutSamples is a purego method to get the number of output samples for a given input sample count in SwrContext.
func (s *SwrContext) SwrGetOutSamples(inSamples ffcommon.FInt) ffcommon.FInt {
	swrGetOutSamplesOnce.Do(func() {
		purego.RegisterLibFunc(&swrGetOutSamples, ffcommon.GetAvswresampleDll(), "swr_get_out_samples")
	})
	return swrGetOutSamples(s, inSamples)
}

/**
 * @}
 *
 * @name Configuration accessors
 * @{
 */

/**
 * Return the @ref LIBSWRESAMPLE_VERSION_INT constant.
 *
 * This is useful to check if the build-time libswresample has the same version
 * as the run-time one.
 *
 * @returns     the unsigned int-typed version
 */
//unsigned swresample_version(void);
var swresampleVersion func() ffcommon.FUnsigned
var swresampleVersionOnce sync.Once

// SwresampleVersion is a purego function to get the swresample library version.
func SwresampleVersion() ffcommon.FUnsigned {
	swresampleVersionOnce.Do(func() {
		purego.RegisterLibFunc(&swresampleVersion, ffcommon.GetAvswresampleDll(), "swresample_version")
	})
	return swresampleVersion()
}

/**
 * Return the swr build-time configuration.
 *
 * @returns     the build-time @c ./configure flags
 */
//const char *swresample_configuration(void);
var swresampleConfiguration func() ffcommon.FConstCharP
var swresampleConfigurationOnce sync.Once

// SwresampleConfiguration is a purego function to get the swresample library configuration.
func SwresampleConfiguration() ffcommon.FConstCharP {
	swresampleConfigurationOnce.Do(func() {
		purego.RegisterLibFunc(&swresampleConfiguration, ffcommon.GetAvswresampleDll(), "swresample_configuration")
	})
	return swresampleConfiguration()
}

/**
 * Return the swr license.
 *
 * @returns     the license of libswresample, determined at build-time
 */
//const char *swresample_license(void);
var swresampleLicense func() ffcommon.FConstCharP
var swresampleLicenseOnce sync.Once

// SwresampleLicense is a purego function to get the swresample library license information.
func SwresampleLicense() ffcommon.FConstCharP {
	swresampleLicenseOnce.Do(func() {
		purego.RegisterLibFunc(&swresampleLicense, ffcommon.GetAvswresampleDll(), "swresample_license")
	})
	return swresampleLicense()
}

/**
 * @}
 *
 * @name AVFrame based API
 * @{
 */

/**
 * Convert the samples in the input AVFrame and write them to the output AVFrame.
 *
 * Input and output AVFrames must have channel_layout, sample_rate and format set.
 *
 * If the output AVFrame does not have the data pointers allocated the nb_samples
 * field will be set using av_frame_get_buffer()
 * is called to allocate the frame.
 *
 * The output AVFrame can be NULL or have fewer allocated samples than required.
 * In this case, any remaining samples not written to the output will be added
 * to an internal FIFO buffer, to be returned at the next call to this function
 * or to swr_convert().
 *
 * If converting sample rate, there may be data remaining in the internal
 * resampling delay buffer. swr_get_delay() tells the number of
 * remaining samples. To get this data as output, call this function or
 * swr_convert() with NULL input.
 *
 * If the SwrContext configuration does not match the output and
 * input AVFrame settings the conversion does not take place and depending on
 * which AVFrame is not matching AVERROR_OUTPUT_CHANGED, AVERROR_INPUT_CHANGED
 * or the result of a bitwise-OR of them is returned.
 *
 * @see swr_delay()
 * @see swr_convert()
 * @see swr_get_delay()
 *
 * @param swr             audio resample context
 * @param output          output AVFrame
 * @param input           input AVFrame
 * @return                0 on success, AVERROR on failure or nonmatching
 *                        configuration.
 */
//int swr_convert_frame(SwrContext *swr,
//AVFrame *output, const AVFrame *input);
type AVFrame = libavutil.AVFrame

var swrConvertFrame func(swr *SwrContext, output, input *AVFrame) ffcommon.FInt
var swrConvertFrameOnce sync.Once

// SwrConvertFrame is a purego method to convert audio samples from one frame to another using SwrContext.
func (swr *SwrContext) SwrConvertFrame(output, input *AVFrame) ffcommon.FInt {
	swrConvertFrameOnce.Do(func() {
		purego.RegisterLibFunc(&swrConvertFrame, ffcommon.GetAvswresampleDll(), "swr_convert_frame")
	})
	return swrConvertFrame(swr, output, input)
}

/**
 * Configure or reconfigure the SwrContext using the information
 * provided by the AVFrames.
 *
 * The original resampling context is reset even on failure.
 * The function calls swr_close() internally if the context is open.
 *
 * @see swr_close();
 *
 * @param swr             audio resample context
 * @param output          output AVFrame
 * @param input           input AVFrame
 * @return                0 on success, AVERROR on failure.
 */
//int swr_config_frame(SwrContext *swr, const AVFrame *out, const AVFrame *in);
var swrConfigFrame func(swr *SwrContext, out, in *AVFrame) ffcommon.FInt
var swrConfigFrameOnce sync.Once

// SwrConfigFrame is a purego method to configure the SwrContext for the given input and output frames.
func (swr *SwrContext) SwrConfigFrame(out, in *AVFrame) ffcommon.FInt {
	swrConfigFrameOnce.Do(func() {
		purego.RegisterLibFunc(&swrConfigFrame, ffcommon.GetAvswresampleDll(), "swr_config_frame")
	})
	return swrConfigFrame(swr, out, in)
}

/**
 * @}
 * @}
 */

//#endif /* SWRESAMPLE_SWRESAMPLE_H */
