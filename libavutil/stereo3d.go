package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2013 Vittorio Giovara <vittorio.giovara@gmail.com>
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

/**
 * @file
 * Stereoscopic video
 */

//#ifndef AVUTIL_STEREO3D_H
//#define AVUTIL_STEREO3D_H
//
//#include <stdint.h>
//
//#include "frame.h"

/**
 * @addtogroup lavu_video
 * @{
 *
 * @defgroup lavu_video_stereo3d Stereo3D types and functions
 * @{
 */

/**
 * @addtogroup lavu_video_stereo3d
 * A stereoscopic video file consists in multiple views embedded in a single
 * frame, usually describing two views of a scene. This file describes all
 * possible codec-independent view arrangements.
 * */

/**
 * List of possible 3D Types
 */
type AVStereo3DType int32

const (
	/**
	 * Video is not stereoscopic (and metadata has to be there).
	 */
	AV_STEREO3D_2D = iota

	/**
	 * Views are next to each other.
	 *
	 * @code{.unparsed}
	 *    LLLLRRRR
	 *    LLLLRRRR
	 *    LLLLRRRR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_SIDEBYSIDE

	/**
	 * Views are on top of each other.
	 *
	 * @code{.unparsed}
	 *    LLLLLLLL
	 *    LLLLLLLL
	 *    RRRRRRRR
	 *    RRRRRRRR
	 * @endcode
	 */
	AV_STEREO3D_TOPBOTTOM

	/**
	 * Views are alternated temporally.
	 *
	 * @code{.unparsed}
	 *     frame0   frame1   frame2   ...
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    LLLLLLLL RRRRRRRR LLLLLLLL
	 *    ...      ...      ...
	 * @endcode
	 */
	AV_STEREO3D_FRAMESEQUENCE

	/**
	 * Views are packed in a checkerboard-like structure per pixel.
	 *
	 * @code{.unparsed}
	 *    LRLRLRLR
	 *    RLRLRLRL
	 *    LRLRLRLR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_CHECKERBOARD

	/**
	 * Views are next to each other, but when upscaling
	 * apply a checkerboard pattern.
	 *
	 * @code{.unparsed}
	 *     LLLLRRRR          L L L L    R R R R
	 *     LLLLRRRR    =>     L L L L  R R R R
	 *     LLLLRRRR          L L L L    R R R R
	 *     LLLLRRRR           L L L L  R R R R
	 * @endcode
	 */
	AV_STEREO3D_SIDEBYSIDE_QUINCUNX

	/**
	 * Views are packed per line, as if interlaced.
	 *
	 * @code{.unparsed}
	 *    LLLLLLLL
	 *    RRRRRRRR
	 *    LLLLLLLL
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_LINES

	/**
	 * Views are packed per column.
	 *
	 * @code{.unparsed}
	 *    LRLRLRLR
	 *    LRLRLRLR
	 *    LRLRLRLR
	 *    ...
	 * @endcode
	 */
	AV_STEREO3D_COLUMNS
)

/**
 * List of possible view types.
 */
type AVStereo3DView int32

const (
	/**
	 * Frame contains two packed views.
	 */
	AV_STEREO3D_VIEW_PACKED = iota

	/**
	 * Frame contains only the left view.
	 */
	AV_STEREO3D_VIEW_LEFT

	/**
	 * Frame contains only the right view.
	 */
	AV_STEREO3D_VIEW_RIGHT
)

/**
 * Inverted views, Right/Bottom represents the left view.
 */
const AV_STEREO3D_FLAG_INVERT = (1 << 0)

/**
 * Stereo 3D type: this structure describes how two videos are packed
 * within a single video surface, with additional information as needed.
 *
 * @note The struct must be allocated with av_stereo3d_alloc() and
 *       its size is not a part of the public ABI.
 */
type AVStereo3D struct {

	/**
	 * How views are packed within the video.
	 */
	Type AVStereo3DType

	/**
	 * Additional information about the frame packing.
	 */
	Flags ffcommon.FInt

	/**
	 * Determines which views are packed.
	 */
	View AVStereo3DView
}

/**
 * Allocate an AVStereo3D structure and set its fields to default values.
 * The resulting struct can be freed using av_freep().
 *
 * @return An AVStereo3D filled with default values or NULL on failure.
 */
//AVStereo3D *av_stereo3d_alloc(void);
var avStereo3DAlloc func() *AVStereo3D
var avStereo3DAllocOnce sync.Once

func AvStereo3dAlloc() *AVStereo3D {
	avStereo3DAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avStereo3DAlloc, ffcommon.GetAvutilDll(), "av_stereo3d_alloc")
	})
	return avStereo3DAlloc()
}

/**
 * Allocate a complete AVFrameSideData and add it to the frame.
 *
 * @param frame The frame which side data is added to.
 *
 * @return The AVStereo3D structure to be filled by caller.
 */
//AVStereo3D *av_stereo3d_create_side_data(AVFrame *frame);
var avStereo3DCreateSideData func(frame *AVFrame) *AVStereo3D
var avStereo3DCreateSideDataOnce sync.Once

func (frame *AVFrame) AvStereo3dCreateSideData() *AVStereo3D {
	avStereo3DCreateSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avStereo3DCreateSideData, ffcommon.GetAvutilDll(), "av_stereo3d_create_side_data")
	})
	return avStereo3DCreateSideData(frame)
}

/**
 * Provide a human-readable name of a given stereo3d type.
 *
 * @param type The input stereo3d type value.
 *
 * @return The name of the stereo3d value, or "unknown".
 */
//const char *av_stereo3d_type_name(unsigned int type);
var avStereo3DTypeName func(type0 ffcommon.FUnsignedInt) ffcommon.FConstCharP
var avStereo3DTypeNameOnce sync.Once

func AvStereo3dTypeName(type0 ffcommon.FUnsignedInt) ffcommon.FConstCharP {
	avStereo3DTypeNameOnce.Do(func() {
		purego.RegisterLibFunc(&avStereo3DTypeName, ffcommon.GetAvutilDll(), "av_stereo3d_type_name")
	})
	return avStereo3DTypeName(type0)
}

/**
 * Get the AVStereo3DType form a human-readable name.
 *
 * @param name The input string.
 *
 * @return The AVStereo3DType value, or -1 if not found.
 */
//int av_stereo3d_from_name(const char *name);
var avStereo3DFromName func(name ffcommon.FConstCharP) ffcommon.FInt
var avStereo3DFromNameOnce sync.Once

func AvStereo3dFromName(name ffcommon.FConstCharP) ffcommon.FInt {
	avStereo3DFromNameOnce.Do(func() {
		purego.RegisterLibFunc(&avStereo3DFromName, ffcommon.GetAvutilDll(), "av_stereo3d_from_name")
	})
	return avStereo3DFromName(name)
}

/**
 * @}
 * @}
 */

//#endif /* AVUTIL_STEREO3D_H */
