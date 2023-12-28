package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
 * Copyright (c) 2008 Peter Ross
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

//#ifndef AVUTIL_CHANNEL_LAYOUT_H
//#define AVUTIL_CHANNEL_LAYOUT_H
//
//#include <stdint.h>

/**
 * @file
 * audio channel layout utility functions
 */

/**
 * @addtogroup lavu_audio
 * @{
 */

/**
 * @defgroup channel_masks Audio channel masks
 *
 * A channel layout is a 64-bits integer with a bit set for every channel.
 * The number of bits set must be equal to the number of channels.
 * The value 0 means that the channel layout is not known.
 * @note this data structure is not powerful enough to handle channels
 * combinations that have the same channel multiple times, such as
 * dual-mono.
 *
 * @{
 */
const AV_CH_FRONT_LEFT = 0x00000001
const AV_CH_FRONT_RIGHT = 0x00000002
const AV_CH_FRONT_CENTER = 0x00000004
const AV_CH_LOW_FREQUENCY = 0x00000008
const AV_CH_BACK_LEFT = 0x00000010
const AV_CH_BACK_RIGHT = 0x00000020
const AV_CH_FRONT_LEFT_OF_CENTER = 0x00000040
const AV_CH_FRONT_RIGHT_OF_CENTER = 0x00000080
const AV_CH_BACK_CENTER = 0x00000100
const AV_CH_SIDE_LEFT = 0x00000200
const AV_CH_SIDE_RIGHT = 0x00000400
const AV_CH_TOP_CENTER = 0x00000800
const AV_CH_TOP_FRONT_LEFT = 0x00001000
const AV_CH_TOP_FRONT_CENTER = 0x00002000
const AV_CH_TOP_FRONT_RIGHT = 0x00004000
const AV_CH_TOP_BACK_LEFT = 0x00008000
const AV_CH_TOP_BACK_CENTER = 0x00010000
const AV_CH_TOP_BACK_RIGHT = 0x00020000
const AV_CH_STEREO_LEFT = 0x20000000  ///< Stereo downmix.
const AV_CH_STEREO_RIGHT = 0x40000000 ///< See AV_CH_STEREO_LEFT.
const AV_CH_WIDE_LEFT = 0x0000000080000000
const AV_CH_WIDE_RIGHT = 0x0000000100000000
const AV_CH_SURROUND_DIRECT_LEFT = 0x0000000200000000
const AV_CH_SURROUND_DIRECT_RIGHT = 0x0000000400000000
const AV_CH_LOW_FREQUENCY_2 = 0x0000000800000000
const AV_CH_TOP_SIDE_LEFT = 0x0000001000000000
const AV_CH_TOP_SIDE_RIGHT = 0x0000002000000000
const AV_CH_BOTTOM_FRONT_CENTER = 0x0000004000000000
const AV_CH_BOTTOM_FRONT_LEFT = 0x0000008000000000
const AV_CH_BOTTOM_FRONT_RIGHT = 0x0000010000000000

/*
  - Channel mask value used for AVCodecContext.request_channel_layout
    to indicate that the user requests the channel order of the decoder output
    to be the native codec channel order.
*/
const AV_CH_LAYOUT_NATIVE = 0x8000000000000000

/**
 * @}
 * @defgroup channel_mask_c Audio channel layouts
 * @{
 * */
