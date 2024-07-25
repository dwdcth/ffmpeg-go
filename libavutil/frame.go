package libavutil

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

/**
 * @file
 * @ingroup lavu_frame
 * reference-counted frame API
 */

//#ifndef AVUTIL_FRAME_H
//#define AVUTIL_FRAME_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "avutil.h"
//#include "buffer.h"
//#include "dict.h"
//#include "rational.h"
//#include "samplefmt.h"
//#include "pixfmt.h"
//#include "version.h"

/**
 * @defgroup lavu_frame AVFrame
 * @ingroup lavu_data
 *
 * @{
 * AVFrame is an abstraction for reference-counted raw multimedia data.
 */
type AVFrameSideDataType int32

const (
	/**
	 * The data is the AVPanScan struct defined in libavcodec.
	 */
	AV_FRAME_DATA_PANSCAN = iota
	/**
	 * ATSC A53 Part 4 Closed Captions.
	 * A53 CC bitstream is stored as uint8_t in AVFrameSideData.data.
	 * The number of bytes of CC data is AVFrameSideData.size.
	 */
	AV_FRAME_DATA_A53_CC
	/**
	 * Stereoscopic 3d metadata.
	 * The data is the AVStereo3D struct defined in libavutil/stereo3d.h.
	 */
	AV_FRAME_DATA_STEREO3D
	/**
	 * The data is the AVMatrixEncoding enum defined in libavutil/channel_layout.h.
	 */
	AV_FRAME_DATA_MATRIXENCODING
	/**
	 * Metadata relevant to a downmix procedure.
	 * The data is the AVDownmixInfo struct defined in libavutil/downmix_info.h.
	 */
	AV_FRAME_DATA_DOWNMIX_INFO
	/**
	 * ReplayGain information in the form of the AVReplayGain struct.
	 */
	AV_FRAME_DATA_REPLAYGAIN
	/**
	 * This side data contains a 3x3 transformation matrix describing an affine
	 * transformation that needs to be applied to the frame for correct
	 * presentation.
	 *
	 * See libavutil/display.h for a detailed description of the data.
	 */
	AV_FRAME_DATA_DISPLAYMATRIX
	/**
	 * Active Format Description data consisting of a single byte as specified
	 * in ETSI TS 101 154 using AVActiveFormatDescription enum.
	 */
	AV_FRAME_DATA_AFD
	/**
	 * Motion vectors exported by some codecs (on demand through the export_mvs
	 * flag set in the libavcodec AVCodecContext flags2 option).
	 * The data is the AVMotionVector struct defined in
	 * libavutil/motion_vector.h.
	 */
	AV_FRAME_DATA_MOTION_VECTORS
	/**
	 * Recommmends skipping the specified number of samples. This is exported
	 * only if the "skip_manual" AVOption is set in libavcodec.
	 * This has the same format as AV_PKT_DATA_SKIP_SAMPLES.
	 * @code
	 * u32le number of samples to skip from start of this packet
	 * u32le number of samples to skip from end of this packet
	 * u8    reason for start skip
	 * u8    reason for end   skip (0=padding silence, 1=convergence)
	 * @endcode
	 */
	AV_FRAME_DATA_SKIP_SAMPLES
	/**
	 * This side data must be associated with an audio frame and corresponds to
	 * enum AVAudioServiceType defined in avcodec.h.
	 */
	AV_FRAME_DATA_AUDIO_SERVICE_TYPE
	/**
	 * Mastering display metadata associated with a video frame. The payload is
	 * an AVMasteringDisplayMetadata type and contains information about the
	 * mastering display color volume.
	 */
	AV_FRAME_DATA_MASTERING_DISPLAY_METADATA
	/**
	 * The GOP timecode in 25 bit timecode format. Data format is 64-bit integer.
	 * This is set on the first frame of a GOP that has a temporal reference of 0.
	 */
	AV_FRAME_DATA_GOP_TIMECODE

	/**
	 * The data represents the AVSphericalMapping structure defined in
	 * libavutil/spherical.h.
	 */
	AV_FRAME_DATA_SPHERICAL

	/**
	 * Content light level (based on CTA-861.3). This payload contains data in
	 * the form of the AVContentLightMetadata struct.
	 */
	AV_FRAME_DATA_CONTENT_LIGHT_LEVEL

	/**
	 * The data contains an ICC profile as an opaque octet buffer following the
	 * format described by ISO 15076-1 with an optional name defined in the
	 * metadata key entry "name".
	 */
	AV_FRAME_DATA_ICC_PROFILE

	//#if FF_API_FRAME_QP
	/**
	 * Implementation-specific description of the format of AV_FRAME_QP_TABLE_DATA.
	 * The contents of this side data are undocumented and internal; use
	 * av_frame_set_qp_table() and av_frame_get_qp_table() to access this in a
	 * meaningful way instead.
	 */
	AV_FRAME_DATA_QP_TABLE_PROPERTIES

	/**
	 * Raw QP table data. Its format is described by
	 * AV_FRAME_DATA_QP_TABLE_PROPERTIES. Use av_frame_set_qp_table() and
	 * av_frame_get_qp_table() to access this instead.
	 */
	AV_FRAME_DATA_QP_TABLE_DATA
	//#endif

	/**
	 * Timecode which conforms to SMPTE ST 12-1. The data is an array of 4 uint32_t
	 * where the first uint32_t describes how many (1-3) of the other timecodes are used.
	 * The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	 * function in libavutil/timecode.h.
	 */
	AV_FRAME_DATA_S12M_TIMECODE

	/**
	 * HDR dynamic metadata associated with a video frame. The payload is
	 * an AVDynamicHDRPlus type and contains information for color
	 * volume transform - application 4 of SMPTE 2094-40:2016 standard.
	 */
	AV_FRAME_DATA_DYNAMIC_HDR_PLUS

	/**
	 * Regions Of Interest, the data is an array of AVRegionOfInterest type, the number of
	 * array element is implied by AVFrameSideData.size / AVRegionOfInterest.self_size.
	 */
	AV_FRAME_DATA_REGIONS_OF_INTEREST

	/**
	 * Encoding parameters for a video frame, as described by AVVideoEncParams.
	 */
	AV_FRAME_DATA_VIDEO_ENC_PARAMS

	/**
	 * User data unregistered metadata associated with a video frame.
	 * This is the H.26[45] UDU SEI message, and shouldn't be used for any other purpose
	 * The data is stored as uint8_t in AVFrameSideData.data which is 16 bytes of
	 * uuid_iso_iec_11578 followed by AVFrameSideData.size - 16 bytes of user_data_payload_byte.
	 */
	AV_FRAME_DATA_SEI_UNREGISTERED

	/**
	 * Film grain parameters for a frame, described by AVFilmGrainParams.
	 * Must be present for every frame which should have film grain applied.
	 */
	AV_FRAME_DATA_FILM_GRAIN_PARAMS
)

type AVActiveFormatDescription int32

