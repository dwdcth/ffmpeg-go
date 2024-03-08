package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2000-2003 Fabrice Bellard
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

//#ifndef AVUTIL_TIME_H
//#define AVUTIL_TIME_H
//
//#include <stdint.h>

/**
 * Get the current time in microseconds.
 */
//int64_t av_gettime(void);
var avGettime func() ffcommon.FInt64T
var avGettimeOnce sync.Once

func AvGettime() ffcommon.FInt64T {
	avGettimeOnce.Do(func() {
		purego.RegisterLibFunc(&avGettime, ffcommon.GetAvutilDll(), "av_gettime")
	})
	return avGettime()
}

/**
 * Get the current time in microseconds since some unspecified starting point.
 * On platforms that support it, the time comes from a monotonic clock
 * This property makes this time source ideal for measuring relative time.
 * The returned values may not be monotonic on platforms where a monotonic
 * clock is not available.
 */
//int64_t av_gettime_relative(void);
var avGettimeRelative func() ffcommon.FInt64T
var avGettimeRelativeOnce sync.Once

func AvGettimeRelative() ffcommon.FInt64T {
	avGettimeRelativeOnce.Do(func() {
		purego.RegisterLibFunc(&avGettimeRelative, ffcommon.GetAvutilDll(), "av_gettime_relative")
	})
	return avGettimeRelative()
}

/**
 * Indicates with a boolean result if the av_gettime_relative() time source
 * is monotonic.
 */
//int av_gettime_relative_is_monotonic(void);
var avGettimeRelativeIsMonotonic func() ffcommon.FInt
var avGettimeRelativeIsMonotonicOnce sync.Once

func AvGettimeRelativeIsMonotonic() ffcommon.FInt {
	avGettimeRelativeIsMonotonicOnce.Do(func() {
		purego.RegisterLibFunc(&avGettimeRelativeIsMonotonic, ffcommon.GetAvutilDll(), "av_gettime_relative_is_monotonic")
	})
	return avGettimeRelativeIsMonotonic()
}

/**
 * Sleep for a period of time.  Although the duration is expressed in
 * microseconds, the actual delay may be rounded to the precision of the
 * system timer.
 *
 * @param  usec Number of microseconds to sleep.
 * @return zero on success or (negative) error code.
 */
//int av_usleep(unsigned usec);
var avUsleep func(usec ffcommon.FUnsigned) ffcommon.FInt
var avUsleepOnce sync.Once

func AvUsleep(usec ffcommon.FUnsigned) ffcommon.FInt {
	avUsleepOnce.Do(func() {
		purego.RegisterLibFunc(&avUsleep, ffcommon.GetAvutilDll(), "av_usleep")
	})
	return avUsleep(usec)
}

//#endif /* AVUTIL_TIME_H */
