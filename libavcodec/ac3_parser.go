package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * AC-3 parser prototypes
 * Copyright (c) 2003 Fabrice Bellard
 * Copyright (c) 2003 Michael Niedermayer
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

//#ifndef AVCODEC_AC3_PARSER_H
//#define AVCODEC_AC3_PARSER_H
//
//#include <stddef.h>
//#include <stdint.h>

/**
 * Extract the bitstream ID and the frame size from AC-3 data.
 */
//int av_ac3_parse_header(const uint8_t *buf, size_t size,
//uint8_t *bitstream_id, uint16_t *frame_size);
var av_ac3_parse_header func(buf *ffcommon.FUint8T, size ffcommon.FSizeT, bitstream_id *ffcommon.FUint8T, frame_size *ffcommon.FUint16T) ffcommon.FInt
var av_ac3_parse_header_once sync.Once

func AvAc3ParseHeader(buf *ffcommon.FUint8T, size ffcommon.FSizeT, bitstream_id *ffcommon.FUint8T, frame_size *ffcommon.FUint16T) (res ffcommon.FInt) {
	av_ac3_parse_header_once.Do(func() {
		purego.RegisterLibFunc(&av_ac3_parse_header, ffcommon.GetAvcodecDll(), "av_ac3_parse_header")
	})
	res = av_ac3_parse_header(buf, size, bitstream_id, frame_size)
	return
}

//#endif /* AVCODEC_AC3_PARSER_H */
