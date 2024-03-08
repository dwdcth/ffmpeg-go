package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Blowfish algorithm
 * Copyright (c) 2012 Samuel Pitoiset
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

//#ifndef AVUTIL_BLOWFISH_H
//#define AVUTIL_BLOWFISH_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_blowfish Blowfish
 * @ingroup lavu_crypto
 * @{
 */

const AV_BF_ROUNDS = 16

type AVBlowfish struct {
	P [AV_BF_ROUNDS + 2]ffcommon.FUint32T
	S [4][256]ffcommon.FUint32T
}

/**
 * Allocate an AVBlowfish context.
 */
//AVBlowfish *av_blowfish_alloc(void);
// purego func
var avBlowfishAlloc func() *AVBlowfish
var avBlowfishAllocOnce sync.Once

func AvBlowfishAlloc() (res *AVBlowfish) {
	avBlowfishAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avBlowfishAlloc, ffcommon.GetAvutilDll(), "av_blowfish_alloc")
	})
	return avBlowfishAlloc()
}

/**
 * Initialize an AVBlowfish context.
 *
 * @param ctx an AVBlowfish context
 * @param key a key
 * @param key_len length of the key
 */
//void av_blowfish_init(struct AVBlowfish *ctx, const uint8_t *key, int key_len);
// purego struct method
var avBlowfishInit func(ctx *AVBlowfish, key *ffcommon.FUint8T, keyLen ffcommon.FInt)
var avBlowfishInitOnce sync.Once

func (ctx *AVBlowfish) AvBlowfishInit(key *ffcommon.FUint8T, keyLen ffcommon.FInt) {
	avBlowfishInitOnce.Do(func() {
		purego.RegisterLibFunc(&avBlowfishInit, ffcommon.GetAvutilDll(), "av_blowfish_init")
	})
	avBlowfishInit(ctx, key, keyLen)
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param xl left four bytes halves of input to be encrypted
 * @param xr right four bytes halves of input to be encrypted
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt_ecb(struct AVBlowfish *ctx, uint32_t *xl, uint32_t *xr,
//int decrypt);
// purego struct method
var avBlowfishCryptEcb func(ctx *AVBlowfish, xl, xr *ffcommon.FUint32T, decrypt ffcommon.FInt)
var avBlowfishCryptEcbOnce sync.Once

func (ctx *AVBlowfish) AvBlowfishCryptEcb(xl, xr *ffcommon.FUint32T, decrypt ffcommon.FInt) {
	avBlowfishCryptEcbOnce.Do(func() {
		purego.RegisterLibFunc(&avBlowfishCryptEcb, ffcommon.GetAvutilDll(), "av_blowfish_crypt_ecb")
	})
	avBlowfishCryptEcb(ctx, xl, xr, decrypt)
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 *
 * @param ctx an AVBlowfish context
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param count number of 8 byte blocks
 * @param iv initialization vector for CBC mode, if NULL ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_blowfish_crypt(struct AVBlowfish *ctx, uint8_t *dst, const uint8_t *src,
//int count, uint8_t *iv, int decrypt);
// purego struct method
var avBlowfishCrypt func(ctx *AVBlowfish, dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt)
var avBlowfishCryptOnce sync.Once

func (ctx *AVBlowfish) AvBlowfishCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	avBlowfishCryptOnce.Do(func() {
		purego.RegisterLibFunc(&avBlowfishCrypt, ffcommon.GetAvutilDll(), "av_blowfish_crypt")
	})
	avBlowfishCrypt(ctx, dst, src, count, iv, decrypt)
}

/**
 * @}
 */

//#endif /* AVUTIL_BLOWFISH_H */
