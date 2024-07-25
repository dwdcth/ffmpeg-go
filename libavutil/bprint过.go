package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2012 Nicolas George
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

//#ifndef AVUTIL_BPRINT_H
//#define AVUTIL_BPRINT_H
//
//#include <stdarg.h>
//
//#include "attributes.h"
//#include "avstring.h"
//
///**
// * Define a structure with extra padding to a fixed size
// * This helps ensuring binary compatibility with future versions.
// */
//
//#define FF_PAD_STRUCTURE(name, size, ...) \
//struct ff_pad_helper_##name { __VA_ARGS__ }; \
//typedef struct name { \
//__VA_ARGS__ \
//char reserved_padding[size - sizeof(struct ff_pad_helper_##name)]; \
//} name;

/**
 * Buffer to print data progressively
 *
 * The string buffer grows as necessary and is always 0-terminated.
 * The content of the string is never accessed, and thus is
 * encoding-agnostic and can even hold binary data.
 *
 * Small buffers are kept in the structure itself, and thus require no
 * memory allocation at all (unless the contents of the buffer is needed
 * after the structure goes out of scope). This is almost as lightweight as
 * declaring a local "char buf[512]".
 *
 * The length of the string can go beyond the allocated size: the buffer is
 * then truncated, but the functions still keep account of the actual total
 * length.
 *
 * In other words, buf->len can be greater than buf->size and records the
 * total length of what would have been to the buffer if there had been
 * enough memory.
 *
 * Append operations do not need to be tested for failure: if a memory
 * allocation fails, data stop being appended to the buffer, but the length
 * is still updated. This situation can be tested with
 * av_bprint_is_complete().
 *
 * The size_max field determines several possible behaviours:
 *
 * size_max = -1 (= UINT_MAX) or any large value will let the buffer be
 * reallocated as necessary, with an amortized linear cost.
 *
 * size_max = 0 prevents writing anything to the buffer: only the total
 * length is computed. The write operations can then possibly be repeated in
 * a buffer with exactly the necessary size
 * (using size_init = size_max = len + 1).
 *
 * size_max = 1 is automatically replaced by the exact size available in the
 * structure itself, thus ensuring no dynamic memory allocation. The
 * internal buffer is large enough to hold a reasonable paragraph of text,
 * such as the current paragraph.
 */

// FF_PAD_STRUCTURE(AVBPrint, 1024,
// char *str;         /**< string so far */
// unsigned len;      /**< length so far */
// unsigned size;     /**< allocated memory */
// unsigned size_max; /**< maximum allocated memory */
// char reserved_internal_buffer[1];
// )
type AVBPrint struct {
	_ [128]uintptr
}

/**
 * Convenience macros for special values for av_bprint_init() size_max
 * parameter.
 */
//#define AV_BPRINT_SIZE_UNLIMITED  ((unsigned)-1)
const AV_BPRINT_SIZE_AUTOMATIC = 1
const AV_BPRINT_SIZE_COUNT_ONLY = 0

/**
 * Init a print buffer.
 *
 * @param buf        buffer to init
 * @param size_init  initial size (including the final 0)
 * @param size_max   maximum size;
 *                   0 means do not write anything, just count the length;
 *                   1 is replaced by the maximum value for automatic storage;
 *                   any large value means that the internal buffer will be
 *                   reallocated as needed up to that limit; -1 is converted to
 *                   UINT_MAX, the largest limit possible.
 *                   Check also AV_BPRINT_SIZE_* macros.
 */
//void av_bprint_init(AVBPrint *buf, unsigned size_init, unsigned size_max);
//todo
var avBprintInit func() ffcommon.FCharP
var avBprintInitOnce sync.Once

func AvBprintInit() ffcommon.FCharP {
	avBprintInitOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintInit, ffcommon.GetAvutilDll(), "av_blowfish_crypt_ecb")
	})
	return avBprintInit()
}

/**
 * Init a print buffer using a pre-existing buffer.
 *
 * The buffer will not be reallocated.
 *
 * @param buf     buffer structure to init
 * @param buffer  byte buffer to use for the string data
 * @param size    size of buffer
 */
//void av_bprint_init_for_buffer(AVBPrint *buf, char *buffer, unsigned size);
//todo
var avBprintInitForBuffer func() ffcommon.FCharP
var avBprintInitForBufferOnce sync.Once

func AvBprintInitForBuffer() ffcommon.FCharP {
	avBprintInitForBufferOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintInitForBuffer, ffcommon.GetAvutilDll(), "av_bprint_init_for_buffer")
	})
	return avBprintInitForBuffer()
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_bprintf(AVBPrint *buf, const char *fmt, ...) av_printf_format(2, 3);
//todo
var avBprintf func() ffcommon.FCharP
var avBprintfOnce sync.Once

func AvBprintf() ffcommon.FCharP {
	avBprintfOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintf, ffcommon.GetAvutilDll(), "av_bprintf")
	})
	return avBprintf()
}

/**
 * Append a formatted string to a print buffer.
 */
//void av_vbprintf(AVBPrint *buf, const char *fmt, va_list vl_arg);
//todo
var avVbprintf func() ffcommon.FCharP
var avVbprintfOnce sync.Once

func AvVbprintf() ffcommon.FCharP {
	avVbprintfOnce.Do(func() {
		purego.RegisterLibFunc(&avVbprintf, ffcommon.GetAvutilDll(), "av_vbprintf")
	})
	return avVbprintf()
}

/**
 * Append char c n times to a print buffer.
 */
