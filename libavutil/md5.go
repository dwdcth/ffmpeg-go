package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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
 * @ingroup lavu_md5
 * Public header for MD5 hash function implementation.
 */

//#ifndef AVUTIL_MD5_H
//#define AVUTIL_MD5_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_md5 MD5
 * @ingroup lavu_hash
 * MD5 hash function implementation.
 *
 * @{
 */

// struct AVMD5;
//
//extern const int av_md5_size;
type AVMD5 struct {
}

/**
 * Allocate an AVMD5 context.
 */
//struct AVMD5 *av_md5_alloc(void);
// purego func
var avMd5Alloc func() *AVMD5
var avMd5AllocOnce sync.Once

func AvMd5Alloc() (res *AVMD5) {
	avMd5AllocOnce.Do(func() {
		purego.RegisterLibFunc(&avMd5Alloc, ffcommon.GetAvutilDll(), "av_md5_alloc")
	})
	res = avMd5Alloc()
	return
}

/**
 * Initialize MD5 hashing.
 *
 * @param ctx pointer to the function context (of size av_md5_size)
 */
//void av_md5_init(struct AVMD5 *ctx);
// purego struct method
var avMd5Init func(ctx *AVMD5)
var avMd5InitOnce sync.Once

func (ctx *AVMD5) AvMd5Init() {
	avMd5InitOnce.Do(func() {
		purego.RegisterLibFunc(&avMd5Init, ffcommon.GetAvutilDll(), "av_md5_init")
	})
	avMd5Init(ctx)
}

/**
 * Update hash value.
 *
 * @param ctx hash function context
 * @param src input data to update hash with
 * @param len input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, int len);
//#else
//void av_md5_update(struct AVMD5 *ctx, const uint8_t *src, size_t len);
//#endif
// purego struct method
var avMd5Update func(ctx *AVMD5, src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT)
var avMd5UpdateOnce sync.Once

func (ctx *AVMD5) AvMd5Update(src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	avMd5UpdateOnce.Do(func() {
		purego.RegisterLibFunc(&avMd5Update, ffcommon.GetAvutilDll(), "av_md5_update")
	})
	avMd5Update(ctx, src, len0)
}

/**
 * Finish hashing and output digest value.
 *
 * @param ctx hash function context
 * @param dst buffer where output digest value is stored
 */
//void av_md5_final(struct AVMD5 *ctx, uint8_t *dst);
// purego struct method
var avMd5Final func(ctx *AVMD5, dst *ffcommon.FUint8T)
var avMd5FinalOnce sync.Once

func (ctx *AVMD5) AvMd5Final(dst *ffcommon.FUint8T) {
	avMd5FinalOnce.Do(func() {
		purego.RegisterLibFunc(&avMd5Final, ffcommon.GetAvutilDll(), "av_md5_final")
	})
	avMd5Final(ctx, dst)
}

/**
 * Hash an array of data.
 *
 * @param dst The output buffer to write the digest into
 * @param src The data to hash
 * @param len The length of the data, in bytes
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_md5_sum(uint8_t *dst, const uint8_t *src, const int len);
//#else
//void av_md5_sum(uint8_t *dst, const uint8_t *src, size_t len);
//#endif
// purego func
var avMd5Sum func(dst, src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT)
var avMd5SumOnce sync.Once

func AvMd5Sum(dst, src *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	avMd5SumOnce.Do(func() {
		purego.RegisterLibFunc(&avMd5Sum, ffcommon.GetAvutilDll(), "av_md5_sum")
	})
	avMd5Sum(dst, src, len0)
}

/**
 * @}
 */

//#endif /* AVUTIL_MD5_H */
