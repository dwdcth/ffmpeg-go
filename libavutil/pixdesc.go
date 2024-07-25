package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * pixel format descriptor
 * Copyright (c) 2009 Michael Niedermayer <michaelni@gmx.at>
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

//#ifndef AVUTIL_PIXDESC_H
//const AVUTIL_PIXDESC_H
//
//#include <inttypes.h>
//
//#include "attributes.h"
//#include "pixfmt.h"
//#include "version.h"

type AVComponentDescriptor struct {

	/**
	 * Which of the 4 planes contains the component.
	 */
	Plane ffcommon.FInt

	/**
	 * Number of elements between 2 horizontally consecutive pixels.
	 * Elements are bits for bitstream formats, bytes otherwise.
	 */
	Step ffcommon.FInt

	/**
	 * Number of elements before the component of the first pixel.
	 * Elements are bits for bitstream formats, bytes otherwise.
	 */
	Offset ffcommon.FInt

	/**
	 * Number of least significant bits that must be shifted away
	 * to get the value.
	 */
	Shift ffcommon.FInt

	/**
	 * Number of bits in the component.
	 */
	Depth ffcommon.FInt

	//#if FF_API_PLUS1_MINUS1
	/** deprecated, use step instead */
	//attribute_deprecated int step_minus1;
	StepMinus1 ffcommon.FInt
	/** deprecated, use depth instead */
	//attribute_deprecated int depth_minus1;
	DepthMinus1 ffcommon.FInt
	/** deprecated, use offset instead */
	//attribute_deprecated int offset_plus1;
	OffsetPlus1 ffcommon.FInt
	//#endif
}

/**
 * Descriptor that unambiguously describes how the bits of a pixel are
 * stored in the up to 4 data planes of an image. It also stores the
 * subsampling factors and number of components.
 *
 * @note This is separate of the colorspace (RGB, YCbCr, YPbPr, JPEG-style YUV
 *       and all the YUV variants) AVPixFmtDescriptor just stores how values
 *       are stored not what these values represent.
 */
type AVPixFmtDescriptor struct {
	Name         ffcommon.FCharPStruct
	NbComponents ffcommon.FUint8T ///< The number of components each pixel has, (1-4)

	/**
	 * Amount to shift the luma width right to find the chroma width.
	 * For YV12 this is 1 for example.
	 * chroma_width = AV_CEIL_RSHIFT(luma_width, log2_chroma_w)
	 * The note above is needed to ensure rounding up.
	 * This value only refers to the chroma components.
	 */
	Log2ChromaW ffcommon.FUint8T

	/**
	 * Amount to shift the luma height right to find the chroma height.
	 * For YV12 this is 1 for example.
	 * chroma_height= AV_CEIL_RSHIFT(luma_height, log2_chroma_h)
	 * The note above is needed to ensure rounding up.
	 * This value only refers to the chroma components.
	 */
	Log2ChromaH ffcommon.FUint8T

	/**
	 * Combination of AV_PIX_FMT_FLAG_... flags.
	 */
	Flags ffcommon.FUint64T

	/**
	 * Parameters that describe how pixels are packed.
	 * If the format has 1 or 2 components, then luma is 0.
	 * If the format has 3 or 4 components:
	 *   if the RGB flag is set then 0 is red, 1 is green and 2 is blue;
	 *   otherwise 0 is luma, 1 is chroma-U and 2 is chroma-V.
	 *
	 * If present, the Alpha channel is always the last component.
	 */
	Comp [4]AVComponentDescriptor

	/**
	 * Alternative comma-separated names.
	 */
	Alias ffcommon.FCharPStruct
}

/**
 * Pixel format is big-endian.
 */
const AV_PIX_FMT_FLAG_BE = (1 << 0)

/**
 * Pixel format has a palette in data[1], values are indexes in this palette.
 */
const AV_PIX_FMT_FLAG_PAL = (1 << 1)

/**
 * All values of a component are bit-wise packed end to end.
 */
const AV_PIX_FMT_FLAG_BITSTREAM = (1 << 2)

/**
 * Pixel format is an HW accelerated format.
 */
const AV_PIX_FMT_FLAG_HWACCEL = (1 << 3)

/**
 * At least one pixel component is not in the first data plane.
 */
const AV_PIX_FMT_FLAG_PLANAR = (1 << 4)

