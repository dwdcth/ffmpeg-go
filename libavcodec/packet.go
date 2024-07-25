package libavcodec

import (
	"sync"
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/v7/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v7/libavutil"
	"github.com/ebitengine/purego"
)

/*
 * AVPacket public API
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

//#ifndef AVCODEC_PACKET_H
//#define AVCODEC_PACKET_H
//
//#include <stddef.h>
//#include <stdint.h>
//
//#include "../libavutil/attributes.h"
//#include "../libavutil/buffer.h"
//#include "../libavutil/dict.h"
//#include "../libavutil/rational.h"
//
//#include "../libavcodec/version.h"

/**
 * @defgroup lavc_packet AVPacket
 *
 * Types and functions for working with AVPacket.
 * @{
 */
type AVPacketSideDataType int32

const (
	/**
	 * An AV_PKT_DATA_PALETTE side data packet contains exactly AVPALETTE_SIZE
	 * bytes worth of palette. This side data signals that a new palette is
	 * present.
	 */
	AV_PKT_DATA_PALETTE = iota

	/**
	 * The AV_PKT_DATA_NEW_EXTRADATA is used to notify the codec or the format
	 * that the extradata buffer was changed and the receiving side should
	 * act upon it appropriately. The new extradata is embedded in the side
	 * data buffer and should be immediately used for processing the current
	 * frame or packet.
	 */
	AV_PKT_DATA_NEW_EXTRADATA

	/**
	 * An AV_PKT_DATA_PARAM_CHANGE side data packet is laid out as follows:
	 * @code
	 * u32le param_flags
	 * if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT)
	 *     s32le channel_count
	 * if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT)
	 *     u64le channel_layout
	 * if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE)
	 *     s32le sample_rate
	 * if (param_flags & AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS)
	 *     s32le width
	 *     s32le height
	 * @endcode
	 */
	AV_PKT_DATA_PARAM_CHANGE

	/**
	 * An AV_PKT_DATA_H263_MB_INFO side data packet contains a number of
	 * structures with info about macroblocks relevant to splitting the
	 * packet into smaller packets on macroblock edges (e.g. as for RFC 2190).
	 * That is, it does not necessarily contain info about all macroblocks,
	 * as long as the distance between macroblocks in the info is smaller
	 * than the target payload size.
	 * Each MB info structure is 12 bytes, and is laid out as follows:
	 * @code
	 * u32le bit offset from the start of the packet
	 * u8    current quantizer at the start of the macroblock
	 * u8    GOB number
	 * u16le macroblock address within the GOB
	 * u8    horizontal MV predictor
	 * u8    vertical MV predictor
	 * u8    horizontal MV predictor for block number 3
	 * u8    vertical MV predictor for block number 3
	 * @endcode
	 */
	AV_PKT_DATA_H263_MB_INFO

	/**
	 * This side data should be associated with an audio stream and contains
	 * ReplayGain information in form of the AVReplayGain struct.
	 */
	AV_PKT_DATA_REPLAYGAIN

	/**
	 * This side data contains a 3x3 transformation matrix describing an affine
	 * transformation that needs to be applied to the decoded video frames for
	 * correct presentation.
	 *
	 * See libavutil/display.h for a detailed description of the data.
	 */
	AV_PKT_DATA_DISPLAYMATRIX

	/**
	 * This side data should be associated with a video stream and contains
	 * Stereoscopic 3D information in form of the AVStereo3D struct.
	 */
	AV_PKT_DATA_STEREO3D

	/**
	 * This side data should be associated with an audio stream and corresponds
	 * to enum AVAudioServiceType.
	 */
	AV_PKT_DATA_AUDIO_SERVICE_TYPE

	/**
	 * This side data contains quality related information from the encoder.
	 * @code
	 * u32le quality factor of the compressed frame. Allowed range is between 1 (good) and FF_LAMBDA_MAX (bad).
	 * u8    picture type
	 * u8    error count
	 * u16   reserved
	 * u64le[error count] sum of squared differences between encoder in and output
	 * @endcode
	 */
	AV_PKT_DATA_QUALITY_STATS

	/**
	 * This side data contains an integer value representing the stream index
	 * of a "fallback" track.  A fallback track indicates an alternate
	 * track to use when the current track can not be decoded for some reason.
	 * e.g. no decoder available for codec.
	 */
	AV_PKT_DATA_FALLBACK_TRACK

	/**
	 * This side data corresponds to the AVCPBProperties struct.
	 */
	AV_PKT_DATA_CPB_PROPERTIES

	/**
	 * Recommmends skipping the specified number of samples
	 * @code
	 * u32le number of samples to skip from start of this packet
	 * u32le number of samples to skip from end of this packet
	 * u8    reason for start skip
	 * u8    reason for end   skip (0=padding silence, 1=convergence)
	 * @endcode
	 */
	AV_PKT_DATA_SKIP_SAMPLES

	/**
	 * An AV_PKT_DATA_JP_DUALMONO side data packet indicates that
	 * the packet may contain "dual mono" audio specific to Japanese DTV
	 * and if it is true, recommends only the selected channel to be used.
	 * @code
	 * u8    selected channels (0=mail/left, 1=sub/right, 2=both)
	 * @endcode
	 */
	AV_PKT_DATA_JP_DUALMONO

	/**
	 * A list of zero terminated key/value strings. There is no end marker for
	 * the list, so it is required to rely on the side data size to stop.
	 */
	AV_PKT_DATA_STRINGS_METADATA

	/**
	 * Subtitle event position
	 * @code
	 * u32le x1
	 * u32le y1
	 * u32le x2
	 * u32le y2
	 * @endcode
	 */
	AV_PKT_DATA_SUBTITLE_POSITION

	/**
	 * Data found in BlockAdditional element of matroska container. There is
	 * no end marker for the data, so it is required to rely on the side data
	 * size to recognize the end. 8 byte id (as found in BlockAddId) followed
	 * by data.
	 */
	AV_PKT_DATA_MATROSKA_BLOCKADDITIONAL

	/**
	 * The optional first identifier line of a WebVTT cue.
	 */
	AV_PKT_DATA_WEBVTT_IDENTIFIER

	/**
	 * The optional settings (rendering instructions) that immediately
	 * follow the timestamp specifier of a WebVTT cue.
	 */
	AV_PKT_DATA_WEBVTT_SETTINGS

	/**
	 * A list of zero terminated key/value strings. There is no end marker for
	 * the list, so it is required to rely on the side data size to stop. This
	 * side data includes updated metadata which appeared in the stream.
	 */
	AV_PKT_DATA_METADATA_UPDATE

	/**
	 * MPEGTS stream ID as uint8_t, this is required to pass the stream ID
	 * information from the demuxer to the corresponding muxer.
	 */
	AV_PKT_DATA_MPEGTS_STREAM_ID

	/**
	 * Mastering display metadata (based on SMPTE-2086:2014). This metadata
	 * should be associated with a video stream and contains data in the form
	 * of the AVMasteringDisplayMetadata struct.
	 */
	AV_PKT_DATA_MASTERING_DISPLAY_METADATA

	/**
	 * This side data should be associated with a video stream and corresponds
	 * to the AVSphericalMapping structure.
	 */
	AV_PKT_DATA_SPHERICAL

	/**
	 * Content light level (based on CTA-861.3). This metadata should be
	 * associated with a video stream and contains data in the form of the
	 * AVContentLightMetadata struct.
	 */
	AV_PKT_DATA_CONTENT_LIGHT_LEVEL

	/**
	 * ATSC A53 Part 4 Closed Captions. This metadata should be associated with
	 * a video stream. A53 CC bitstream is stored as uint8_t in AVPacketSideData.data.
	 * The number of bytes of CC data is AVPacketSideData.size.
	 */
	AV_PKT_DATA_A53_CC

	/**
	 * This side data is encryption initialization data.
	 * The format is not part of ABI, use av_encryption_init_info_* methods to
	 * access.
	 */
	AV_PKT_DATA_ENCRYPTION_INIT_INFO

	/**
	 * This side data contains encryption info for how to decrypt the packet.
	 * The format is not part of ABI, use av_encryption_info_* methods to access.
	 */
	AV_PKT_DATA_ENCRYPTION_INFO

	/**
	 * Active Format Description data consisting of a single byte as specified
	 * in ETSI TS 101 154 using AVActiveFormatDescription enum.
	 */
	AV_PKT_DATA_AFD

	/**
	 * Producer Reference Time data corresponding to the AVProducerReferenceTime struct,
	 * usually exported by some encoders (on demand through the prft flag set in the
	 * AVCodecContext export_side_data field).
	 */
	AV_PKT_DATA_PRFT

	/**
	 * ICC profile data consisting of an opaque octet buffer following the
	 * format described by ISO 15076-1.
	 */
	AV_PKT_DATA_ICC_PROFILE

	/**
	 * DOVI configuration
	 * ref:
	 * dolby-vision-bitstreams-within-the-iso-base-media-file-format-v2.1.2, section 2.2
	 * dolby-vision-bitstreams-in-mpeg-2-transport-stream-multiplex-v1.2, section 3.3
	 * Tags are stored in struct AVDOVIDecoderConfigurationRecord.
	 */
	AV_PKT_DATA_DOVI_CONF

	/**
	 * Timecode which conforms to SMPTE ST 12-1:2014. The data is an array of 4 uint32_t
	 * where the first uint32_t describes how many (1-3) of the other timecodes are used.
	 * The timecode format is described in the documentation of av_timecode_get_smpte_from_framenum()
	 * function in libavutil/timecode.h.
	 */
	AV_PKT_DATA_S12M_TIMECODE

	/**
	 * The number of side data types.
	 * This is not part of the public API/ABI in the sense that it may
	 * change when new side data types are added.
	 * This must stay the last enum value.
	 * If its value becomes huge, some code using it
	 * needs to be updated as it assumes it to be smaller than other limits.
	 */
	AV_PKT_DATA_NB
)

