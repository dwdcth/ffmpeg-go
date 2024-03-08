package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (C) 2013 Reimar Döffinger <Reimar.Doeffinger@gmx.de>
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
 * @ingroup lavu_murmur3
 * Public header for MurmurHash3 hash function implementation.
 */

//#ifndef AVUTIL_MURMUR3_H
//#define AVUTIL_MURMUR3_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "version.h"

/**
 * @defgroup lavu_murmur3 Murmur3
 * @ingroup lavu_hash
 * MurmurHash3 hash function implementation.
 *
 * MurmurHash3 is a non-cryptographic hash function, of which three
 * incompatible versions were created by its inventor Austin Appleby:
 *
 * - 32-bit output
 * - 128-bit output for 32-bit platforms
 * - 128-bit output for 64-bit platforms
 *
 * FFmpeg only implements the last variant: 128-bit output designed for 64-bit
 * platforms. Even though the hash function was designed for 64-bit platforms,
 * the function in reality works on 32-bit systems too, only with reduced
 * performance.
 *
 * @anchor lavu_murmur3_seedinfo
 * By design, MurmurHash3 requires a seed to operate. In response to this,
 * libavutil provides two functions for hash initiation, one that requires a
 * seed (av_murmur3_init_seeded()) and one that uses a fixed arbitrary integer
 * as the seed, and therefore does not (av_murmur3_init()).
 *
 * To make hashes comparable, you should provide the same seed for all calls to
 * this hash function -- if you are supplying one yourself, that is.
 *
 * @{
 */

/**
 * Allocate an AVMurMur3 hash context.
 *
 * @return Uninitialized hash context or `NULL` in case of error
 */
//struct AVMurMur3 *av_murmur3_alloc(void);
type AVMurMur3 struct {
}

// purego func
var avMurmur3Alloc func() *AVMurMur3
var avMurmur3AllocOnce sync.Once

func AvMurmur3Alloc() (res *AVMurMur3) {
	avMurmur3AllocOnce.Do(func() {
		purego.RegisterLibFunc(&avMurmur3Alloc, ffcommon.GetAvutilDll(), "av_murmur3_alloc")
	})
	res = avMurmur3Alloc()
	return
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context with a seed.
 *
 * @param[out] c    Hash context
 * @param[in]  seed Random seed
 *
 * @see av_murmur3_init()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init_seeded(struct AVMurMur3 *c, uint64_t seed);
// purego struct method
var avMurmur3InitSeeded func(c *AVMurMur3, seed ffcommon.FUint64T)
var avMurmur3InitSeededOnce sync.Once

func (c *AVMurMur3) AvMurmur3InitSeeded(seed ffcommon.FUint64T) {
	avMurmur3InitSeededOnce.Do(func() {
		purego.RegisterLibFunc(&avMurmur3InitSeeded, ffcommon.GetAvutilDll(), "av_murmur3_init_seeded")
	})
	avMurmur3InitSeeded(c, seed)
}

/**
 * Initialize or reinitialize an AVMurMur3 hash context.
 *
 * Equivalent to av_murmur3_init_seeded() with a built-in seed.
 *
 * @param[out] c    Hash context
 *
 * @see av_murmur3_init_seeded()
 * @see @ref lavu_murmur3_seedinfo "Detailed description" on a discussion of
 * seeds for MurmurHash3.
 */
//void av_murmur3_init(struct AVMurMur3 *c);
var avMurmur3Init func(c *AVMurMur3)
var avMurmur3InitOnce sync.Once

func (c *AVMurMur3) AvMurmur3Init() {
	avMurmur3InitOnce.Do(func() {
		purego.RegisterLibFunc(&avMurmur3Init, ffcommon.GetAvutilDll(), "av_murmur3_init")
	})
	avMurmur3Init(c)
}

/**
 * Update hash context with new data.
 *
 * @param[out] c    Hash context
 * @param[in]  src  Input data to update hash with
 * @param[in]  len  Number of bytes to read from `src`
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, int len);
//#else
//void av_murmur3_update(struct AVMurMur3 *c, const uint8_t *src, size_t len);
//#endif
// purego struct method
// purego struct method
var avMurmur3Update func(c *AVMurMur3, src *ffcommon.FUint8T, len0 ffcommon.FIntOrSizeT) ffcommon.FCharP
var avMurmur3UpdateOnce sync.Once

func (c *AVMurMur3) AvMurmur3Update(src *ffcommon.FUint8T, len0 ffcommon.FIntOrSizeT) (res ffcommon.FCharP) {
	avMurmur3UpdateOnce.Do(func() {
		purego.RegisterLibFunc(&avMurmur3Update, ffcommon.GetAvutilDll(), "av_murmur3_update")
	})
	res = avMurmur3Update(c, src, len0)
	return
}

/**
 * Finish hashing and output digest value.
 *
 * @param[in,out] c    Hash context
 * @param[out]    dst  Buffer where output digest value is stored
 */
//void av_murmur3_final(struct AVMurMur3 *c, uint8_t dst[16]);
// purego struct method
var avMurmur3Final func(c *AVMurMur3, dst *[16]ffcommon.FUint8T)
var avMurmur3FinalOnce sync.Once

func (c *AVMurMur3) AvMurmur3Final(dst [16]ffcommon.FUint8T) {
	avMurmur3FinalOnce.Do(func() {
		purego.RegisterLibFunc(&avMurmur3Final, ffcommon.GetAvutilDll(), "av_murmur3_final")
	})
	avMurmur3Final(c, &dst)
}

/**
 * @}
 */

//#endif /* AVUTIL_MURMUR3_H */
