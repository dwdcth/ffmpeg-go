package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
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

//#ifndef AVUTIL_FILE_H
//#define AVUTIL_FILE_H
//
//#include <stdint.h>
//
//#include "avutil.h"

/**
 * @file
 * Misc file utilities.
 */

/**
 * Read the file with name filename, and put its content in a newly
 * allocated buffer or map it with mmap() when available.
 * In case of success set *bufptr to the read or mmapped buffer, and
 * *size to the size in bytes of the buffer in *bufptr.
 * Unlike mmap this function succeeds with zero sized files, in this
 * case *bufptr will be set to NULL and *size will be set to 0.
 * The returned buffer must be released with av_file_unmap().
 *
 * @param log_offset loglevel offset used for logging
 * @param log_ctx context used for logging
 * @return a non negative number in case of success, a negative value
 * corresponding to an AVERROR error code in case of failure
 */
//av_warn_unused_result
//int av_file_map(const char *filename, uint8_t **bufptr, size_t *size,
//int log_offset, void *log_ctx);
// purego func
var avFileMap func(filename ffcommon.FConstCharP, bufptr **ffcommon.FUint8T, size *ffcommon.FSizeT,
	logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) ffcommon.FInt
var avFileMapOnce sync.Once

func AvFileMap(filename ffcommon.FConstCharP, bufptr **ffcommon.FUint8T, size *ffcommon.FSizeT,
	logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) (res ffcommon.FInt) {
	avFileMapOnce.Do(func() {
		purego.RegisterLibFunc(&avFileMap, ffcommon.GetAvutilDll(), "av_file_map")
	})
	res = avFileMap(filename, bufptr, size, logOffset, logCtx)
	return
}

/**
 * Unmap or free the buffer bufptr created by av_file_map().
 *
 * @param size size in bytes of bufptr, must be the same as returned
 * by av_file_map()
 */
//void av_file_unmap(uint8_t *bufptr, size_t size);
// purego func
var avFileUnmap func(bufptr *ffcommon.FUint8T, size ffcommon.FSizeT)
var avFileUnmapOnce sync.Once

func AvFileUnmap(bufptr *ffcommon.FUint8T, size ffcommon.FSizeT) {
	avFileUnmapOnce.Do(func() {
		purego.RegisterLibFunc(&avFileUnmap, ffcommon.GetAvutilDll(), "av_file_unmap")
	})
	avFileUnmap(bufptr, size)
}

/**
 * Wrapper to work around the lack of mkstemp() on mingw.
 * Also, tries to create file in /tmp first, if possible.
 * *prefix can be a character constant; *filename will be allocated internally.
 * @return file descriptor of opened file (or negative value corresponding to an
 * AVERROR code on error)
 * and opened file name in **filename.
 * @note On very old libcs it is necessary to set a secure umask before
 *       calling this, av_tempfile() can't call umask itself as it is used in
 *       libraries and could interfere with the calling application.
 * @deprecated as fd numbers cannot be passed saftely between libs on some platforms
 */
//int av_tempfile(const char *prefix, char **filename, int log_offset, void *log_ctx);
// purego func
var avTempfile func(prefix ffcommon.FConstCharP, filename *ffcommon.FBuf, logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) ffcommon.FInt
var avTempfileOnce sync.Once

func AvTempfile(prefix ffcommon.FConstCharP, filename *ffcommon.FBuf, logOffset ffcommon.FInt, logCtx ffcommon.FVoidP) (res ffcommon.FInt) {
	avTempfileOnce.Do(func() {
		purego.RegisterLibFunc(&avTempfile, ffcommon.GetAvutilDll(), "av_tempfile")
	})
	res = avTempfile(prefix, filename, logOffset, logCtx)
	return
}

//#endif /* AVUTIL_FILE_H */