const AV_PKT_DATA_QUALITY_FACTOR = AV_PKT_DATA_QUALITY_STATS //DEPRECATED

type AVPacketSideData struct {
	Data *ffcommon.FUint8T
	//#if FF_API_BUFFER_SIZE_T
	//int      size;
	//#else
	//size_t   size;
	//#endif
	Size ffcommon.FIntOrSizeT
	Type AVPacketSideDataType
}

/**
 * This structure stores compressed data. It is typically exported by demuxers
 * and then passed as input to decoders, or received as output from encoders and
 * then passed to muxers.
 *
 * For video, it should typically contain one compressed frame. For audio it may
 * contain several compressed frames. Encoders are allowed to output empty
 * packets, with no compressed data, containing only side data
 * (e.g. to update some stream parameters at the end of encoding).
 *
 * The semantics of data ownership depends on the buf field.
 * If it is set, the packet data is dynamically allocated and is
 * valid indefinitely until a call to av_packet_unref() reduces the
 * reference count to 0.
 *
 * If the buf field is not set av_packet_ref() would make a copy instead
 * of increasing the reference count.
 *
 * The side data is always allocated with av_malloc(), copied by
 * av_packet_ref() and freed by av_packet_unref().
 *
 * sizeof(AVPacket) being a part of the public ABI is deprecated. once
 * av_init_packet() is removed, new packets will only be able to be allocated
 * with av_packet_alloc(), and new fields may be added to the end of the struct
 * with a minor bump.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 * @see av_packet_unref
 */
