package libswscale

import (
	"sync"
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/dwdcth/ffmpeg-go/libavutil"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (C) 2001-2011 Michael Niedermayer <michaelni@gmx.at>
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

//#ifndef SWSCALE_SWSCALE_H
//const SWSCALE_SWSCALE_H
//
///**
// * @file
// * @ingroup libsws
// * external API header
// */
//
//#include <stdint.h>
//
//#include "../libavutil/avutil.h"
//#include "../libavutil/log.h"
//#include "../libavutil/pixfmt.h"
//#include "version.h"

/**
 * @defgroup libsws libswscale
 * Color conversion and scaling library.
 *
 * @{
 *
 * Return the LIBSWSCALE_VERSION_INT constant.
 */
//unsigned swscale_version(void);
//todo
var swscale_version func() ffcommon.FCharP
var swscale_versionOnce sync.Once

func SwscaleVersion() ffcommon.FCharP {
	swscale_versionOnce.Do(func() {
		purego.RegisterLibFunc(&swscale_version, ffcommon.GetAvswscaleDll(), "swscale_version")
	})

	return swscale_version()
}

/**
 * Return the libswscale build-time configuration.
 */
//const char *swscale_configuration(void);
var swscaleConfiguration func() ffcommon.FConstCharP
var swscaleConfigurationOnce sync.Once

func SwscaleConfiguration() ffcommon.FConstCharP {
	swscaleConfigurationOnce.Do(func() {
		purego.RegisterLibFunc(&swscaleConfiguration, ffcommon.GetAvswscaleDll(), "swscale_configuration")
	})

	return swscaleConfiguration()
}

/**
 * Return the libswscale license.
 */
//const char *swscale_license(void);
var swscaleLicense func() ffcommon.FConstCharP
var swscaleLicenseOnce sync.Once

func SwscaleLicense() ffcommon.FConstCharP {
	swscaleLicenseOnce.Do(func() {
		purego.RegisterLibFunc(&swscaleLicense, ffcommon.GetAvswscaleDll(), "swscale_license")
	})

	return swscaleLicense()
}

/* values for the flags, the stuff on the command line is different */
const SWS_FAST_BILINEAR = 1
const SWS_BILINEAR = 2
const SWS_BICUBIC = 4
const SWS_X = 8
const SWS_POINT = 0x10
const SWS_AREA = 0x20
const SWS_BICUBLIN = 0x40
const SWS_GAUSS = 0x80
const SWS_SINC = 0x100
const SWS_LANCZOS = 0x200
const SWS_SPLINE = 0x400

const SWS_SRC_V_CHR_DROP_MASK = 0x30000
const SWS_SRC_V_CHR_DROP_SHIFT = 16

const SWS_PARAM_DEFAULT = 123456

const SWS_PRINT_INFO = 0x1000

// the following 3 flags are not completely implemented
// internal chrominance subsampling info
const SWS_FULL_CHR_H_INT = 0x2000

// input subsampling info
const SWS_FULL_CHR_H_INP = 0x4000
const SWS_DIRECT_BGR = 0x8000
const SWS_ACCURATE_RND = 0x40000
const SWS_BITEXACT = 0x80000
const SWS_ERROR_DIFFUSION = 0x800000

const SWS_MAX_REDUCE_CUTOFF = 0.002

const SWS_CS_ITU709 = 1
const SWS_CS_FCC = 4
const SWS_CS_ITU601 = 5
const SWS_CS_ITU624 = 5
const SWS_CS_SMPTE170M = 5
const SWS_CS_SMPTE240M = 7
const SWS_CS_DEFAULT = 5
const SWS_CS_BT2020 = 9

/**
 * Return a pointer to yuv<->rgb coefficients for the given colorspace
 * suitable for sws_setColorspaceDetails().
 *
 * @param colorspace One of the SWS_CS_* macros. If invalid,
 * SWS_CS_DEFAULT is used.
 */
//const int *sws_getCoefficients(int colorspace);
var swsGetCoefficients func(colorspace ffcommon.FInt) *ffcommon.FInt
var swsGetCoefficientsOnce sync.Once

func SwsGetCoefficients(colorspace ffcommon.FInt) *ffcommon.FInt {
	swsGetCoefficientsOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetCoefficients, ffcommon.GetAvswscaleDll(), "sws_getCoefficients")
	})

	return swsGetCoefficients(colorspace)

	// return (*ffcommon.FInt)(unsafe.Pointer(t))
}

