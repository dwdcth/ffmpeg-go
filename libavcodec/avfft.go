package libavcodec

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
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

//#ifndef AVCODEC_AVFFT_H
//#define AVCODEC_AVFFT_H

/**
 * @file
 * @ingroup lavc_fft
 * FFT functions
 */

/**
 * @defgroup lavc_fft FFT functions
 * @ingroup lavc_misc
 *
 * @{
 */

//type  FFTSample=ffcommon.FFloat

type FFTComplex struct {
	Re, Im ffcommon.FFloat
}

// typedef struct FFTContext FFTContext;
type FFTContext struct {
}

/**
 * Set up a complex FFT.
 * @param nbits           log2 of the length of the input array
 * @param inverse         if 0 perform the forward transform, if 1 perform the inverse
 */
//FFTContext *av_fft_init(int nbits, int inverse);
var avFftInitFunc func(nbits, inverse ffcommon.FInt) *FFTContext
var avFftInitFuncOnce sync.Once

func AvFftInit(nbits, inverse ffcommon.FInt) (res *FFTContext) {
	avFftInitFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avFftInitFunc, ffcommon.GetAvcodecDll(), "av_fft_init")
	})

	res = avFftInitFunc(nbits, inverse)
	return
}

/**
 * Do the permutation needed BEFORE calling ff_fft_calc().
 */
//void av_fft_permute(FFTContext *s, FFTComplex *z);
var avFftPermuteFunc func(s *FFTContext, z *FFTComplex)
var avFftPermuteFuncOnce sync.Once

func (s *FFTContext) AvFftPermute(z *FFTComplex) {
	avFftPermuteFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avFftPermuteFunc, ffcommon.GetAvcodecDll(), "av_fft_permute")
	})

	avFftPermuteFunc(s, z)
}

/**
 * Do a complex FFT with the parameters defined in av_fft_init(). The
 * input data must be permuted before. No 1.0/sqrt(n) normalization is done.
 */
//void av_fft_calc(FFTContext *s, FFTComplex *z);
var avFftCalcFunc func(s *FFTContext, z *FFTComplex)
var avFftCalcFuncOnce sync.Once

func (s *FFTContext) AvFftCalc(z *FFTComplex) {
	avFftCalcFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avFftCalcFunc, ffcommon.GetAvcodecDll(), "av_fft_calc")
	})

	avFftCalcFunc(s, z)
}

// void av_fft_end(FFTContext *s);
var avFftEndFunc func(s *FFTContext)
var avFftEndFuncOnce sync.Once

func (s *FFTContext) AvFftEnd() {
	avFftEndFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avFftEndFunc, ffcommon.GetAvcodecDll(), "av_fft_end")
	})

	avFftEndFunc(s)
}

// FFTContext *av_mdct_init(int nbits, int inverse, double scale);
var avMdctInitFunc func(nbits, inverse ffcommon.FInt, scale ffcommon.FDouble) *FFTContext
var avMdctInitFuncOnce sync.Once

func AvMdctInit(nbits, inverse ffcommon.FInt, scale ffcommon.FDouble) (res *FFTContext) {
	avMdctInitFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avMdctInitFunc, ffcommon.GetAvcodecDll(), "av_mdct_init")
	})

	res = avMdctInitFunc(nbits, inverse, scale)
	return
}

// void av_imdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
var avImdctCalcFunc func(s *FFTContext, output, input *ffcommon.FFTSample)
var avImdctCalcFuncOnce sync.Once

func (s *FFTContext) AvImdctCalc(output, input *ffcommon.FFTSample) {
	avImdctCalcFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avImdctCalcFunc, ffcommon.GetAvcodecDll(), "av_imdct_calc")
	})

	avImdctCalcFunc(s, output, input)
}

// void av_imdct_half(FFTContext *s, FFTSample *output, const FFTSample *input);
var avImdctHalfFunc func(s *FFTContext, output, input *ffcommon.FFTSample)
var avImdctHalfFuncOnce sync.Once

func (s *FFTContext) AvImdctHalf(output, input *ffcommon.FFTSample) {
	avImdctHalfFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avImdctHalfFunc, ffcommon.GetAvcodecDll(), "av_imdct_half")
	})

	avImdctHalfFunc(s, output, input)
}

// void av_mdct_calc(FFTContext *s, FFTSample *output, const FFTSample *input);
var avMdctCalcFunc func(s *FFTContext, output, input *ffcommon.FFTSample)
var avMdctCalcFuncOnce sync.Once