/**
 * The pixel format contains RGB-like data (as opposed to YUV/grayscale).
 */
const AV_PIX_FMT_FLAG_RGB = (1 << 5)

//#if FF_API_PSEUDOPAL
/**
 * The pixel format is "pseudo-paletted". This means that it contains a
 * fixed palette in the 2nd plane but the palette is fixed/constant for each
 * PIX_FMT. This allows interpreting the data as if it was PAL8, which can
 * in some cases be simpler. Or the data can be interpreted purely based on
 * the pixel format without using the palette.
 * An example of a pseudo-paletted format is AV_PIX_FMT_GRAY8
 *
 * @deprecated This flag is deprecated, and will be removed. When it is removed,
 * the extra palette allocation in AVFrame.data[1] is removed as well. Only
 * actual paletted formats (as indicated by AV_PIX_FMT_FLAG_PAL) will have a
 * palette. Starting with FFmpeg versions which have this flag deprecated, the
 * extra "pseudo" palette is already ignored, and API users are not required to
 * allocate a palette for AV_PIX_FMT_FLAG_PSEUDOPAL formats (it was required
 * before the deprecation, though).
 */
const AV_PIX_FMT_FLAG_PSEUDOPAL = (1 << 6)

//#endif

/**
 * The pixel format has an alpha channel. This is set on all formats that
 * support alpha in some way, including AV_PIX_FMT_PAL8. The alpha is always
 * straight, never pre-multiplied.
 *
 * If a codec or a filter does not support alpha, it should set all alpha to
 * opaque, or use the equivalent pixel formats without alpha component, e.g.
 * AV_PIX_FMT_RGB0 (or AV_PIX_FMT_RGB24 etc.) instead of AV_PIX_FMT_RGBA.
 */
const AV_PIX_FMT_FLAG_ALPHA = (1 << 7)

/**
 * The pixel format is following a Bayer pattern
 */
const AV_PIX_FMT_FLAG_BAYER = (1 << 8)

/**
 * The pixel format contains IEEE-754 floating point values. Precision (double,
 * single, or half) should be determined by the pixel size (64, 32, or 16 bits).
 */
const AV_PIX_FMT_FLAG_FLOAT = (1 << 9)

/**
 * Return the number of bits per pixel used by the pixel format
 * described by pixdesc. Note that this is not the same as the number
 * of bits per sample.
 *
 * The returned number of bits refers to the number of bits actually
 * used for storing the pixel information, that is padding bits are
 * not counted.
 */
//int av_get_bits_per_pixel(const AVPixFmtDescriptor *pixdesc);
var avGetBitsPerPixel func(pixdesc *AVPixFmtDescriptor) ffcommon.FInt
var avGetBitsPerPixelOnce sync.Once

func (pixdesc *AVPixFmtDescriptor) AvGetBitsPerPixel() ffcommon.FInt {
	avGetBitsPerPixelOnce.Do(func() {
		purego.RegisterLibFunc(&avGetBitsPerPixel, ffcommon.GetAvutilDll(), "av_get_bits_per_pixel")
	})
	return avGetBitsPerPixel(pixdesc)
}

/**
 * Return the number of bits per pixel for the pixel format
 * described by pixdesc, including any padding or unused bits.
 */
//int av_get_padded_bits_per_pixel(const AVPixFmtDescriptor *pixdesc);
var avGetPaddedBitsPerPixel func(pixdesc *AVPixFmtDescriptor) ffcommon.FInt
var avGetPaddedBitsPerPixelOnce sync.Once

func (pixdesc *AVPixFmtDescriptor) AvGetPaddedBitsPerPixel() ffcommon.FInt {
	avGetPaddedBitsPerPixelOnce.Do(func() {
		purego.RegisterLibFunc(&avGetPaddedBitsPerPixel, ffcommon.GetAvutilDll(), "av_get_padded_bits_per_pixel")
	})
	return avGetPaddedBitsPerPixel(pixdesc)
}

/**
 * @return a pixel format descriptor for provided pixel format or NULL if
 * this pixel format is unknown.
 */
//const AVPixFmtDescriptor *av_pix_fmt_desc_get(enum AVPixelFormat pix_fmt);
var avPixFmtDescGet func(pix_fmt AVPixelFormat) *AVPixFmtDescriptor
var avPixFmtDescGetOnce sync.Once