// when used for filters they must have an odd number of elements
// coeffs cannot be shared between vectors
type SwsVector struct {
	Coeff  *ffcommon.FDouble ///< pointer to the list of coefficients
	Length ffcommon.FInt     ///< number of coefficients in the vector
}

// vectors can be shared
type SwsFilter struct {
	LumH *SwsVector
	LumV *SwsVector
	ChrH *SwsVector
	ChrV *SwsVector
}

// struct SwsContext;
type SwsContext struct {
}

/**
 * Return a positive value if pix_fmt is a supported input format, 0
 * otherwise.
 */
//int sws_isSupportedInput(enum AVPixelFormat pix_fmt);
type AVPixelFormat = libavutil.AVPixelFormat

var swsIsSupportedInput func(pix_fmt AVPixelFormat) ffcommon.FInt
var swsIsSupportedInputOnce sync.Once

func SwsIsSupportedInput(pix_fmt AVPixelFormat) ffcommon.FInt {
	swsIsSupportedInputOnce.Do(func() {
		purego.RegisterLibFunc(&swsIsSupportedInput, ffcommon.GetAvswscaleDll(), "sws_isSupportedInput")
	})

	t := uintptr(swsIsSupportedInput(pix_fmt))
	if t == 0 {
		return ffcommon.FInt(0)
	}

	return ffcommon.FInt(t)
}

/**
 * Return a positive value if pix_fmt is a supported output format, 0
 * otherwise.
 */
//int sws_isSupportedOutput(enum AVPixelFormat pix_fmt);
var swsIsSupportedOutput func(pix_fmt AVPixelFormat) ffcommon.FInt
var swsIsSupportedOutputOnce sync.Once

func SwsIsSupportedOutput(pix_fmt AVPixelFormat) ffcommon.FInt {
	swsIsSupportedOutputOnce.Do(func() {
		purego.RegisterLibFunc(&swsIsSupportedOutput, ffcommon.GetAvswscaleDll(), "sws_isSupportedOutput")
	})

	t := uintptr(swsIsSupportedOutput(pix_fmt))
	if t == 0 {
		return ffcommon.FInt(0)
	}

	return ffcommon.FInt(t)
}

/**
 * @param[in]  pix_fmt the pixel format
 * @return a positive value if an endianness conversion for pix_fmt is
 * supported, 0 otherwise.
 */
//int sws_isSupportedEndiannessConversion(enum AVPixelFormat pix_fmt);
var swsIsSupportedEndiannessConversion func(pix_fmt AVPixelFormat) ffcommon.FInt
var swsIsSupportedEndiannessConversionOnce sync.Once

func SwsIsSupportedEndiannessConversion(pix_fmt AVPixelFormat) ffcommon.FInt {
	swsIsSupportedEndiannessConversionOnce.Do(func() {
		purego.RegisterLibFunc(&swsIsSupportedEndiannessConversion, ffcommon.GetAvswscaleDll(), "sws_isSupportedEndiannessConversion")
	})

	t := uintptr(swsIsSupportedEndiannessConversion(pix_fmt))
	if t == 0 {
		return ffcommon.FInt(0)
	}

	return ffcommon.FInt(t)
}

/**
 * Allocate an empty SwsContext. This must be filled and passed to
 * sws_init_context(). For filling see AVOptions, options.c and
 * sws_setColorspaceDetails().
 */
//struct SwsContext *sws_alloc_context(void);
var swsAllocContext func() *SwsContext
var swsAllocContextOnce sync.Once

func SwsAllocContext() *SwsContext {
	swsAllocContextOnce.Do(func() {
		purego.RegisterLibFunc(&swsAllocContext, ffcommon.GetAvswscaleDll(), "sws_alloc_context")
	})

	return swsAllocContext()

}

/**
 * Initialize the swscaler context sws_context.
 *
 * @return zero or positive value on success, a negative value on
 * error
 */