const (
	AV_AFD_SAME         = 8
	AV_AFD_4_3          = 9
	AV_AFD_16_9         = 10
	AV_AFD_14_9         = 11
	AV_AFD_4_3_SP_14_9  = 13
	AV_AFD_16_9_SP_14_9 = 14
	AV_AFD_SP_4_3       = 15
)

/**
 * Structure to hold side data for an AVFrame.
 *
 * sizeof(AVFrameSideData) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 */
type AVFrameSideData struct {
	Type AVFrameSideDataType
	Data *ffcommon.FUint8T
	//#if FF_API_BUFFER_SIZE_T
	//int      size;
	//#else
	//size_t   size;
	//#endif
	Size     ffcommon.FIntOrSizeT
	Metadata *AVDictionary
	Buf      *AVBufferRef
}

/**
 * Structure describing a single Region Of Interest.
 *
 * When multiple regions are defined in a single side-data block, they
 * should be ordered from most to least important - some encoders are only
 * capable of supporting a limited number of distinct regions, so will have
 * to truncate the list.
 *
 * When overlapping regions are defined, the first region containing a given
 * area of the frame applies.
 */
type AVRegionOfInterest struct {

	/**
	 * Must be set to the size of this data structure (that is,
	 * sizeof(AVRegionOfInterest)).
	 */
	SelfSize ffcommon.FUint8T
	/**
	 * Distance in pixels from the top edge of the frame to the top and
	 * bottom edges and from the left edge of the frame to the left and
	 * right edges of the rectangle defining this region of interest.
	 *
	 * The constraints on a region are encoder dependent, so the region
	 * actually affected may be slightly larger for alignment or other
	 * reasons.
	 */
	Top    ffcommon.FInt
	Bottom ffcommon.FInt
	Left   ffcommon.FInt
	Right  ffcommon.FInt
	/**
	 * Quantisation offset.
	 *
	 * Must be in the range -1 to +1.  A value of zero indicates no quality
	 * change.  A negative value asks for better quality (less quantisation),
	 * while a positive value asks for worse quality (greater quantisation).
	 *
	 * The range is calibrated so that the extreme values indicate the
	 * largest possible offset - if the rest of the frame is encoded with the
	 * worst possible quality, an offset of -1 indicates that this region
	 * should be encoded with the best possible quality anyway.  Intermediate
	 * values are then interpolated in some codec-dependent way.
	 *
	 * For example, in 10-bit H.264 the quantisation parameter varies between
	 * -12 and 51.  A typical qoffset value of -1/10 therefore indicates that
	 * this region should be encoded with a QP around one-tenth of the full
	 * range better than the rest of the frame.  So, if most of the frame
	 * were to be encoded with a QP of around 30, this region would get a QP
	 * of around 24 (an offset of approximately -1/10 * (51 - -12) = -6.3).
	 * An extreme value of -1 would indicate that this region should be
	 * encoded with the best possible quality regardless of the treatment of
	 * the rest of the frame - that is, should be encoded at a QP of -12.
	 */
	Qoffset AVRational
}

/**
 * This structure describes decoded (raw) audio or video data.
 *
 * AVFrame must be allocated using av_frame_alloc(). Note that this only
 * allocates the AVFrame itself, the buffers for the data must be managed
 * through other means (see below).
 * AVFrame must be freed with av_frame_free().
 *
 * AVFrame is typically allocated once and then reused multiple times to hold
 * different data (e.g. a single AVFrame to hold frames received from a
 * decoder). In such a case, av_frame_unref() will free any references held by
 * the frame and reset it to its original clean state before it
 * is reused again.
 *
 * The data described by an AVFrame is usually reference counted through the
 * AVBuffer API. The underlying buffer references are stored in AVFrame.buf /
 * AVFrame.extended_buf. An AVFrame is considered to be reference counted if at
 * least one reference is set, i.e. if AVFrame.buf[0] != NULL. In such a case,
 * every single data plane must be contained in one of the buffers in
 * AVFrame.buf or AVFrame.extended_buf.
 * There may be a single buffer for all the data, or one separate buffer for
 * each plane, or anything in between.
 *
 * sizeof(AVFrame) is not a part of the public ABI, so new fields may be added
 * to the end with a minor bump.
 *
 * Fields can be accessed through AVOptions, the name string used, matches the
 * C structure field name for fields accessible through AVOptions. The AVClass
 * for AVFrame can be obtained from avcodec_get_frame_class()
 */
const AV_NUM_DATA_POINTERS = 8

