package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * copyright (c) 2007 Michael Niedermayer <michaelni@gmx.at>
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

//#ifndef AVUTIL_AES_H
//#define AVUTIL_AES_H
//
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"
//
///**
// * @defgroup lavu_aes AES
// * @ingroup lavu_crypto
// * @{
// */
//
//extern const int av_aes_size;

// struct AVAES;
type AVAES struct {
}

/**
 * Allocate an AVAES context.
 */
//struct AVAES *av_aes_alloc(void);
var avAesAlloc func() *AVAES
var avAesAllocOnce sync.Once

func AvAesAlloc() *AVAES {
	avAesAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avAesAlloc, ffcommon.GetAvutilDll(), "av_aes_alloc")
	})
	return avAesAlloc()
}

/**
 * Initialize an AVAES context.
 * @param key_bits 128, 192 or 256
 * @param decrypt 0 for encryption, 1 for decryption
 */
//int av_aes_init(struct AVAES *a, const uint8_t *key, int key_bits, int decrypt);
var avAesInit func(a *AVAES, key *ffcommon.FUint8T, key_bits, decrypt ffcommon.FUint) ffcommon.FInt
var avAesInitOnce sync.Once

func (a *AVAES) AvAesInit(key *ffcommon.FUint8T, key_bits, decrypt ffcommon.FUint) ffcommon.FInt {
	avAesInitOnce.Do(func() {
		purego.RegisterLibFunc(&avAesInit, ffcommon.GetAvutilDll(), "av_aes_init")
	})
	return avAesInit(a, key, key_bits, decrypt)
}

/**
 * Encrypt or decrypt a buffer using a previously initialized context.
 * @param count number of 16 byte blocks
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_aes_crypt(struct AVAES *a, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
var avAesCrypt func(a *AVAES, dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt)
var avAesCryptOnce sync.Once

func (a *AVAES) AvAesCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	avAesCryptOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCrypt, ffcommon.GetAvutilDll(), "av_aes_crypt")
	})
	avAesCrypt(a, dst, src, count, iv, decrypt)
}

/**
 * @}
 */

//#endif /* AVUTIL_AES_H */