const AV_CH_LAYOUT_MONO = (AV_CH_FRONT_CENTER)
const AV_CH_LAYOUT_STEREO = (AV_CH_FRONT_LEFT | AV_CH_FRONT_RIGHT)
const AV_CH_LAYOUT_2POINT1 = (AV_CH_LAYOUT_STEREO | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_2_1 = (AV_CH_LAYOUT_STEREO | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_SURROUND = (AV_CH_LAYOUT_STEREO | AV_CH_FRONT_CENTER)
const AV_CH_LAYOUT_3POINT1 = (AV_CH_LAYOUT_SURROUND | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_4POINT0 = (AV_CH_LAYOUT_SURROUND | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_4POINT1 = (AV_CH_LAYOUT_4POINT0 | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_2_2 = (AV_CH_LAYOUT_STEREO | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT)
const AV_CH_LAYOUT_QUAD = (AV_CH_LAYOUT_STEREO | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
const AV_CH_LAYOUT_5POINT0 = (AV_CH_LAYOUT_SURROUND | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT)
const AV_CH_LAYOUT_5POINT1 = (AV_CH_LAYOUT_5POINT0 | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_5POINT0_BACK = (AV_CH_LAYOUT_SURROUND | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
const AV_CH_LAYOUT_5POINT1_BACK = (AV_CH_LAYOUT_5POINT0_BACK | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_6POINT0 = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_6POINT0_FRONT = (AV_CH_LAYOUT_2_2 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
const AV_CH_LAYOUT_HEXAGONAL = (AV_CH_LAYOUT_5POINT0_BACK | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_6POINT1 = (AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_6POINT1_BACK = (AV_CH_LAYOUT_5POINT1_BACK | AV_CH_BACK_CENTER)
const AV_CH_LAYOUT_6POINT1_FRONT = (AV_CH_LAYOUT_6POINT0_FRONT | AV_CH_LOW_FREQUENCY)
const AV_CH_LAYOUT_7POINT0 = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
const AV_CH_LAYOUT_7POINT0_FRONT = (AV_CH_LAYOUT_5POINT0 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
const AV_CH_LAYOUT_7POINT1 = (AV_CH_LAYOUT_5POINT1 | AV_CH_BACK_LEFT | AV_CH_BACK_RIGHT)
const AV_CH_LAYOUT_7POINT1_WIDE = (AV_CH_LAYOUT_5POINT1 | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
const AV_CH_LAYOUT_7POINT1_WIDE_BACK = (AV_CH_LAYOUT_5POINT1_BACK | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER)
const AV_CH_LAYOUT_OCTAGONAL = (AV_CH_LAYOUT_5POINT0 | AV_CH_BACK_LEFT | AV_CH_BACK_CENTER | AV_CH_BACK_RIGHT)
const AV_CH_LAYOUT_HEXADECAGONAL = (AV_CH_LAYOUT_OCTAGONAL | AV_CH_WIDE_LEFT | AV_CH_WIDE_RIGHT | AV_CH_TOP_BACK_LEFT | AV_CH_TOP_BACK_RIGHT | AV_CH_TOP_BACK_CENTER | AV_CH_TOP_FRONT_CENTER | AV_CH_TOP_FRONT_LEFT | AV_CH_TOP_FRONT_RIGHT)
const AV_CH_LAYOUT_STEREO_DOWNMIX = (AV_CH_STEREO_LEFT | AV_CH_STEREO_RIGHT)
const AV_CH_LAYOUT_22POINT2 = (AV_CH_LAYOUT_5POINT1_BACK | AV_CH_FRONT_LEFT_OF_CENTER | AV_CH_FRONT_RIGHT_OF_CENTER | AV_CH_BACK_CENTER | AV_CH_LOW_FREQUENCY_2 | AV_CH_SIDE_LEFT | AV_CH_SIDE_RIGHT | AV_CH_TOP_FRONT_LEFT | AV_CH_TOP_FRONT_RIGHT | AV_CH_TOP_FRONT_CENTER | AV_CH_TOP_CENTER | AV_CH_TOP_BACK_LEFT | AV_CH_TOP_BACK_RIGHT | AV_CH_TOP_SIDE_LEFT | AV_CH_TOP_SIDE_RIGHT | AV_CH_TOP_BACK_CENTER | AV_CH_BOTTOM_FRONT_CENTER | AV_CH_BOTTOM_FRONT_LEFT | AV_CH_BOTTOM_FRONT_RIGHT)

type AVMatrixEncoding = int32

const (
	AV_MATRIX_ENCODING_NONE = iota
	AV_MATRIX_ENCODING_DOLBY
	AV_MATRIX_ENCODING_DPLII
	AV_MATRIX_ENCODING_DPLIIX
	AV_MATRIX_ENCODING_DPLIIZ
	AV_MATRIX_ENCODING_DOLBYEX
	AV_MATRIX_ENCODING_DOLBYHEADPHONE
	AV_MATRIX_ENCODING_NB
)

/**
 * Return a channel layout id that matches name, or 0 if no match is found.
 *
 * name can be one or several of the following notations,
 * separated by '+' or '|':
 * - the name of an usual channel layout (mono, stereo, 4.0, quad, 5.0,
 *   5.0(side), 5.1, 5.1(side), 7.1, 7.1(wide), downmix);
 * - the name of a single channel (FL, FR, FC, LFE, BL, BR, FLC, FRC, BC,
 *   SL, SR, TC, TFL, TFC, TFR, TBL, TBC, TBR, DL, DR);
 * - a number of channels, in decimal, followed by 'c', yielding
 *   the default channel layout for that number of channels (@see
 *   av_get_default_channel_layout);
 * - a channel layout mask, in hexadecimal starting with "0x" (see the
 *   AV_CH_* macros).
 *
 * Example: "stereo+FC" = "2c+FC" = "2c+1c" = "0x7"
 */
//uint64_t av_get_channel_layout(const char *name);
var avGetChannelLayout func(name ffcommon.FConstCharP) ffcommon.FUint64T

var avGetChannelLayoutOnce sync.Once

func AvGetChannelLayout(name ffcommon.FConstCharP) ffcommon.FUint64T {
	avGetChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelLayout, ffcommon.GetAvutilDll(), "av_get_channel_layout")
	})
	return avGetChannelLayout(name)
}

/**
 * Return a channel layout and the number of channels based on the specified name.
 *
 * This function is similar to (@see av_get_channel_layout), but can also parse
 * unknown channel layout specifications.
 *
 * @param[in]  name             channel layout specification string
 * @param[out] channel_layout   parsed channel layout (0 if unknown)
 * @param[out] nb_channels      number of channels
 *
 * @return 0 on success, AVERROR(EINVAL) if the parsing fails.
 */
//int av_get_extended_channel_layout(const char *name, uint64_t* channel_layout, int* nb_channels);
var avGetExtendedChannelLayout func(name ffcommon.FConstCharP, channel_layout *ffcommon.FUint64T, nb_channels *ffcommon.FUint) ffcommon.FInt

var avGetExtendedChannelLayoutOnce sync.Once

func AvGetExtendedChannelLayout(name ffcommon.FConstCharP, channel_layout *ffcommon.FUint64T, nb_channels *ffcommon.FUint) ffcommon.FInt {
	avGetExtendedChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avGetExtendedChannelLayout, ffcommon.GetAvutilDll(), "av_get_extended_channel_layout")
	})
	return avGetExtendedChannelLayout(name, channel_layout, nb_channels)
}

/**
 * Return a description of a channel layout.
 * If nb_channels is <= 0, it is guessed from the channel_layout.
 *
 * @param buf put here the string containing the channel layout
 * @param buf_size size in bytes of the buffer
 */
//void av_get_channel_layout_string(char *buf, int buf_size, int nb_channels, uint64_t channel_layout);
var avGetChannelLayoutString func(buf ffcommon.FBuf, buf_size, nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T)

var avGetChannelLayoutStringOnce sync.Once

func AvGetChannelLayoutString(buf ffcommon.FBuf, buf_size, nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T) {
	avGetChannelLayoutStringOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelLayoutString, ffcommon.GetAvutilDll(), "av_get_channel_layout_string")
	})
	avGetChannelLayoutString(buf, buf_size, nb_channels, channel_layout)
}

//struct AVBPrint;

/**
 * Append a description of a channel layout to a bprint buffer.
 */
//void av_bprint_channel_layout(struct AVBPrint *bp, int nb_channels, uint64_t channel_layout);
var avBprintChannelLayout func(bp *AVBPrint, nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T)

var avBprintChannelLayoutOnce sync.Once

func (bp *AVBPrint) AvBprintChannelLayout(nb_channels ffcommon.FInt, channel_layout ffcommon.FUint64T) {
	avBprintChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintChannelLayout, ffcommon.GetAvutilDll(), "av_bprint_channel_layout")
	})
	avBprintChannelLayout(bp, nb_channels, channel_layout)
}

/**
 * Return the number of channels in the channel layout.
 */
//int av_get_channel_layout_nb_channels(uint64_t channel_layout);
var avGetChannelLayoutNbChannels func(channel_layout ffcommon.FUint64T) ffcommon.FInt

var avGetChannelLayoutNbChannelsOnce sync.Once

func AvGetChannelLayoutNbChannels(channel_layout ffcommon.FUint64T) ffcommon.FInt {
	avGetChannelLayoutNbChannelsOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelLayoutNbChannels, ffcommon.GetAvutilDll(), "av_get_channel_layout_nb_channels")
	})
	return avGetChannelLayoutNbChannels(channel_layout)
}