type AVFrame struct {

	/**
	 * pointer to the picture/channel planes.
	 * This might be different from the first allocated byte
	 *
	 * Some decoders access areas outside 0,0 - width,height, please
	 * see avcodec_align_dimensions2(). Some filters and swscale can read
	 * up to 16 bytes beyond the planes, if these filters are to be used,
	 * then 16 extra bytes must be allocated.
	 *
	 * NOTE: Except for hwaccel formats, pointers not needed by the format
	 * MUST be set to NULL.
	 */
	Data [8] /*AV_NUM_DATA_POINTERS*/ *ffcommon.FUint8T

	/**
	 * For video, size in bytes of each picture line.
	 * For audio, size in bytes of each plane.
	 *
	 * For audio, only linesize[0] may be set. For planar audio, each channel
	 * plane must be the same size.
	 *
	 * For video the linesizes should be multiples of the CPUs alignment
	 * preference, this is 16 or 32 for modern desktop CPUs.
	 * Some code requires such alignment other code can be slower without
	 * correct alignment, for yet other it makes no difference.
	 *
	 * @note The linesize may be larger than the size of usable data -- there
	 * may be extra padding present for performance reasons.
	 */
	Linesize [8] /*AV_NUM_DATA_POINTERS*/ ffcommon.FInt

	/**
	 * pointers to the data planes/channels.
	 *
	 * For video, this should simply point to data[].
	 *
	 * For planar audio, each channel has a separate data pointer, and
	 * linesize[0] contains the size of each channel buffer.
	 * For packed audio, there is just one data pointer, and linesize[0]
	 * contains the total size of the buffer for all channels.
	 *
	 * Note: Both data and extended_data should always be set in a valid frame,
	 * but for planar audio with more channels that can fit in data,
	 * extended_data must be used in order to access all channels.
	 */
	ExtendedData **ffcommon.FUint8T

	/**
	 * @name Video dimensions
	 * Video frames only. The coded dimensions (in pixels) of the video frame,
	 * i.e. the size of the rectangle that contains some well-defined values.
	 *
	 * @note The part of the frame intended for display/presentation is further
	 * restricted by the @ref cropping "Cropping rectangle".
	 * @{
	 */
	Width, Height ffcommon.FInt
	/**
	 * @}
	 */

	/**
	 * number of audio samples (per channel) described by this frame
	 */
	NbSamples ffcommon.FInt

	/**
	 * format of the frame, -1 if unknown or unset
	 * Values correspond to enum AVPixelFormat for video frames,
	 * enum AVSampleFormat for audio)
	 */
	Format ffcommon.FInt

	/**
	 * 1 -> keyframe, 0-> not
	 */
	KeyFrame ffcommon.FInt

	/**
	 * Picture type of the frame.
	 */
	PictType AVPictureType

	/**
	 * Sample aspect ratio for the video frame, 0/1 if unknown/unspecified.
	 */
	SampleAspectRatio AVRational

	/**
	 * Presentation timestamp in time_base units (time when frame should be shown to user).
	 */
	Pts ffcommon.FInt64T

	//#if FF_API_PKT_PTS
	/**
	 * PTS copied from the AVPacket that was decoded to produce this frame.
	 * @deprecated use the pts field instead
	 */
	//attribute_deprecated
	//int64_t pkt_pts;
	PktPts ffcommon.FInt64T
	//#endif

	/**
	 * DTS copied from the AVPacket that triggered returning this frame. (if frame threading isn't used)
	 * This is also the Presentation time of this AVFrame calculated from
	 * only AVPacket.dts values without pts values.
	 */
	PktDts ffcommon.FInt64T

	/**
	 * picture number in bitstream order
	 */
	CodedPictureNumber ffcommon.FInt
	/**
	 * picture number in display order
	 */
	DisplayPictureNumber ffcommon.FInt

	/**
	 * quality (between 1 (good) and FF_LAMBDA_MAX (bad))
	 */
	Quality ffcommon.FInt

	/**
	 * for some private data of the user
	 */
	Opaque ffcommon.FVoidP

	//#if FF_API_ERROR_FRAME
	///**
	// * @deprecated unused
	// */
	//attribute_deprecated
	Error [8] /*AV_NUM_DATA_POINTERS*/ ffcommon.FUint64T
	//#endif

	/**
	 * When decoding, this signals how much the picture must be delayed.
	 * extra_delay = repeat_pict / (2*fps)
	 */
	RepeatPict ffcommon.FInt

	/**
	 * The content of the picture is interlaced.
	 */
	InterlacedFrame ffcommon.FInt

	/**
	 * If the content is interlaced, is top field displayed first.
	 */
	TopFieldFirst ffcommon.FInt

	/**
	 * Tell user application that palette has changed from previous frame.
	 */
	PaletteHasChanged ffcommon.FInt

	/**
	 * reordered opaque 64 bits (generally an integer or a double precision float
	 * PTS but can be anything).
	 * The user sets AVCodecContext.reordered_opaque to represent the input at
	 * that time,
	 * the decoder reorders values as needed and sets AVFrame.reordered_opaque
	 * to exactly one of the values provided by the user through AVCodecContext.reordered_opaque
	 */
	ReorderedOpaque ffcommon.FInt64T

	/**
	 * Sample rate of the audio data.
	 */
	SampleRate ffcommon.FInt

	/**
	 * Channel layout of the audio data.
	 */
	ChannelLayout ffcommon.FUint64T

	/**
	 * AVBuffer references backing the data for this frame. If all elements of
	 * this array are NULL, then this frame is not reference counted. This array
	 * must be filled contiguously -- if buf[i] is non-NULL then buf[j] must
	 * also be non-NULL for all j < i.
	 *
	 * There may be at most one AVBuffer per data plane, so for video this array
	 * always contains all the references. For planar audio with more than
	 * AV_NUM_DATA_POINTERS channels, there may be more buffers than can fit in
	 * this array. Then the extra AVBufferRef pointers are stored in the
	 * extended_buf array.
	 */
	Buf [8] /*AV_NUM_DATA_POINTERS*/ *AVBufferRef

	/**
	 * For planar audio which requires more than AV_NUM_DATA_POINTERS
	 * AVBufferRef pointers, this array will hold all the references which
	 * cannot fit into AVFrame.buf.
	 *
	 * Note that this is different from AVFrame.extended_data, which always
	 * contains all the pointers. This array only contains the extra pointers,
	 * which cannot fit into AVFrame.buf.
	 *
	 * This array is always allocated using av_malloc() by whoever constructs
	 * the frame. It is freed in av_frame_unref().
	 */
	ExtendedBuf **AVBufferRef
	/**
	 * Number of elements in extended_buf.
	 */
	NbExtendedBuf ffcommon.FInt

	SideData   **AVFrameSideData
	NbSideData ffcommon.FInt

	/**
	 * @defgroup lavu_frame_flags AV_FRAME_FLAGS
	 * @ingroup lavu_frame
	 * Flags describing additional frame properties.
	 *
	 * @{
	 */

	/**
	 * The frame data may be corrupted, e.g. due to decoding errors.
	 */
	//const AV_FRAME_FLAG_CORRUPT   =    (1 << 0)
	/**
	 * A flag to mark the frames which need to be decoded, but shouldn't be output.
	 */
	//const AV_FRAME_FLAG_DISCARD =  (1 << 2)
	/**
	 * @}
	 */

	/**
	 * Frame flags, a combination of @ref lavu_frame_flags
	 */
	Flags ffcommon.FInt

	/**
	 * MPEG vs JPEG YUV range.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	ColorRange AVColorRange

	ColorPrimaries AVColorPrimaries

	ColorTrc AVColorTransferCharacteristic

	/**
	 * YUV colorspace type.
	 * - encoding: Set by user
	 * - decoding: Set by libavcodec
	 */
	Colorspace AVColorSpace

	ChromaLocation AVChromaLocation

	/**
	 * frame timestamp estimated using various heuristics, in stream time base
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	BestEffortTimestamp ffcommon.FInt64T

	/**
	 * reordered pos from the last AVPacket that has been input into the decoder
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	PktPos ffcommon.FInt64T

	/**
	 * duration of the corresponding packet, expressed in
	 * AVStream->time_base units, 0 if unknown.
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	PktDuration ffcommon.FInt64T

	/**
	 * metadata.
	 * - encoding: Set by user.
	 * - decoding: Set by libavcodec.
	 */
	Metadata *AVDictionary

	/**
	 * decode error flags of the frame, set to a combination of
	 * FF_DECODE_ERROR_xxx flags if the decoder produced a frame, but there
	 * were errors during the decoding.
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	DecodeErrorFlags ffcommon.FInt
	//const FF_DECODE_ERROR_INVALID_BITSTREAM =  1
	//const FF_DECODE_ERROR_MISSING_REFERENCE  = 2
	//const FF_DECODE_ERROR_CONCEALMENT_ACTIVE = 4
	//const FF_DECODE_ERROR_DECODE_SLICES   =    8

	/**
	 * number of audio channels, only used for audio.
	 * - encoding: unused
	 * - decoding: Read by user.
	 */
	Channels ffcommon.FInt

	/**
	 * size of the corresponding packet containing the compressed
	 * frame.
	 * It is set to a negative value if unknown.
	 * - encoding: unused
	 * - decoding: set by libavcodec, read by user.
	 */
	PktSize ffcommon.FInt

	//#if FF_API_FRAME_QP
	/**
	 * QP table
	 */
	//attribute_deprecated
	//int8_t *qscale_table;
	QscaleTable *ffcommon.FInt8T
	/**
	 * QP store stride
	 */
	//attribute_deprecated
	//int qstride;
	Qstride ffcommon.FInt

	//attribute_deprecated
	//int qscale_type;
	QscaleType ffcommon.FInt

	//attribute_deprecated
	//AVBufferRef *qp_table_buf;
	QpTableBuf *AVBufferRef
	//#endif
	/**
	 * For hwaccel-format frames, this should be a reference to the
	 * AVHWFramesContext describing the frame.
	 */
	HwFramesCtx *AVBufferRef

	/**
	 * AVBufferRef for free use by the API user. FFmpeg will never check the
	 * contents of the buffer ref. FFmpeg calls av_buffer_unref() on it when
	 * the frame is unreferenced. av_frame_copy_props() calls create a new
	 * reference with av_buffer_ref() for the target frame's opaque_ref field.
	 *
	 * This is unrelated to the opaque field, although it serves a similar
	 * purpose.
	 */
	OpaqueRef *AVBufferRef

	/**
	 * @anchor cropping
	 * @name Cropping
	 * Video frames only. The number of pixels to discard from the the
	 * top/bottom/left/right border of the frame to obtain the sub-rectangle of
	 * the frame intended for presentation.
	 * @{
	 */
	CropTop    ffcommon.FSizeT
	CropBottom ffcommon.FSizeT
	CropLeft   ffcommon.FSizeT
	CropRight  ffcommon.FSizeT
	/**
	 * @}
	 */

	/**
	 * AVBufferRef for internal use by a single libav* library.
	 * Must not be used to transfer data between libraries.
	 * Has to be NULL when ownership of the frame leaves the respective library.
	 *
	 * Code outside the FFmpeg libs should never check or change the contents of the buffer ref.
	 *
	 * FFmpeg calls av_buffer_unref() on it when the frame is unreferenced.
	 * av_frame_copy_props() calls create a new reference with av_buffer_ref()
	 * for the target frame's private_ref field.
	 */
	PrivateRef *AVBufferRef
}

