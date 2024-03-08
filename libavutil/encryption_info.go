package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/**
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

//#ifndef AVUTIL_ENCRYPTION_INFO_H
//#define AVUTIL_ENCRYPTION_INFO_H
//
//#include <stddef.h>
//#include <stdint.h>

type AVSubsampleEncryptionInfo struct {

	/** The number of bytes that are clear. */
	BytesOfClearData ffcommon.FUnsignedInt

	/**
	 * The number of bytes that are protected.  If using pattern encryption,
	 * the pattern applies to only the protected bytes; if not using pattern
	 * encryption, all these bytes are encrypted.
	 */
	BytesOfProtectedData ffcommon.FUnsignedInt
}

/**
 * This describes encryption info for a packet.  This contains frame-specific
 * info for how to decrypt the packet before passing it to the decoder.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInfo struct {

	/** The fourcc encryption scheme, in big-endian byte order. */
	Scheme ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are encrypted.
	 */
	CryptByteBlock ffcommon.FUint32T

	/**
	 * Only used for pattern encryption.  This is the number of 16-byte blocks
	 * that are clear.
	 */
	SkipByteBlock ffcommon.FUint32T

	/**
	 * The ID of the key used to encrypt the packet.  This should always be
	 * 16 bytes long, but may be changed in the future.
	 */
	KeyId     *ffcommon.FUint8T
	KeyIdSize ffcommon.FUint32T

	/**
	 * The initialization vector.  This may have been zero-filled to be the
	 * correct block size.  This should always be 16 bytes long, but may be
	 * changed in the future.
	 */
	Iv     *ffcommon.FUint8T
	IvSize ffcommon.FUint32T

	/**
	 * An array of subsample encryption info specifying how parts of the sample
	 * are encrypted.  If there are no subsamples, then the whole sample is
	 * encrypted.
	 */
	Subsamples     *AVSubsampleEncryptionInfo
	SubsampleCount ffcommon.FUint32T
}

/**
 * This describes info used to initialize an encryption key system.
 *
 * The size of this struct is not part of the public ABI.
 */
type AVEncryptionInitInfo struct {

	/**
	 * A unique identifier for the key system this is for, can be NULL if it
	 * is not known.  This should always be 16 bytes, but may change in the
	 * future.
	 */
	SystemId     *ffcommon.FUint8T
	SystemIdSize ffcommon.FUint32T

	/**
	 * An array of key IDs this initialization data is for.  All IDs are the
	 * same length.  Can be NULL if there are no known key IDs.
	 */
	KeyIds **ffcommon.FUint8T
	/** The number of key IDs. */
	NumKeyIds ffcommon.FUint32T
	/**
	 * The number of bytes in each key ID.  This should always be 16, but may
	 * change in the future.
	 */
	KeyIdSize ffcommon.FUint32T

	/**
	 * Key-system specific initialization data.  This data is copied directly
	 * from the file and the format depends on the specific key system.  This
	 * can be NULL if there is no initialization data; in that case, there
	 * will be at least one key ID.
	 */
	Data     *ffcommon.FUint8T
	DataSize ffcommon.FUint32T
	/**
	 * An optional pointer to the next initialization info in the list.
	 */
	Next *AVEncryptionInitInfo
}

/**
 * Allocates an AVEncryptionInfo structure and sub-pointers to hold the given
 * number of subsamples.  This will allocate pointers for the key ID, IV,
 * and subsample entries, set the size members, and zero-initialize the rest.
 *
 * @param subsample_count The number of subsamples.
 * @param key_id_size The number of bytes in the key ID, should be 16.
 * @param iv_size The number of bytes in the IV, should be 16.
 *
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_alloc(uint32_t subsample_count, uint32_t key_id_size, uint32_t iv_size);
var avEncryptionInfoAlloc func(subsampleCount, keyIDSize, ivSize ffcommon.FUint32T) *AVEncryptionInfo

var avEncryptionInfoAllocOnce sync.Once

func AvEncryptionInfoAlloc(subsampleCount, keyIDSize, ivSize ffcommon.FUint32T) *AVEncryptionInfo {
	avEncryptionInfoAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInfoAlloc, ffcommon.GetAvutilDll(), "av_encryption_info_alloc")
	})
	return avEncryptionInfoAlloc(subsampleCount, keyIDSize, ivSize)
}

/**
 * Allocates an AVEncryptionInfo structure with a copy of the given data.
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_clone(const AVEncryptionInfo *info);
var avEncryptionInfoClone func(info *AVEncryptionInfo) *AVEncryptionInfo
var avEncryptionInfoCloneOnce sync.Once

func (info *AVEncryptionInfo) AvEncryptionInfoClone() *AVEncryptionInfo {
	avEncryptionInfoCloneOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInfoClone, ffcommon.GetAvutilDll(), "av_encryption_info_clone")
	})
	return avEncryptionInfoClone(info)
}

/**
 * Frees the given encryption info object.  This MUST NOT be used to free the
 * side-data data pointer, that should use normal side-data methods.
 */
//void av_encryption_info_free(AVEncryptionInfo *info);
var avEncryptionInfoFree func(info *AVEncryptionInfo)
var avEncryptionInfoFreeOnce sync.Once

func (info *AVEncryptionInfo) AvEncryptionInfoFree() {
	avEncryptionInfoFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInfoFree, ffcommon.GetAvutilDll(), "av_encryption_info_free")
	})
	avEncryptionInfoFree(info)
}

/**
 * Creates a copy of the AVEncryptionInfo that is contained in the given side
 * data.  The resulting object should be passed to av_encryption_info_free()
 * when done.
 *
 * @return The new AVEncryptionInfo structure, or NULL on error.
 */
