package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * Copyright (c) 2007 Mans Rullgard
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

//#ifndef AVUTIL_AVSTRING_H
//#define AVUTIL_AVSTRING_H
//
//#include <stddef.h>
//#include <stdint.h>
//#include "attributes.h"
//#include "version.h"

/**
 * @addtogroup lavu_string
 * @{
 */

/**
 * Return non-zero if pfx is a prefix of str. If it is, *ptr is set to
 * the address of the first character in str after the prefix.
 *
 * @param str input string
 * @param pfx prefix to test
 * @param ptr updated if the prefix is matched inside str
 * @return non-zero if the prefix matches, zero otherwise
 */
//int av_strstart(const char *str, const char *pfx, const char **ptr);
//todo
var avStrstart func() ffcommon.FCharP
var avStrstartOnce sync.Once

func AvStrstart() (res ffcommon.FCharP) {
	avStrstartOnce.Do(func() {
		purego.RegisterLibFunc(&avStrstart, ffcommon.GetAvutilDll(), "av_strstart")
	})
	res = avStrstart()
	return
}

/**
 * Return non-zero if pfx is a prefix of str independent of case. If
 * it is, *ptr is set to the address of the first character in str
 * after the prefix.
 *
 * @param str input string
 * @param pfx prefix to test
 * @param ptr updated if the prefix is matched inside str
 * @return non-zero if the prefix matches, zero otherwise
 */
//int av_stristart(const char *str, const char *pfx, const char **ptr);
//todo
var avStristart func() ffcommon.FCharP
var avStristartOnce sync.Once

func AvStristart() (res ffcommon.FCharP) {
	avStristartOnce.Do(func() {
		purego.RegisterLibFunc(&avStristart, ffcommon.GetAvutilDll(), "av_stristart")
	})
	res = avStristart()
	return
}

/**
 * Locate the first case-independent occurrence in the string haystack
 * of the string needle.  A zero-length string needle is considered to
 * match at the start of haystack.
 *
 * This function is a case-insensitive version of the standard strstr().
 *
 * @param haystack string to search in
 * @param needle   string to search for
 * @return         pointer to the located match within haystack
 *                 or a null pointer if no match
 */
//char *av_stristr(const char *haystack, const char *needle);
//todo
var avStristr func() ffcommon.FCharP
var avStristrOnce sync.Once

func AvStristr() (res ffcommon.FCharP) {
	avStristrOnce.Do(func() {
		purego.RegisterLibFunc(&avStristr, ffcommon.GetAvutilDll(), "av_stristr")
	})
	res = avStristr()
	return
}

/**
 * Locate the first occurrence of the string needle in the string haystack
 * where not more than hay_length characters are searched. A zero-length
 * string needle is considered to match at the start of haystack.
 *
 * This function is a length-limited version of the standard strstr().
 *
 * @param haystack   string to search in
 * @param needle     string to search for
 * @param hay_length length of string to search in
 * @return           pointer to the located match within haystack
 *                   or a null pointer if no match
 */
//char *av_strnstr(const char *haystack, const char *needle, size_t hay_length);
//todo
var avStrnstr func() ffcommon.FCharP
var avStrnstrOnce sync.Once

func AvStrnstr() (res ffcommon.FCharP) {
	avStrnstrOnce.Do(func() {
		purego.RegisterLibFunc(&avStrnstr, ffcommon.GetAvutilDll(), "av_strnstr")
	})
	res = avStrnstr()
	return
}

/**
 * Copy the string src to dst, but no more than size - 1 bytes, and
 * null-terminate dst.
 *
 * This function is the same as BSD strlcpy().
 *
 * @param dst destination buffer
 * @param src source string
 * @param size size of destination buffer
 * @return the length of src
 *
 * @warning since the return value is the length of src, src absolutely
 * _must_ be a properly 0-terminated string, otherwise this will read beyond
 * the end of the buffer and possibly crash.
 */
//size_t av_strlcpy(char *dst, const char *src, size_t size);
//todo
var avStrlcpy func() ffcommon.FCharP
var avStrlcpyOnce sync.Once

