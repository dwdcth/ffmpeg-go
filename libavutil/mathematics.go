package libavutil

import (
	"sync"
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * copyright (c) 2005-2012 Michael Niedermayer <michaelni@gmx.at>
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
 * @addtogroup lavu_math
 * Mathematical utilities for working with timestamp and time base.
 */

//#ifndef AVUTIL_MATHEMATICS_H
//const AVUTIL_MATHEMATICS_H
//
//#include <stdint.h>
//#include <math.h>
//#include "attributes.h"
//#include "rational.h"
//#include "intfloat.h"

// #ifndef M_E
const M_E = 2.7182818284590452354 /* e */
// #endif
// #ifndef M_LN2
const M_LN2 = 0.69314718055994530942 /* log_e 2 */
// #endif
// #ifndef M_LN10
const M_LN10 = 2.30258509299404568402 /* log_e 10 */
// #endif
// #ifndef M_LOG2_10
const M_LOG2_10 = 3.32192809488736234787 /* log_2 10 */
// #endif
// #ifndef M_PHI
const M_PHI = 1.61803398874989484820 /* phi / golden ratio */
// #endif
// #ifndef M_PI
const M_PI = 3.14159265358979323846 /* pi */
// #endif
// #ifndef M_PI_2
const M_PI_2 = 1.57079632679489661923 /* pi/2 */
// #endif
// #ifndef M_SQRT1_2
const M_SQRT1_2 = 0.70710678118654752440 /* 1/sqrt(2) */
// #endif
// #ifndef M_SQRT2
const M_SQRT2 = 1.41421356237309504880 /* sqrt(2) */
//#endif
//#ifndef NAN
//const NAN        =    av_int2float(0x7fc00000)//todo
//#endif
//#ifndef INFINITY
//const INFINITY     =  av_int2float(0x7f800000)//todo
//#endif

/**
 * @addtogroup lavu_math
 *
 * @{
 */

/**
 * Rounding methods.
 */
type AVRounding int32

const (
	AV_ROUND_ZERO     = 0 ///< Round toward zero.
	AV_ROUND_INF      = 1 ///< Round away from zero.
	AV_ROUND_DOWN     = 2 ///< Round toward -infinity.
	AV_ROUND_UP       = 3 ///< Round toward +infinity.
	AV_ROUND_NEAR_INF = 5 ///< Round to nearest and halfway cases away from zero.
	/**
	 * Flag telling rescaling functions to pass `INT64_MIN`/`MAX` through
	 * unchanged, avoiding special cases for #AV_NOPTS_VALUE.
	 *
	 * Unlike other values of the enumeration AVRounding, this value is a
	 * bitmask that must be used in conjunction with another value of the
	 * enumeration through a bitwise OR, in order to set behavior for normal
	 * cases.
	 *
	 * @code{.c}
	 * av_rescale_rnd(3, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling 3:
	 * //     Calculating 3 * 1 / 2
	 * //     3 / 2 is rounded up to 2
	 * //     => 2
	 *
	 * av_rescale_rnd(AV_NOPTS_VALUE, 1, 2, AV_ROUND_UP | AV_ROUND_PASS_MINMAX);
	 * // Rescaling AV_NOPTS_VALUE:
	 * //     AV_NOPTS_VALUE == INT64_MIN
	 * //     AV_NOPTS_VALUE is passed through
	 * //     => AV_NOPTS_VALUE
	 * @endcode
	 */
	AV_ROUND_PASS_MINMAX = 8192
)

/**
 * Compute the greatest common divisor of two integer operands.
 *
 * @param a,b Operands
 * @return GCD of a and b up to sign; if a >= 0 and b >= 0, return value is >= 0;
 * if a == 0 and b == 0, returns 0.
 */
//int64_t av_const av_gcd(int64_t a, int64_t b);
// purego func
var avGcd func(a, b ffcommon.FInt64T) ffcommon.FInt64T
var avGcdOnce sync.Once

func AvGcd(a, b ffcommon.FInt64T) (res ffcommon.FInt64T) {
	avGcdOnce.Do(func() {
		purego.RegisterLibFunc(&avGcd, ffcommon.GetAvutilDll(), "av_gcd")
	})
	return avGcd(a, b)
}

/**
 * Rescale a 64-bit integer with rounding to nearest.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow.
 *
 * This function is equivalent to av_rescale_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale_rnd(), av_rescale_q(), av_rescale_q_rnd()
 */
//int64_t av_rescale(int64_t a, int64_t b, int64_t c) av_const;
// purego func
// purego func
var avRescale func(a, b, c ffcommon.FInt64T) ffcommon.FInt64T
var avRescaleOnce sync.Once

func AvRescale(a, b, c ffcommon.FInt64T) (res ffcommon.FInt64T) {
	avRescaleOnce.Do(func() {
		purego.RegisterLibFunc(&avRescale, ffcommon.GetAvutilDll(), "av_rescale")
	})
	return avRescale(a, b, c)
}

/**
 * Rescale a 64-bit integer with specified rounding.
 *
 * The operation is mathematically equivalent to `a * b / c`, but writing that
 * directly can overflow, and does not support different rounding methods.
 *
 * @see av_rescale(), av_rescale_q(), av_rescale_q_rnd()
 */
//int64_t av_rescale_rnd(int64_t a, int64_t b, int64_t c, enum AVRounding rnd) av_const;
// purego func
var avRescaleRnd func(a, b, c ffcommon.FInt64T, rnd AVRounding) ffcommon.FInt64T
var avRescaleRndOnce sync.Once

func AvRescaleRnd(a, b, c ffcommon.FInt64T, rnd AVRounding) (res ffcommon.FInt64T) {
	avRescaleRndOnce.Do(func() {
		purego.RegisterLibFunc(&avRescaleRnd, ffcommon.GetAvutilDll(), "av_rescale_rnd")
	})
	return avRescaleRnd(a, b, c, rnd)
}

/**
 * Rescale a 64-bit integer by 2 rational numbers.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * This function is equivalent to av_rescale_q_rnd() with #AV_ROUND_NEAR_INF.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q_rnd()
 */
//int64_t av_rescale_q(int64_t a, AVRational bq, AVRational cq) av_const;
// purego func
var avRescaleQ func(a, bq, cq uintptr) ffcommon.FInt64T
var avRescaleQOnce sync.Once

func AvRescaleQ(a ffcommon.FInt64T, bq, cq AVRational) (res ffcommon.FInt64T) {
	avRescaleQOnce.Do(func() {
		purego.RegisterLibFunc(&avRescaleQ, ffcommon.GetAvutilDll(), "av_rescale_q")
	})
	return avRescaleQ(
		uintptr(a),
		*(*uintptr)(unsafe.Pointer(&bq)), // 特殊处理，不然返回结果是错误的
		*(*uintptr)(unsafe.Pointer(&cq)), // 特殊处理，不然返回结果是错误的
	)
}

/**
 * Rescale a 64-bit integer by 2 rational numbers with specified rounding.
 *
 * The operation is mathematically equivalent to `a * bq / cq`.
 *
 * @see av_rescale(), av_rescale_rnd(), av_rescale_q()
 */
//int64_t av_rescale_q_rnd(int64_t a, AVRational bq, AVRational cq,
//enum AVRounding rnd) av_const;
// purego func
var avRescaleQRnd func(a, bq, cq, rnd uintptr) ffcommon.FInt64T
var avRescaleQRndOnce sync.Once

func AvRescaleQRnd(a ffcommon.FInt64T, bq, cq AVRational, rnd AVRounding) (res ffcommon.FInt64T) {
	avRescaleQRndOnce.Do(func() {
		purego.RegisterLibFunc(&avRescaleQRnd, ffcommon.GetAvutilDll(), "av_rescale_q_rnd")
	})
	res = avRescaleQRnd(
		uintptr(a),
		*(*uintptr)(unsafe.Pointer(&bq)), // 特殊处理，不然返回结果是错误的
		*(*uintptr)(unsafe.Pointer(&cq)), // 特殊处理，不然返回结果是错误的
		uintptr(rnd),
	)
	return
}

/**
 * Compare two timestamps each in its own time base.
 *
 * @return One of the following values:
 *         - -1 if `ts_a` is before `ts_b`
 *         - 1 if `ts_a` is after `ts_b`
 *         - 0 if they represent the same position
 *
 * @warning
 * The result of the function is undefined if one of the timestamps is outside
 * the `int64_t` range when represented in the other's timebase.
 */
//int av_compare_ts(int64_t ts_a, AVRational tb_a, int64_t ts_b, AVRational tb_b);
// purego func
var avCompareTs func(ts_a, tb_a, ts_b, tb_b uintptr) ffcommon.FInt
var avCompareTsOnce sync.Once

func AvCompareTs(ts_a ffcommon.FInt64T, tb_a AVRational, ts_b ffcommon.FInt64T, tb_b AVRational) (res ffcommon.FInt) {
	avCompareTsOnce.Do(func() {
		purego.RegisterLibFunc(&avCompareTs, ffcommon.GetAvutilDll(), "av_compare_ts")
	})
	return avCompareTs(uintptr(ts_a),
		*(*uintptr)(unsafe.Pointer(&tb_a)),
		uintptr(ts_b),
		*(*uintptr)(unsafe.Pointer(&tb_b)),
	)
}

/**
 * Compare the remainders of two integer operands divided by a common divisor.
 *
 * In other words, compare the least significant `log2(mod)` bits of integers
 * `a` and `b`.
 *
 * @code{.c}
 * av_compare_mod(0x11, 0x02, 0x10) < 0 // since 0x11 % 0x10  (0x1) < 0x02 % 0x10  (0x2)
 * av_compare_mod(0x11, 0x02, 0x20) > 0 // since 0x11 % 0x20 (0x11) > 0x02 % 0x20 (0x02)
 * @endcode
 *
 * @param a,b Operands
 * @param mod Divisor; must be a power of 2
 * @return
 *         - a negative value if `a % mod < b % mod`
 *         - a positive value if `a % mod > b % mod`
 *         - zero             if `a % mod == b % mod`
 */
//int64_t av_compare_mod(uint64_t a, uint64_t b, uint64_t mod);
// purego func
var avCompareMod func(a, b, mod ffcommon.FInt64T) ffcommon.FInt64T
var avCompareModOnce sync.Once

func AvCompareMod(a, b, mod ffcommon.FInt64T) (res ffcommon.FInt64T) {
	avCompareModOnce.Do(func() {
		purego.RegisterLibFunc(&avCompareMod, ffcommon.GetAvutilDll(), "av_compare_mod")
	})
	return avCompareMod(a, b, mod)
}

/**
 * Rescale a timestamp while preserving known durations.
 *
 * This function is designed to be called per audio packet to scale the input
 * timestamp to a different time base. Compared to a simple av_rescale_q()
 * call, this function is robust against possible inconsistent frame durations.
 *
 * The `last` parameter is a state variable that must be preserved for all
 * subsequent calls for the same stream. For the first call, `*last` should be
 * initialized to #AV_NOPTS_VALUE.
 *
 * @param[in]     in_tb    Input time base
 * @param[in]     in_ts    Input timestamp
 * @param[in]     fs_tb    Duration time base; typically this is finer-grained
 *                         (greater) than `in_tb` and `out_tb`
 * @param[in]     duration Duration till the next call to this function (i.e.
 *                         duration of the current packet/frame)
 * @param[in,out] last     Pointer to a timestamp expressed in terms of
 *                         `fs_tb`, acting as a state variable
 * @param[in]     out_tb   Output timebase
 * @return        Timestamp expressed in terms of `out_tb`
 *
 * @note In the context of this function, "duration" is in term of samples, not
 *       seconds.
 */
//int64_t av_rescale_delta(AVRational in_tb, int64_t in_ts,  AVRational fs_tb, int duration, int64_t *last, AVRational out_tb);
// purego func
var avRescaleDelta func(inTb AVRational, inTs ffcommon.FInt64T, fsTb AVRational, duration ffcommon.FInt, last *ffcommon.FInt64T, outTb AVRational) ffcommon.FInt64T
var avRescaleDeltaOnce sync.Once

func AvRescaleDelta(inTb AVRational, inTs ffcommon.FInt64T, fsTb AVRational, duration ffcommon.FInt, last *ffcommon.FInt64T, outTb AVRational) (res ffcommon.FInt64T) {
	avRescaleDeltaOnce.Do(func() {
		purego.RegisterLibFunc(&avRescaleDelta, ffcommon.GetAvutilDll(), "av_rescale_delta")
	})
	return avRescaleDelta(inTb, inTs, fsTb, duration, last, outTb)
}

/**
 * Add a value to a timestamp.
 *
 * This function guarantees that when the same value is repeatly added that
 * no accumulation of rounding errors occurs.
 *
 * @param[in] ts     Input timestamp
 * @param[in] ts_tb  Input timestamp time base
 * @param[in] inc    Value to be added
 * @param[in] inc_tb Time base of `inc`
 */
//int64_t av_add_stable(AVRational ts_tb, int64_t ts, AVRational inc_tb, int64_t inc);
// purego func
var avAddStable func(tsTb AVRational, ts ffcommon.FInt64T, incTb AVRational, inc ffcommon.FInt64T) ffcommon.FInt64T
var avAddStableOnce sync.Once

func AvAddStable(tsTb AVRational, ts ffcommon.FInt64T, incTb AVRational, inc ffcommon.FInt64T) (res ffcommon.FInt64T) {
	avAddStableOnce.Do(func() {
		purego.RegisterLibFunc(&avAddStable, ffcommon.GetAvutilDll(), "av_add_stable")
	})
	return avAddStable(tsTb, ts, incTb, inc)
}

/**
 * @}
 */

//#endif /* AVUTIL_MATHEMATICS_H */