//#if FF_API_FRAME_GET_SET
/**
 * Accessors for some AVFrame fields. These used to be provided for ABI
 * compatibility, and do not need to be used anymore.
 */
//attribute_deprecated
//int64_t av_frame_get_best_effort_timestamp(const AVFrame *frame);
var avFrameGetBestEffortTimestamp func(frame *AVFrame) ffcommon.FInt64T
var avFrameGetBestEffortTimestampOnce sync.Once

func (frame *AVFrame) AvFrameGetBestEffortTimestamp() ffcommon.FInt64T {
	avFrameGetBestEffortTimestampOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetBestEffortTimestamp,
			ffcommon.GetAvutilDll(),
			"av_frame_get_best_effort_timestamp",
		)
	})
	if avFrameGetBestEffortTimestamp != nil {
		return avFrameGetBestEffortTimestamp(frame)
	}
	return ffcommon.FInt64T(0)
}

// attribute_deprecated
// void    av_frame_set_best_effort_timestamp(AVFrame *frame, int64_t val);
var avFrameSetBestEffortTimestamp func(frame *AVFrame, val ffcommon.FInt64T)
var avFrameSetBestEffortTimestampOnce sync.Once

func (frame *AVFrame) av_frame_set_best_effort_timestamp(val ffcommon.FInt64T) {
	avFrameSetBestEffortTimestampOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetBestEffortTimestamp,
			ffcommon.GetAvutilDll(),
			"av_frame_set_best_effort_timestamp",
		)
	})
	if avFrameSetBestEffortTimestamp != nil {
		avFrameSetBestEffortTimestamp(frame, val)
	}
}

// attribute_deprecated
// int64_t av_frame_get_pkt_duration         (const AVFrame *frame);
var avFrameGetPktDuration func(frame *AVFrame) ffcommon.FInt64T
var avFrameGetPktDurationOnce sync.Once

func (frame *AVFrame) AvFrameGetPktDuration() ffcommon.FInt64T {
	avFrameGetPktDurationOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetPktDuration,
			ffcommon.GetAvutilDll(),
			"av_frame_get_pkt_duration",
		)
	})
	if avFrameGetPktDuration != nil {
		return avFrameGetPktDuration(frame)
	}
	return ffcommon.FInt64T(0)
}

// attribute_deprecated
// void    av_frame_set_pkt_duration         (AVFrame *frame, int64_t val);
var avFrameSetPktDuration func(frame *AVFrame, val ffcommon.FInt64T)
var avFrameSetPktDurationOnce sync.Once

func (frame *AVFrame) AvFrameSetPktDuration(val ffcommon.FInt64T) {
	avFrameSetPktDurationOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetPktDuration,
			ffcommon.GetAvutilDll(),
			"av_frame_set_pkt_duration",
		)
	})
	if avFrameSetPktDuration != nil {
		avFrameSetPktDuration(frame, val)
	}
}

// attribute_deprecated
// int64_t av_frame_get_pkt_pos              (const AVFrame *frame);
var avFrameGetPktPos func(frame *AVFrame) ffcommon.FInt64T
var avFrameGetPktPosOnce sync.Once

func (frame *AVFrame) AvFrameGetPktPos() ffcommon.FInt64T {
	avFrameGetPktPosOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetPktPos,
			ffcommon.GetAvutilDll(),
			"av_frame_get_pkt_pos",
		)
	})
	if avFrameGetPktPos != nil {
		return avFrameGetPktPos(frame)
	}
	return ffcommon.FInt64T(0)
}

// attribute_deprecated
// void    av_frame_set_pkt_pos              (AVFrame *frame, int64_t val);
var avFrameSetPktPos func(frame *AVFrame, val ffcommon.FInt64T)
var avFrameSetPktPosOnce sync.Once

func (frame *AVFrame) AvFrameSetPktPos(val ffcommon.FInt64T) {
	avFrameSetPktPosOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetPktPos,
			ffcommon.GetAvutilDll(),
			"av_frame_set_pkt_pos",
		)
	})
	if avFrameSetPktPos != nil {
		avFrameSetPktPos(frame, val)
	}
}

// attribute_deprecated
// int64_t av_frame_get_channel_layout       (const AVFrame *frame);
var avFrameGetChannelLayout func(frame *AVFrame) ffcommon.FInt64T
var avFrameGetChannelLayoutOnce sync.Once

func (frame *AVFrame) AvFrameGetChannelLayout() ffcommon.FInt64T {
	avFrameGetChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetChannelLayout,
			ffcommon.GetAvutilDll(),
			"av_frame_get_channel_layout",
		)
	})
	if avFrameGetChannelLayout != nil {
		return avFrameGetChannelLayout(frame)
	}
	return ffcommon.FInt64T(0)
}