func AvPixFmtDescGet(pix_fmt AVPixelFormat) *AVPixFmtDescriptor {
	avPixFmtDescGetOnce.Do(func() {
		purego.RegisterLibFunc(&avPixFmtDescGet, ffcommon.GetAvutilDll(), "av_pix_fmt_desc_get")
	})
	return avPixFmtDescGet(pix_fmt)
}

/**
 * Iterate over all pixel format descriptors known to libavutil.
 *
 * @param prev previous descriptor. NULL to get the first descriptor.
 *
 * @return next descriptor or NULL after the last descriptor
 */
//const AVPixFmtDescriptor *av_pix_fmt_desc_next(const AVPixFmtDescriptor *prev);
var avPixFmtDescNext func(prev *AVPixFmtDescriptor) *AVPixFmtDescriptor
var avPixFmtDescNextOnce sync.Once

func (prev *AVPixFmtDescriptor) AvPixFmtDescNext() *AVPixFmtDescriptor {
	avPixFmtDescNextOnce.Do(func() {
		purego.RegisterLibFunc(&avPixFmtDescNext, ffcommon.GetAvutilDll(), "av_pix_fmt_desc_next")
	})
	return avPixFmtDescNext(prev)
}

/**
 * @return an AVPixelFormat id described by desc, or AV_PIX_FMT_NONE if desc
 * is not a valid pointer to a pixel format descriptor.
 */
//enum AVPixelFormat av_pix_fmt_desc_get_id(const AVPixFmtDescriptor *desc);
var avPixFmtDescGetId func(desc *AVPixFmtDescriptor) AVPixelFormat
var avPixFmtDescGetIdOnce sync.Once

func (desc *AVPixFmtDescriptor) AvPixFmtDescGetId() AVPixelFormat {
	avPixFmtDescGetIdOnce.Do(func() {
		purego.RegisterLibFunc(&avPixFmtDescGetId, ffcommon.GetAvutilDll(), "av_pix_fmt_desc_get_id")
	})
	return avPixFmtDescGetId(desc)
}

/**
 * Utility function to access log2_chroma_w log2_chroma_h from
 * the pixel format AVPixFmtDescriptor.
 *
 * @param[in]  pix_fmt the pixel format
 * @param[out] h_shift store log2_chroma_w (horizontal/width shift)
 * @param[out] v_shift store log2_chroma_h (vertical/height shift)
 *
 * @return 0 on success, AVERROR(ENOSYS) on invalid or unknown pixel format
 */
//int av_pix_fmt_get_chroma_sub_sample(enum AVPixelFormat pix_fmt,
//int *h_shift, int *v_shift);
var avPixFmtGetChromaSubSample func(pix_fmt AVPixelFormat, h_shift, v_shift *ffcommon.FInt) ffcommon.FInt
var avPixFmtGetChromaSubSampleOnce sync.Once

func AvPixFmtGetChromaSubSample(pix_fmt AVPixelFormat, h_shift, v_shift *ffcommon.FInt) ffcommon.FInt {
	avPixFmtGetChromaSubSampleOnce.Do(func() {
		purego.RegisterLibFunc(&avPixFmtGetChromaSubSample, ffcommon.GetAvutilDll(), "av_pix_fmt_get_chroma_sub_sample")
	})
	return avPixFmtGetChromaSubSample(pix_fmt, h_shift, v_shift)
}

/**
 * @return number of planes in pix_fmt, a negative AVERROR if pix_fmt is not a
 * valid pixel format.
 */
//int av_pix_fmt_count_planes(enum AVPixelFormat pix_fmt);
var avPixFmtCountPlanes func(pix_fmt AVPixelFormat) ffcommon.FInt
var avPixFmtCountPlanesOnce sync.Once

func AvPixFmtCountPlanes(pix_fmt AVPixelFormat) ffcommon.FInt {
	avPixFmtCountPlanesOnce.Do(func() {
		purego.RegisterLibFunc(&avPixFmtCountPlanes, ffcommon.GetAvutilDll(), "av_pix_fmt_count_planes")
	})
	return avPixFmtCountPlanes(pix_fmt)
}

/**
 * @return the name for provided color range or NULL if unknown.
 */