/**
 * Return default channel layout for a given number of channels.
 */
//int64_t av_get_default_channel_layout(int nb_channels);
var avGetDefaultChannelLayout func(nb_channels ffcommon.FInt) ffcommon.FInt64T

var avGetDefaultChannelLayoutOnce sync.Once

func AvGetDefaultChannelLayout(nb_channels ffcommon.FInt) ffcommon.FInt64T {
	avGetDefaultChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avGetDefaultChannelLayout, ffcommon.GetAvutilDll(), "av_get_default_channel_layout")
	})
	return avGetDefaultChannelLayout(nb_channels)
}

/**
 * Get the index of a channel in channel_layout.
 *
 * @param channel a channel layout describing exactly one channel which must be
 *                present in channel_layout.
 *
 * @return index of channel in channel_layout on success, a negative AVERROR
 *         on error.
 */
//int av_get_channel_layout_channel_index(uint64_t channel_layout,
//uint64_t channel);
var avGetChannelLayoutChannelIndex func(channel_layout, channel ffcommon.FUint64T) ffcommon.FInt

var avGetChannelLayoutChannelIndexOnce sync.Once

func AvGetChannelLayoutChannelIndex(channel_layout, channel ffcommon.FUint64T) ffcommon.FInt {
	avGetChannelLayoutChannelIndexOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelLayoutChannelIndex, ffcommon.GetAvutilDll(), "av_get_channel_layout_channel_index")
	})
	return avGetChannelLayoutChannelIndex(channel_layout, channel)
}