//av_warn_unused_result
//int sws_init_context(struct SwsContext *sws_context, SwsFilter *srcFilter, SwsFilter *dstFilter);
var swsInitContext func(sws_context *SwsContext, srcFilter, dstFilter *SwsFilter) ffcommon.FInt
var swsInitContextOnce sync.Once

func (sws_context *SwsContext) SwsInitContext(srcFilter, dstFilter *SwsFilter) ffcommon.FInt {
	swsInitContextOnce.Do(func() {
		purego.RegisterLibFunc(&swsInitContext, ffcommon.GetAvswscaleDll(), "sws_init_context")
	})

	t := uintptr(swsInitContext(sws_context, srcFilter, dstFilter))
	if t == 0 {
		return ffcommon.FInt(0)
	}

	return ffcommon.FInt(t)
}

/**
 * Free the swscaler context swsContext.
 * If swsContext is NULL, then does nothing.
 */
//void sws_freeContext(struct SwsContext *swsContext);
var swsFreeContext func(swsContext *SwsContext)
var swsFreeContextOnce sync.Once

func (swsContext *SwsContext) SwsFreeContext() {
	swsFreeContextOnce.Do(func() {
		purego.RegisterLibFunc(&swsFreeContext, ffcommon.GetAvswscaleDll(), "sws_freeContext")
	})

	swsFreeContext(swsContext)
}

/**
 * Allocate and return an SwsContext. You need it to perform
 * scaling/conversion operations using sws_scale().
 *
 * @param srcW the width of the source image
 * @param srcH the height of the source image
 * @param srcFormat the source image format
 * @param dstW the width of the destination image
 * @param dstH the height of the destination image
 * @param dstFormat the destination image format
 * @param flags specify which algorithm and options to use for rescaling
 * @param param extra parameters to tune the used scaler
 *              For SWS_BICUBIC param[0] and [1] tune the shape of the basis
 *              function, param[0] tunes f(1) and param[1] f´(1)
 *              For SWS_GAUSS param[0] tunes the exponent and thus cutoff
 *              frequency
 *              For SWS_LANCZOS param[0] tunes the width of the window function
 * @return a pointer to an allocated context, or NULL in case of error
 * @note this function is to be removed after a saner alternative is
 *       written
 */
//struct SwsContext *sws_getContext(int srcW, int srcH, enum AVPixelFormat srcFormat,
//int dstW, int dstH, enum AVPixelFormat dstFormat,
//int flags, SwsFilter *srcFilter,
//SwsFilter *dstFilter, const double *param);
//var swsGetContext func(srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat,
//	dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat,
//	flags ffcommon.FInt, srcFilter, dstFilter *SwsFilter, param *ffcommon.FDouble) *SwsContext
var swsGetContext func(srcW, srcH, srcFormat,
	dstW, dstH, dstFormat,
	flags, srcFilter, dstFilter, param uintptr) *SwsContext
var swsGetContextOnce sync.Once

func SwsGetContext(srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat,
	dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat,
	flags ffcommon.FInt, srcFilter, dstFilter *SwsFilter, param *ffcommon.FDouble) *SwsContext {
	swsGetContextOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetContext, ffcommon.GetAvswscaleDll(), "sws_getContext")
	})

	return swsGetContext(uintptr(srcW), uintptr(srcH), uintptr(srcFormat),
		uintptr(dstW), uintptr(dstH), uintptr(dstFormat),
		uintptr(flags),
		uintptr(unsafe.Pointer(srcFilter)),
		uintptr(unsafe.Pointer(dstFilter)),
		uintptr(unsafe.Pointer(param)))
	//return swsGetContext(srcW, srcH, srcFormat, dstW, dstH, dstFormat, flags, srcFilter, dstFilter, param)

}

/**
 * Scale the image slice in srcSlice and put the resulting scaled
 * slice in the image in dst. A slice is a sequence of consecutive
 * rows in an image.
 *
 * Slices have to be provided in sequential order, either in
 * top-bottom or bottom-top order. If slices are provided in
 * non-sequential order the behavior of the function is undefined.
 *
 * @param c         the scaling context previously created with
 *                  sws_getContext()
 * @param srcSlice  the array containing the pointers to the planes of
 *                  the source slice
 * @param srcStride the array containing the strides for each plane of
 *                  the source image
 * @param srcSliceY the position in the source image of the slice to
 *                  process, that is the number (counted starting from
 *                  zero) in the image of the first row of the slice
 * @param srcSliceH the height of the source slice, that is the number
 *                  of rows in the slice
 * @param dst       the array containing the pointers to the planes of
 *                  the destination image
 * @param dstStride the array containing the strides for each plane of
 *                  the destination image
 * @return          the height of the output slice
 */