//const char *av_color_range_name(enum AVColorRange range);
var avColorRangeName func(range0 AVColorRange) ffcommon.FConstCharP
var avColorRangeNameOnce sync.Once

func AvColorRangeName(range0 AVColorRange) ffcommon.FConstCharP {
	avColorRangeNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorRangeName, ffcommon.GetAvutilDll(), "av_color_range_name")
	})
	return avColorRangeName(range0)
}

/**
 * @return the AVColorRange value for name or an AVError if not found.
 */
//int av_color_range_from_name(const char *name);
var avColorRangeFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avColorRangeFromNameOnce sync.Once

func AvColorRangeFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avColorRangeFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorRangeFromName, ffcommon.GetAvutilDll(), "av_color_range_from_name")
	})
	return avColorRangeFromName(name)
}

/**
 * @return the name for provided color primaries or NULL if unknown.
 */
//const char *av_color_primaries_name(enum AVColorPrimaries primaries);
var avColorPrimariesName func(primaries AVColorPrimaries) ffcommon.FConstCharP
var avColorPrimariesNameOnce sync.Once

func AvColorPrimariesName(primaries AVColorPrimaries) ffcommon.FConstCharP {
	avColorPrimariesNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorPrimariesName, ffcommon.GetAvutilDll(), "av_color_primaries_name")
	})
	return avColorPrimariesName(primaries)
}

/**
 * @return the AVColorPrimaries value for name or an AVError if not found.
 */
//int av_color_primaries_from_name(const char *name);
var avColorPrimariesFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avColorPrimariesFromNameOnce sync.Once

func AvColorPrimariesFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avColorPrimariesFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorPrimariesFromName, ffcommon.GetAvutilDll(), "av_color_primaries_from_name")
	})
	return avColorPrimariesFromName(name)
}

/**
 * @return the name for provided color transfer or NULL if unknown.
 */
//const char *av_color_transfer_name(enum AVColorTransferCharacteristic transfer);
var avColorTransferName func(transfer AVColorTransferCharacteristic) ffcommon.FConstCharP
var avColorTransferNameOnce sync.Once

func AvColorTransferName(transfer AVColorTransferCharacteristic) ffcommon.FConstCharP {
	avColorTransferNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorTransferName, ffcommon.GetAvutilDll(), "av_color_transfer_name")
	})
	return avColorTransferName(transfer)
}

/**
 * @return the AVColorTransferCharacteristic value for name or an AVError if not found.
 */
//int av_color_transfer_from_name(const char *name);
var avColorTransferFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avColorTransferFromNameOnce sync.Once

func AvColorTransferFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avColorTransferFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorTransferFromName, ffcommon.GetAvutilDll(), "av_color_transfer_from_name")
	})
	return avColorTransferFromName(name)
}

/**
 * @return the name for provided color space or NULL if unknown.
 */
//const char *av_color_space_name(enum AVColorSpace space);
var avColorSpaceName func(space AVColorSpace) ffcommon.FConstCharP
var avColorSpaceNameOnce sync.Once

func AvColorSpaceName(space AVColorSpace) ffcommon.FConstCharP {
	avColorSpaceNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorSpaceName, ffcommon.GetAvutilDll(), "av_color_space_name")
	})
	return avColorSpaceName(space)
}

/**
 * @return the AVColorSpace value for name or an AVError if not found.
 */
//int av_color_space_from_name(const char *name);
var avColorSpaceFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avColorSpaceFromNameOnce sync.Once

func AvColorSpaceFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avColorSpaceFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avColorSpaceFromName, ffcommon.GetAvutilDll(), "av_color_space_from_name")
	})
	return avColorSpaceFromName(name)
}

/**
 * @return the name for provided chroma location or NULL if unknown.
 */
//const char *av_chroma_location_name(enum AVChromaLocation location);
var avChromaLocationName func(location AVChromaLocation) ffcommon.FConstCharP
var avChromaLocationNameOnce sync.Once

func AvChromaLocationName(location AVChromaLocation) ffcommon.FConstCharP {
	avChromaLocationNameOnce.Do(func() {
		purego.RegisterLibFunc(&avChromaLocationName, ffcommon.GetAvutilDll(), "av_chroma_location_name")
	})
	return avChromaLocationName(location)
}

/**
 * @return the AVChromaLocation value for name or an AVError if not found.
 */