func AvStrlcpy() (res ffcommon.FCharP) {
	avStrlcpyOnce.Do(func() {
		purego.RegisterLibFunc(&avStrlcpy, ffcommon.GetAvutilDll(), "av_strlcpy")
	})
	res = avStrlcpy()
	return
}

/**
 * Append the string src to the string dst, but to a total length of
 * no more than size - 1 bytes, and null-terminate dst.
 *
 * This function is similar to BSD strlcat(), but differs when
 * size <= strlen(dst).
 *
 * @param dst destination buffer
 * @param src source string
 * @param size size of destination buffer
 * @return the total length of src and dst
 *
 * @warning since the return value use the length of src and dst, these
 * absolutely _must_ be a properly 0-terminated strings, otherwise this
 * will read beyond the end of the buffer and possibly crash.
 */
//size_t av_strlcat(char *dst, const char *src, size_t size);
//todo
var avStrlcat func() ffcommon.FCharP
var avStrlcatOnce sync.Once

func AvStrlcat() (res ffcommon.FCharP) {
	avStrlcatOnce.Do(func() {
		purego.RegisterLibFunc(&avStrlcat, ffcommon.GetAvutilDll(), "av_strlcat")
	})
	res = avStrlcat()
	return
}

/**
 * Append output to a string, according to a format. Never write out of
 * the destination buffer, and always put a terminating 0 within
 * the buffer.
 * @param dst destination buffer (string to which the output is
 *  appended)
 * @param size total size of the destination buffer
 * @param fmt printf-compatible format string, specifying how the
 *  following parameters are used
 * @return the length of the string that would have been generated
 *  if enough space had been available
 */
//size_t av_strlcatf(char *dst, size_t size, const char *fmt, ...) av_printf_format(3, 4);
//todo
var avStrlcatf func() ffcommon.FCharP
var avStrlcatfOnce sync.Once

func AvStrlcatf() (res ffcommon.FCharP) {
	avStrlcatfOnce.Do(func() {
		purego.RegisterLibFunc(&avStrlcatf, ffcommon.GetAvutilDll(), "av_strlcatf")
	})
	res = avStrlcatf()
	return
}

/**
 * Get the count of continuous non zero chars starting from the beginning.
 *
 * @param len maximum number of characters to check in the string, that
 *            is the maximum value which is returned by the function
 */
//static inline size_t av_strnlen(const char *s, size_t len)
//{
//size_t i;
//for (i = 0; i < len && s[i]; i++)
//;
//return i;
//}
//todo
var avStrnlen func() ffcommon.FCharP
var avStrnlenOnce sync.Once

func AvStrnlen() (res ffcommon.FCharP) {
	avStrnlenOnce.Do(func() {
		purego.RegisterLibFunc(&avStrnlen, ffcommon.GetAvutilDll(), "av_strnlen")
	})
	res = avStrnlen()
	return
}

/**
 * Print arguments following specified format into a large enough auto
 * allocated buffer. It is similar to GNU asprintf().
 * @param fmt printf-compatible format string, specifying how the
 *            following parameters are used.
 * @return the allocated string
 * @note You have to free the string yourself with av_free().
 */
//char *av_asprintf(const char *fmt, ...) av_printf_format(1, 2);
//todo
var avAsprintf func() ffcommon.FCharP
var avAsprintfOnce sync.Once

func AvAsprintf() (res ffcommon.FCharP) {
	avAsprintfOnce.Do(func() {
		purego.RegisterLibFunc(&avAsprintf, ffcommon.GetAvutilDll(), "av_asprintf")
	})
	res = avAsprintf()
	return
}

//#if FF_API_D2STR
/**
 * Convert a number to an av_malloced string.
 * @deprecated  use av_asprintf() with "%f" or a more specific format
 */
//attribute_deprecated
//char *av_d2str(double d);
//#endif
//todo
var avD2str func() ffcommon.FCharP
var avD2strOnce sync.Once

func AvD2str() (res ffcommon.FCharP) {
	avD2strOnce.Do(func() {
		purego.RegisterLibFunc(&avD2str, ffcommon.GetAvutilDll(), "av_d2str")
	})
	res = avD2str()
	return
}

