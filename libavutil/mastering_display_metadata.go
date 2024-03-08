package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2016 Neil Birkbeck <neil.birkbeck@gmail.com>
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

//#ifndef AVUTIL_MASTERING_DISPLAY_METADATA_H
//#define AVUTIL_MASTERING_DISPLAY_METADATA_H
//
//#include "frame.h"
//#include "rational.h"

/**
 * Mastering display metadata capable of representing the color volume of
 * the display used to master the content (SMPTE 2086:2014).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_mastering_display_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVMasteringDisplayMetadata struct {

	/**
	 * CIE 1931 xy chromaticity coords of color primaries (r, g, b order).
	 */
	DisplayPrimaries [3][2]AVRational

	/**
	 * CIE 1931 xy chromaticity coords of white point.
	 */
	WhitePoint [2]AVRational

	/**
	 * Min luminance of mastering display (cd/m^2).
	 */
	MinLuminance AVRational

	/**
	 * Max luminance of mastering display (cd/m^2).
	 */
	MaxLuminance AVRational

	/**
	 * Flag indicating whether the display primaries (and white point) are set.
	 */
	HasPrimaries ffcommon.FInt

	/**
	 * Flag indicating whether the luminance (min_ and max_) have been set.
	 */
	HasLuminance ffcommon.FInt
}

/**
 * Allocate an AVMasteringDisplayMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVMasteringDisplayMetadata filled with default values or NULL
 *         on failure.
 */
//AVMasteringDisplayMetadata *av_mastering_display_metadata_alloc(void);
// purego func
var avMasteringDisplayMetadataAlloc func() *AVMasteringDisplayMetadata
var avMasteringDisplayMetadataAllocOnce sync.Once

func AvMasteringDisplayMetadataAlloc() (res *AVMasteringDisplayMetadata) {
	avMasteringDisplayMetadataAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avMasteringDisplayMetadataAlloc, ffcommon.GetAvutilDll(), "av_mastering_display_metadata_alloc")
	})
	res = avMasteringDisplayMetadataAlloc()
	return
}

/**
 * Allocate a complete AVMasteringDisplayMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVMasteringDisplayMetadata structure to be filled by caller.
 */
//AVMasteringDisplayMetadata *av_mastering_display_metadata_create_side_data(AVFrame *frame);
// purego struct method
var avMasteringDisplayMetadataCreateSideData func(frame *AVFrame) *AVMasteringDisplayMetadata
var avMasteringDisplayMetadataCreateSideDataOnce sync.Once

func (frame *AVFrame) AvMasteringDisplayMetadataCreateSideData() (res *AVMasteringDisplayMetadata) {
	avMasteringDisplayMetadataCreateSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avMasteringDisplayMetadataCreateSideData, ffcommon.GetAvutilDll(), "av_mastering_display_metadata_create_side_data")
	})
	res = avMasteringDisplayMetadataCreateSideData(frame)
	return
}

/**
 * Content light level needed by to transmit HDR over HDMI (CTA-861.3).
 *
 * To be used as payload of a AVFrameSideData or AVPacketSideData with the
 * appropriate type.
 *
 * @note The struct should be allocated with av_content_light_metadata_alloc()
 *       and its size is not a part of the public ABI.
 */
type AVContentLightMetadata struct {

	/**
	 * Max content light level (cd/m^2).
	 */
	MaxCLL ffcommon.FUnsigned

	/**
	 * Max average light level per frame (cd/m^2).
	 */
	MaxFALL ffcommon.FUnsigned
}

/**
 * Allocate an AVContentLightMetadata structure and set its fields to
 * default values. The resulting struct can be freed using av_freep().
 *
 * @return An AVContentLightMetadata filled with default values or NULL
 *         on failure.
 */
//AVContentLightMetadata *av_content_light_metadata_alloc(size_t *size);
// purego func
var avContentLightMetadataAlloc func(size ffcommon.FSizeT) *AVContentLightMetadata
var avContentLightMetadataAllocOnce sync.Once

func AvContentLightMetadataAlloc(size ffcommon.FSizeT) (res *AVContentLightMetadata) {
	avContentLightMetadataAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avContentLightMetadataAlloc, ffcommon.GetAvutilDll(), "av_content_light_metadata_alloc")
	})
	res = avContentLightMetadataAlloc(size)
	return
}

/**
 * Allocate a complete AVContentLightMetadata and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVContentLightMetadata structure to be filled by caller.
 */
//AVContentLightMetadata *av_content_light_metadata_create_side_data(AVFrame *frame);
// purego struct method
var avContentLightMetadataCreateSideData func(frame *AVFrame) *AVContentLightMetadata
var avContentLightMetadataCreateSideDataOnce sync.Once

func (frame *AVFrame) AvContentLightMetadataCreateSideData() (res *AVContentLightMetadata) {
	avContentLightMetadataCreateSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avContentLightMetadataCreateSideData, ffcommon.GetAvutilDll(), "av_content_light_metadata_create_side_data")
	})
	res = avContentLightMetadataCreateSideData(frame)
	return
}

//#endif /* AVUTIL_MASTERING_DISPLAY_METADATA_H */