// attribute_deprecated
// void    av_frame_set_channel_layout       (AVFrame *frame, int64_t val);
var avFrameSetChannelLayout func(frame *AVFrame, val ffcommon.FInt64T)
var avFrameSetChannelLayoutOnce sync.Once

func (frame *AVFrame) AvFrameSetChannelLayout(val ffcommon.FInt64T) {
	avFrameSetChannelLayoutOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetChannelLayout,
			ffcommon.GetAvutilDll(),
			"av_frame_set_channel_layout",
		)
	})
	if avFrameSetChannelLayout != nil {
		avFrameSetChannelLayout(frame, val)
	}
}

// attribute_deprecated
// int     av_frame_get_channels             (const AVFrame *frame);
var avFrameGetChannels func(frame *AVFrame) ffcommon.FInt
var avFrameGetChannelsOnce sync.Once

func (frame *AVFrame) AvFrameGetChannels() ffcommon.FInt {
	avFrameGetChannelsOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetChannels,
			ffcommon.GetAvutilDll(),
			"av_frame_get_channels",
		)
	})
	if avFrameGetChannels != nil {
		return avFrameGetChannels(frame)
	}
	return ffcommon.FInt(0)
}

// attribute_deprecated
// void    av_frame_set_channels             (AVFrame *frame, int     val);
var avFrameSetChannels func(frame *AVFrame, val ffcommon.FInt)
var avFrameSetChannelsOnce sync.Once

func (frame *AVFrame) AvFrameSetChannels(val ffcommon.FInt) {
	avFrameSetChannelsOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetChannels,
			ffcommon.GetAvutilDll(),
			"av_frame_set_channels",
		)
	})
	if avFrameSetChannels != nil {
		avFrameSetChannels(frame, val)
	}
}

// attribute_deprecated
// int     av_frame_get_sample_rate          (const AVFrame *frame);
var avFrameGetSampleRate func(frame *AVFrame) ffcommon.FInt
var avFrameGetSampleRateOnce sync.Once

func (frame *AVFrame) AvFrameGetSampleRate() ffcommon.FInt {
	avFrameGetSampleRateOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetSampleRate,
			ffcommon.GetAvutilDll(),
			"av_frame_get_sample_rate",
		)
	})
	if avFrameGetSampleRate != nil {
		return avFrameGetSampleRate(frame)
	}
	return ffcommon.FInt(0)
}

// attribute_deprecated
// void    av_frame_set_sample_rate          (AVFrame *frame, int     val);
var avFrameSetSampleRate func(frame *AVFrame, val ffcommon.FInt)
var avFrameSetSampleRateOnce sync.Once

func (frame *AVFrame) AvFrameSetSampleRate(val ffcommon.FInt) {
	avFrameSetSampleRateOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetSampleRate,
			ffcommon.GetAvutilDll(),
			"av_frame_set_sample_rate",
		)
	})
	if avFrameSetSampleRate != nil {
		avFrameSetSampleRate(frame, val)
	}
}

// attribute_deprecated
// AVDictionary *av_frame_get_metadata       (const AVFrame *frame);
var avFrameGetMetadata func(frame *AVFrame) *AVDictionary
var avFrameGetMetadataOnce sync.Once

func (frame *AVFrame) AvFrameGetMetadata() *AVDictionary {
	avFrameGetMetadataOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetMetadata,
			ffcommon.GetAvutilDll(),
			"av_frame_get_metadata",
		)
	})
	if avFrameGetMetadata != nil {
		return avFrameGetMetadata(frame)
	}
	return nil
}

// attribute_deprecated
// void          av_frame_set_metadata       (AVFrame *frame, AVDictionary *val);
var avFrameSetMetadata func(frame *AVFrame, val *AVDictionary)
var avFrameSetMetadataOnce sync.Once

func (frame *AVFrame) AvFrameSetMetadata(val *AVDictionary) {
	avFrameSetMetadataOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetMetadata,
			ffcommon.GetAvutilDll(),
			"av_frame_set_metadata",
		)
	})
	if avFrameSetMetadata != nil {
		avFrameSetMetadata(frame, val)
	}
}

// attribute_deprecated
// int     av_frame_get_decode_error_flags   (const AVFrame *frame);
var avFrameGetDecodeErrorFlags func(frame *AVFrame) ffcommon.FInt
var avFrameGetDecodeErrorFlagsOnce sync.Once

func (frame *AVFrame) AvFrameGetDecodeErrorFlags() ffcommon.FInt {
	avFrameGetDecodeErrorFlagsOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetDecodeErrorFlags,
			ffcommon.GetAvutilDll(),
			"av_frame_get_decode_error_flags",
		)
	})
	if avFrameGetDecodeErrorFlags != nil {
		return avFrameGetDecodeErrorFlags(frame)
	}
	return ffcommon.FInt(0)
}

// attribute_deprecated
// void    av_frame_set_decode_error_flags   (AVFrame *frame, int     val);
var avFrameSetDecodeErrorFlags func(frame *AVFrame, val ffcommon.FInt)
var avFrameSetDecodeErrorFlagsOnce sync.Once

func (frame *AVFrame) AvFrameSetDecodeErrorFlags(val ffcommon.FInt) {
	avFrameSetDecodeErrorFlagsOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetDecodeErrorFlags,
			ffcommon.GetAvutilDll(),
			"av_frame_set_decode_error_flags",
		)
	})
	if avFrameSetDecodeErrorFlags != nil {
		avFrameSetDecodeErrorFlags(frame, val)
	}
}

// attribute_deprecated
// int     av_frame_get_pkt_size(const AVFrame *frame);
var avFrameGetPktSize func(frame *AVFrame) ffcommon.FInt
var avFrameGetPktSizeOnce sync.Once

func (frame *AVFrame) AvFrameGetPktSize() ffcommon.FInt {
	avFrameGetPktSizeOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetPktSize,
			ffcommon.GetAvutilDll(),
			"av_frame_get_pkt_size",
		)
	})
	if avFrameGetPktSize != nil {
		return avFrameGetPktSize(frame)
	}
	return ffcommon.FInt(0)
}

// attribute_deprecated
// void    av_frame_set_pkt_size(AVFrame *frame, int val);
var avFrameSetPktSize func(frame *AVFrame, val ffcommon.FInt)
var avFrameSetPktSizeOnce sync.Once

func (frame *AVFrame) AvFrameSetPktSize(val ffcommon.FInt) {
	avFrameSetPktSizeOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetPktSize,
			ffcommon.GetAvutilDll(),
			"av_frame_set_pkt_size",
		)
	})
	if avFrameSetPktSize != nil {
		avFrameSetPktSize(frame, val)
	}
}

// #if FF_API_FRAME_QP
// attribute_deprecated
// int8_t *av_frame_get_qp_table(AVFrame *f, int *stride, int *type);
var avFrameGetQpTable func(f *AVFrame, stride, type0 *ffcommon.FInt) *ffcommon.FInt8T
var avFrameGetQpTableOnce sync.Once