/**
 * Unescape the given string until a non escaped terminating char,
 * and return the token corresponding to the unescaped string.
 *
 * The normal \ and ' escaping is supported. Leading and trailing
 * whitespaces are removed, unless they are escaped with '\' or are
 * enclosed between ''.
 *
 * @param buf the buffer to parse, buf will be updated to point to the
 * terminating char
 * @param term a 0-terminated list of terminating chars
 * @return the malloced unescaped string, which must be av_freed by
 * the user, NULL in case of allocation failure
 */
//char *av_get_token(const char **buf, const char *term);
//todo
var avGetToken func() ffcommon.FCharP
var avGetTokenOnce sync.Once

func AvGetToken() (res ffcommon.FCharP) {
	avGetTokenOnce.Do(func() {
		purego.RegisterLibFunc(&avGetToken, ffcommon.GetAvutilDll(), "av_get_token")
	})
	res = avGetToken()
	return
}

/**
 * Split the string into several tokens which can be accessed by
 * successive calls to av_strtok().
 *
 * A token is defined as a sequence of characters not belonging to the
 * set specified in delim.
 *
 * On the first call to av_strtok(), s should point to the string to
 * parse, and the value of saveptr is ignored. In subsequent calls, s
 * should be NULL, and saveptr should be unchanged since the previous
 * call.
 *
 * This function is similar to strtok_r() defined in POSIX.1.
 *
 * @param s the string to parse, may be NULL
 * @param delim 0-terminated list of token delimiters, must be non-NULL
 * @param saveptr user-provided pointer which points to stored
 * information necessary for av_strtok() to continue scanning the same
 * string. saveptr is updated to point to the next character after the
 * first delimiter found, or to NULL if the string was terminated
 * @return the found token, or NULL when no token is found
 */
//char *av_strtok(char *s, const char *delim, char **saveptr);
//todo
var avStrtok func() ffcommon.FCharP
var avStrtokOnce sync.Once

func AvStrtok() (res ffcommon.FCharP) {
	avStrtokOnce.Do(func() {
		purego.RegisterLibFunc(&avStrtok, ffcommon.GetAvutilDll(), "av_strtok")
	})
	res = avStrtok()
	return
}

/**
 * Locale-independent conversion of ASCII isdigit.
 */
//static inline av_const int av_isdigit(int c)
//{
//return c >= '0' && c <= '9';
//}
//todo
var avIsdigit func() ffcommon.FCharP
var avIsdigitOnce sync.Once

func AvIsdigit() (res ffcommon.FCharP) {
	avIsdigitOnce.Do(func() {
		purego.RegisterLibFunc(&avIsdigit, ffcommon.GetAvutilDll(), "av_isdigit")
	})
	res = avIsdigit()
	return
}

/**
 * Locale-independent conversion of ASCII isgraph.
 */
//static inline av_const int av_isgraph(int c)
//{
//return c > 32 && c < 127;
//}
//todo
var avIsgraph func() ffcommon.FCharP
var avIsgraphOnce sync.Once

func AvIsgraph() (res ffcommon.FCharP) {
	avIsgraphOnce.Do(func() {
		purego.RegisterLibFunc(&avIsgraph, ffcommon.GetAvutilDll(), "av_isgraph")
	})
	res = avIsgraph()
	return
}

/**
 * Locale-independent conversion of ASCII isspace.
 */
//static inline av_const int av_isspace(int c)
//{
//return c == ' ' || c == '\f' || c == '\n' || c == '\r' || c == '\t' ||
//c == '\v';
//}
//todo
var avIsspace func() ffcommon.FCharP
var avIsspaceOnce sync.Once

func AvIsspace() (res ffcommon.FCharP) {
	avIsspaceOnce.Do(func() {
		purego.RegisterLibFunc(&avIsspace, ffcommon.GetAvutilDll(), "av_isspace")
	})
	res = avIsspace()
	return
}

/**
 * Locale-independent conversion of ASCII characters to uppercase.
 */
//static inline av_const int av_toupper(int c)
//{
//if (c >= 'a' && c <= 'z')
//c ^= 0x20;
//return c;
//}
//todo
var avToupper func() ffcommon.FCharP
var avToupperOnce sync.Once