type AVBufferRef = libavutil.AVBufferRef
type AVPacket struct {

	/**
	 * A reference to the reference-counted buffer where the packet data is
	 * stored.
	 * May be NULL, then the packet data is not reference-counted.
	 */
	Buf *AVBufferRef
	/**
	 * Presentation timestamp in AVStream->time_base units; the time at which
	 * the decompressed packet will be presented to the user.
	 * Can be AV_NOPTS_VALUE if it is not stored in the file.
	 * pts MUST be larger or equal to dts as presentation cannot happen before
	 * decompression, unless one wants to view hex dumps. Some formats misuse
	 * the terms dts and pts/cts to mean something different. Such timestamps
	 * must be converted to true pts/dts before they are stored in AVPacket.
	 */
	Pts ffcommon.FInt64T
	/**
	 * Decompression timestamp in AVStream->time_base units; the time at which
	 * the packet is decompressed.
	 * Can be AV_NOPTS_VALUE if it is not stored in the file.
	 */
	Dts         ffcommon.FInt64T
	Data        *ffcommon.FUint8T
	Size        ffcommon.FUint
	StreamIndex ffcommon.FUint
	/**
	 * A combination of AV_PKT_FLAG values
	 */
	Flags ffcommon.FUint
	/**
	 * Additional packet data that can be provided by the container.
	 * Packet can contain several types of side information.
	 */
	SideData      *AVPacketSideData
	SideDataElems ffcommon.FInt

	/**
	 * Duration of this packet in AVStream->time_base units, 0 if unknown.
	 * Equals next_pts - this_pts in presentation order.
	 */
	Duration ffcommon.FInt64T

	Pos ffcommon.FInt64T ///< byte position in stream, -1 if unknown

	//#if FF_API_CONVERGENCE_DURATION
	/**
	 * @deprecated Same as the duration field, but as int64_t. This was required
	 * for Matroska subtitles, whose duration values could overflow when the
	 * duration field was still an int.
	 */
	//attribute_deprecated
	//int64_t convergence_duration;
	ConvergenceDuration ffcommon.FInt64T
	//#endif
}

// #if FF_API_INIT_PACKET
// attribute_deprecated
type AVPacketList struct {
	Pkt  AVPacket
	Next *AVPacketList
}

