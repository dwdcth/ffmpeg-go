package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Audio FIFO
 * Copyright (c) 2012 Justin Ruggles <justin.ruggles@gmail.com>
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
 * Audio FIFO Buffer
 */

//#ifndef AVUTIL_AUDIO_FIFO_H
//#define AVUTIL_AUDIO_FIFO_H
//
//#include "avutil.h"
//#include "fifo.h"
//#include "samplefmt.h"

/**
 * @addtogroup lavu_audio
 * @{
 *
 * @defgroup lavu_audiofifo Audio FIFO Buffer
 * @{
 */

/**
 * Context for an Audio FIFO Buffer.
 *
 * - Operates at the sample level rather than the byte level.
 * - Supports multiple channels with either planar or packed sample format.
 * - Automatic reallocation when writing to a full buffer.
 */
//typedef struct AVAudioFifo AVAudioFifo;
type AVAudioFifo struct {
}

/**
 * Free an AVAudioFifo.
 *
 * @param af  AVAudioFifo to free
 */
//void av_audio_fifo_free(AVAudioFifo *af);
var avAudioFifoFree func(af *AVAudioFifo)

var avAudioFifoFreeOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoFree() {
	avAudioFifoFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoFree, ffcommon.GetAvutilDll(), "av_audio_fifo_free")
	})
	avAudioFifoFree(af)
}

/**
 * Allocate an AVAudioFifo.
 *
 * @param sample_fmt  sample format
 * @param channels    number of channels
 * @param nb_samples  initial allocation size, in samples
 * @return            newly allocated AVAudioFifo, or NULL on error
 */
//AVAudioFifo *av_audio_fifo_alloc(enum AVSampleFormat sample_fmt, int channels,
//int nb_samples);
var avAudioFifoAlloc func(sample_fmt AVSampleFormat, channels, nb_samples ffcommon.FInt) *AVAudioFifo

var avAudioFifoAllocOnce sync.Once

func AvAudioFifoAlloc(sample_fmt AVSampleFormat, channels, nb_samples ffcommon.FInt) *AVAudioFifo {
	avAudioFifoAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoAlloc, ffcommon.GetAvutilDll(), "av_audio_fifo_alloc")
	})
	return avAudioFifoAlloc(sample_fmt, channels, nb_samples)
}

/**
 * Reallocate an AVAudioFifo.
 *
 * @param af          AVAudioFifo to reallocate
 * @param nb_samples  new allocation size, in samples
 * @return            0 if OK, or negative AVERROR code on failure
 */
//av_warn_unused_result
//int av_audio_fifo_realloc(AVAudioFifo *af, int nb_samples);
var avAudioFifoRealloc func(af *AVAudioFifo, nb_samples ffcommon.FInt) ffcommon.FInt

var avAudioFifoReallocOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoRealloc(nb_samples ffcommon.FInt) ffcommon.FInt {
	avAudioFifoReallocOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoRealloc, ffcommon.GetAvutilDll(), "av_audio_fifo_realloc")
	})
	return avAudioFifoRealloc(af, nb_samples)
}

/**
 * Write data to an AVAudioFifo.
 *
 * The AVAudioFifo will be reallocated automatically if the available space
 * is less than nb_samples.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to write to
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to write
 * @return            number of samples actually written, or negative AVERROR
 *                    code on failure. If successful, the number of samples
 *                    actually written will always be nb_samples.
 */
//int av_audio_fifo_write(AVAudioFifo *af, void **data, int nb_samples);
var avAudioFifoWrite func(af *AVAudioFifo, data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt

var avAudioFifoWriteOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoWrite(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt {
	avAudioFifoWriteOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoWrite, ffcommon.GetAvutilDll(), "av_audio_fifo_write")
	})
	return avAudioFifoWrite(af, data, nb_samples)
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_peek(AVAudioFifo *af, void **data, int nb_samples);
var avAudioFifoPeek func(af *AVAudioFifo, data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt

var avAudioFifoPeekOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoPeek(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt {
	avAudioFifoPeekOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoPeek, ffcommon.GetAvutilDll(), "av_audio_fifo_peek")
	})
	return avAudioFifoPeek(af, data, nb_samples)
}