//int sws_scale(struct SwsContext *c, const uint8_t *const srcSlice[],
//const int srcStride[], int srcSliceY, int srcSliceH,
//uint8_t *const dst[], const int dstStride[]);
var swsScale func(c *SwsContext, srcSlice **ffcommon.FUint8T, srcStride *ffcommon.FInt, srcSliceY, srcSliceH ffcommon.FUint, dst **ffcommon.FUint8T, dstStride *ffcommon.FInt) ffcommon.FInt
var swsScaleOnce sync.Once

func (c *SwsContext) SwsScale(srcSlice **ffcommon.FUint8T, srcStride *ffcommon.FInt, srcSliceY, srcSliceH ffcommon.FUint, dst **ffcommon.FUint8T, dstStride *ffcommon.FInt) ffcommon.FInt {
	swsScaleOnce.Do(func() {
		purego.RegisterLibFunc(&swsScale, ffcommon.GetAvswscaleDll(), "sws_scale")
	})

	t := uintptr(swsScale(c, srcSlice, srcStride, srcSliceY, srcSliceH, dst, dstStride))
	return ffcommon.FInt(t)
}

/**
 * @param dstRange flag indicating the while-black range of the output (1=jpeg / 0=mpeg)
 * @param srcRange flag indicating the while-black range of the input (1=jpeg / 0=mpeg)
 * @param table the yuv2rgb coefficients describing the output yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param inv_table the yuv2rgb coefficients describing the input yuv space, normally ff_yuv2rgb_coeffs[x]
 * @param brightness 16.16 fixed point brightness correction
 * @param contrast 16.16 fixed point contrast correction
 * @param saturation 16.16 fixed point saturation correction
 * @return -1 if not supported
 */
//int sws_setColorspaceDetails(struct SwsContext *c, const int inv_table[4],
//int srcRange, const int table[4], int dstRange,
//int brightness, int contrast, int saturation);
var swsSetColorspaceDetails func(c *SwsContext, inv_table [4]*ffcommon.FInt, srcRange ffcommon.FInt, table [4]ffcommon.FInt, dstRange, brightness, contrast, saturation ffcommon.FInt) ffcommon.FInt
var swsSetColorspaceDetailsOnce sync.Once

func (c *SwsContext) SwsSetColorspaceDetails(inv_table [4]*ffcommon.FInt, srcRange ffcommon.FInt, table [4]ffcommon.FInt, dstRange, brightness, contrast, saturation ffcommon.FInt) ffcommon.FInt {
	swsSetColorspaceDetailsOnce.Do(func() {
		purego.RegisterLibFunc(&swsSetColorspaceDetails, ffcommon.GetAvswscaleDll(), "sws_setColorspaceDetails")
	})

	t := uintptr(swsSetColorspaceDetails(c, inv_table, srcRange, table, dstRange, brightness, contrast, saturation))
	return ffcommon.FInt(t)
}

/**
 * @return -1 if not supported
 */
//int sws_getColorspaceDetails(struct SwsContext *c, int **inv_table,
//int *srcRange, int **table, int *dstRange,
//int *brightness, int *contrast, int *saturation);
var swsGetColorspaceDetails func(c *SwsContext, inv_table **ffcommon.FInt, srcRange *ffcommon.FInt, table **ffcommon.FInt, dstRange, brightness, contrast, saturation *ffcommon.FInt) ffcommon.FInt
var swsGetColorspaceDetailsOnce sync.Once

func (c *SwsContext) SwsGetColorspaceDetails(inv_table **ffcommon.FInt, srcRange *ffcommon.FInt, table **ffcommon.FInt, dstRange, brightness, contrast, saturation *ffcommon.FInt) ffcommon.FInt {
	swsGetColorspaceDetailsOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetColorspaceDetails, ffcommon.GetAvswscaleDll(), "sws_getColorspaceDetails")
	})

	t := uintptr(swsGetColorspaceDetails(c, inv_table, srcRange, table, dstRange, brightness, contrast, saturation))
	return ffcommon.FInt(t)
}