//int av_chroma_location_from_name(const char *name);
var avChromaLocationFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avChromaLocationFromNameOnce sync.Once

func AvChromaLocationFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avChromaLocationFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avChromaLocationFromName, ffcommon.GetAvutilDll(), "av_chroma_location_from_name")
	})
	return avChromaLocationFromName(name)
}

/**
 * Return the pixel format corresponding to name.
 *
 * If there is no pixel format with name name, then looks for a
 * pixel format with the name corresponding to the native endian
 * format of name.
 * For example in a little-endian system, first looks for "gray16",
 * then for "gray16le".
 *
 * Finally if no pixel format has been found, returns AV_PIX_FMT_NONE.
 */
//enum AVPixelFormat av_get_pix_fmt(const char *name);
var avGetPixFmt func(name ffcommon.FConstCharP) AVPixelFormat
var avGetPixFmtOnce sync.Once

func AvGetPixFmt(name ffcommon.FConstCharP) AVPixelFormat {
	avGetPixFmtOnce.Do(func() {
		purego.RegisterLibFunc(&avGetPixFmt, ffcommon.GetAvutilDll(), "av_get_pix_fmt")
	})
	return avGetPixFmt(name)
}

/**
 * Return the short name for a pixel format, NULL in case pix_fmt is
 * unknown.
 *
 * @see av_get_pix_fmt(), av_get_pix_fmt_string()
 */
//const char *av_get_pix_fmt_name(enum AVPixelFormat pix_fmt);
var avGetPixFmtName func(pix_fmt AVPixelFormat) ffcommon.FConstCharP
var avGetPixFmtNameOnce sync.Once

func AvGetPixFmtName(pix_fmt AVPixelFormat) ffcommon.FConstCharP {
	avGetPixFmtNameOnce.Do(func() {
		purego.RegisterLibFunc(&avGetPixFmtName, ffcommon.GetAvutilDll(), "av_get_pix_fmt_name")
	})
	return avGetPixFmtName(pix_fmt)
}

/**
 * Print in buf the string corresponding to the pixel format with
 * number pix_fmt, or a header if pix_fmt is negative.
 *
 * @param buf the buffer where to write the string
 * @param buf_size the size of buf
 * @param pix_fmt the number of the pixel format to print the
 * corresponding info string, or a negative value to print the
 * corresponding header.
 */
//char *av_get_pix_fmt_string(char *buf, int buf_size,
//enum AVPixelFormat pix_fmt);
var avGetPixFmtString func(buf ffcommon.FCharP, buf_size ffcommon.FInt, pix_fmt AVPixelFormat) ffcommon.FConstCharP
var avGetPixFmtStringOnce sync.Once

func AvGetPixFmtString(buf ffcommon.FCharP, buf_size ffcommon.FInt, pix_fmt AVPixelFormat) ffcommon.FConstCharP {
	avGetPixFmtStringOnce.Do(func() {
		purego.RegisterLibFunc(&avGetPixFmtString, ffcommon.GetAvutilDll(), "av_get_pix_fmt_string")
	})
	return avGetPixFmtString(buf, buf_size, pix_fmt)
}

/**
 * Read a line from an image, and write the values of the
 * pixel format component c to dst.
 *
 * @param data the array containing the pointers to the planes of the image
 * @param linesize the array containing the linesizes of the image
 * @param desc the pixel format descriptor for the image
 * @param x the horizontal coordinate of the first pixel to read
 * @param y the vertical coordinate of the first pixel to read
 * @param w the width of the line to read, that is the number of
 * values to write to dst
 * @param read_pal_component if not zero and the format is a paletted
 * format writes the values corresponding to the palette
 * component c in data[1] to dst, rather than the palette indexes in
 * data[0]. The behavior is undefined if the format is not paletted.
 * @param dst_element_size size of elements in dst array (2 or 4 byte)
 */
//void av_read_image_line2(void *dst, const uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w, int read_pal_component,
//int dst_element_size);
var avReadImageLine2 func(dst ffcommon.FVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, read_pal_component, dst_element_size ffcommon.FInt)

var avReadImageLine2Once sync.Once