/**
 * Get the channel with the given index in channel_layout.
 */
//uint64_t av_channel_layout_extract_channel(uint64_t channel_layout, int index);
var avChannelLayoutExtractChannel func(channel_layout ffcommon.FUint64T, index ffcommon.FInt) ffcommon.FUint64T

var avChannelLayoutExtractChannelOnce sync.Once

func AvChannelLayoutExtractChannel(channel_layout ffcommon.FUint64T, index ffcommon.FInt) ffcommon.FUint64T {
	avChannelLayoutExtractChannelOnce.Do(func() {
		purego.RegisterLibFunc(&avChannelLayoutExtractChannel, ffcommon.GetAvutilDll(), "av_channel_layout_extract_channel")
	})
	return avChannelLayoutExtractChannel(channel_layout, index)
}

/**
 * Get the name of a given channel.
 *
 * @return channel name on success, NULL on error.
 */
//const char *av_get_channel_name(uint64_t channel);
var avGetChannelName func(channel ffcommon.FUint64T) ffcommon.FConstCharP

var avGetChannelNameOnce sync.Once

func AvGetChannelName(channel ffcommon.FUint64T) ffcommon.FConstCharP {
	avGetChannelNameOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelName, ffcommon.GetAvutilDll(), "av_get_channel_name")
	})
	return avGetChannelName(channel)
}

/**
 * Get the description of a given channel.
 *
 * @param channel  a channel layout with a single channel
 * @return  channel description on success, NULL on error
 */
//const char *av_get_channel_description(uint64_t channel);
var avGetChannelDescription func(channel ffcommon.FUint64T) ffcommon.FConstCharP

var avGetChannelDescriptionOnce sync.Once

func AvGetChannelDescription(channel ffcommon.FUint64T) ffcommon.FConstCharP {
	avGetChannelDescriptionOnce.Do(func() {
		purego.RegisterLibFunc(&avGetChannelDescription, ffcommon.GetAvutilDll(), "av_get_channel_description")
	})
	return avGetChannelDescription(channel)
}

/**
 * Get the value and name of a standard channel layout.
 *
 * @param[in]  index   index in an internal list, starting at 0
 * @param[out] layout  channel layout mask
 * @param[out] name    name of the layout
 * @return  0  if the layout exists,
 *          <0 if index is beyond the limits
 */
//int av_get_standard_channel_layout(unsigned index, uint64_t *layout,
//const char **name);
var avGetStandardChannelLayout func(index ffcommon.FUnsigned, layout *ffcommon.FUint64T, name *ffcommon.FBuf) ffcommon.FInt

var avGetStandardChannelLayoutOnce sync.Once

func AvGetStandardChannelLayout(index ffcommon.FUnsigned, layout *ffcommon.FUint64T, name *ffcommon.FBuf) ffcommon.FInt {
	avGetStandardChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(&avGetStandardChannelLayout, ffcommon.GetAvutilDll(), "av_get_standard_channel_layout")
	})
	return avGetStandardChannelLayout(index, layout, name)
}

/**
 * @}
 * @}
 */

//#endif /* AVUTIL_CHANNEL_LAYOUT_H */