//#endif

const AV_PKT_FLAG_KEY = 0x0001     ///< The packet contains a keyframe
const AV_PKT_FLAG_CORRUPT = 0x0002 ///< The packet content is corrupted
/**
 * Flag is used to discard packets which are required to maintain valid
 * decoder state but are not required for output and should be dropped
 * after decoding.
 **/
const AV_PKT_FLAG_DISCARD = 0x0004

/**
 * The packet comes from a trusted source.
 *
 * Otherwise-unsafe constructs such as arbitrary pointers to data
 * outside the packet may be followed.
 */
const AV_PKT_FLAG_TRUSTED = 0x0008

/**
 * Flag is used to indicate packets that contain frames that can
 * be discarded by the decoder.  I.e. Non-reference frames.
 */
const AV_PKT_FLAG_DISPOSABLE = 0x0010

type AVSideDataParamChangeFlags = int32

const (
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_COUNT  = 0x0001
	AV_SIDE_DATA_PARAM_CHANGE_CHANNEL_LAYOUT = 0x0002
	AV_SIDE_DATA_PARAM_CHANGE_SAMPLE_RATE    = 0x0004
	AV_SIDE_DATA_PARAM_CHANGE_DIMENSIONS     = 0x0008
)

/**
 * Allocate an AVPacket and set its fields to default values.  The resulting
 * struct must be freed using av_packet_free().
 *
 * @return An AVPacket filled with default values or NULL on failure.
 *
 * @note this only allocates the AVPacket itself, not the data buffers. Those
 * must be allocated through other means such as av_new_packet.
 *
 * @see av_new_packet
 */
//AVPacket *av_packet_alloc(void);
var avPacketAlloc func() *AVPacket
var avPacketAllocOnce sync.Once

func AvPacketAlloc() *AVPacket {
	avPacketAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketAlloc, ffcommon.GetAvcodecDll(), "av_packet_alloc")
	})
	return avPacketAlloc()
}

/**
 * Create a new packet that references the same data as src.
 *
 * This is a shortcut for av_packet_alloc()+av_packet_ref().
 *
 * @return newly created AVPacket on success, NULL on error.
 *
 * @see av_packet_alloc
 * @see av_packet_ref
 */
//AVPacket *av_packet_clone(const AVPacket *src);
var avPacketClone func(src *AVPacket) *AVPacket
var avPacketCloneOnce sync.Once

func (src *AVPacket) AvPacketClone() *AVPacket {
	avPacketCloneOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketClone, ffcommon.GetAvcodecDll(), "av_packet_clone")
	})
	return avPacketClone(src)
}

/**
 * Free the packet, if the packet is reference counted, it will be
 * unreferenced first.
 *
 * @param pkt packet to be freed. The pointer will be set to NULL.
 * @note passing NULL is a no-op.
 */
//void av_packet_free(AVPacket **pkt);
var avPacketFree func(pkt **AVPacket)
var avPacketFreeOnce sync.Once

func AvPacketFree(pkt **AVPacket) {
	avPacketFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketFree, ffcommon.GetAvcodecDll(), "av_packet_free")
	})
	avPacketFree(pkt)
}

//#if FF_API_INIT_PACKET
/**
  * Initialize optional fields of a packet with default values.
  *
  * Note, this does not touch the data and size members, which have to be
  * initialized separately.
  *
  * @param pkt packet
  *
  * @see av_packet_alloc
  * @see av_packet_unref
  *
  * @deprecated This function is deprecated. Once it's removed,
                sizeof(AVPacket) will not be a part of the ABI anymore.
*/
//attribute_deprecated
//void av_init_packet(AVPacket *pkt);
var avInitPacket func(pkt *AVPacket)
var avInitPacketOnce sync.Once

func (pkt *AVPacket) AvInitPacket() {
	avInitPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avInitPacket, ffcommon.GetAvcodecDll(), "av_init_packet")
	})
	avInitPacket(pkt)
}

//#endif

/**
 * Allocate the payload of a packet and initialize its fields with
 * default values.
 *
 * @param pkt packet
 * @param size wanted payload size
 * @return 0 if OK, AVERROR_xxx otherwise
 */
//int av_new_packet(AVPacket *pkt, int size);
var avNewPacket func(pkt *AVPacket, size ffcommon.FInt) ffcommon.FInt
var avNewPacketOnce sync.Once

func (pkt *AVPacket) AvNewPacket(size ffcommon.FInt) ffcommon.FInt {
	avNewPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avNewPacket, ffcommon.GetAvcodecDll(), "av_new_packet")
	})
	return avNewPacket(pkt, size)
}

