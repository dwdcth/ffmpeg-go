package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
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
 * Public dictionary API.
 * @deprecated
 *  AVDictionary is provided for compatibility with libav. It is both in
 *  implementation as well as API inefficient. It does not scale and is
 *  extremely slow with large dictionaries.
 *  It is recommended that new code uses our tree container from tree.c/h
 *  where applicable, which uses AVL trees to achieve O(log n) performance.
 */

//#ifndef AVUTIL_DICT_H
//const AVUTIL_DICT_H
//
//#include <stdint.h>
//
//#include "version.h"

/**
  * @addtogroup lavu_dict AVDictionary
  * @ingroup lavu_data
  *
  * @brief Simple key:value store
  *
  * @{
  * Dictionaries are used for storing key:value pairs. To create
  * an AVDictionary, simply pass an address of a NULL pointer to
  * av_dict_set(). NULL can be used as an empty dictionary wherever
  * a pointer to an AVDictionary is required.
  * Use av_dict_get() to retrieve an entry or iterate over all
  * entries and finally av_dict_free() to free the dictionary
  * and all its contents.
  *
  @code
    AVDictionary *d = NULL;           // "create" an empty dictionary
    AVDictionaryEntry *t = NULL;

    av_dict_set(&d, "foo", "bar", 0); // add an entry

    char *k = av_strdup("key");       // if your strings are already allocated,
    char *v = av_strdup("value");     // you can avoid copying them like this
    av_dict_set(&d, k, v, AV_DICT_DONT_STRDUP_KEY | AV_DICT_DONT_STRDUP_VAL);

    while (t = av_dict_get(d, "", t, AV_DICT_IGNORE_SUFFIX)) {
        <....>                             // iterate over all entries in d
    }
    av_dict_free(&d);
  @endcode
*/

const AV_DICT_MATCH_CASE = 1    /**< Only get an entry with exact-case key match. Only relevant in av_dict_get(). */
const AV_DICT_IGNORE_SUFFIX = 2 /**< Return first entry in a dictionary whose first part corresponds to the search key,
  ignoring the suffix of the found key string. Only relevant in av_dict_get(). */
const AV_DICT_DONT_STRDUP_KEY = 4 /**< Take ownership of a key that's been
  allocated with av_malloc() or another memory allocation function. */
const AV_DICT_DONT_STRDUP_VAL = 8 /**< Take ownership of a value that's been
  allocated with av_malloc() or another memory allocation function. */
const AV_DICT_DONT_OVERWRITE = 16 ///< Don't overwrite existing entries.
const AV_DICT_APPEND = 32         /**< If the entry already exists, append to it.  Note that no
  delimiter is added, the strings are simply concatenated. */
const AV_DICT_MULTIKEY = 64 /**< Allow to store several equal keys in the dictionary */

type AVDictionaryEntry struct {
	Key   ffcommon.FCharPStruct
	Value ffcommon.FCharPStruct
}

// typedef struct AVDictionary AVDictionary;
type AVDictionary struct {
}

/**
 * Get a dictionary entry with matching key.
 *
 * The returned entry key or value must not be changed, or it will
 * cause undefined behavior.
 *
 * To iterate through all the dictionary entries, you can set the matching key
 * to the null string "" and set the AV_DICT_IGNORE_SUFFIX flag.
 *
 * @param prev Set to the previous matching element to find the next.
 *             If set to NULL the first matching element is returned.
 * @param key matching key
 * @param flags a collection of AV_DICT_* flags controlling how the entry is retrieved
 * @return found entry or NULL in case no matching entry was found in the dictionary
 */
//AVDictionaryEntry *av_dict_get(const AVDictionary *m, const char *key,
//const AVDictionaryEntry *prev, int flags);
var avDictGet func(m *AVDictionary, key ffcommon.FConstCharP, prev *AVDictionaryEntry, flags ffcommon.FInt) *AVDictionaryEntry
var avDictGetOnce sync.Once

