package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (C) 2007 Michael Niedermayer <michaelni@gmx.at>
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
 * @ingroup lavu_sha
 * Public header for SHA-1 & SHA-256 hash function implementations.
 */

//#ifndef AVUTIL_SHA_H
//#define AVUTIL_SHA_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

/**
 * @defgroup lavu_sha SHA
 * @ingroup lavu_hash
 * SHA-1 and SHA-256 (Secure Hash Algorithm) hash function implementations.
 *
 * This module supports the following SHA hash functions:
 *
 * - SHA-1: 160 bits
 * - SHA-224: 224 bits, as a variant of SHA-2
 * - SHA-256: 256 bits, as a variant of SHA-2
 *
 * @see For SHA-384, SHA-512, and variants thereof, see @ref lavu_sha512.
 *
 * @{
 */

//extern const int av_sha_size;

// struct AVSHA;
type AVSHA struct {
}

/**
 * Allocate an AVSHA context.
 */
//struct AVSHA *av_sha_alloc(void);
// purego func
var avShaAlloc func() *AVRC4
var avShaAllocOnce sync.Once

func AvShaAlloc() (res *AVRC4) {
	avShaAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avShaAlloc, ffcommon.GetAvutilDll(), "av_sha_alloc")
	})
	res = avShaAlloc()
	return
}

/**
 * Initialize SHA-1 or SHA-2 hashing.
 *
 * @param context pointer to the function context (of size av_sha_size)
 * @param bits    number of bits in digest (SHA-1 - 160 bits, SHA-2 224 or 256 bits)
 * @return        zero if initialization succeeded, -1 otherwise
 */
//int av_sha_init(struct AVSHA* context, int bits);
// purego struct method
var avShaInit func(context *AVSHA, key *ffcommon.FUint8T, bits ffcommon.FInt) ffcommon.FInt
var avShaInitOnce sync.Once

func (context *AVSHA) AvShaInit(key *ffcommon.FUint8T, bits ffcommon.FInt) (res ffcommon.FInt) {
	avShaInitOnce.Do(func() {
		purego.RegisterLibFunc(&avShaInit, ffcommon.GetAvutilDll(), "av_sha_init")
	})
	res = avShaInit(context, key, bits)
	return
}

/**
 * Update hash value.
 *
 * @param ctx     hash function context
 * @param data    input data to update hash with
 * @param len     input data length
 */
//#if FF_API_CRYPTO_SIZE_T
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, unsigned int len);
//#else
//void av_sha_update(struct AVSHA *ctx, const uint8_t *data, size_t len);
//#endif
// purego struct method
var avShaUpdate func(context *AVSHA, data *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT)
var avShaUpdateOnce sync.Once

func (context *AVSHA) AvShaUpdate(data *ffcommon.FUint8T, len0 ffcommon.FUnsignedIntOrSizeT) {
	avShaUpdateOnce.Do(func() {
		purego.RegisterLibFunc(&avShaUpdate, ffcommon.GetAvutilDll(), "av_sha_update")
	})
	avShaUpdate(context, data, len0)
}

/**
 * Finish hashing and output digest value.
 *
 * @param context hash function context
 * @param digest  buffer where output digest value is stored
 */
//void av_sha_final(struct AVSHA* context, uint8_t *digest);
// purego struct method
var avShaFinal func(context *AVSHA, digest *ffcommon.FUint8T)
var avShaFinalOnce sync.Once

func (context *AVSHA) AvShaFinal(digest *ffcommon.FUint8T) {
	avShaFinalOnce.Do(func() {
		purego.RegisterLibFunc(&avShaFinal, ffcommon.GetAvutilDll(), "av_sha_final")
	})
	avShaFinal(context, digest)
}

/**
 * @}
 */

//#endif /* AVUTIL_SHA_H */