func AvReadImageLine2(dst ffcommon.FVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, read_pal_component, dst_element_size ffcommon.FInt) {
	avReadImageLine2Once.Do(func() {
		purego.RegisterLibFunc(
			&avReadImageLine2,
			ffcommon.GetAvutilDll(),
			"av_read_image_line2",
		)
	})
	if avReadImageLine2 != nil {
		avReadImageLine2(dst, data, linesize, desc, x, y, c, w, read_pal_component, dst_element_size)
	}
}

// void av_read_image_line(uint16_t *dst, const uint8_t *data[4],
// const int linesize[4], const AVPixFmtDescriptor *desc,
// int x, int y, int c, int w, int read_pal_component);
var avReadImageLine func(dst *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]*ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, read_pal_component ffcommon.FInt)

var avReadImageLineOnce sync.Once

func AvReadImageLine(dst *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]*ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, read_pal_component ffcommon.FInt) {
	avReadImageLineOnce.Do(func() {
		purego.RegisterLibFunc(
			&avReadImageLine,
			ffcommon.GetAvutilDll(),
			"av_read_image_line",
		)
	})
	if avReadImageLine != nil {
		avReadImageLine(dst, data, linesize, desc, x, y, c, w, read_pal_component)
	}
}

/**
 * Write the values from src to the pixel format component c of an
 * image line.
 *
 * @param src array containing the values to write
 * @param data the array containing the pointers to the planes of the
 * image to write into. It is supposed to be zeroed.
 * @param linesize the array containing the linesizes of the image
 * @param desc the pixel format descriptor for the image
 * @param x the horizontal coordinate of the first pixel to write
 * @param y the vertical coordinate of the first pixel to write
 * @param w the width of the line to write, that is the number of
 * values to write to the image line
 * @param src_element_size size of elements in src array (2 or 4 byte)
 */
//void av_write_image_line2(const void *src, uint8_t *data[4],
//const int linesize[4], const AVPixFmtDescriptor *desc,
//int x, int y, int c, int w, int src_element_size);
var avWriteImageLine2 func(src ffcommon.FConstVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, src_element_size ffcommon.FInt)

var avWriteImageLine2Once sync.Once

func AvWriteImageLine2(src ffcommon.FConstVoidP, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w, src_element_size ffcommon.FInt) {
	avWriteImageLine2Once.Do(func() {
		purego.RegisterLibFunc(
			&avWriteImageLine2,
			ffcommon.GetAvutilDll(),
			"av_write_image_line2",
		)
	})
	if avWriteImageLine2 != nil {
		avWriteImageLine2(src, data, linesize, desc, x, y, c, w, src_element_size)
	}
}

// void av_write_image_line(const uint16_t *src, uint8_t *data[4],
// const int linesize[4], const AVPixFmtDescriptor *desc,
// int x, int y, int c, int w);
var avWriteImageLine func(src *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w ffcommon.FInt)

var avWriteImageLineOnce sync.Once

func AvWriteImageLine(src *ffcommon.FUint16T, data [4]*ffcommon.FUint8T,
	linesize [4]ffcommon.FInt, desc *AVPixFmtDescriptor,
	x, y, c, w ffcommon.FInt) {
	avWriteImageLineOnce.Do(func() {
		purego.RegisterLibFunc(
			&avWriteImageLine,
			ffcommon.GetAvutilDll(),
			"av_write_image_line",
		)
	})
	if avWriteImageLine != nil {
		avWriteImageLine(src, data, linesize, desc, x, y, c, w)
	}
}

/**
 * Utility function to swap the endianness of a pixel format.
 *
 * @param[in]  pix_fmt the pixel format
 *
 * @return pixel format with swapped endianness if it exists,
 * otherwise AV_PIX_FMT_NONE
 */
//enum AVPixelFormat av_pix_fmt_swap_endianness(enum AVPixelFormat pix_fmt);
var avPixFmtSwapEndianness func(pix_fmt AVPixelFormat) AVPixelFormat
var avPixFmtSwapEndiannessOnce sync.Once

func AvPixFmtSwapEndianness(pix_fmt AVPixelFormat) AVPixelFormat {
	avPixFmtSwapEndiannessOnce.Do(func() {
		purego.RegisterLibFunc(
			&avPixFmtSwapEndianness,
			ffcommon.GetAvutilDll(),
			"av_pix_fmt_swap_endianness",
		)
	})
	if avPixFmtSwapEndianness != nil {
		return avPixFmtSwapEndianness(pix_fmt)
	}
	return pix_fmt
}