func (m *AVDictionary) AvDictGet(key ffcommon.FConstCharP, prev *AVDictionaryEntry, flags ffcommon.FInt) (res *AVDictionaryEntry) {
	avDictGetOnce.Do(func() {
		purego.RegisterLibFunc(&avDictGet, ffcommon.GetAvutilDll(), "av_dict_get")
	})
	return avDictGet(m, key, prev, flags)
}

/**
 * Get number of entries in dictionary.
 *
 * @param m dictionary
 * @return  number of entries in dictionary
 */
//int av_dict_count(const AVDictionary *m);
var avDictCount func(m *AVDictionary) ffcommon.FInt
var avDictCountOnce sync.Once

func (m *AVDictionary) AvDictCount() (res ffcommon.FInt) {
	avDictCountOnce.Do(func() {
		purego.RegisterLibFunc(&avDictCount, ffcommon.GetAvutilDll(), "av_dict_count")
	})
	return avDictCount(m)
}

/**
 * Set the given entry in *pm, overwriting an existing entry.
 *
 * Note: If AV_DICT_DONT_STRDUP_KEY or AV_DICT_DONT_STRDUP_VAL is set,
 * these arguments will be freed on error.
 *
 * Warning: Adding a new entry to a dictionary invalidates all existing entries
 * previously returned with av_dict_get.
 *
 * @param pm pointer to a pointer to a dictionary struct. If *pm is NULL
 * a dictionary struct is allocated and put in *pm.
 * @param key entry key to add to *pm (will either be av_strduped or added as a new key depending on flags)
 * @param value entry value to add to *pm (will be av_strduped or added as a new key depending on flags).
 *        Passing a NULL value will cause an existing entry to be deleted.
 * @return >= 0 on success otherwise an error code <0
 */
//int av_dict_set(AVDictionary **pm, const char *key, const char *value, int flags);
var avDictSet func(pm **AVDictionary, key, value ffcommon.FConstCharP, flags ffcommon.FInt) ffcommon.FInt
var avDictSetOnce sync.Once

func AvDictSet(pm **AVDictionary, key, value ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	avDictSetOnce.Do(func() {
		purego.RegisterLibFunc(&avDictSet, ffcommon.GetAvutilDll(), "av_dict_set")
	})
	return avDictSet(pm, key, value, flags)
}

/**
 * Convenience wrapper for av_dict_set that converts the value to a string
 * and stores it.
 *
 * Note: If AV_DICT_DONT_STRDUP_KEY is set, key will be freed on error.
 */
//int av_dict_set_int(AVDictionary **pm, const char *key, int64_t value, int flags);
var avDictSetInt func(pm **AVDictionary, key ffcommon.FConstCharP, value ffcommon.FInt64T, flags ffcommon.FInt) ffcommon.FInt
var avDictSetIntOnce sync.Once

func AvDictSetInt(pm **AVDictionary, key ffcommon.FConstCharP, value ffcommon.FInt64T, flags ffcommon.FInt) (res ffcommon.FInt) {
	avDictSetIntOnce.Do(func() {
		purego.RegisterLibFunc(&avDictSetInt, ffcommon.GetAvutilDll(), "av_dict_set_int")
	})
	return avDictSetInt(pm, key, value, flags)
}

/**
 * Parse the key/value pairs list and add the parsed entries to a dictionary.
 *
 * In case of failure, all the successfully set entries are stored in
 * *pm. You may need to manually free the created dictionary.
 *
 * @param key_val_sep  a 0-terminated list of characters used to separate
 *                     key from value
 * @param pairs_sep    a 0-terminated list of characters used to separate
 *                     two pairs from each other
 * @param flags        flags to use when adding to dictionary.
 *                     AV_DICT_DONT_STRDUP_KEY and AV_DICT_DONT_STRDUP_VAL
 *                     are ignored since the key/value tokens will always
 *                     be duplicated.
 * @return             0 on success, negative AVERROR code on failure
 */