/**
 * Reduce packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param size new size
 */
//void av_shrink_packet(AVPacket *pkt, int size);
var avShrinkPacket func(pkt *AVPacket, size ffcommon.FInt)
var avShrinkPacketOnce sync.Once

func (pkt *AVPacket) AvShrinkPacket(size ffcommon.FInt) {
	avShrinkPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avShrinkPacket, ffcommon.GetAvcodecDll(), "av_shrink_packet")
	})
	avShrinkPacket(pkt, size)
}

/**
 * Increase packet size, correctly zeroing padding
 *
 * @param pkt packet
 * @param grow_by number of bytes by which to increase the size of the packet
 */
//int av_grow_packet(AVPacket *pkt, int grow_by);
var avGrowPacket func(pkt *AVPacket, size ffcommon.FInt) ffcommon.FInt
var avGrowPacketOnce sync.Once

func (pkt *AVPacket) AvGrowPacket(size ffcommon.FInt) ffcommon.FInt {
	avGrowPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avGrowPacket, ffcommon.GetAvcodecDll(), "av_grow_packet")
	})
	return avGrowPacket(pkt, size)
}

/**
 * Initialize a reference-counted packet from av_malloc()ed data.
 *
 * @param pkt packet to be initialized. This function will set the data, size,
 *        and buf fields, all others are left untouched.
 * @param data Data allocated by av_malloc() to be used as packet data. If this
 *        function returns successfully, the data is owned by the underlying AVBuffer.
 *        The caller may not access the data through other means.
 * @param size size of data in bytes, without the padding. I.e. the full buffer
 *        size is assumed to be size + AV_INPUT_BUFFER_PADDING_SIZE.
 *
 * @return 0 on success, a negative AVERROR on error
 */
//int av_packet_from_data(AVPacket *pkt, uint8_t *data, int size);
var avPacketFromData func(pkt *AVPacket, data *ffcommon.FUint8T, size ffcommon.FInt) ffcommon.FInt
var avPacketFromDataOnce sync.Once

func (pkt *AVPacket) AvPacketFromData(data *ffcommon.FUint8T, size ffcommon.FInt) ffcommon.FInt {
	avPacketFromDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketFromData, ffcommon.GetAvcodecDll(), "av_packet_from_data")
	})
	return avPacketFromData(pkt, data, size)
}

//#if FF_API_AVPACKET_OLD_API
/**
 * @warning This is a hack - the packet memory allocation stuff is broken. The
 * packet is allocated if it was not really allocated.
 *
 * @deprecated Use av_packet_ref or av_packet_make_refcounted
 */
//attribute_deprecated
//int av_dup_packet(AVPacket *pkt);
var avDupPacket func(pkt *AVPacket) ffcommon.FInt
var avDupPacketOnce sync.Once

func (pkt *AVPacket) AvDupPacket() ffcommon.FInt {
	avDupPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avDupPacket, ffcommon.GetAvcodecDll(), "av_dup_packet")
	})
	return avDupPacket(pkt)
}

/**
 * Copy packet, including contents
 *
 * @return 0 on success, negative AVERROR on fail
 *
 * @deprecated Use av_packet_ref
 */
//attribute_deprecated
//int av_copy_packet(AVPacket *dst, const AVPacket *src);
var avCopyPacket func(dst, src *AVPacket) ffcommon.FInt
var avCopyPacketOnce sync.Once

func (dst *AVPacket) AvCopyPacket(src *AVPacket) ffcommon.FInt {
	avCopyPacketOnce.Do(func() {
		purego.RegisterLibFunc(&avCopyPacket, ffcommon.GetAvcodecDll(), "av_copy_packet")
	})
	return avCopyPacket(dst, src)
}

/**
 * Copy packet side data
 *
 * @return 0 on success, negative AVERROR on fail
 *
 * @deprecated Use av_packet_copy_props
 */
//attribute_deprecated
//int av_copy_packet_side_data(AVPacket *dst, const AVPacket *src);
var avCopyPacketSideData func(dst, src *AVPacket) ffcommon.FInt
var avCopyPacketSideDataOnce sync.Once

func (dst *AVPacket) AvCopyPacketSideData(src *AVPacket) ffcommon.FInt {
	avCopyPacketSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avCopyPacketSideData, ffcommon.GetAvcodecDll(), "av_copy_packet_side_data")
	})
	return avCopyPacketSideData(dst, src)
}

/**
 * Free a packet.
 *
 * @deprecated Use av_packet_unref
 *
 * @param pkt packet to free
 */
//attribute_deprecated
//void av_free_packet(AVPacket *pkt);
var avFreePacket func(pkt *AVPacket)
var avFreePacketOnce sync.Once