/**
 * Allocate and return an uninitialized vector with length coefficients.
 */
//SwsVector *sws_allocVec(int length);
var swsAllocVec func(length ffcommon.FInt) *SwsVector
var swsAllocVecOnce sync.Once

func SwsAllocVec(length ffcommon.FInt) *SwsVector {
	swsAllocVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsAllocVec, ffcommon.GetAvswscaleDll(), "sws_allocVec")
	})

	return swsAllocVec(length)

}

/**
 * Return a normalized Gaussian curve used to filter stuff
 * quality = 3 is high quality, lower is lower quality.
 */
//SwsVector *sws_getGaussianVec(double variance, double quality);
var swsGetGaussianVec func(variance, quality *ffcommon.FDouble) *SwsVector
var swsGetGaussianVecOnce sync.Once

func SwsGetGaussianVec(variance, quality *ffcommon.FDouble) *SwsVector {
	swsGetGaussianVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetGaussianVec, ffcommon.GetAvswscaleDll(), "sws_getGaussianVec")
	})

	return swsGetGaussianVec(variance, quality)

}

/**
 * Scale all the coefficients of a by the scalar value.
 */
//void sws_scaleVec(SwsVector *a, double scalar);
var swsScaleVec func(a *SwsVector, scalar *ffcommon.FDouble)
var swsScaleVecOnce sync.Once

func (a *SwsVector) SwsScaleVec(scalar *ffcommon.FDouble) {
	swsScaleVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsScaleVec, ffcommon.GetAvswscaleDll(), "sws_scaleVec")
	})

	swsScaleVec(a, scalar)
}

/**
 * Scale all the coefficients of a so that their sum equals height.
 */
//void sws_normalizeVec(SwsVector *a, double height);
var swsNormalizeVec func(a *SwsVector, height *ffcommon.FDouble)
var swsNormalizeVecOnce sync.Once

func (a *SwsVector) SwsNormalizeVec(height *ffcommon.FDouble) {
	swsNormalizeVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsNormalizeVec, ffcommon.GetAvswscaleDll(), "sws_normalizeVec")
	})

	swsNormalizeVec(a, height)
}

// #if FF_API_SWS_VECTOR
// attribute_deprecated SwsVector *sws_getConstVec(double c, int length);
var swsGetConstVec func(c *ffcommon.FDouble, length ffcommon.FInt) *SwsVector
var swsGetConstVecOnce sync.Once

func SwsGetConstVec(c *ffcommon.FDouble, length ffcommon.FInt) *SwsVector {
	swsGetConstVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetConstVec, ffcommon.GetAvswscaleDll(), "sws_getConstVec")
	})

	return swsGetConstVec(c, length)
	// if t == 0 {
	// 	return nil
	// }

	// return (*SwsVector)(unsafe.Pointer(t))
}

// attribute_deprecated SwsVector *sws_getIdentityVec(void);
var swsGetIdentityVec func() *SwsVector
var swsGetIdentityVecOnce sync.Once

func SwsGetIdentityVec() *SwsVector {
	swsGetIdentityVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetIdentityVec, ffcommon.GetAvswscaleDll(), "sws_getIdentityVec")
	})

	return swsGetIdentityVec()
	// if t == 0 {
	// 	return nil
	// }

	// return (*SwsVector)(unsafe.Pointer(t))
}

// attribute_deprecated void sws_convVec(SwsVector *a, SwsVector *b);
var swsConvVec func(a, b *SwsVector)
var swsConvVecOnce sync.Once

func (a *SwsVector) SwsConvVec(b *SwsVector) {
	swsConvVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsConvVec, ffcommon.GetAvswscaleDll(), "sws_convVec")
	})

	swsConvVec(a, b)
}

// attribute_deprecated void sws_addVec(SwsVector *a, SwsVector *b);
var swsAddVec func(a, b *SwsVector)
var swsAddVecOnce sync.Once