func AvToupper() (res ffcommon.FCharP) {
	avToupperOnce.Do(func() {
		purego.RegisterLibFunc(&avToupper, ffcommon.GetAvutilDll(), "av_toupper")
	})
	res = avToupper()
	return
}

/**
 * Locale-independent conversion of ASCII characters to lowercase.
 */
//static inline av_const int av_tolower(int c)
//{
//if (c >= 'A' && c <= 'Z')
//c ^= 0x20;
//return c;
//}
//todo
var avTolower func() ffcommon.FCharP
var avTolowerOnce sync.Once

func AvTolower() (res ffcommon.FCharP) {
	avTolowerOnce.Do(func() {
		purego.RegisterLibFunc(&avTolower, ffcommon.GetAvutilDll(), "av_tolower")
	})
	res = avTolower()
	return
}

/**
 * Locale-independent conversion of ASCII isxdigit.
 */
//static inline av_const int av_isxdigit(int c)
//{
//c = av_tolower(c);
//return av_isdigit(c) || (c >= 'a' && c <= 'f');
//}
//todo
var avIsxdigit func() ffcommon.FCharP
var avIsxdigitOnce sync.Once

func AvIsxdigit() (res ffcommon.FCharP) {
	avIsxdigitOnce.Do(func() {
		purego.RegisterLibFunc(&avIsxdigit, ffcommon.GetAvutilDll(), "av_isxdigit")
	})
	res = avIsxdigit()
	return
}

/**
 * Locale-independent case-insensitive compare.
 * @note This means only ASCII-range characters are case-insensitive
 */
//int av_strcasecmp(const char *a, const char *b);
//todo
var avStrcasecmp func() ffcommon.FCharP
var avStrcasecmpOnce sync.Once

func AvStrcasecmp() (res ffcommon.FCharP) {
	avStrcasecmpOnce.Do(func() {
		purego.RegisterLibFunc(&avStrcasecmp, ffcommon.GetAvutilDll(), "av_strcasecmp")
	})
	res = avStrcasecmp()
	return
}

/**
 * Locale-independent case-insensitive compare.
 * @note This means only ASCII-range characters are case-insensitive
 */
//int av_strncasecmp(const char *a, const char *b, size_t n);
//todo
var avStrncasecmp func() ffcommon.FCharP
var avStrncasecmpOnce sync.Once

func AvStrncasecmp() (res ffcommon.FCharP) {
	avStrncasecmpOnce.Do(func() {
		purego.RegisterLibFunc(&avStrncasecmp, ffcommon.GetAvutilDll(), "av_strncasecmp")
	})
	res = avStrncasecmp()
	return
}

/**
 * Locale-independent strings replace.
 * @note This means only ASCII-range characters are replace
 */
//char *av_strireplace(const char *str, const char *from, const char *to);
//todo
var avStrireplace func() ffcommon.FCharP
var avStrireplaceOnce sync.Once

func AvStrireplace() (res ffcommon.FCharP) {
	avStrireplaceOnce.Do(func() {
		purego.RegisterLibFunc(&avStrireplace, ffcommon.GetAvutilDll(), "av_strireplace")
	})
	res = avStrireplace()
	return
}

/**
 * Thread safe basename.
 * @param path the string to parse, on DOS both \ and / are considered separators.
 * @return pointer to the basename substring.
 * If path does not contain a slash, the function returns a copy of path.
 * If path is a NULL pointer or points to an empty string, a pointer
 * to a string "." is returned.
 */
//const char *av_basename(const char *path);
//todo
var avBasename func() ffcommon.FCharP
var avBasenameOnce sync.Once

func AvBasename() (res ffcommon.FCharP) {
	avBasenameOnce.Do(func() {
		purego.RegisterLibFunc(&avBasename, ffcommon.GetAvutilDll(), "av_basename")
	})
	res = avBasename()
	return
}

/**
 * Thread safe dirname.
 * @param path the string to parse, on DOS both \ and / are considered separators.
 * @return A pointer to a string that's the parent directory of path.
 * If path is a NULL pointer or points to an empty string, a pointer
 * to a string "." is returned.
 * @note the function may modify the contents of the path, so copies should be passed.
 */