func (pkt *AVPacket) AvFreePacket() {
	avFreePacketOnce.Do(func() {
		purego.RegisterLibFunc(&avFreePacket, ffcommon.GetAvcodecDll(), "av_free_packet")
	})
	avFreePacket(pkt)
}

//#endif
/**
 * Allocate new information of a packet.
 *
 * @param pkt packet
 * @param type side information type
 * @param size side information size
 * @return pointer to fresh allocated data or NULL otherwise
 */
//uint8_t* av_packet_new_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
var avPacketNewSideData func(pkt *AVPacket, type0 AVPacketSideDataType, size ffcommon.FIntOrSizeT) ffcommon.FInt
var avPacketNewSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketNewSideData(type0 AVPacketSideDataType, size ffcommon.FIntOrSizeT) ffcommon.FInt {
	avPacketNewSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketNewSideData, ffcommon.GetAvcodecDll(), "av_packet_new_side_data")
	})
	return avPacketNewSideData(pkt, type0, size)
}

/**
 * Wrap an existing array as a packet side data.
 *
 * @param pkt packet
 * @param type side information type
 * @param data the side data array. It must be allocated with the av_malloc()
 *             family of functions. The ownership of the data is transferred to
 *             pkt.
 * @param size side information size
 * @return a non-negative number on success, a negative AVERROR code on
 *         failure. On failure, the packet is unchanged and the data remains
 *         owned by the caller.
 */
//int av_packet_add_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//uint8_t *data, size_t size);
var avPacketAddSideData func(pkt *AVPacket, type0 AVPacketSideDataType, data *ffcommon.FUint8T, size ffcommon.FSizeT) ffcommon.FInt
var avPacketAddSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketAddSideData(type0 AVPacketSideDataType, data *ffcommon.FUint8T, size ffcommon.FSizeT) ffcommon.FInt {
	avPacketAddSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketAddSideData, ffcommon.GetAvcodecDll(), "av_packet_add_side_data")
	})
	return avPacketAddSideData(pkt, type0, data, size)
}

/**
 * Shrink the already allocated side data buffer
 *
 * @param pkt packet
 * @param type side information type
 * @param size new side information size
 * @return 0 on success, < 0 on failure
 */
//int av_packet_shrink_side_data(AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int size);
//#else
//size_t size);
//#endif
var avPacketShrinkSideData func(pkt *AVPacket, type0 AVPacketSideDataType, data *ffcommon.FUint8T, size ffcommon.FIntOrSizeT) ffcommon.FInt
var avPacketShrinkSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketShrinkSideData(type0 AVPacketSideDataType, data *ffcommon.FUint8T, size ffcommon.FIntOrSizeT) ffcommon.FInt {
	avPacketShrinkSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketShrinkSideData, ffcommon.GetAvcodecDll(), "av_packet_shrink_side_data")
	})
	return avPacketShrinkSideData(pkt, type0, data, size)
}

/**
 * Get side information from packet.
 *
 * @param pkt packet
 * @param type desired side information type
 * @param size If supplied, *size will be set to the size of the side data
 *             or to zero if the desired side data is not present.
 * @return pointer to data if present or NULL otherwise
 */
//uint8_t* av_packet_get_side_data(const AVPacket *pkt, enum AVPacketSideDataType type,
//#if FF_API_BUFFER_SIZE_T
//int *size);
//#else
//size_t *size);
var avPacketGetSideData func(pkt *AVPacket, type0 AVPacketSideDataType, data *ffcommon.FUint8T, size *ffcommon.FIntOrSizeT) ffcommon.FInt
var avPacketGetSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketGetSideData(type0 AVPacketSideDataType, data *ffcommon.FUint8T, size *ffcommon.FIntOrSizeT) ffcommon.FInt {
	avPacketGetSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketGetSideData, ffcommon.GetAvcodecDll(), "av_packet_get_side_data")
	})
	return avPacketGetSideData(pkt, type0, data, size)
}

//#endif

// #if FF_API_MERGE_SD_API
// attribute_deprecated
// int av_packet_merge_side_data(AVPacket *pkt);
var avPacketMergeSideData func(pkt *AVPacket) ffcommon.FInt
var avPacketMergeSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketMergeSideData() ffcommon.FInt {
	avPacketMergeSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketMergeSideData, ffcommon.GetAvcodecDll(), "av_packet_merge_side_data")
	})
	return avPacketMergeSideData(pkt)
}

