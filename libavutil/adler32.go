package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * copyright (c) 2006 Mans Rullgard
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
 * @ingroup lavu_adler32
 * Public header for Adler-32 hash function implementation.
 */

//#ifndef AVUTIL_ADLER32_H
//#define AVUTIL_ADLER32_H
//
//#include <stddef.h>
//#include <stdint.h>
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_adler32 Adler-32
 * @ingroup lavu_hash
 * Adler-32 hash function implementation.
 *
 * @{
 */

//#if FF_API_CRYPTO_SIZE_T
//typedef unsigned long AVAdler;
//#else
//typedef uint32_t AVAdler;
//#endif

/**
 * Calculate the Adler32 checksum of a buffer.
 *
 * Passing the return value to a subsequent av_adler32_update() call
 * allows the checksum of multiple buffers to be calculated as though
 * they were concatenated.
 *
 * @param adler initial checksum value
 * @param buf   pointer to input buffer
 * @param len   size of input buffer
 * @return      updated checksum
 */
//AVAdler av_adler32_update(AVAdler adler, const uint8_t *buf,
//#if FF_API_CRYPTO_SIZE_T
//unsigned int len) av_pure;
//#else
//size_t len) av_pure;
//#endif
var avAdler32Update func(adler ffcommon.FAVAdler, buf *ffcommon.FUint8T, av_pure ffcommon.FUnsignedIntOrSizeT) ffcommon.FAVAdler
var avAdler32UpdateOnce sync.Once

func AvAdler32Update(adler ffcommon.FAVAdler, buf *ffcommon.FUint8T, av_pure ffcommon.FUnsignedIntOrSizeT) ffcommon.FAVAdler {
	avAdler32UpdateOnce.Do(func() {
		purego.RegisterLibFunc(&avAdler32Update, ffcommon.GetAvutilDll(), "av_adler32_update")
	})
	return avAdler32Update(adler, buf, av_pure)
}

/**
 * @}
 */

//#endif /* AVUTIL_ADLER32_H */