//int av_dict_parse_string(AVDictionary **pm, const char *str,
//const char *key_val_sep, const char *pairs_sep,
//int flags);
var avDictParseString func(pm **AVDictionary, str, key_val_sep, pairs_sep ffcommon.FConstCharP, flags ffcommon.FInt) ffcommon.FInt
var avDictParseStringOnce sync.Once

func AvDictParseString(pm **AVDictionary, str, key_val_sep, pairs_sep ffcommon.FConstCharP, flags ffcommon.FInt) (res ffcommon.FInt) {
	avDictParseStringOnce.Do(func() {
		purego.RegisterLibFunc(&avDictParseString, ffcommon.GetAvutilDll(), "av_dict_parse_string")
	})
	return avDictParseString(pm, str, key_val_sep, pairs_sep, flags)
}

/**
 * Copy entries from one AVDictionary struct into another.
 * @param dst pointer to a pointer to a AVDictionary struct. If *dst is NULL,
 *            this function will allocate a struct for you and put it in *dst
 * @param src pointer to source AVDictionary struct
 * @param flags flags to use when setting entries in *dst
 * @note metadata is read using the AV_DICT_IGNORE_SUFFIX flag
 * @return 0 on success, negative AVERROR code on failure. If dst was allocated
 *           by this function, callers should free the associated memory.
 */
//int av_dict_copy(AVDictionary **dst, const AVDictionary *src, int flags);
var avDictCopy func(dst **AVDictionary, src *AVDictionary, flags ffcommon.FInt) ffcommon.FInt
var avDictCopyOnce sync.Once

func AvDictCopy(dst **AVDictionary, src *AVDictionary, flags ffcommon.FInt) (res ffcommon.FInt) {
	avDictCopyOnce.Do(func() {
		purego.RegisterLibFunc(&avDictCopy, ffcommon.GetAvutilDll(), "av_dict_copy")
	})
	return avDictCopy(dst, src, flags)
}

/**
 * Free all the memory allocated for an AVDictionary struct
 * and all keys and values.
 */
//void av_dict_free(AVDictionary **m);
var avDictFree func(m **AVDictionary)
var avDictFreeOnce sync.Once

func AvDictFree(m **AVDictionary) {
	avDictFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avDictFree, ffcommon.GetAvutilDll(), "av_dict_free")
	})
	avDictFree(m)
}

/**
 * Get dictionary entries as a string.
 *
 * Create a string containing dictionary's entries.
 * Such string may be passed back to av_dict_parse_string().
 * @note String is escaped with backslashes ('\').
 *
 * @param[in]  m             dictionary
 * @param[out] buffer        Pointer to buffer that will be allocated with string containg entries.
 *                           Buffer must be freed by the caller when is no longer needed.
 * @param[in]  key_val_sep   character used to separate key from value
 * @param[in]  pairs_sep     character used to separate two pairs from each other
 * @return                   >= 0 on success, negative on error
 * @warning Separators cannot be neither '\\' nor '\0'. They also cannot be the same.
 */
//int av_dict_get_string(const AVDictionary *m, char **buffer,
//const char key_val_sep, const char pairs_sep);
var avDictGetString func(m *AVDictionary, buffer *ffcommon.FBuf, key_val_sep, pairs_sep ffcommon.FConstCharP) ffcommon.FInt
var avDictGetStringOnce sync.Once

func (m *AVDictionary) AvDictGetString(buffer *ffcommon.FBuf, key_val_sep, pairs_sep ffcommon.FConstCharP) (res ffcommon.FInt) {
	avDictGetStringOnce.Do(func() {
		purego.RegisterLibFunc(&avDictGetString, ffcommon.GetAvutilDll(), "av_dict_get_string")
	})
	return avDictGetString(m, buffer, key_val_sep, pairs_sep)
}

/**
 * @}
 */

//#endif /* AVUTIL_DICT_H */