// attribute_deprecated
// int av_packet_split_side_data(AVPacket *pkt);
var avPacketSplitSideData func(pkt *AVPacket) ffcommon.FInt
var avPacketSplitSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketSplitSideData() ffcommon.FInt {
	avPacketSplitSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketSplitSideData, ffcommon.GetAvcodecDll(), "av_packet_split_side_data")
	})
	return avPacketSplitSideData(pkt)
}

//#endif

// const char *av_packet_side_data_name(enum AVPacketSideDataType type);
var avPacketSideDataName func(type0 AVPacketSideDataType) string
var avPacketSideDataNameOnce sync.Once

func AvPacketSideDataName(type0 AVPacketSideDataType) string {
	avPacketSideDataNameOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketSideDataName, ffcommon.GetAvcodecDll(), "av_packet_side_data_name")
	})
	return avPacketSideDataName(type0)
}

/**
 * Pack a dictionary for use in side_data.
 *
 * @param dict The dictionary to pack.
 * @param size pointer to store the size of the returned data
 * @return pointer to data if successful, NULL otherwise
 */
//#if FF_API_BUFFER_SIZE_T
//uint8_t *av_packet_pack_dictionary(AVDictionary *dict, int *size);
//#else
//uint8_t *av_packet_pack_dictionary(AVDictionary *dict, size_t *size);
//#endif
var avPacketPackDictionary func(dict *AVDictionary, size *ffcommon.FIntOrSizeT) *ffcommon.FUint8T
var avPacketPackDictionaryOnce sync.Once

func AvPacketPackDictionary(dict *AVDictionary, size *ffcommon.FIntOrSizeT) *ffcommon.FUint8T {
	avPacketPackDictionaryOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketPackDictionary, ffcommon.GetAvcodecDll(), "av_packet_pack_dictionary")
	})
	return avPacketPackDictionary(dict, size)
}

/**
 * Unpack a dictionary from side_data.
 *
 * @param data data from side_data
 * @param size size of the data
 * @param dict the metadata storage dictionary
 * @return 0 on success, < 0 on failure
 */
//#if FF_API_BUFFER_SIZE_T
//int av_packet_unpack_dictionary(const uint8_t *data, int size, AVDictionary **dict);
//#else
//int av_packet_unpack_dictionary(const uint8_t *data, size_t size,
//AVDictionary **dict);
//#endif
var avPacketUnpackDictionary func(data *ffcommon.FUint8T, size ffcommon.FIntOrSizeT, dict **AVDictionary) ffcommon.FInt
var avPacketUnpackDictionaryOnce sync.Once

func AvPacketUnpackDictionary(data *ffcommon.FUint8T, size ffcommon.FIntOrSizeT, dict **AVDictionary) ffcommon.FInt {
	avPacketUnpackDictionaryOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketUnpackDictionary, ffcommon.GetAvcodecDll(), "av_packet_unpack_dictionary")
	})
	return avPacketUnpackDictionary(data, size, dict)
}

/**
 * Convenience function to free all the side data stored.
 * All the other fields stay untouched.
 *
 * @param pkt packet
 */
//void av_packet_free_side_data(AVPacket *pkt);
var avPacketFreeSideData func(pkt *AVPacket)
var avPacketFreeSideDataOnce sync.Once

func (pkt *AVPacket) AvPacketFreeSideData() {
	avPacketFreeSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketFreeSideData, ffcommon.GetAvcodecDll(), "av_packet_free_side_data")
	})
	avPacketFreeSideData(pkt)
}

/**
 * Setup a new reference to the data described by a given packet
 *
 * If src is reference-counted, setup dst as a new reference to the
 * buffer in src. Otherwise allocate a new buffer in dst and copy the
 * data from src into it.
 *
 * All the other fields are copied from src.
 *
 * @see av_packet_unref
 *
 * @param dst Destination packet. Will be completely overwritten.
 * @param src Source packet
 *
 * @return 0 on success, a negative AVERROR on error. On error, dst
 *         will be blank (as if returned by av_packet_alloc()).
 */
//int av_packet_ref(AVPacket *dst, const AVPacket *src);
var avPacketRef func(dst, src *AVPacket) ffcommon.FInt
var avPacketRefOnce sync.Once

func AvPacketRef(dst, src *AVPacket) ffcommon.FInt {
	avPacketRefOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketRef, ffcommon.GetAvcodecDll(), "av_packet_ref")
	})
	return avPacketRef(dst, src)
}

/**
 * Wipe the packet.
 *
 * Unreference the buffer referenced by the packet and reset the
 * remaining packet fields to their default values.
 *
 * @param pkt The packet to be unreferenced.
 */