const FF_LOSS_RESOLUTION = 0x0001 /**< loss due to resolution change */
const FF_LOSS_DEPTH = 0x0002      /**< loss due to color depth change */
const FF_LOSS_COLORSPACE = 0x0004 /**< loss due to color space conversion */
const FF_LOSS_ALPHA = 0x0008      /**< loss of alpha bits */
const FF_LOSS_COLORQUANT = 0x0010 /**< loss due to color quantization */
const FF_LOSS_CHROMA = 0x0020     /**< loss of chroma (e.g. RGB to gray conversion) */

/**
 * Compute what kind of losses will occur when converting from one specific
 * pixel format to another.
 * When converting from one pixel format to another, information loss may occur.
 * For example, when converting from RGB24 to GRAY, the color information will
 * be lost. Similarly, other losses occur when converting from some formats to
 * other formats. These losses can involve loss of chroma, but also loss of
 * resolution, loss of color depth, loss due to the color space conversion, loss
 * of the alpha bits or loss due to color quantization.
 * av_get_fix_fmt_loss() informs you about the various types of losses
 * which will occur when converting from one pixel format to another.
 *
 * @param[in] dst_pix_fmt destination pixel format
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @return Combination of flags informing you what kind of losses will occur
 * (maximum loss for an invalid dst_pix_fmt).
 */
//int av_get_pix_fmt_loss(enum AVPixelFormat dst_pix_fmt,
//enum AVPixelFormat src_pix_fmt,
//int has_alpha);
var avGetPixFmtLoss func(dst_pix_fmt, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt) ffcommon.FInt
var avGetPixFmtLossOnce sync.Once

func AvGetPixFmtLoss(dst_pix_fmt, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt) ffcommon.FInt {
	avGetPixFmtLossOnce.Do(func() {
		purego.RegisterLibFunc(
			&avGetPixFmtLoss,
			ffcommon.GetAvutilDll(),
			"av_get_pix_fmt_loss",
		)
	})
	if avGetPixFmtLoss != nil {
		return avGetPixFmtLoss(dst_pix_fmt, src_pix_fmt, has_alpha)
	}
	return ffcommon.FInt(0)
}

/**
 * Compute what kind of losses will occur when converting from one specific
 * pixel format to another.
 * When converting from one pixel format to another, information loss may occur.
 * For example, when converting from RGB24 to GRAY, the color information will
 * be lost. Similarly, other losses occur when converting from some formats to
 * other formats. These losses can involve loss of chroma, but also loss of
 * resolution, loss of color depth, loss due to the color space conversion, loss
 * of the alpha bits or loss due to color quantization.
 * av_get_fix_fmt_loss() informs you about the various types of losses
 * which will occur when converting from one pixel format to another.
 *
 * @param[in] dst_pix_fmt destination pixel format
 * @param[in] src_pix_fmt source pixel format
 * @param[in] has_alpha Whether the source pixel format alpha channel is used.
 * @return Combination of flags informing you what kind of losses will occur
 * (maximum loss for an invalid dst_pix_fmt).
 */
//enum AVPixelFormat av_find_best_pix_fmt_of_2(enum AVPixelFormat dst_pix_fmt1, enum AVPixelFormat dst_pix_fmt2,
//enum AVPixelFormat src_pix_fmt, int has_alpha, int *loss_ptr);
var avFindBestPixFmtOf2 func(dst_pix_fmt1, dst_pix_fmt2, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) AVPixelFormat
var avFindBestPixFmtOf2Once sync.Once

func AvFindBestPixFmtOf2(dst_pix_fmt1, dst_pix_fmt2, src_pix_fmt AVPixelFormat, has_alpha ffcommon.FInt, loss_ptr *ffcommon.FInt) AVPixelFormat {
	avFindBestPixFmtOf2Once.Do(func() {
		purego.RegisterLibFunc(
			&avFindBestPixFmtOf2,
			ffcommon.GetAvutilDll(),
			"av_find_best_pix_fmt_of_2",
		)
	})
	if avFindBestPixFmtOf2 != nil {
		return avFindBestPixFmtOf2(dst_pix_fmt1, dst_pix_fmt2, src_pix_fmt, has_alpha, loss_ptr)
	}
	return AVPixelFormat(0)
}

//#endif /* AVUTIL_PIXDESC_H */
