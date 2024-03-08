package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * DES encryption/decryption
 * Copyright (c) 2007 Reimar Doeffinger
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

//#ifndef AVUTIL_DES_H
//#define AVUTIL_DES_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_des DES
 * @ingroup lavu_crypto
 * @{
 */

type AVDES struct {
	RoundKeys [3][16]ffcommon.FUint64T
	TripleDes ffcommon.FInt
}

/**
 * Allocate an AVDES context.
 */
//AVDES *av_des_alloc(void);
// purego func
var avDesAlloc func() *AVDES
var avDesAllocOnce sync.Once

func AvDesAlloc() (res *AVDES) {
	avDesAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avDesAlloc, ffcommon.GetAvutilDll(), "av_des_alloc")
	})
	res = avDesAlloc()
	return
}

/**
 * @brief Initializes an AVDES context.
 *
 * @param key_bits must be 64 or 192
 * @param decrypt 0 for encryption/CBC-MAC, 1 for decryption
 * @return zero on success, negative value otherwise
 */
//int av_des_init(struct AVDES *d, const uint8_t *key, int key_bits, int decrypt);
// purego struct method
var avDesInit func(d *AVDES, key *ffcommon.FUint8T, keyBits, decrypt ffcommon.FUint) ffcommon.FInt
var avDesInitOnce sync.Once

func (d *AVDES) AvDesInit(key *ffcommon.FUint8T, keyBits, decrypt ffcommon.FUint) ffcommon.FInt {
	avDesInitOnce.Do(func() {
		purego.RegisterLibFunc(&avDesInit, ffcommon.GetAvutilDll(), "av_des_init")
	})
	return avDesInit(d, key, keyBits, decrypt)
}

/**
 * @brief Encrypts / decrypts using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 * @param iv initialization vector for CBC mode, if NULL then ECB will be used,
 *           must be 8-byte aligned
 * @param decrypt 0 for encryption, 1 for decryption
 */
//void av_des_crypt(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
var avDesCrypt func(d *AVDES, dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt)
var avDesCryptOnce sync.Once

func (d *AVDES) AvDesCrypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	avDesCryptOnce.Do(func() {
		purego.RegisterLibFunc(&avDesCrypt, ffcommon.GetAvutilDll(), "av_des_crypt")
	})
	avDesCrypt(d, dst, src, count, iv, decrypt)
}

/**
 * @brief Calculates CBC-MAC using the DES algorithm.
 *
 * @param count number of 8 byte blocks
 * @param dst destination array, can be equal to src, must be 8-byte aligned
 * @param src source array, can be equal to dst, must be 8-byte aligned, may be NULL
 */
//void av_des_mac(struct AVDES *d, uint8_t *dst, const uint8_t *src, int count);

var avDesMac func(d *AVDES, dst, src *ffcommon.FUint8T, count ffcommon.FInt)
var avDesMacOnce sync.Once

func (d *AVDES) AvDesMac(dst, src *ffcommon.FUint8T, count ffcommon.FInt) {
	avDesMacOnce.Do(func() {
		purego.RegisterLibFunc(&avDesMac, ffcommon.GetAvutilDll(), "av_des_mac")
	})
	avDesMac(d, dst, src, count)
}

/**
 * @}
 */

//#endif /* AVUTIL_DES_H */
