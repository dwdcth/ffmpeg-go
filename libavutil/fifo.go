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

/**
 * @file
 * a very simple circular buffer FIFO implementation
 */

//#ifndef AVUTIL_FIFO_H
//#define AVUTIL_FIFO_H
//
//#include <stdint.h>
//#include "avutil.h"
//#include "attributes.h"

type AVFifoBuffer struct {
	Buffer          *ffcommon.FUint8T
	Rptr, Wptr, End *ffcommon.FUint8T
	Rndx, Wndx      ffcommon.FUint32T
}

/**
 * Initialize an AVFifoBuffer.
 * @param size of FIFO
 * @return AVFifoBuffer or NULL in case of memory allocation failure
 */
//AVFifoBuffer *av_fifo_alloc(unsigned int size);
var avFifoAlloc func(size ffcommon.FUnsignedInt) *AVFifoBuffer

var avFifoAllocOnce sync.Once

func AvFifoAlloc(size ffcommon.FUnsignedInt) *AVFifoBuffer {
	avFifoAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoAlloc, ffcommon.GetAvutilDll(), "av_fifo_alloc")
	})
	return avFifoAlloc(size)
}

/**
 * Initialize an AVFifoBuffer.
 * @param nmemb number of elements
 * @param size  size of the single element
 * @return AVFifoBuffer or NULL in case of memory allocation failure
 */
//AVFifoBuffer *av_fifo_alloc_array(size_t nmemb, size_t size);
var avFifoAllocArray func(nmemb, size ffcommon.FSizeT) *AVFifoBuffer

var avFifoAllocArrayOnce sync.Once

func AvFifoAllocArray(nmemb, size ffcommon.FSizeT) *AVFifoBuffer {
	avFifoAllocArrayOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoAllocArray, ffcommon.GetAvutilDll(), "av_fifo_alloc_array")
	})
	return avFifoAllocArray(nmemb, size)
}

/**
 * Free an AVFifoBuffer.
 * @param f AVFifoBuffer to free
 */
//void av_fifo_free(AVFifoBuffer *f);
var avFifoFree func(f *AVFifoBuffer)

var avFifoFreeOnce sync.Once

func (f *AVFifoBuffer) AvFifoFree() {
	avFifoFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoFree, ffcommon.GetAvutilDll(), "av_fifo_free")
	})
	avFifoFree(f)
}

/**
 * Free an AVFifoBuffer and reset pointer to NULL.
 * @param f AVFifoBuffer to free
 */
//void av_fifo_freep(AVFifoBuffer **f);
var avFifoFreep func(f **AVFifoBuffer)

var avFifoFreepOnce sync.Once

func AvFifoFreep(f **AVFifoBuffer) {
	avFifoFreepOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoFreep, ffcommon.GetAvutilDll(), "av_fifo_freep")
	})
	avFifoFreep(f)
}

/**
 * Reset the AVFifoBuffer to the state right after av_fifo_alloc, in particular it is emptied.
 * @param f AVFifoBuffer to reset
 */
//void av_fifo_reset(AVFifoBuffer *f);
var avFifoReset func(f *AVFifoBuffer)

var avFifoResetOnce sync.Once

func (f *AVFifoBuffer) AvFifoReset() {
	avFifoResetOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoReset, ffcommon.GetAvutilDll(), "av_fifo_reset")
	})
	avFifoReset(f)
}

/**
 * Return the amount of data in bytes in the AVFifoBuffer, that is the
 * amount of data you can read from it.
 * @param f AVFifoBuffer to read from
 * @return size
 */
//int av_fifo_size(const AVFifoBuffer *f);
var avFifoSize func(f *AVFifoBuffer) ffcommon.FInt

var avFifoSizeOnce sync.Once

func (f *AVFifoBuffer) AvFifoSize() ffcommon.FInt {
	avFifoSizeOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoSize, ffcommon.GetAvutilDll(), "av_fifo_size")
	})
	return avFifoSize(f)
}