func (a *SwsVector) SwsAddVec(b *SwsVector) {
	swsAddVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsAddVec, ffcommon.GetAvswscaleDll(), "sws_addVec")
	})

	swsAddVec(a, b)
}

// attribute_deprecated void sws_subVec(SwsVector *a, SwsVector *b);
var swsSubVec func(a, b *SwsVector)
var swsSubVecOnce sync.Once

func (a *SwsVector) SwsSubVec(b *SwsVector) {
	swsSubVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsSubVec, ffcommon.GetAvswscaleDll(), "sws_subVec")
	})

	swsSubVec(a, b)
}

// attribute_deprecated void sws_shiftVec(SwsVector *a, int shift);
var swsShiftVec func(a *SwsVector, shift ffcommon.FInt)
var swsShiftVecOnce sync.Once

func (a *SwsVector) SwsShiftVec(shift ffcommon.FInt) {
	swsShiftVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsShiftVec, ffcommon.GetAvswscaleDll(), "sws_shiftVec")
	})

	swsShiftVec(a, shift)
}

// attribute_deprecated SwsVector *sws_cloneVec(SwsVector *a);
var swsCloneVec func(a *SwsVector) *SwsVector
var swsCloneVecOnce sync.Once

func (a *SwsVector) SwsCloneVec() *SwsVector {
	swsCloneVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsCloneVec, ffcommon.GetAvswscaleDll(), "sws_cloneVec")
	})

	return swsCloneVec(a)
	// if t == 0 {
	// 	return nil
	// }

	// return (*SwsVector)(unsafe.Pointer(t))
}

// attribute_deprecated void sws_printVec2(SwsVector *a, AVClass *log_ctx, int log_level);
var swsPrintVec2 func(a *SwsVector, log_ctx *AVClass, log_level ffcommon.FInt)
var swsPrintVec2Once sync.Once

func (a *SwsVector) SwsPrintVec2(log_ctx *AVClass, log_level ffcommon.FInt) {
	swsPrintVec2Once.Do(func() {
		purego.RegisterLibFunc(&swsPrintVec2, ffcommon.GetAvswscaleDll(), "sws_printVec2")
	})

	swsPrintVec2(a, log_ctx, log_level)
}

//#endif

// void sws_freeVec(SwsVector *a);

var swsFreeVec func(a *SwsVector)
var swsFreeVecOnce sync.Once

func (a *SwsVector) SwsFreeVec() {
	swsFreeVecOnce.Do(func() {
		purego.RegisterLibFunc(&swsFreeVec, ffcommon.GetAvswscaleDll(), "sws_freeVec")
	})

	swsFreeVec(a)
}

// SwsFilter *sws_getDefaultFilter(float lumaGBlur, float chromaGBlur,
// float lumaSharpen, float chromaSharpen,
// float chromaHShift, float chromaVShift,
// int verbose);
var swsGetDefaultFilter func(lumaGBlur, chromaGBlur, lumaSharpen, chromaSharpen, chromaHShift, chromaVShift *ffcommon.FFloat, verbose ffcommon.FInt) *SwsFilter
var swsGetDefaultFilterOnce sync.Once

func SwsGetDefaultFilter(lumaGBlur, chromaGBlur, lumaSharpen, chromaSharpen, chromaHShift, chromaVShift *ffcommon.FFloat, verbose ffcommon.FInt) *SwsFilter {
	swsGetDefaultFilterOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetDefaultFilter, ffcommon.GetAvswscaleDll(), "sws_getDefaultFilter")
	})

	return swsGetDefaultFilter(lumaGBlur, chromaGBlur, lumaSharpen, chromaSharpen, chromaHShift, chromaVShift, verbose)
	// if t == 0 {
	//     return nil
	// }

	// return (*SwsFilter)(unsafe.Pointer(t))
}

// void sws_freeFilter(SwsFilter *filter);
var swsFreeFilter func(filter *SwsFilter)
var swsFreeFilterOnce sync.Once

func (filter *SwsFilter) SwsFreeFilter() {
	swsFreeFilterOnce.Do(func() {
		purego.RegisterLibFunc(&swsFreeFilter, ffcommon.GetAvswscaleDll(), "sws_freeFilter")
	})

	swsFreeFilter(filter)
}