func (s *FFTContext) AvMdctCalc(output, input *ffcommon.FFTSample) {
	avMdctCalcFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avMdctCalcFunc, ffcommon.GetAvcodecDll(), "av_mdct_calc")
	})

	avMdctCalcFunc(s, output, input)
}

// void av_mdct_end(FFTContext *s);
var avMdctEndFunc func(s *FFTContext)
var avMdctEndFuncOnce sync.Once

func (s *FFTContext) AvMdctEnd(output, input *ffcommon.FFTSample) {
	avMdctEndFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avMdctEndFunc, ffcommon.GetAvcodecDll(), "av_mdct_end")
	})

	avMdctEndFunc(s)
}

/* Real Discrete Fourier Transform */
type RDFTransformType = int32

const (
	DFT_R2C = iota
	IDFT_C2R
	IDFT_R2C
	DFT_C2R
)

// typedef struct RDFTContext RDFTContext;
type RDFTContext struct {
}

/**
 * Set up a real FFT.
 * @param nbits           log2 of the length of the input array
 * @param trans           the type of transform
 */
//RDFTContext *av_rdft_init(int nbits, enum RDFTransformType trans);
var avRdftInitFunc func(nbits ffcommon.FInt, trans RDFTransformType) *RDFTContext
var avRdftInitFuncOnce sync.Once

func AvRdftInit(nbits ffcommon.FInt, trans RDFTransformType) (res *RDFTContext) {
	avRdftInitFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avRdftInitFunc, ffcommon.GetAvcodecDll(), "av_rdft_init")
	})

	res = avRdftInitFunc(nbits, trans)
	return
}

// void av_rdft_calc(RDFTContext *s, FFTSample *data);
var avRdftCalcFunc func(s *RDFTContext, data *ffcommon.FFTSample)
var avRdftCalcFuncOnce sync.Once

func (s *RDFTContext) AvRdftCalc(data *ffcommon.FFTSample) {
	avRdftCalcFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avRdftCalcFunc, ffcommon.GetAvcodecDll(), "av_rdft_calc")
	})

	avRdftCalcFunc(s, data)
}

// void av_rdft_end(RDFTContext *s);
var avRdftEndFunc func(s *RDFTContext)
var avRdftEndFuncOnce sync.Once

func (s *RDFTContext) AvRdftEnd() {
	avRdftEndFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avRdftEndFunc, ffcommon.GetAvcodecDll(), "av_rdft_end")
	})

	avRdftEndFunc(s)
}

/* Discrete Cosine Transform */

// typedef struct DCTContext DCTContext;
type DCTContext struct {
}
type DCTTransformType = int32

const (
	DCT_II = iota
	DCT_III
	DCT_I
	DST_I
)

/**
 * Set up DCT.
 *
 * @param nbits           size of the input array:
 *                        (1 << nbits)     for DCT-II, DCT-III and DST-I
 *                        (1 << nbits) + 1 for DCT-I
 * @param type            the type of transform
 *
 * @note the first element of the input of DST-I is ignored
 */
//DCTContext *av_dct_init(int nbits, enum DCTTransformType type);
var avDctInitFunc func(nbits ffcommon.FInt, type0 DCTTransformType) *DCTContext
var avDctInitFuncOnce sync.Once

func AvDctInit(nbits ffcommon.FInt, type0 DCTTransformType) (res *DCTContext) {
	avDctInitFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avDctInitFunc, ffcommon.GetAvcodecDll(), "av_dct_init")
	})

	res = avDctInitFunc(nbits, type0)
	return
}

// void av_dct_calc(DCTContext *s, FFTSample *data);
var avDctCalcFunc func(s *DCTContext, data *ffcommon.FFTSample)
var avDctCalcFuncOnce sync.Once

func (s *DCTContext) AvDctCalc(data *ffcommon.FFTSample) {
	avDctCalcFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avDctCalcFunc, ffcommon.GetAvcodecDll(), "av_dct_calc")
	})

	avDctCalcFunc(s, data)
}

// void av_dct_end (DCTContext *s);
var avDctEndFunc func(s *DCTContext)
var avDctEndFuncOnce sync.Once

func (s *DCTContext) AvDctEnd() {
	avDctEndFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avDctEndFunc, ffcommon.GetAvcodecDll(), "av_dct_end")
	})

	avDctEndFunc(s)
}

/**
 * @}
 */

//#endif /* AVCODEC_AVFFT_H */