/**
 * Return the amount of space in bytes in the AVFifoBuffer, that is the
 * amount of data you can write into it.
 * @param f AVFifoBuffer to write into
 * @return size
 */
//int av_fifo_space(const AVFifoBuffer *f);
var avFifoSpace func(f *AVFifoBuffer) ffcommon.FInt

var avFifoSpaceOnce sync.Once

func (f *AVFifoBuffer) AvFifoSpace() ffcommon.FInt {
	avFifoSpaceOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoSpace, ffcommon.GetAvutilDll(), "av_fifo_space")
	})
	return avFifoSpace(f)
}

/**
 * Feed data at specific position from an AVFifoBuffer to a user-supplied callback.
 * Similar as av_fifo_gereric_read but without discarding data.
 * @param f AVFifoBuffer to read from
 * @param offset offset from current read position
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_peek_at(AVFifoBuffer *f, void *dest, int offset, int buf_size, void (*func)(void*, void*, int));
var avFifoGenericPeekAt func(f *AVFifoBuffer, dest ffcommon.FVoidP, offset, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt

var avFifoGenericPeekAtOnce sync.Once

func (f *AVFifoBuffer) AvFifoGenericPeekAt(dest ffcommon.FVoidP, offset, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt {
	avFifoGenericPeekAtOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoGenericPeekAt, ffcommon.GetAvutilDll(), "av_fifo_generic_peek_at")
	})
	return avFifoGenericPeekAt(f, dest, offset, buf_size, func0)
}

/**
 * Feed data from an AVFifoBuffer to a user-supplied callback.
 * Similar as av_fifo_gereric_read but without discarding data.
 * @param f AVFifoBuffer to read from
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_peek(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
var avFifoGenericPeek func(f *AVFifoBuffer, dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt

var avFifoGenericPeekOnce sync.Once

func (f *AVFifoBuffer) AvFifoGenericPeek(dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt {
	avFifoGenericPeekOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoGenericPeek, ffcommon.GetAvutilDll(), "av_fifo_generic_peek")
	})
	return avFifoGenericPeek(f, dest, buf_size, func0)
}

/**
 * Feed data from an AVFifoBuffer to a user-supplied callback.
 * @param f AVFifoBuffer to read from
 * @param buf_size number of bytes to read
 * @param func generic read function
 * @param dest data destination
 */
//int av_fifo_generic_read(AVFifoBuffer *f, void *dest, int buf_size, void (*func)(void*, void*, int));
var avFifoGenericRead func(f *AVFifoBuffer, dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt

var avFifoGenericReadOnce sync.Once

func (f *AVFifoBuffer) AvFifoGenericRead(dest ffcommon.FVoidP, buf_size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt {
	avFifoGenericReadOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoGenericRead, ffcommon.GetAvutilDll(), "av_fifo_generic_read")
	})
	return avFifoGenericRead(f, dest, buf_size, func0)
}

/**
 * Feed data from a user-supplied callback to an AVFifoBuffer.
 * @param f AVFifoBuffer to write to
 * @param src data source; non-const since it may be used as a
 * modifiable context by the function defined in func
 * @param size number of bytes to write
 * @param func generic write function; the first parameter is src,
 * the second is dest_buf, the third is dest_buf_size.
 * func must return the number of bytes written to dest_buf, or <= 0 to
 * indicate no more data available to write.
 * If func is NULL, src is interpreted as a simple byte array for source data.
 * @return the number of bytes written to the FIFO
 */
//int av_fifo_generic_write(AVFifoBuffer *f, void *src, int size, int (*func)(void*, void*, int));
var avFifoGenericWrite func(f *AVFifoBuffer, src ffcommon.FVoidP, size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt

var avFifoGenericWriteOnce sync.Once

func (f *AVFifoBuffer) AvFifoGenericWrite(src ffcommon.FVoidP, size ffcommon.FInt, func0 func(ffcommon.FVoidP, ffcommon.FVoidP, ffcommon.FInt) uintptr) ffcommon.FInt {
	avFifoGenericWriteOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoGenericWrite, ffcommon.GetAvutilDll(), "av_fifo_generic_write")
	})
	return avFifoGenericWrite(f, src, size, func0)
}

