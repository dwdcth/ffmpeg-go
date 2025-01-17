package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
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

/**
 * @file
 * A public API for Vorbis parsing
 *
 * Determines the duration for each packet.
 */

//#ifndef AVCODEC_VORBIS_PARSER_H
//#define AVCODEC_VORBIS_PARSER_H
//
//#include <stdint.h>

// typedef struct AVVorbisParseContext AVVorbisParseContext;
type AVVorbisParseContext struct {
}

/**
 * Allocate and initialize the Vorbis parser using headers in the extradata.
 */
//AVVorbisParseContext *av_vorbis_parse_init(const uint8_t *extradata,
//                                           int extradata_size);
var avVorbisParseInit func(extradata *ffcommon.FUint8T, extradata_size ffcommon.FInt) *AVVorbisParseContext
var avVorbisParseInitOnce sync.Once

func AvVorbisParseInit(extradata *ffcommon.FUint8T, extradata_size ffcommon.FInt) *AVVorbisParseContext {
	avVorbisParseInitOnce.Do(func() {
		purego.RegisterLibFunc(&avVorbisParseInit, ffcommon.GetAvcodecDll(), "av_vorbis_parse_init")
	})
	return avVorbisParseInit(extradata, extradata_size)
}

/**
 * Free the parser and everything associated with it.
 */
//void av_vorbis_parse_free(AVVorbisParseContext **s);
var avVorbisParseFree func(s **AVVorbisParseContext)
var avVorbisParseFreeOnce sync.Once

func AvVorbisParseFree(s **AVVorbisParseContext) {
	avVorbisParseFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avVorbisParseFree, ffcommon.GetAvcodecDll(), "av_vorbis_parse_free")
	})
	avVorbisParseFree(s)
}

const VORBIS_FLAG_HEADER = 0x00000001
const VORBIS_FLAG_COMMENT = 0x00000002
const VORBIS_FLAG_SETUP = 0x00000004

/**
 * Get the duration for a Vorbis packet.
 *
 * If @p flags is @c NULL,
 * special frames are considered invalid.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 * @param flags    flags for special frames
 */
//int av_vorbis_parse_frame_flags(AVVorbisParseContext *s, const uint8_t *buf,
//                                int buf_size, int *flags);
var avVorbisParseFrameFlags func(s *AVVorbisParseContext, buf *ffcommon.FUint8T, buf_size ffcommon.FInt, flags *ffcommon.FInt) ffcommon.FInt
var avVorbisParseFrameFlagsOnce sync.Once

func (s *AVVorbisParseContext) AvVorbisParseFrameFlags(buf *ffcommon.FUint8T, buf_size ffcommon.FInt, flags *ffcommon.FInt) ffcommon.FInt {
	avVorbisParseFrameFlagsOnce.Do(func() {
		purego.RegisterLibFunc(&avVorbisParseFrameFlags, ffcommon.GetAvcodecDll(), "av_vorbis_parse_frame_flags")
	})
	return avVorbisParseFrameFlags(s, buf, buf_size, flags)
}

/**
 * Get the duration for a Vorbis packet.
 *
 * @param s        Vorbis parser context
 * @param buf      buffer containing a Vorbis frame
 * @param buf_size size of the buffer
 */
//int av_vorbis_parse_frame(AVVorbisParseContext *s, const uint8_t *buf,
//                          int buf_size);
var avVorbisParseFrame func(s *AVVorbisParseContext, buf *ffcommon.FUint8T, buf_size ffcommon.FInt) ffcommon.FInt
var avVorbisParseFrameOnce sync.Once

func (s *AVVorbisParseContext) AvVorbisParseFrame(buf *ffcommon.FUint8T, buf_size ffcommon.FInt) ffcommon.FInt {
	avVorbisParseFrameOnce.Do(func() {
		purego.RegisterLibFunc(&avVorbisParseFrame, ffcommon.GetAvcodecDll(), "av_vorbis_parse_frame")
	})
	return avVorbisParseFrame(s, buf, buf_size)
}

// void av_vorbis_parse_reset(AVVorbisParseContext *s);
var avVorbisParseReset func(s *AVVorbisParseContext)
var avVorbisParseResetOnce sync.Once

func (s *AVVorbisParseContext) AvVorbisParseReset() {
	avVorbisParseResetOnce.Do(func() {
		purego.RegisterLibFunc(&avVorbisParseReset, ffcommon.GetAvcodecDll(), "av_vorbis_parse_reset")
	})
	avVorbisParseReset(s)
}

//#endif /* AVCODEC_VORBIS_PARSER_H */