func (f *AVFrame) AvFrameGetQpTable(stride, type0 *ffcommon.FInt) *ffcommon.FInt8T {
	avFrameGetQpTableOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetQpTable,
			ffcommon.GetAvutilDll(),
			"av_frame_get_qp_table",
		)
	})
	if avFrameGetQpTable != nil {
		return avFrameGetQpTable(f, stride, type0)
	}
	return nil
}

// attribute_deprecated
// int av_frame_set_qp_table(AVFrame *f, AVBufferRef *buf, int stride, int type);
var avFrameSetQpTable func(f *AVFrame, buf *AVBufferRef, stride, type0 ffcommon.FInt) *ffcommon.FInt8T
var avFrameSetQpTableOnce sync.Once

func (f *AVFrame) AvFrameSetQpTable(buf *AVBufferRef, stride, type0 ffcommon.FInt) *ffcommon.FInt8T {
	avFrameSetQpTableOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetQpTable,
			ffcommon.GetAvutilDll(),
			"av_frame_set_qp_table",
		)
	})
	if avFrameSetQpTable != nil {
		return avFrameSetQpTable(f, buf, stride, type0)
	}
	return nil
}

// #endif
// attribute_deprecated
// enum AVColorSpace av_frame_get_colorspace(const AVFrame *frame);
var avFrameGetColorspace func(frame *AVFrame) AVColorSpace
var avFrameGetColorspaceOnce sync.Once

func (frame *AVFrame) AvFrameGetColorspace() AVColorSpace {
	avFrameGetColorspaceOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetColorspace,
			ffcommon.GetAvutilDll(),
			"av_frame_get_colorspace",
		)
	})
	if avFrameGetColorspace != nil {
		return avFrameGetColorspace(frame)
	}
	return AVColorSpace(0)
}

// attribute_deprecated
// void    av_frame_set_colorspace(AVFrame *frame, enum AVColorSpace val);
var avFrameSetColorspace func(frame *AVFrame, val AVColorSpace)
var avFrameSetColorspaceOnce sync.Once

func (frame *AVFrame) AvFrameSetColorspace(val AVColorSpace) {
	avFrameSetColorspaceOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetColorspace,
			ffcommon.GetAvutilDll(),
			"av_frame_set_colorspace",
		)
	})
	if avFrameSetColorspace != nil {
		avFrameSetColorspace(frame, val)
	}
}

// attribute_deprecated
// enum AVColorRange av_frame_get_color_range(const AVFrame *frame);
var avFrameGetColorRange func(frame *AVFrame) AVColorRange
var avFrameGetColorRangeOnce sync.Once

func (frame *AVFrame) AvFrameGetColorRange() AVColorRange {
	avFrameGetColorRangeOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetColorRange,
			ffcommon.GetAvutilDll(),
			"av_frame_get_color_range",
		)
	})
	if avFrameGetColorRange != nil {
		return avFrameGetColorRange(frame)
	}
	return AVColorRange(0)
}

// attribute_deprecated
// void    av_frame_set_color_range(AVFrame *frame, enum AVColorRange val);
var avFrameSetColorRange func(frame *AVFrame, val AVColorRange)
var avFrameSetColorRangeOnce sync.Once

func (frame *AVFrame) AvFrameSetColorRange(val AVColorRange) {
	avFrameSetColorRangeOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSetColorRange,
			ffcommon.GetAvutilDll(),
			"av_frame_set_color_range",
		)
	})
	if avFrameSetColorRange != nil {
		avFrameSetColorRange(frame, val)
	}
}

//#endif

/**
 * Get the name of a colorspace.
 * @return a static string identifying the colorspace; can be NULL.
 */
//const char *av_get_colorspace_name(enum AVColorSpace val);
var avGetColorspaceName func(val AVColorSpace) ffcommon.FCharP
var avGetColorspaceNameOnce sync.Once

func AvGetColorspaceName(val AVColorSpace) ffcommon.FCharP {
	avGetColorspaceNameOnce.Do(func() {
		purego.RegisterLibFunc(
			&avGetColorspaceName,
			ffcommon.GetAvutilDll(),
			"av_get_colorspace_name",
		)
	})

	return avGetColorspaceName(val)

}

/**
 * Allocate an AVFrame and set its fields to default values.  The resulting
 * struct must be freed using av_frame_free().
 *
 * @return An AVFrame filled with default values or NULL on failure.
 *
 * @note this only allocates the AVFrame itself, not the data buffers. Those
 * must be allocated through other means, e.g. with av_frame_get_buffer() or
 * manually.
 */
//AVFrame *av_frame_alloc(void);
var avFrameAlloc func() *AVFrame
var avFrameAllocOnce sync.Once

func AvFrameAlloc() *AVFrame {
	avFrameAllocOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameAlloc,
			ffcommon.GetAvutilDll(),
			"av_frame_alloc",
		)
	})
	if avFrameAlloc != nil {
		return avFrameAlloc()
	}
	return nil
}

/**
 * Free the frame and any dynamically allocated objects in it,
 * e.g. extended_data. If the frame is reference counted, it will be
 * unreferenced first.
 *
 * @param frame frame to be freed. The pointer will be set to NULL.
 */
//void av_frame_free(AVFrame **frame);
var avFrameFree func(frame **AVFrame)
var avFrameFreeOnce sync.Once

func AvFrameFree(frame **AVFrame) {
	avFrameFreeOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameFree,
			ffcommon.GetAvutilDll(),
			"av_frame_free",
		)
	})
	if avFrameFree != nil {
		avFrameFree(frame)
	}
}

/**
 * Set up a new reference to the data described by the source frame.
 *
 * Copy frame properties from src to dst and create a new reference for each
 * AVBufferRef from src.
 *
 * If src is not reference counted, new buffers are allocated and the data is
 * copied.
 *
 * @warning: dst MUST have been either unreferenced with av_frame_unref(dst),
 *           or newly allocated with av_frame_alloc() before calling this
 *           function, or undefined behavior will occur.
 *
 * @return 0 on success, a negative AVERROR on error
 */
//int av_frame_ref(AVFrame *dst, const AVFrame *src);
var avFrameRef func(dst, src *AVFrame) ffcommon.FInt
var avFrameRefOnce sync.Once

func AvFrameRef(dst, src *AVFrame) ffcommon.FInt {
	avFrameRefOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameRef,
			ffcommon.GetAvutilDll(),
			"av_frame_ref",
		)
	})
	if avFrameRef != nil {
		return avFrameRef(dst, src)
	}
	return ffcommon.FInt(0)
}

/**
 * Create a new frame that references the same data as src.
 *
 * This is a shortcut for av_frame_alloc()+av_frame_ref().
 *
 * @return newly created AVFrame on success, NULL on error.
 */
//AVFrame *av_frame_clone(const AVFrame *src);
var avFrameClone func(src *AVFrame) *AVFrame
var avFrameCloneOnce sync.Once