//void av_bprint_chars(AVBPrint *buf, char c, unsigned n);
//todo
var avBprintChars func() ffcommon.FCharP
var avBprintCharsOnce sync.Once

func AvBprintChars() ffcommon.FCharP {
	avBprintCharsOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintChars, ffcommon.GetAvutilDll(), "av_bprint_chars")
	})
	return avBprintChars()
}

/**
 * Append data to a print buffer.
 *
 * param buf  bprint buffer to use
 * param data pointer to data
 * param size size of data
 */
//void av_bprint_append_data(AVBPrint *buf, const char *data, unsigned size);
//todo
var avBprintAppendData func() ffcommon.FCharP
var avBprintAppendDataOnce sync.Once

func AvBprintAppendData() ffcommon.FCharP {
	avBprintAppendDataOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintAppendData, ffcommon.GetAvutilDll(), "av_bprint_append_data")
	})
	return avBprintAppendData()
}

// struct tm;
type Tm struct {
}

/**
 * Append a formatted date and time to a print buffer.
 *
 * param buf  bprint buffer to use
 * param fmt  date and time format string, see strftime()
 * param tm   broken-down time structure to translate
 *
 * @note due to poor design of the standard strftime function, it may
 * produce poor results if the format string expands to a very long text and
 * the bprint buffer is near the limit stated by the size_max option.
 */
//void av_bprint_strftime(AVBPrint *buf, const char *fmt, const struct tm *tm);
//todo
var avBprintStrftime func() ffcommon.FCharP
var avBprintStrftimeOnce sync.Once

func AvBprintStrftime() ffcommon.FCharP {
	avBprintStrftimeOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintStrftime, ffcommon.GetAvutilDll(), "av_bprint_strftime")
	})
	return avBprintStrftime()
}

/**
 * Allocate bytes in the buffer for external use.
 *
 * @param[in]  buf          buffer structure
 * @param[in]  size         required size
 * @param[out] mem          pointer to the memory area
 * @param[out] actual_size  size of the memory area after allocation;
 *                          can be larger or smaller than size
 */
//void av_bprint_get_buffer(AVBPrint *buf, unsigned size,
//unsigned char **mem, unsigned *actual_size);
//todo
var avBprintGetBuffer func() ffcommon.FCharP
var avBprintGetBufferOnce sync.Once

func AvBprintGetBuffer() ffcommon.FCharP {
	avBprintGetBufferOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintGetBuffer, ffcommon.GetAvutilDll(), "av_bprint_get_buffer")
	})
	return avBprintGetBuffer()
}

/**
 * Reset the string to "" but keep internal allocated data.
 */
//void av_bprint_clear(AVBPrint *buf);
//todo
var avBprintClear func() ffcommon.FCharP
var avBprintClearOnce sync.Once

func AvBprintClear() ffcommon.FCharP {
	avBprintClearOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintClear, ffcommon.GetAvutilDll(), "av_bprint_clear")
	})
	return avBprintClear()
}

/**
 * Test if the print buffer is complete (not truncated).
 *
 * It may have been truncated due to a memory allocation failure
 * or the size_max limit (compare size and size_max if necessary).
 */
//static inline int av_bprint_is_complete(const AVBPrint *buf)
//{
//return buf->len < buf->size;
//}
//todo
var avBprintIsComplete func() ffcommon.FCharP
var avBprintIsCompleteOnce sync.Once

func AvBprintIsComplete() ffcommon.FCharP {
	avBprintIsCompleteOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintIsComplete, ffcommon.GetAvutilDll(), "av_bprint_is_complete")
	})
	return avBprintIsComplete()
}

/**
 * Finalize a print buffer.
 *
 * The print buffer can no longer be used afterwards,
 * but the len and size fields are still valid.
 *
 * @arg[out] ret_str  if not NULL, used to return a permanent copy of the
 *                    buffer contents, or NULL if memory allocation fails;
 *                    if NULL, the buffer is discarded and freed
 * @return  0 for success or error code (probably AVERROR(ENOMEM))
 */
//int av_bprint_finalize(AVBPrint *buf, char **ret_str);
//todo
var avBprintFinalize func() ffcommon.FCharP
var avBprintFinalizeOnce sync.Once

func AvBprintFinalize() ffcommon.FCharP {
	avBprintFinalizeOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintFinalize, ffcommon.GetAvutilDll(), "av_bprint_finalize")
	})
	return avBprintFinalize()
}

/**
 * Escape the content in src and append it to dstbuf.
 *
 * @param dstbuf        already inited destination bprint buffer
 * @param src           string containing the text to escape
 * @param special_chars string containing the special characters which
 *                      need to be escaped, can be NULL
 * @param mode          escape mode to employ, see AV_ESCAPE_MODE_* macros.
 *                      Any unknown value for mode will be considered equivalent to
 *                      AV_ESCAPE_MODE_BACKSLASH, but this behaviour can change without
 *                      notice.
 * @param flags         flags which control how to escape, see AV_ESCAPE_FLAG_* macros
 */
//void av_bprint_escape(AVBPrint *dstbuf, const char *src, const char *special_chars,
//enum AVEscapeMode mode, int flags);
//todo
var avBprintEscape func() ffcommon.FCharP
var avBprintEscapeOnce sync.Once

func AvBprintEscape() ffcommon.FCharP {
	avBprintEscapeOnce.Do(func() {
		purego.RegisterLibFunc(&avBprintEscape, ffcommon.GetAvutilDll(), "av_bprint_escape")
	})
	return avBprintEscape()
}

//#endif /* AVUTIL_BPRINT_H */
