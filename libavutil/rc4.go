package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * RC4 encryption/decryption/pseudo-random number generator
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

//#ifndef AVUTIL_RC4_H
//#define AVUTIL_RC4_H
//
//#include <stdint.h>

/**
 * @defgroup lavu_rc4 RC4
 * @ingroup lavu_crypto
 * @{
 */

type AVRC4 struct {
	State [256]ffcommon.FUint8T
	X, Y  ffcommon.FInt
}

/**
 * Allocate an AVRC4 context.
 */
//AVRC4 *av_rc4_alloc(void);
// purego func
var avRc4Alloc func() *AVRC4
var avRc4AllocOnce sync.Once

func AvRc4Alloc() (res *AVRC4) {
	avRc4AllocOnce.Do(func() {
		purego.RegisterLibFunc(&avRc4Alloc, ffcommon.GetAvutilDll(), "av_rc4_alloc")
	})
	res = avRc4Alloc()
	return
}

/**
 * @brief Initializes an AVRC4 context.
 *
 * @param key_bits must be a multiple of 8
 * @param decrypt 0 for encryption, 1 for decryption, currently has no effect
 * @return zero on success, negative value otherwise
 */
//int av_rc4_init(struct AVRC4 *d, const uint8_t *key, int key_bits, int decrypt);
// purego struct method
var avRc4Init func(d *AVRC4, key *ffcommon.FUint8T, keyBits, decrypt ffcommon.FInt) ffcommon.FInt
var avRc4InitOnce sync.Once

func (d *AVRC4) AvRc4Init(key *ffcommon.FUint8T, keyBits, decrypt ffcommon.FInt) (res ffcommon.FInt) {
	avRc4InitOnce.Do(func() {
		purego.RegisterLibFunc(&avRc4Init, ffcommon.GetAvutilDll(), "av_rc4_init")
	})
	res = avRc4Init(d, key, keyBits, decrypt)
	return
}

/**
 * @brief Encrypts / decrypts using the RC4 algorithm.
 *
 * @param count number of bytes
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst, may be NULL
 * @param iv not (yet) used for RC4, should be NULL
 * @param decrypt 0 for encryption, 1 for decryption, not (yet) used
 */
//void av_rc4_crypt(struct AVRC4 *d, uint8_t *dst, const uint8_t *src, int count, uint8_t *iv, int decrypt);
// purego struct method
var avRc4Crypt func(d *AVRC4, dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt)
var avRc4CryptOnce sync.Once

func (d *AVRC4) AvRc4Crypt(dst, src *ffcommon.FUint8T, count ffcommon.FInt, iv *ffcommon.FUint8T, decrypt ffcommon.FInt) {
	avRc4CryptOnce.Do(func() {
		purego.RegisterLibFunc(&avRc4Crypt, ffcommon.GetAvutilDll(), "av_rc4_crypt")
	})
	avRc4Crypt(d, dst, src, count, iv, decrypt)
}

/**
 * @}
 */

//#endif /* AVUTIL_RC4_H */
