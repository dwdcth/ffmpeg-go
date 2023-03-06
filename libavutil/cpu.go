package libavutil

import (
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
)

/*
 * Copyright (c) 2000, 2001, 2002 Fabrice Bellard
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

//#ifndef AVUTIL_CPU_H
//#define AVUTIL_CPU_H
//
//#include <stddef.h>
//
//#include "attributes.h"

const AV_CPU_FLAG_FORCE = 0x80000000 /* force usage of selected flags (OR) */

/* lower 16 bits - CPU features */
const AV_CPU_FLAG_MMX = 0x0001          ///< standard MMX
const AV_CPU_FLAG_MMXEXT = 0x0002       ///< SSE integer functions or AMD MMX ext
const AV_CPU_FLAG_MMX2 = 0x0002         ///< SSE integer functions or AMD MMX ext
const AV_CPU_FLAG_3DNOW = 0x0004        ///< AMD 3DNOW
const AV_CPU_FLAG_SSE = 0x0008          ///< SSE functions
const AV_CPU_FLAG_SSE2 = 0x0010         ///< PIV SSE2 functions
const AV_CPU_FLAG_SSE2SLOW = 0x40000000 ///< SSE2 supported, but usually not faster
// /< than regular MMX/SSE (e.g. Core1)
const AV_CPU_FLAG_3DNOWEXT = 0x0020     ///< AMD 3DNowExt
const AV_CPU_FLAG_SSE3 = 0x0040         ///< Prescott SSE3 functions
const AV_CPU_FLAG_SSE3SLOW = 0x20000000 ///< SSE3 supported, but usually not faster
// /< than regular MMX/SSE (e.g. Core1)
const AV_CPU_FLAG_SSSE3 = 0x0080        ///< Conroe SSSE3 functions
const AV_CPU_FLAG_SSSE3SLOW = 0x4000000 ///< SSSE3 supported, but usually not faster
const AV_CPU_FLAG_ATOM = 0x10000000     ///< Atom processor, some SSSE3 instructions are slower
const AV_CPU_FLAG_SSE4 = 0x0100         ///< Penryn SSE4.1 functions
const AV_CPU_FLAG_SSE42 = 0x0200        ///< Nehalem SSE4.2 functions
const AV_CPU_FLAG_AESNI = 0x80000       ///< Advanced Encryption Standard functions
const AV_CPU_FLAG_AVX = 0x4000          ///< AVX functions: requires OS support even if YMM registers aren't used
const AV_CPU_FLAG_AVXSLOW = 0x8000000   ///< AVX supported, but slow when using YMM registers (e.g. Bulldozer)
const AV_CPU_FLAG_XOP = 0x0400          ///< Bulldozer XOP functions
const AV_CPU_FLAG_FMA4 = 0x0800         ///< Bulldozer FMA4 functions
const AV_CPU_FLAG_CMOV = 0x1000         ///< supports cmov instruction
const AV_CPU_FLAG_AVX2 = 0x8000         ///< AVX2 functions: requires OS support even if YMM registers aren't used
const AV_CPU_FLAG_FMA3 = 0x10000        ///< Haswell FMA3 functions
const AV_CPU_FLAG_BMI1 = 0x20000        ///< Bit Manipulation Instruction Set 1
const AV_CPU_FLAG_BMI2 = 0x40000        ///< Bit Manipulation Instruction Set 2
const AV_CPU_FLAG_AVX512 = 0x100000     ///< AVX-512 functions: requires OS support even if YMM/ZMM registers aren't used

const AV_CPU_FLAG_ALTIVEC = 0x0001 ///< standard
const AV_CPU_FLAG_VSX = 0x0002     ///< ISA 2.06
const AV_CPU_FLAG_POWER8 = 0x0004  ///< ISA 2.07

const AV_CPU_FLAG_ARMV5TE = (1 << 0)
const AV_CPU_FLAG_ARMV6 = (1 << 1)
const AV_CPU_FLAG_ARMV6T2 = (1 << 2)
const AV_CPU_FLAG_VFP = (1 << 3)
const AV_CPU_FLAG_VFPV3 = (1 << 4)
const AV_CPU_FLAG_NEON = (1 << 5)
const AV_CPU_FLAG_ARMV8 = (1 << 6)
const AV_CPU_FLAG_VFP_VM = (1 << 7) ///< VFPv2 vector mode, deprecated in ARMv7-A and unavailable in various CPUs implementations
const AV_CPU_FLAG_SETEND = (1 << 16)

const AV_CPU_FLAG_MMI = (1 << 0)
const AV_CPU_FLAG_MSA = (1 << 1)

/**
 * Return the flags which specify extensions supported by the CPU.
 * The returned value is affected by av_force_cpu_flags() if that was used
 * before. So av_get_cpu_flags() can easily be used in an application to
 * detect the enabled cpu flags.
 */
//int av_get_cpu_flags(void);
func AvGetCpuFlags() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_get_cpu_flags").Call()
	res = ffcommon.FInt(t)
	return
}

/**
 * Disables cpu detection and forces the specified flags.
 * -1 is a special case that disables forcing of specific flags.
 */
//void av_force_cpu_flags(int flags);
func AvForceCpuFlags(flags ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_force_cpu_flags").Call(
		uintptr(flags),
	)
}

/**
 * Set a mask on flags returned by av_get_cpu_flags().
 * This function is mainly useful for testing.
 * Please use av_force_cpu_flags() and av_get_cpu_flags() instead which are more flexible
 */
//attribute_deprecated void av_set_cpu_flags_mask(int mask);
func AvSetCpuFlagsMask(mask ffcommon.FInt) {
	ffcommon.GetAvutilDll().NewProc("av_set_cpu_flags_mask").Call(
		uintptr(mask),
	)
}

/**
 * Parse CPU flags from a string.
 *
 * The returned flags contain the specified flags as well as related unspecified flags.
 *
 * This function exists only for compatibility with libav.
 * Please use av_parse_cpu_caps() when possible.
 * @return a combination of AV_CPU_* flags, negative on error.
 */
//attribute_deprecated
//int av_parse_cpu_flags(const char *s);
func AvParseCpuFlags(s ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_parse_cpu_flags").Call(
		ffcommon.UintPtrFromString(s),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Parse CPU caps from a string and update the given AV_CPU_* flags based on that.
 *
 * @return negative on error.
 */
//int av_parse_cpu_caps(unsigned *flags, const char *s);
func AvParseCpuCaps(flag *ffcommon.FUnsigned, s ffcommon.FConstCharP) (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_parse_cpu_caps").Call(
		uintptr(unsafe.Pointer(flag)),
		ffcommon.UintPtrFromString(s),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * @return the number of logical CPU cores present.
 */
//int av_cpu_count(void);
func AvCpuCount() (res ffcommon.FInt) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_cpu_count").Call()
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the maximum data alignment that may be required by FFmpeg.
 *
 * Note that this is affected by the build configuration and the CPU flags mask,
 * so e.g. if the CPU supports AVX, but libavutil has been built with
 * --disable-avx or the AV_CPU_FLAG_AVX flag has been disabled through
 *  av_set_cpu_flags_mask(), then this function will behave as if AVX is not
 *  present.
 */
//size_t av_cpu_max_align(void);
func AvCpuMaxAlign() (res ffcommon.FSizeT) {
	t, _, _ := ffcommon.GetAvutilDll().NewProc("av_cpu_max_align").Call()
	res = ffcommon.FSizeT(t)
	return
}

//#endif /* AVUTIL_CPU_H */