//const char *av_dirname(char *path);
//todo
var avDirname func() ffcommon.FCharP
var avDirnameOnce sync.Once

func AvDirname() (res ffcommon.FCharP) {
	avDirnameOnce.Do(func() {
		purego.RegisterLibFunc(&avDirname, ffcommon.GetAvutilDll(), "av_dirname")
	})
	res = avDirname()
	return
}

/**
 * Match instances of a name in a comma-separated list of names.
 * List entries are checked from the start to the end of the names list,
 * the first match ends further processing. If an entry prefixed with '-'
 * matches, then 0 is returned. The "ALL" list entry is considered to
 * match all names.
 *
 * @param name  Name to look for.
 * @param names List of names.
 * @return 1 on match, 0 otherwise.
 */
//int av_match_name(const char *name, const char *names);
//todo
var avMatchName func() ffcommon.FCharP
var avMatchNameOnce sync.Once

func AvMatchName() (res ffcommon.FCharP) {
	avMatchNameOnce.Do(func() {
		purego.RegisterLibFunc(&avMatchName, ffcommon.GetAvutilDll(), "av_match_name")
	})
	res = avMatchName()
	return
}

/**
 * Append path component to the existing path.
 * Path separator '/' is placed between when needed.
 * Resulting string have to be freed with av_free().
 * @param path      base path
 * @param component component to be appended
 * @return new path or NULL on error.
 */
//char *av_append_path_component(const char *path, const char *component);
//todo
var avAppendPathComponent func() ffcommon.FCharP
var avAppendPathComponentOnce sync.Once

func AvAppendPathComponent() (res ffcommon.FCharP) {
	avAppendPathComponentOnce.Do(func() {
		purego.RegisterLibFunc(&avAppendPathComponent, ffcommon.GetAvutilDll(), "av_append_path_component")
	})
	res = avAppendPathComponent()
	return
}

type AVEscapeMode = int32

const (
	AV_ESCAPE_MODE_AUTO      = iota ///< Use auto-selected escaping mode.
	AV_ESCAPE_MODE_BACKSLASH        ///< Use backslash escaping.
	AV_ESCAPE_MODE_QUOTE            ///< Use single-quote escaping.
	AV_ESCAPE_MODE_XML              ///< Use XML non-markup character data escaping.
)

/**
 * Consider spaces special and escape them even in the middle of the
 * string.
 *
 * This is equivalent to adding the whitespace characters to the special
 * characters lists, except it is guaranteed to use the exact same list
 * of whitespace characters as the rest of libavutil.
 */
const AV_ESCAPE_FLAG_WHITESPACE = (1 << 0)

/**
 * Escape only specified special characters.
 * Without this flag, escape also any characters that may be considered
 * special by av_get_token(), such as the single quote.
 */
const AV_ESCAPE_FLAG_STRICT = (1 << 1)

/**
 * Within AV_ESCAPE_MODE_XML, additionally escape single quotes for single
 * quoted attributes.
 */
const AV_ESCAPE_FLAG_XML_SINGLE_QUOTES = (1 << 2)

/**
 * Within AV_ESCAPE_MODE_XML, additionally escape double quotes for double
 * quoted attributes.
 */
const AV_ESCAPE_FLAG_XML_DOUBLE_QUOTES = (1 << 3)

/**
 * Escape string in src, and put the escaped string in an allocated
 * string in *dst, which must be freed with av_free().
 *
 * @param dst           pointer where an allocated string is put
 * @param src           string to escape, must be non-NULL
 * @param special_chars string containing the special characters which
 *                      need to be escaped, can be NULL
 * @param mode          escape mode to employ, see AV_ESCAPE_MODE_* macros.
 *                      Any unknown value for mode will be considered equivalent to
 *                      AV_ESCAPE_MODE_BACKSLASH, but this behaviour can change without
 *                      notice.
 * @param flags         flags which control how to escape, see AV_ESCAPE_FLAG_ macros
 * @return the length of the allocated string, or a negative error code in case of error
 * @see av_bprint_escape()
 */