/**
 * Check if context can be reused, otherwise reallocate a new one.
 *
 * If context is NULL, just calls sws_getContext() to get a new
 * context. Otherwise, checks if the parameters are the ones already
 * saved in context. If that is the case, returns the current
 * context. Otherwise, frees context and gets a new context with
 * the new parameters.
 *
 * Be warned that srcFilter and dstFilter are not checked, they
 * are assumed to remain the same.
 */
//struct SwsContext *sws_getCachedContext(struct SwsContext *context,
//int srcW, int srcH, enum AVPixelFormat srcFormat,
//int dstW, int dstH, enum AVPixelFormat dstFormat,
//int flags, SwsFilter *srcFilter,
//SwsFilter *dstFilter, const double *param);
var swsGetCachedContext func(context *SwsContext, srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat, dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat, flags ffcommon.FInt, srcFilter, dstFilter *SwsFilter, param *ffcommon.FDouble) *SwsContext
var swsGetCachedContextOnce sync.Once

func (context *SwsContext) SwsGetCachedContext(srcW, srcH ffcommon.FInt, srcFormat AVPixelFormat, dstW, dstH ffcommon.FInt, dstFormat AVPixelFormat, flags ffcommon.FInt, srcFilter, dstFilter *SwsFilter, param *ffcommon.FDouble) *SwsContext {
	swsGetCachedContextOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetCachedContext, ffcommon.GetAvswscaleDll(), "sws_getCachedContext")
	})

	return swsGetCachedContext(context, srcW, srcH, srcFormat, dstW, dstH, dstFormat, flags, srcFilter, dstFilter, param)
	// if t == 0 {
	//     return nil
	// }

	// return (*SwsContext)(unsafe.Pointer(t))
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
 *
 * The output frame will have the same packed format as the palette.
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
//void sws_convertPalette8ToPacked32(const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette);
var swsConvertPalette8ToPacked32 func(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T)
var swsConvertPalette8ToPacked32Once sync.Once

func SwsConvertPalette8ToPacked32(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T) {
	swsConvertPalette8ToPacked32Once.Do(func() {
		purego.RegisterLibFunc(&swsConvertPalette8ToPacked32, ffcommon.GetAvswscaleDll(), "sws_convertPalette8ToPacked32")
	})

	swsConvertPalette8ToPacked32(src, dst, num_pixels, palette)
}

/**
 * Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
 *
 * With the palette format "ABCD", the destination frame ends up with the format "ABC".
 *
 * @param src        source frame buffer
 * @param dst        destination frame buffer
 * @param num_pixels number of pixels to convert
 * @param palette    array with [256] entries, which must match color arrangement (RGB or BGR) of src
 */
//void sws_convertPalette8ToPacked24(const uint8_t *src, uint8_t *dst, int num_pixels, const uint8_t *palette);
var swsConvertPalette8ToPacked24 func(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T)
var swsConvertPalette8ToPacked24Once sync.Once

func SwsConvertPalette8ToPacked24(src, dst *ffcommon.FUint8T, num_pixels ffcommon.FInt, palette *ffcommon.FUint8T) {
	swsConvertPalette8ToPacked24Once.Do(func() {
		purego.RegisterLibFunc(&swsConvertPalette8ToPacked24, ffcommon.GetAvswscaleDll(), "sws_convertPalette8ToPacked24")
	})

	swsConvertPalette8ToPacked24(src, dst, num_pixels, palette)
}

/**
 * Get the AVClass for swsContext. It can be used in combination with
 * AV_OPT_SEARCH_FAKE_OBJ for examining options.
 *
 * @see av_opt_find().
 */
//const AVClass *sws_get_class(void);
type AVClass = libavutil.AVClass

var swsGetClass func() *AVClass
var swsGetClassOnce sync.Once

func SwsGetClass() *AVClass {
	swsGetClassOnce.Do(func() {
		purego.RegisterLibFunc(&swsGetClass, ffcommon.GetAvswscaleDll(), "sws_get_class")
	})

	return swsGetClass()
	// if t == 0 {
	// 	return nil
	// }

	// return (*AVClass)(unsafe.Pointer(t))
}

/**
 * @}
 */

//#endif /* SWSCALE_SWSCALE_H */