/**
 * Resize an AVFifoBuffer.
 * In case of reallocation failure, the old FIFO is kept unchanged.
 *
 * @param f AVFifoBuffer to resize
 * @param size new AVFifoBuffer size in bytes
 * @return <0 for failure, >=0 otherwise
 */
//int av_fifo_realloc2(AVFifoBuffer *f, unsigned int size);
var avFifoRealloc2 func(f *AVFifoBuffer, size ffcommon.FUnsignedInt) ffcommon.FInt

var avFifoRealloc2Once sync.Once

func (f *AVFifoBuffer) AvFifoRealloc2(size ffcommon.FUnsignedInt) ffcommon.FInt {
	avFifoRealloc2Once.Do(func() {
		purego.RegisterLibFunc(&avFifoRealloc2, ffcommon.GetAvutilDll(), "av_fifo_realloc2")
	})
	return avFifoRealloc2(f, size)
}

/**
 * Enlarge an AVFifoBuffer.
 * In case of reallocation failure, the old FIFO is kept unchanged.
 * The new fifo size may be larger than the requested size.
 *
 * @param f AVFifoBuffer to resize
 * @param additional_space the amount of space in bytes to allocate in addition to av_fifo_size()
 * @return <0 for failure, >=0 otherwise
 */
//int av_fifo_grow(AVFifoBuffer *f, unsigned int additional_space);
var avFifoGrow func(f *AVFifoBuffer, additional_space ffcommon.FUnsignedInt) ffcommon.FInt

var avFifoGrowOnce sync.Once

func (f *AVFifoBuffer) AvFifoGrow(additional_space ffcommon.FUnsignedInt) ffcommon.FInt {
	avFifoGrowOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoGrow, ffcommon.GetAvutilDll(), "av_fifo_grow")
	})
	return avFifoGrow(f, additional_space)
}

/**
 * Read and discard the specified amount of data from an AVFifoBuffer.
 * @param f AVFifoBuffer to read from
 * @param size amount of data to read in bytes
 */
//void av_fifo_drain(AVFifoBuffer *f, int size);
var avFifoDrain func(f *AVFifoBuffer, size ffcommon.FInt)

var avFifoDrainOnce sync.Once

func (f *AVFifoBuffer) AvFifoDrain(size ffcommon.FInt) {
	avFifoDrainOnce.Do(func() {
		purego.RegisterLibFunc(&avFifoDrain, ffcommon.GetAvutilDll(), "av_fifo_drain")
	})
	avFifoDrain(f, size)
}

/**
 * Return a pointer to the data stored in a FIFO buffer at a certain offset.
 * The FIFO buffer is not modified.
 *
 * @param f    AVFifoBuffer to peek at, f must be non-NULL
 * @param offs an offset in bytes, its absolute value must be less
 *             than the used buffer size or the returned pointer will
 *             point outside to the buffer data.
 *             The used buffer size can be checked with av_fifo_size().
 */
//static inline uint8_t *av_fifo_peek2(const AVFifoBuffer *f, int offs)
//{
//uint8_t *ptr = f->rptr + offs;
//if (ptr >= f->end)
//ptr = f->buffer + (ptr - f->end);
//else if (ptr < f->buffer)
//ptr = f->end - (f->buffer - ptr);
//return ptr;
//}
//todo
var avFifoPeek2 func() ffcommon.FCharP
var avFifoPeek2Once sync.Once

func AvFifoPeek2() ffcommon.FCharP {
	avFifoPeek2Once.Do(func() {
		purego.RegisterLibFunc(&avFifoPeek2, ffcommon.GetAvutilDll(), "av_fifo_peek2")
	})
	return avFifoPeek2()
}

//#endif /* AVUTIL_FIFO_H */