//av_warn_unused_result
//int av_escape(char **dst, const char *src, const char *special_chars,
//enum AVEscapeMode mode, int flags);
//todo
var avEscape func() ffcommon.FCharP
var avEscapeOnce sync.Once

func AvEscape() (res ffcommon.FCharP) {
	avEscapeOnce.Do(func() {
		purego.RegisterLibFunc(&avEscape, ffcommon.GetAvutilDll(), "av_escape")
	})
	res = avEscape()
	return
}

const AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES = 1          ///< accept codepoints over 0x10FFFF
const AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS = 2             ///< accept non-characters - 0xFFFE and 0xFFFF
const AV_UTF8_FLAG_ACCEPT_SURROGATES = 4                 ///< accept UTF-16 surrogates codes
const AV_UTF8_FLAG_EXCLUDE_XML_INVALID_CONTROL_CODES = 8 ///< exclude control codes not accepted by XML

const AV_UTF8_FLAG_ACCEPT_ALL = AV_UTF8_FLAG_ACCEPT_INVALID_BIG_CODES | AV_UTF8_FLAG_ACCEPT_NON_CHARACTERS | AV_UTF8_FLAG_ACCEPT_SURROGATES

/**
 * Read and decode a single UTF-8 code point (character) from the
 * buffer in *buf, and update *buf to point to the next byte to
 * decode.
 *
 * In case of an invalid byte sequence, the pointer will be updated to
 * the next byte after the invalid sequence and the function will
 * return an error code.
 *
 * Depending on the specified flags, the function will also fail in
 * case the decoded code point does not belong to a valid range.
 *
 * @note For speed-relevant code a carefully implemented use of
 * GET_UTF8() may be preferred.
 *
 * @param codep   pointer used to return the parsed code in case of success.
 *                The value in *codep is set even in case the range check fails.
 * @param bufp    pointer to the address the first byte of the sequence
 *                to decode, updated by the function to point to the
 *                byte next after the decoded sequence
 * @param buf_end pointer to the end of the buffer, points to the next
 *                byte past the last in the buffer. This is used to
 *                avoid buffer overreads (in case of an unfinished
 *                UTF-8 sequence towards the end of the buffer).
 * @param flags   a collection of AV_UTF8_FLAG_* flags
 * @return >= 0 in case a sequence was successfully read, a negative
 * value in case of invalid sequence
 */
//av_warn_unused_result
//int av_utf8_decode(int32_t *codep, const uint8_t **bufp, const uint8_t *buf_end,
//unsigned int flags);
//todo
var avUtf8Decode func() ffcommon.FCharP
var avUtf8DecodeOnce sync.Once

func AvUtf8Decode() (res ffcommon.FCharP) {
	avUtf8DecodeOnce.Do(func() {
		purego.RegisterLibFunc(&avUtf8Decode, ffcommon.GetAvutilDll(), "av_utf8_decode")
	})
	res = avUtf8Decode()
	return
}

/**
 * Check if a name is in a list.
 * @returns 0 if not found, or the 1 based index where it has been found in the
 *            list.
 */
//int av_match_list(const char *name, const char *list, char separator);
//todo
var avMatchList func() ffcommon.FCharP
var avMatchListOnce sync.Once

func AvMatchList() (res ffcommon.FCharP) {
	avMatchListOnce.Do(func() {
		purego.RegisterLibFunc(&avMatchList, ffcommon.GetAvutilDll(), "av_match_list")
	})
	res = avMatchList()
	return
}

/**
 * See libc sscanf manual for more information.
 * Locale-independent sscanf implementation.
 */
//int av_sscanf(const char *string, const char *format, ...);
//todo
var avSscanf func() ffcommon.FCharP
var avSscanfOnce sync.Once

func AvSscanf() (res ffcommon.FCharP) {
	avSscanfOnce.Do(func() {
		purego.RegisterLibFunc(&avSscanf, ffcommon.GetAvutilDll(), "av_sscanf")
	})
	res = avSscanf()
	return
}

/**
 * @}
 */

//#endif /* AVUTIL_AVSTRING_H */