func (src *AVFrame) AvFrameClone() *AVFrame {
	avFrameCloneOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameClone,
			ffcommon.GetAvutilDll(),
			"av_frame_clone",
		)
	})
	if avFrameClone != nil {
		return avFrameClone(src)
	}
	return nil
}

/**
 * Unreference all the buffers referenced by frame and reset the frame fields.
 */
//void av_frame_unref(AVFrame *frame);
var avFrameUnref func(src *AVFrame)
var avFrameUnrefOnce sync.Once

func (src *AVFrame) AvFrameUnref() {
	avFrameUnrefOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameUnref,
			ffcommon.GetAvutilDll(),
			"av_frame_unref",
		)
	})
	if avFrameUnref != nil {
		avFrameUnref(src)
	}
}

/**
 * Move everything contained in src to dst and reset src.
 *
 * @warning: dst is not unreferenced, but directly overwritten without reading
 *           or deallocating its contents. Call av_frame_unref(dst) manually
 *           before calling this function to ensure that no memory is leaked.
 */
//void av_frame_move_ref(AVFrame *dst, AVFrame *src);
var avFrameMoveRef func(dst, src *AVFrame) ffcommon.FCharP
var avFrameMoveRefOnce sync.Once

func AvFrameMoveRef(dst, src *AVFrame) ffcommon.FCharP {
	avFrameMoveRefOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameMoveRef,
			ffcommon.GetAvutilDll(),
			"av_frame_move_ref",
		)
	})
	return avFrameMoveRef(dst, src)
}

/**
 * Allocate new buffer(s) for audio or video data.
 *
 * The following fields must be set on frame before calling this function:
 * - format (pixel format for video, sample format for audio)
 * - width and height for video
 * - nb_samples and channel_layout for audio
 *
 * This function will fill AVFrame.data and AVFrame.buf arrays and, if
 * necessary, allocate and fill AVFrame.extended_data and AVFrame.extended_buf.
 * For planar formats, one buffer will be allocated for each plane.
 *
 * @warning: if frame already has been allocated, calling this function will
 *           leak memory. In addition, undefined behavior can occur in certain
 *           cases.
 *
 * @param frame frame in which to store the new buffers.
 * @param align Required buffer size alignment. If equal to 0, alignment will be
 *              chosen automatically for the current CPU. It is highly
 *              recommended to pass 0 here unless you know what you are doing.
 *
 * @return 0 on success, a negative AVERROR on error.
 */
//int av_frame_get_buffer(AVFrame *frame, int align);
var avFrameGetBuffer func(frame *AVFrame, align ffcommon.FInt) ffcommon.FInt
var avFrameGetBufferOnce sync.Once

func (frame *AVFrame) AvFrameGetBuffer(align ffcommon.FInt) ffcommon.FInt {
	avFrameGetBufferOnce.Do(func() {
		purego.RegisterLibFunc(&avFrameGetBuffer, ffcommon.GetAvutilDll(), "av_frame_get_buffer")
	})
	return avFrameGetBuffer(frame, align)
}

/**
 * Check if the frame data is writable.
 *
 * @return A positive value if the frame data is writable (which is true if and
 * only if each of the underlying buffers has only one reference, namely the one
 * stored in this frame). Return 0 otherwise.
 *
 * If 1 is returned the answer is valid until av_buffer_ref() is called on any
 * of the underlying AVBufferRefs (e.g. through av_frame_ref() or directly).
 *
 * @see av_frame_make_writable(), av_buffer_is_writable()
 */
//int av_frame_is_writable(AVFrame *frame);
var avFrameIsWritable func(frame *AVFrame) ffcommon.FInt
var avFrameIsWritableOnce sync.Once

func (frame *AVFrame) AvFrameIsWritable() ffcommon.FInt {
	avFrameIsWritableOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameIsWritable,
			ffcommon.GetAvutilDll(),
			"av_frame_is_writable",
		)
	})
	if avFrameIsWritable != nil {
		return avFrameIsWritable(frame)
	}
	return ffcommon.FInt(0)
}

/**
 * Ensure that the frame data is writable, avoiding data copy if possible.
 *
 * Do nothing if the frame is writable, allocate new buffers and copy the data
 * if it is not.
 *
 * @return 0 on success, a negative AVERROR on error.
 *
 * @see av_frame_is_writable(), av_buffer_is_writable(),
 * av_buffer_make_writable()
 */
//int av_frame_make_writable(AVFrame *frame);
var avFrameMakeWritable func(frame *AVFrame) ffcommon.FInt
var avFrameMakeWritableOnce sync.Once

func (frame *AVFrame) AvFrameMakeWritable() ffcommon.FInt {
	avFrameMakeWritableOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameMakeWritable,
			ffcommon.GetAvutilDll(),
			"av_frame_make_writable",
		)
	})
	if avFrameMakeWritable != nil {
		return avFrameMakeWritable(frame)
	}
	return ffcommon.FInt(0)
}

/**
 * Copy the frame data from src to dst.
 *
 * This function does not allocate anything, dst must be already initialized and
 * allocated with the same parameters as src.
 *
 * This function only copies the frame data (i.e. the contents of the data /
 * extended data arrays), not any other properties.
 *
 * @return >= 0 on success, a negative AVERROR on error.
 */
//int av_frame_copy(AVFrame *dst, const AVFrame *src);
var avFrameCopy func(dst, src *AVFrame) ffcommon.FInt
var avFrameCopyOnce sync.Once

func AvFrameCopy(dst, src *AVFrame) ffcommon.FInt {
	avFrameCopyOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameCopy,
			ffcommon.GetAvutilDll(),
			"av_frame_copy",
		)
	})
	if avFrameCopy != nil {
		return avFrameCopy(dst, src)
	}
	return ffcommon.FInt(0)
}

/**
 * Copy only "metadata" fields from src to dst.
 *
 * Metadata for the purpose of this function are those fields that do not affect
 * the data layout in the buffers.  E.g. pts, sample rate (for audio) or sample
 * aspect ratio (for video), but not width/height or channel layout.
 * Side data is also copied.
 */
//int av_frame_copy_props(AVFrame *dst, const AVFrame *src);
var avFrameCopyProps func(dst, src *AVFrame) ffcommon.FInt
var avFrameCopyPropsOnce sync.Once

func AvFrameCopyProps(dst, src *AVFrame) ffcommon.FInt {
	avFrameCopyPropsOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameCopyProps,
			ffcommon.GetAvutilDll(),
			"av_frame_copy_props",
		)
	})
	if avFrameCopyProps != nil {
		return avFrameCopyProps(dst, src)
	}
	return ffcommon.FInt(0)
}

/**
 * Get the buffer reference a given data plane is stored in.
 *
 * @param plane index of the data plane of interest in frame->extended_data.
 *
 * @return the buffer reference that contains the plane or NULL if the input
 * frame is not valid.
 */
//AVBufferRef *av_frame_get_plane_buffer(AVFrame *frame, int plane);
var avFrameGetPlaneBuffer func(frame *AVFrame, plane ffcommon.FInt) *AVBufferRef
var avFrameGetPlaneBufferOnce sync.Once