/**
 * Peek data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to peek
 * @param offset      offset from current read position
 * @return            number of samples actually peek, or negative AVERROR code
 *                    on failure. The number of samples actually peek will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_peek_at(AVAudioFifo *af, void **data, int nb_samples, int offset);
var avAudioFifoPeekAt func(af *AVAudioFifo, data *ffcommon.FVoidP, nb_samples, offset ffcommon.FInt) ffcommon.FInt

var avAudioFifoPeekAtOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoPeekAt(data *ffcommon.FVoidP, nb_samples, offset ffcommon.FInt) ffcommon.FInt {
	avAudioFifoPeekAtOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoPeekAt, ffcommon.GetAvutilDll(), "av_audio_fifo_peek_at")
	})
	return avAudioFifoPeekAt(af, data, nb_samples, offset)
}

/**
 * Read data from an AVAudioFifo.
 *
 * @see enum AVSampleFormat
 * The documentation for AVSampleFormat describes the data layout.
 *
 * @param af          AVAudioFifo to read from
 * @param data        audio data plane pointers
 * @param nb_samples  number of samples to read
 * @return            number of samples actually read, or negative AVERROR code
 *                    on failure. The number of samples actually read will not
 *                    be greater than nb_samples, and will only be less than
 *                    nb_samples if av_audio_fifo_size is less than nb_samples.
 */
//int av_audio_fifo_read(AVAudioFifo *af, void **data, int nb_samples);
var avAudioFifoRead func(af *AVAudioFifo, data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt

var avAudioFifoReadOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoRead(data *ffcommon.FVoidP, nb_samples ffcommon.FInt) ffcommon.FInt {
	avAudioFifoReadOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoRead, ffcommon.GetAvutilDll(), "av_audio_fifo_read")
	})
	return avAudioFifoRead(af, data, nb_samples)
}

/**
 * Drain data from an AVAudioFifo.
 *
 * Removes the data without reading it.
 *
 * @param af          AVAudioFifo to drain
 * @param nb_samples  number of samples to drain
 * @return            0 if OK, or negative AVERROR code on failure
 */
//int av_audio_fifo_drain(AVAudioFifo *af, int nb_samples);
var avAudioFifoDrain func(af *AVAudioFifo, nb_samples ffcommon.FInt) ffcommon.FInt

var avAudioFifoDrainOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoDrain(nb_samples ffcommon.FInt) ffcommon.FInt {
	avAudioFifoDrainOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoDrain, ffcommon.GetAvutilDll(), "av_audio_fifo_drain")
	})
	return avAudioFifoDrain(af, nb_samples)
}

/**
 * Reset the AVAudioFifo buffer.
 *
 * This empties all data in the buffer.
 *
 * @param af  AVAudioFifo to reset
 */
//void av_audio_fifo_reset(AVAudioFifo *af);
var avAudioFifoReset func(af *AVAudioFifo)

var avAudioFifoResetOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoReset() {
	avAudioFifoResetOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoReset, ffcommon.GetAvutilDll(), "av_audio_fifo_reset")
	})
	avAudioFifoReset(af)
}

/**
 * Get the current number of samples in the AVAudioFifo available for reading.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for reading
 */
//int av_audio_fifo_size(AVAudioFifo *af);
var avAudioFifoSize func(af *AVAudioFifo) ffcommon.FInt

var avAudioFifoSizeOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoSize() ffcommon.FInt {
	avAudioFifoSizeOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoSize, ffcommon.GetAvutilDll(), "av_audio_fifo_size")
	})
	return avAudioFifoSize(af)
}

/**
 * Get the current number of samples in the AVAudioFifo available for writing.
 *
 * @param af  the AVAudioFifo to query
 * @return    number of samples available for writing
 */
//int av_audio_fifo_space(AVAudioFifo *af);
var avAudioFifoSpace func(af *AVAudioFifo) ffcommon.FInt

var avAudioFifoSpaceOnce sync.Once

func (af *AVAudioFifo) AvAudioFifoSpace() ffcommon.FInt {
	avAudioFifoSpaceOnce.Do(func() {
		purego.RegisterLibFunc(&avAudioFifoSpace, ffcommon.GetAvutilDll(), "av_audio_fifo_space")
	})
	return avAudioFifoSpace(af)
}

/**
 * @}
 * @}
 */

//#endif /* AVUTIL_AUDIO_FIFO_H */