//void av_packet_unref(AVPacket *pkt);
var avPacketUnref func(pkt *AVPacket)
var avPacketUnrefOnce sync.Once

func (pkt *AVPacket) AvPacketUnref() {
	avPacketUnrefOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketUnref, ffcommon.GetAvcodecDll(), "av_packet_unref")
	})
	avPacketUnref(pkt)
}

/**
 * Move every field in src to dst and reset src.
 *
 * @see av_packet_unref
 *
 * @param src Source packet, will be reset
 * @param dst Destination packet
 */
//void av_packet_move_ref(AVPacket *dst, AVPacket *src);
var avPacketMoveRef func(dst, src *AVPacket)
var avPacketMoveRefOnce sync.Once

func AvPacketMoveRef(dst, src *AVPacket) {
	avPacketMoveRefOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketMoveRef, ffcommon.GetAvcodecDll(), "av_packet_move_ref")
	})
	avPacketMoveRef(dst, src)
}

/**
 * Copy only "properties" fields from src to dst.
 *
 * Properties for the purpose of this function are all the fields
 * beside those related to the packet data (buf, data, size)
 *
 * @param dst Destination packet
 * @param src Source packet
 *
 * @return 0 on success AVERROR on failure.
 */
//int av_packet_copy_props(AVPacket *dst, const AVPacket *src);
var avPacketCopyProps func(dst, src *AVPacket) ffcommon.FInt
var avPacketCopyPropsOnce sync.Once

func AvPacketCopyProps(dst, src *AVPacket) ffcommon.FInt {
	avPacketCopyPropsOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketCopyProps, ffcommon.GetAvcodecDll(), "av_packet_copy_props")
	})
	return avPacketCopyProps(dst, src)
}

/**
 * Ensure the data described by a given packet is reference counted.
 *
 * @note This function does not ensure that the reference will be writable.
 *       Use av_packet_make_writable instead for that purpose.
 *
 * @see av_packet_ref
 * @see av_packet_make_writable
 *
 * @param pkt packet whose data should be made reference counted.
 *
 * @return 0 on success, a negative AVERROR on error. On failure, the
 *         packet is unchanged.
 */
//int av_packet_make_refcounted(AVPacket *pkt);
var avPacketMakeRefcounted func(pkt *AVPacket) ffcommon.FInt
var avPacketMakeRefcountedOnce sync.Once

func (pkt *AVPacket) AvPacketMakeRefcounted() ffcommon.FInt {
	avPacketMakeRefcountedOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketMakeRefcounted, ffcommon.GetAvcodecDll(), "av_packet_make_refcounted")
	})
	return avPacketMakeRefcounted(pkt)
}

/**
 * Create a writable reference for the data described by a given packet,
 * avoiding data copy if possible.
 *
 * @param pkt Packet whose data should be made writable.
 *
 * @return 0 on success, a negative AVERROR on failure. On failure, the
 *         packet is unchanged.
 */
//int av_packet_make_writable(AVPacket *pkt);
var avPacketMakeWritable func(pkt *AVPacket) ffcommon.FInt
var avPacketMakeWritableOnce sync.Once

func (pkt *AVPacket) AvPacketMakeWritable() ffcommon.FInt {
	avPacketMakeWritableOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketMakeWritable, ffcommon.GetAvcodecDll(), "av_packet_make_writable")
	})
	return avPacketMakeWritable(pkt)
}

/**
 * Convert valid timing fields (timestamps / durations) in a packet from one
 * timebase to another. Timestamps with unknown values (AV_NOPTS_VALUE) will be
 * ignored.
 *
 * @param pkt packet on which the conversion will be performed
 * @param tb_src source timebase, in which the timing fields in pkt are
 *               expressed
 * @param tb_dst destination timebase, to which the timing fields will be
 *               converted
 */
//void av_packet_rescale_ts(AVPacket *pkt, AVRational tb_src, AVRational tb_dst);
var avPacketRescaleTs func(pkt uintptr, tb_src, tb_dst uintptr)
var avPacketRescaleTsOnce sync.Once

func (pkt *AVPacket) AvPacketRescaleTs(tb_src, tb_dst AVRational) {
	avPacketRescaleTsOnce.Do(func() {
		purego.RegisterLibFunc(&avPacketRescaleTs, ffcommon.GetAvcodecDll(), "av_packet_rescale_ts")
	})
	avPacketRescaleTs(
		uintptr(unsafe.Pointer(pkt)),
		uintptr(unsafe.Pointer(&tb_src)),
		uintptr(unsafe.Pointer(&tb_dst)),
	)
}

/**
 * @}
 */

//#endif // AVCODEC_PACKET_H