func (frame *AVFrame) AvFrameGetPlaneBuffer(plane ffcommon.FInt) *AVBufferRef {
	avFrameGetPlaneBufferOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetPlaneBuffer,
			ffcommon.GetAvutilDll(),
			"av_frame_get_plane_buffer",
		)
	})
	if avFrameGetPlaneBuffer != nil {
		return avFrameGetPlaneBuffer(frame, plane)
	}
	return nil
}

/**
 * Add a new side data to a frame.
 *
 * @param frame a frame to which the side data should be added
 * @param type type of the added side data
 * @param size size of the side data
 *
 * @return newly added side data on success, NULL on error
 */
//AVFrameSideData *av_frame_new_side_data(AVFrame *frame,
//enum AVFrameSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
var avFrameNewSideData func(frame *AVFrame, type0 AVFrameSideDataType, size ffcommon.FIntOrSizeT) *AVFrameSideData
var avFrameNewSideDataOnce sync.Once

func (frame *AVFrame) AvFrameNewSideData(type0 AVFrameSideDataType, size ffcommon.FIntOrSizeT) *AVFrameSideData {
	avFrameNewSideDataOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameNewSideData,
			ffcommon.GetAvutilDll(),
			"av_frame_new_side_data",
		)
	})
	if avFrameNewSideData != nil {
		return avFrameNewSideData(frame, type0, size)
	}
	return nil
}

/**
 * Add a new side data to a frame from an existing AVBufferRef
 *
 * @param frame a frame to which the side data should be added
 * @param type  the type of the added side data
 * @param buf   an AVBufferRef to add as side data. The ownership of
 *              the reference is transferred to the frame.
 *
 * @return newly added side data on success, NULL on error. On failure
 *         the frame is unchanged and the AVBufferRef remains owned by
 *         the caller.
 */
//AVFrameSideData *av_frame_new_side_data_from_buf(AVFrame *frame,
//enum AVFrameSideDataType type,
//AVBufferRef *buf);
var avFrameNewSideDataFromBuf func(frame *AVFrame, type0 AVFrameSideDataType, buf *AVBufferRef) *AVFrameSideData
var avFrameNewSideDataFromBufOnce sync.Once

func (frame *AVFrame) AvFrameNewSideDataFromBuf(type0 AVFrameSideDataType, buf *AVBufferRef) *AVFrameSideData {
	avFrameNewSideDataFromBufOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameNewSideDataFromBuf,
			ffcommon.GetAvutilDll(),
			"av_frame_new_side_data_from_buf",
		)
	})
	if avFrameNewSideDataFromBuf != nil {
		return avFrameNewSideDataFromBuf(frame, type0, buf)
	}
	return nil
}

/**
 * @return a pointer to the side data of a given type on success, NULL if there
 * is no side data with such type in this frame.
 */
//AVFrameSideData *av_frame_get_side_data(const AVFrame *frame,
//enum AVFrameSideDataType type);
var avFrameGetSideData func(frame *AVFrame, type0 AVFrameSideDataType) *AVFrameSideData
var avFrameGetSideDataOnce sync.Once

func (frame *AVFrame) AvFrameGetSideData(type0 AVFrameSideDataType) *AVFrameSideData {
	avFrameGetSideDataOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameGetSideData,
			ffcommon.GetAvutilDll(),
			"av_frame_get_side_data",
		)
	})
	if avFrameGetSideData != nil {
		return avFrameGetSideData(frame, type0)
	}
	return nil
}

/**
 * Remove and free all side data instances of the given type.
 */
//void av_frame_remove_side_data(AVFrame *frame, enum AVFrameSideDataType type);
var avFrameRemoveSideData func(frame *AVFrame, type0 AVFrameSideDataType)
var avFrameRemoveSideDataOnce sync.Once

func (frame *AVFrame) AvFrameRemoveSideData(type0 AVFrameSideDataType) {
	avFrameRemoveSideDataOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameRemoveSideData,
			ffcommon.GetAvutilDll(),
			"av_frame_remove_side_data",
		)
	})
	if avFrameRemoveSideData != nil {
		avFrameRemoveSideData(frame, type0)
	}
}

/**
 * Flags for frame cropping.
 */
//enum {
/**
 * Apply the maximum possible cropping, even if it requires setting the
 * AVFrame.data[] entries to unaligned pointers. Passing unaligned data
 * to FFmpeg API is generally not allowed, and causes undefined behavior
 * (such as crashes). You can pass unaligned data only to FFmpeg APIs that
 * are explicitly documented to accept it. Use this flag only if you
 * absolutely know what you are doing.
 */
const AV_FRAME_CROP_UNALIGNED = 1 << 0

//};

/**
 * Crop the given video AVFrame according to its crop_left/crop_top/crop_right/
 * crop_bottom fields. If cropping is successful, the function will adjust the
 * data pointers and the width/height fields, and set the crop fields to 0.
 *
 * In all cases, the cropping boundaries will be rounded to the inherent
 * alignment of the pixel format. In some cases, such as for opaque hwaccel
 * formats, the left/top cropping is ignored. The crop fields are set to 0 even
 * if the cropping was rounded or ignored.
 *
 * @param frame the frame which should be cropped
 * @param flags Some combination of AV_FRAME_CROP_* flags, or 0.
 *
 * @return >= 0 on success, a negative AVERROR on error. If the cropping fields
 * were invalid, AVERROR(ERANGE) is returned, and nothing is changed.
 */
//int av_frame_apply_cropping(AVFrame *frame, int flags);
var avFrameApplyCropping func(frame *AVFrame, flags ffcommon.FInt) ffcommon.FInt
var avFrameApplyCroppingOnce sync.Once

func (frame *AVFrame) AvFrameApplyCropping(flags ffcommon.FInt) ffcommon.FInt {
	avFrameApplyCroppingOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameApplyCropping,
			ffcommon.GetAvutilDll(),
			"av_frame_apply_cropping",
		)
	})
	if avFrameApplyCropping != nil {
		return avFrameApplyCropping(frame, flags)
	}
	return 0 // Return a default value if the function is not registered
}

/**
 * @return a string identifying the side data type
 */
//const char *av_frame_side_data_name(enum AVFrameSideDataType type);
var avFrameSideDataName func(type0 AVFrameSideDataType) ffcommon.FCharP
var avFrameSideDataNameOnce sync.Once

func AvFrameSideDataName(type0 AVFrameSideDataType) ffcommon.FCharP {
	avFrameSideDataNameOnce.Do(func() {
		purego.RegisterLibFunc(
			&avFrameSideDataName,
			ffcommon.GetAvutilDll(),
			"av_frame_side_data_name",
		)
	})
	if avFrameSideDataName != nil {
		return avFrameSideDataName(type0)
	}
	return ""
}

/**
 * @}
 */

//#endif /* AVUTIL_FRAME_H */
