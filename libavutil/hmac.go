package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (C) 2012 Martin Storsjo
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

//#ifndef AVUTIL_HMAC_H
//#define AVUTIL_HMAC_H
//
//#include <stdint.h>
//
//#include "version.h"
/**
 * @defgroup lavu_hmac HMAC
 * @ingroup lavu_crypto
 * @{
 */
type AVHMACType = int32

const (
	AV_HMAC_MD5 = iota
	AV_HMAC_SHA1
	AV_HMAC_SHA224
	AV_HMAC_SHA256
	AV_HMAC_SHA384
	AV_HMAC_SHA512
)

// typedef struct AVHMAC AVHMAC;
type AVHMAC struct {
}

/**
 * Allocate an AVHMAC context.
 * @param type The hash function used for the HMAC.
 */
//AVHMAC *av_hmac_alloc(enum AVHMACType type);
// purego func
var avHmacAlloc func(type0 AVHMACType) *AVHMAC
var avHmacAllocOnce sync.Once

func AvHmacAlloc(type0 AVHMACType) (res *AVHMAC) {
	avHmacAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacAlloc, ffcommon.GetAvutilDll(), "av_hmac_alloc")
	})
	res = avHmacAlloc(type0)
	return
}

/**
 * Free an AVHMAC context.
 * @param ctx The context to free, may be NULL
 */
//void av_hmac_free(AVHMAC *ctx);
// purego struct method
var avHmacFree func(ctx *AVHMAC)
var avHmacFreeOnce sync.Once

func (ctx *AVHMAC) AvHmacFree() {
	avHmacFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacFree, ffcommon.GetAvutilDll(), "av_hmac_free")
	})
	avHmacFree(ctx)
}

/**
 * Initialize an AVHMAC context with an authentication key.
 * @param ctx    The HMAC context
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 */
//void av_hmac_init(AVHMAC *ctx, const uint8_t *key, unsigned int keylen);
// purego struct method
var avHmacInit func(ctx *AVHMAC, key *ffcommon.FUint8T, keylen ffcommon.FInt)
var avHmacInitOnce sync.Once

func (ctx *AVHMAC) AvHmacInit(key *ffcommon.FUint8T, keylen ffcommon.FInt) {
	avHmacInitOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacInit, ffcommon.GetAvutilDll(), "av_hmac_init")
	})
	avHmacInit(ctx, key, keylen)
}

/**
 * Hash data with the HMAC.
 * @param ctx  The HMAC context
 * @param data The data to hash
 * @param len  The length of the data, in bytes
 */
//void av_hmac_update(AVHMAC *ctx, const uint8_t *data, unsigned int len);
// purego struct method
var avHmacUpdate func(ctx *AVHMAC, data *ffcommon.FUint8T, keylen ffcommon.FInt)
var avHmacUpdateOnce sync.Once

func (ctx *AVHMAC) AvHmacUpdate(data *ffcommon.FUint8T, keylen ffcommon.FInt) {
	avHmacUpdateOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacUpdate, ffcommon.GetAvutilDll(), "av_hmac_update")
	})
	avHmacUpdate(ctx, data, keylen)
}

/**
 * Finish hashing and output the HMAC digest.
 * @param ctx    The HMAC context
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_final(AVHMAC *ctx, uint8_t *out, unsigned int outlen);
// purego struct method
var avHmacFinal func(ctx *AVHMAC, out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) ffcommon.FInt
var avHmacFinalOnce sync.Once

func (ctx *AVHMAC) AvHmacFinal(out *ffcommon.FUint8T, outlen ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	avHmacFinalOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacFinal, ffcommon.GetAvutilDll(), "av_hmac_final")
	})
	res = avHmacFinal(ctx, out, outlen)
	return
}

/**
 * Hash an array of data with a key.
 * @param ctx    The HMAC context
 * @param data   The data to hash
 * @param len    The length of the data, in bytes
 * @param key    The authentication key
 * @param keylen The length of the key, in bytes
 * @param out    The output buffer to write the digest into
 * @param outlen The length of the out buffer, in bytes
 * @return       The number of bytes written to out, or a negative error code.
 */
//int av_hmac_calc(AVHMAC *ctx, const uint8_t *data, unsigned int len,
//const uint8_t *key, unsigned int keylen,
//uint8_t *out, unsigned int outlen);
// purego struct method
var avHmacCalc func(ctx *AVHMAC, data *ffcommon.FUint8T, len0, key *ffcommon.FUnsignedInt,
	out *ffcommon.FUint8T, outlen *ffcommon.FUnsignedInt) ffcommon.FInt
var avHmacCalcOnce sync.Once

func (ctx *AVHMAC) AvHmacCalc(data *ffcommon.FUint8T, len0, key *ffcommon.FUnsignedInt,
	out *ffcommon.FUint8T, outlen *ffcommon.FUnsignedInt) (res ffcommon.FInt) {
	avHmacCalcOnce.Do(func() {
		purego.RegisterLibFunc(&avHmacCalc, ffcommon.GetAvutilDll(), "av_hmac_calc")
	})
	res = avHmacCalc(ctx, data, len0, key, out, outlen)
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_HMAC_H */
