package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * AES-CTR cipher
 * Copyright (c) 2015 Eran Kornblau <erankor at gmail dot com>
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

//#ifndef AVUTIL_AES_CTR_H
//#define AVUTIL_AES_CTR_H
//
//#include <stdint.h>
//
//#include "attributes.h"
//#include "version.h"

const AES_CTR_KEY_SIZE = (16)
const AES_CTR_IV_SIZE = (8)

// struct AVAESCTR;
type AVAESCTR struct {
}

/**
 * Allocate an AVAESCTR context.
 */
//struct AVAESCTR *av_aes_ctr_alloc(void);
// purego func
var avAesCtrAlloc func() *AVAESCTR
var avAesCtrAllocOnce sync.Once

func AvAesCtrAlloc() (res *AVAESCTR) {
	avAesCtrAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrAlloc, ffcommon.GetAvutilDll(), "av_aes_ctr_alloc")
	})
	return avAesCtrAlloc()
}

/**
 * Initialize an AVAESCTR context.
 * @param key encryption key, must have a length of AES_CTR_KEY_SIZE
 */
//int av_aes_ctr_init(struct AVAESCTR *a, const uint8_t *key);
// purego struct method
var avAesCtrInit func(a *AVAESCTR, key *ffcommon.FUint8T) ffcommon.FInt
var avAesCtrInitOnce sync.Once

func (a *AVAESCTR) AvAesCtrInit(key *ffcommon.FUint8T) (res ffcommon.FInt) {
	avAesCtrInitOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrInit, ffcommon.GetAvutilDll(), "av_aes_ctr_init")
	})
	return avAesCtrInit(a, key)
}

/**
 * Release an AVAESCTR context.
 */
//void av_aes_ctr_free(struct AVAESCTR *a);
// purego struct method
var avAesCtrFree func(a *AVAESCTR)
var avAesCtrFreeOnce sync.Once

func (a *AVAESCTR) AvAesCtrFree() {
	avAesCtrFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrFree, ffcommon.GetAvutilDll(), "av_aes_ctr_free")
	})
	avAesCtrFree(a)
}

/**
 * Process a buffer using a previously initialized context.
 * @param dst destination array, can be equal to src
 * @param src source array, can be equal to dst
 * @param size the size of src and dst
 */
//void av_aes_ctr_crypt(struct AVAESCTR *a, uint8_t *dst, const uint8_t *src, int size);
// purego struct method
var avAesCtrCrypt func(a *AVAESCTR, dst, src *ffcommon.FUint8T, size ffcommon.FUint)
var avAesCtrCryptOnce sync.Once

func (a *AVAESCTR) AvAesCtrCrypt(dst, src *ffcommon.FUint8T, size ffcommon.FUint) {
	avAesCtrCryptOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrCrypt, ffcommon.GetAvutilDll(), "av_aes_ctr_crypt")
	})
	avAesCtrCrypt(a, dst, src, size)
}

/**
 * Get the current iv
 */
//const uint8_t* av_aes_ctr_get_iv(struct AVAESCTR *a);
// purego struct method
var avAesCtrGetIv func(a *AVAESCTR) *ffcommon.FUint8T
var avAesCtrGetIvOnce sync.Once

func (a *AVAESCTR) AvAesCtrGetIv() (res *ffcommon.FUint8T) {
	avAesCtrGetIvOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrGetIv, ffcommon.GetAvutilDll(), "av_aes_ctr_get_iv")
	})
	return avAesCtrGetIv(a)
}

/**
 * Generate a random iv
 */
//void av_aes_ctr_set_random_iv(struct AVAESCTR *a);
// purego struct method
var avAesCtrSetRandomIv func(a *AVAESCTR)
var avAesCtrSetRandomIvOnce sync.Once

func (a *AVAESCTR) AvAesCtrSetRandomIv() {
	avAesCtrSetRandomIvOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrSetRandomIv, ffcommon.GetAvutilDll(), "av_aes_ctr_set_random_iv")
	})
	avAesCtrSetRandomIv(a)
}

/**
 * Forcefully change the 8-byte iv
 */
//void av_aes_ctr_set_iv(struct AVAESCTR *a, const uint8_t* iv);
// purego struct method
var avAesCtrSetIv func(a *AVAESCTR, iv *ffcommon.FUint8T)
var avAesCtrSetIvOnce sync.Once

func (a *AVAESCTR) AvAesCtrSetIv(iv *ffcommon.FUint8T) {
	avAesCtrSetIvOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrSetIv, ffcommon.GetAvutilDll(), "av_aes_ctr_set_iv")
	})
	avAesCtrSetIv(a, iv)
}

/**
 * Forcefully change the "full" 16-byte iv, including the counter
 */
//void av_aes_ctr_set_full_iv(struct AVAESCTR *a, const uint8_t* iv);
// purego struct method
var avAesCtrSetFullIv func(a *AVAESCTR, iv *ffcommon.FUint8T)
var avAesCtrSetFullIvOnce sync.Once

func (a *AVAESCTR) AvAesCtrSetFullIv(iv *ffcommon.FUint8T) {
	avAesCtrSetFullIvOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrSetFullIv, ffcommon.GetAvutilDll(), "av_aes_ctr_set_full_iv")
	})
	avAesCtrSetFullIv(a, iv)
}

/**
 * Increment the top 64 bit of the iv (performed after each frame)
 */
//void av_aes_ctr_increment_iv(struct AVAESCTR *a);
// purego struct method
var avAesCtrIncrementIv func(a *AVAESCTR)
var avAesCtrIncrementIvOnce sync.Once

func (a *AVAESCTR) AvAesCtrIncrementIv() {
	avAesCtrIncrementIvOnce.Do(func() {
		purego.RegisterLibFunc(&avAesCtrIncrementIv, ffcommon.GetAvutilDll(), "av_aes_ctr_increment_iv")
	})
	avAesCtrIncrementIv(a)
}

/**
 * @}
 */

//#endif /* AVUTIL_AES_CTR_H */
