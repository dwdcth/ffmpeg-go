package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2011 Mans Rullgard
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

//#ifndef AVUTIL_INTFLOAT_H
//#define AVUTIL_INTFLOAT_H
//
//#include <stdint.h>
//#include "attributes.h"
//
//union av_intfloat32 {
//uint32_t i;
//float    f;
//};
//
//union av_intfloat64 {
//uint64_t i;
//double   f;
//};

/**
 * Reinterpret a 32-bit integer as a float.
 */
//static av_always_inline float av_int2float(uint32_t i)
//{
//union av_intfloat32 v;
//v.i = i;
//return v.f;
//}
//todo
// purego func
var avInt2float func() string
var avInt2floatOnce sync.Once

func AvInt2float() (res ffcommon.FCharP) {
	avInt2floatOnce.Do(func() {
		purego.RegisterLibFunc(&avInt2float, ffcommon.GetAvutilDll(), "av_int2float")
	})
	res = avInt2float()
	return
}

/**
 * Reinterpret a float as a 32-bit integer.
 */
//static av_always_inline uint32_t av_float2int(float f)
//{
//union av_intfloat32 v;
//v.f = f;
//return v.i;
//}
//todo
// purego func
var avFloat2int func() string
var avFloat2intOnce sync.Once

func AvFloat2int() (res ffcommon.FCharP) {
	avFloat2intOnce.Do(func() {
		purego.RegisterLibFunc(&avFloat2int, ffcommon.GetAvutilDll(), "av_float2int")
	})
	res = avFloat2int()
	return
}

/**
 * Reinterpret a 64-bit integer as a double.
 */
//static av_always_inline double av_int2double(uint64_t i)
//{
//union av_intfloat64 v;
//v.i = i;
//return v.f;
//}
//todo
// purego func
var avInt2double func() string
var avInt2doubleOnce sync.Once

func AvInt2double() (res ffcommon.FCharP) {
	avInt2doubleOnce.Do(func() {
		purego.RegisterLibFunc(&avInt2double, ffcommon.GetAvutilDll(), "av_int2double")
	})
	res = avInt2double()
	return
}

/**
 * Reinterpret a double as a 64-bit integer.
 */
//static av_always_inline uint64_t av_double2int(double f)
//{
//union av_intfloat64 v;
//v.f = f;
//return v.i;
//}
//todo
// purego func
var avDouble2int func() string
var avDouble2intOnce sync.Once

func AvDouble2int() (res ffcommon.FCharP) {
	avDouble2intOnce.Do(func() {
		purego.RegisterLibFunc(&avDouble2int, ffcommon.GetAvutilDll(), "av_double2int")
	})
	res = avDouble2int()
	return
}

//#endif /* AVUTIL_INTFLOAT_H */