//AVEncryptionInfo *av_encryption_info_get_side_data(const uint8_t *side_data, size_t side_data_size);
// purego func
var avEncryptionInfoGetSideData func(sideData ffcommon.FConstCharP, sideDataSize ffcommon.FSizeT) *AVEncryptionInfo
var avEncryptionInfoGetSideDataOnce sync.Once

func AvEncryptionInfoGetSideData(sideData ffcommon.FConstCharP, sideDataSize ffcommon.FSizeT) (res *AVEncryptionInfo) {
	avEncryptionInfoGetSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInfoGetSideData, ffcommon.GetAvutilDll(), "av_encryption_info_get_side_data")
	})
	return avEncryptionInfoGetSideData(sideData, sideDataSize)
}

/**
 * Allocates and initializes side data that holds a copy of the given encryption
 * info.  The resulting pointer should be either freed using av_free or given
 * to av_packet_add_side_data().
 *
 * @return The new side-data pointer, or NULL.
 */
//uint8_t *av_encryption_info_add_side_data(
//const AVEncryptionInfo *info, size_t *side_data_size);
// purego struct method
var avEncryptionInfoAddSideData func(info *AVEncryptionInfo, sideDataSize ffcommon.FSizeT) *ffcommon.FUint8T
var avEncryptionInfoAddSideDataOnce sync.Once

func (info *AVEncryptionInfo) AvEncryptionInfo_AddSideData(sideDataSize ffcommon.FSizeT) (res *ffcommon.FUint8T) {
	avEncryptionInfoAddSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInfoAddSideData, ffcommon.GetAvutilDll(), "av_encryption_info_add_side_data")
	})
	return avEncryptionInfoAddSideData(info, sideDataSize)
}

/**
 * Allocates an AVEncryptionInitInfo structure and sub-pointers to hold the
 * given sizes.  This will allocate pointers and set all the fields.
 *
 * @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
//AVEncryptionInitInfo *av_encryption_init_info_alloc(
//uint32_t system_id_size, uint32_t num_key_ids, uint32_t key_id_size, uint32_t data_size);
// purego func
var avEncryptionInitInfoAlloc func(systemIDSize, numKeyIDs, keyIDSize, dataSize ffcommon.FUint32T) *AVEncryptionInitInfo
var avEncryptionInitInfoAllocOnce sync.Once

func AvEncryptionInitInfoAlloc(systemIDSize, numKeyIDs, keyIDSize, dataSize ffcommon.FUint32T) (res *AVEncryptionInitInfo) {
	avEncryptionInitInfoAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInitInfoAlloc, ffcommon.GetAvutilDll(), "av_encryption_init_info_alloc")
	})
	return avEncryptionInitInfoAlloc(systemIDSize, numKeyIDs, keyIDSize, dataSize)
}

/**
 * Frees the given encryption init info object.  This MUST NOT be used to free
 * the side-data data pointer, that should use normal side-data methods.
 */
//void av_encryption_init_info_free(AVEncryptionInitInfo* info);
// purego struct method
var avEncryptionInitInfoFree func(info *AVEncryptionInfo)
var avEncryptionInitInfoFreeOnce sync.Once

func (info *AVEncryptionInfo) AvEncryptionInitInfoFree() {
	avEncryptionInitInfoFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInitInfoFree, ffcommon.GetAvutilDll(), "av_encryption_init_info_free")
	})
	avEncryptionInitInfoFree(info)
}

/**
 * Creates a copy of the AVEncryptionInitInfo that is contained in the given
 * side data.  The resulting object should be passed to
 * av_encryption_init_info_free() when done.
 *
 * @return The new AVEncryptionInitInfo structure, or NULL on error.
 */
//AVEncryptionInitInfo *av_encryption_init_info_get_side_data(
//const uint8_t* side_data, size_t side_data_size);
// purego func
var avEncryptionInitInfoGetSideData func(sideData *ffcommon.FUint8T, sideDataSize ffcommon.FSizeT) *AVEncryptionInitInfo
var avEncryptionInitInfoGetSideDataOnce sync.Once

func AvEncryptionInitInfoGetSideData(sideData *ffcommon.FUint8T, sideDataSize ffcommon.FSizeT) (res *AVEncryptionInitInfo) {
	avEncryptionInitInfoGetSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInitInfoGetSideData, ffcommon.GetAvutilDll(), "av_encryption_init_info_get_side_data")
	})
	return avEncryptionInitInfoGetSideData(sideData, sideDataSize)
}

/**
 * Allocates and initializes side data that holds a copy of the given encryption
 * init info.  The resulting pointer should be either freed using av_free or
 * given to av_packet_add_side_data().
 *
 * @return The new side-data pointer, or NULL.
 */
//uint8_t *av_encryption_init_info_add_side_data(
//const AVEncryptionInitInfo *info, size_t *side_data_size);
// purego struct method
var avEncryptionInitInfoAddSideData func(info *AVEncryptionInfo, sideDataSize *ffcommon.FSizeT) *ffcommon.FUint8T
var avEncryptionInitInfoAddSideDataOnce sync.Once

func (info *AVEncryptionInfo) AvEncryptionInitInfoAddSideData(sideDataSize *ffcommon.FSizeT) (res *ffcommon.FUint8T) {
	avEncryptionInitInfoAddSideDataOnce.Do(func() {
		purego.RegisterLibFunc(&avEncryptionInitInfoAddSideData, ffcommon.GetAvutilDll(), "av_encryption_init_info_add_side_data")
	})
	return avEncryptionInitInfoAddSideData(info, sideDataSize)
}

//#endif /* AVUTIL_ENCRYPTION_INFO_H */
