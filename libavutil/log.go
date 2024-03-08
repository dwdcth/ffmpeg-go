package libavutil

import (
	"sync"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/ebitengine/purego"
)

/*
 * copyright (c) 2006 Michael Niedermayer <michaelni@gmx.at>
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

// #ifndef AVUTIL_LOG_H
// #define AVUTIL_LOG_H
//
// #include <stdarg.h>
// #include "avutil.h"
// #include "attributes.h"
// #include "version.h"
type AVClassCategory int32

const (
	AV_CLASS_CATEGORY_NA = iota
	AV_CLASS_CATEGORY_INPUT
	AV_CLASS_CATEGORY_OUTPUT
	AV_CLASS_CATEGORY_MUXER
	AV_CLASS_CATEGORY_DEMUXER
	AV_CLASS_CATEGORY_ENCODER
	AV_CLASS_CATEGORY_DECODER
	AV_CLASS_CATEGORY_FILTER
	AV_CLASS_CATEGORY_BITSTREAM_FILTER
	AV_CLASS_CATEGORY_SWSCALER
	AV_CLASS_CATEGORY_SWRESAMPLER
)
const (
	AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT = 40 + iota
	AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT
	AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT
	AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT
	AV_CLASS_CATEGORY_DEVICE_OUTPUT
	AV_CLASS_CATEGORY_DEVICE_INPUT
	AV_CLASS_CATEGORY_NB ///< not part of ABI/API
)

//#define AV_IS_INPUT_DEVICE(category) \
//(((category) == AV_CLASS_CATEGORY_DEVICE_VIDEO_INPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_AUDIO_INPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_INPUT))
//
//#define AV_IS_OUTPUT_DEVICE(category) \
//(((category) == AV_CLASS_CATEGORY_DEVICE_VIDEO_OUTPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_AUDIO_OUTPUT) || \
//((category) == AV_CLASS_CATEGORY_DEVICE_OUTPUT))

//struct AVOptionRanges;
//type AVOptionRanges struct {
//
//}
/**
 * Describe the class of an AVClass context structure. That is an
 * arbitrary struct of which the first field is a pointer to an
 * AVClass struct (e.g. AVCodecContext, AVFormatContext etc.).
 */
type AVClass struct {

	/**
	 * The name of the class; usually it is the same name as the
	 * context structure type to which the AVClass is associated.
	 */
	ClassName ffcommon.FCharPStruct

	/**
	 * A pointer to a function which returns the name of a context
	 * instance ctx associated with the class.
	 */
	//const char* (*item_name)(void* ctx);
	ItemName uintptr

	/**
	 * a pointer to the first option specified in the class if any or NULL
	 *
	 * @see av_set_default_options()
	 */
	Option *AVOption

	/**
	 * LIBAVUTIL_VERSION with which this structure was created.
	 * This is used to allow fields to be added without requiring major
	 * version bumps everywhere.
	 */

	Version ffcommon.FInt

	/**
	 * Offset in the structure where log_level_offset is stored.
	 * 0 means there is no such variable
	 */
	LogLevelOffsetOffset ffcommon.FInt

	/**
	 * Offset in the structure where a pointer to the parent context for
	 * logging is stored. For example a decoder could pass its AVCodecContext
	 * to eval as such a parent context, which an av_log() implementation
	 * could then leverage to display the parent context.
	 * The offset can be NULL.
	 */
	ParentLogContextOffset ffcommon.FInt

	/**
	 * Return next AVOptions-enabled child or NULL
	 */
	//void* (*child_next)(void *obj, void *prev);
	ChildNext uintptr

	//#if FF_API_CHILD_CLASS_NEXT
	///**
	// * Return an AVClass corresponding to the next potential
	// * AVOptions-enabled child.
	// *
	// * The difference between child_next and this is that
	// * child_next iterates over _already existing_ objects, while
	// * child_class_next iterates over _all possible_ children.
	// */
	//attribute_deprecated
	//const struct AVClass* (*child_class_next)(const struct AVClass *prev);
	ChildClassNext uintptr
	//#endif

	/**
	 * Category used for visualization (like color)
	 * This is only set if the category is equal for all objects using this class.
	 * available since version (51 << 16 | 56 << 8 | 100)
	 */
	Category AVClassCategory

	/**
	 * Callback to return the category.
	 * available since version (51 << 16 | 59 << 8 | 100)
	 */
	//VClassCategory (*get_category)(void* ctx);
	GetCategory uintptr

	/**
	 * Callback to return the supported/allowed ranges.
	 * available since version (52.12)
	 */
	//int (*query_ranges)(struct AVOptionRanges **, void *obj, const char *key, int flags);
	QueryRanges uintptr

	/**
	 * Iterate over the AVClasses corresponding to potential AVOptions-enabled
	 * children.
	 *
	 * @param iter pointer to opaque iteration state. The caller must initialize
	 *             *iter to NULL before the first call.
	 * @return AVClass for the next AVOptions-enabled child or NULL if there are
	 *         no more such children.
	 *
	 * @note The difference between child_next and this is that child_next
	 *       iterates over _already existing_ objects, while child_class_iterate
	 *       iterates over _all possible_ children.
	 */
	//const struct AVClass* (*child_class_iterate)(void **iter);
	ChildClassIterate uintptr
}

/**
 * @addtogroup lavu_log
 *
 * @{
 *
 * @defgroup lavu_log_constants Logging Constants
 *
 * @{
 */

/**
 * Print no output.
 */
const AV_LOG_QUIET = -8

/**
 * Something went really wrong and we will crash now.
 */
const AV_LOG_PANIC = 0

/**
 * Something went wrong and recovery is not possible.
 * For example, no header was found for a format which depends
 * on headers or an illegal combination of parameters is used.
 */
const AV_LOG_FATAL = 8

/**
 * Something went wrong and cannot losslessly be recovered.
 * However, not all future data is affected.
 */
const AV_LOG_ERROR = 16

/**
 * Something somehow does not look correct. This may or may not
 * lead to problems. An example would be the use of '-vstrict -2'.
 */
const AV_LOG_WARNING = 24

/**
 * Standard information.
 */
const AV_LOG_INFO = 32

/**
 * Detailed information.
 */
const AV_LOG_VERBOSE = 40

/**
 * Stuff which is only useful for libav* developers.
 */
const AV_LOG_DEBUG = 48

/**
 * Extremely verbose debugging, useful for libav* development.
 */
const AV_LOG_TRACE = 56

const AV_LOG_MAX_OFFSET = (AV_LOG_TRACE - AV_LOG_QUIET)

/**
 * @}
 */

/**
  * Sets additional colors for extended debugging sessions.
  * @code
    av_log(ctx, AV_LOG_DEBUG|AV_LOG_C(134), "Message in purple\n");
    @endcode
  * Requires 256color terminal support. Uses outside debugging is not
  * recommended.
*/
//const AV_LOG_C(x) ((x) << 8)

/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 */
//void av_log(void *avcl, int level, const char *fmt, ...) av_printf_format(3, 4);
var avLog func(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ...ffcommon.FConstCharP)
var avLogOnceS sync.Once

func AvLog(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ...ffcommon.FConstCharP) {
	avLogOnceS.Do(func() {
		purego.RegisterLibFunc(&avLog, ffcommon.GetAvutilDll(), "av_log")
	})
	// uintptrs := []uintptr{uintptr(avcl), uintptr(level)}
	// for i := 0; i < len(fmt0); i++ {
	// 	uintptrs = append(uintptrs, ffcommon.UintPtrFromString(fmt0[i]))
	// }
	// avLog(uintptrs...)
	avLog(avcl, level, fmt0...)
}

/**
 * Send the specified message to the log once with the initial_level and then with
 * the subsequent_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct or NULL if general log.
 * @param initial_level importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant" for the first occurance.
 * @param subsequent_level importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant" after the first occurance.
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param state a variable to keep trak of if a message has already been printed
 *        this must be initialized to 0 before the first use. The same state
 *        must not be accessed by 2 Threads simultaneously.
 */
//void av_log_once(void* avcl, int initial_level, int subsequent_level, int *state, const char *fmt, ...) av_printf_format(5, 6);
var avLogOnce func(avcl ffcommon.FVoidP, initial_level, subsequent_level ffcommon.FInt, state *ffcommon.FInt, fmt0 ...ffcommon.FConstCharP)
var avLogOnceOnce sync.Once

func AvLogOnce(avcl ffcommon.FVoidP, initial_level, subsequent_level ffcommon.FInt, state *ffcommon.FInt, fmt0 ...ffcommon.FConstCharP) {
	avLogOnceOnce.Do(func() {
		purego.RegisterLibFunc(&avLogOnce, ffcommon.GetAvutilDll(), "av_log_once")
	})
	// uintptrs := []uintptr{uintptr(avcl), uintptr(initial_level), uintptr(subsequent_level), uintptr(unsafe.Pointer(state))}
	// for i := 0; i < len(fmt0); i++ {
	// 	uintptrs = append(uintptrs, ffcommon.UintPtrFromString(fmt0[i]))
	// }
	// avLogOnce(uintptrs...)
	avLogOnce(avcl, initial_level, subsequent_level, state, fmt0...)
}

/**
 * Send the specified message to the log if the level is less than or equal
 * to the current av_log_level. By default, all logging messages are sent to
 * stderr. This behavior can be altered by setting a different logging callback
 * function.
 * @see av_log_set_callback
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
//void av_vlog(void *avcl, int level, const char *fmt, va_list vl);
var avVlog func(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList)
var avVlogOnce sync.Once

func AvVlog(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList) {
	avVlogOnce.Do(func() {
		purego.RegisterLibFunc(&avVlog, ffcommon.GetAvutilDll(), "av_vlog")
	})
	avVlog(avcl, level, fmt0, vl)
}

/**
 * Get the current log level
 *
 * @see lavu_log_constants
 *
 * @return Current log level
 */
//int av_log_get_level(void);
var avLogGetLevel func() ffcommon.FInt
var avLogGetLevelOnce sync.Once

func AvLogGetLevel() ffcommon.FInt {
	avLogGetLevelOnce.Do(func() {
		purego.RegisterLibFunc(&avLogGetLevel, ffcommon.GetAvutilDll(), "av_log_get_level")
	})
	return avLogGetLevel()
}

/**
 * Set the log level
 *
 * @see lavu_log_constants
 *
 * @param level Logging level
 */
//void av_log_set_level(int level);
var avLogSetLevel func(level ffcommon.FInt)
var avLogSetLevelOnce sync.Once

func AvLogSetLevel(level ffcommon.FInt) {
	avLogSetLevelOnce.Do(func() {
		purego.RegisterLibFunc(&avLogSetLevel, ffcommon.GetAvutilDll(), "av_log_set_level")
	})
	avLogSetLevel(level)
}

/**
 * Set the logging callback
 *
 * @note The callback must be thread safe, even if the application does not use
 *       threads itself as some codecs are multithreaded.
 *
 * @see av_log_default_callback
 *
 * @param callback A logging function with a compatible signature.
 */
//void av_log_set_callback(void (*callback)(void*, int, const char*, va_list));
var avLogSetCallback func(callback func(ffcommon.FVoidP, ffcommon.FInt, ffcommon.FCharPStruct, ffcommon.FVaList) uintptr)
var avLogSetCallbackOnce sync.Once

func AvLogSetCallback(callback func(ffcommon.FVoidP, ffcommon.FInt, ffcommon.FCharPStruct, ffcommon.FVaList) uintptr) {
	avLogSetCallbackOnce.Do(func() {
		purego.RegisterLibFunc(&avLogSetCallback, ffcommon.GetAvutilDll(), "av_log_set_callback")
	})
	avLogSetCallback(callback)
}

/**
 * Default logging callback
 *
 * It prints the message to stderr, optionally colorizing it.
 *
 * @param avcl A pointer to an arbitrary struct of which the first field is a
 *        pointer to an AVClass struct.
 * @param level The importance level of the message expressed using a @ref
 *        lavu_log_constants "Logging Constant".
 * @param fmt The format string (printf-compatible) that specifies how
 *        subsequent arguments are converted to output.
 * @param vl The arguments referenced by the format string.
 */
//void av_log_default_callback(void *avcl, int level, const char *fmt,
//va_list vl);
var avLogDefaultCallback func(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList)
var avLogDefaultCallbackOnce sync.Once

func AvLogDefaultCallback(avcl ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList) {
	avLogDefaultCallbackOnce.Do(func() {
		purego.RegisterLibFunc(&avLogDefaultCallback, ffcommon.GetAvutilDll(), "av_log_default_callback")
	})
	avLogDefaultCallback(avcl, level, fmt0, vl)
}

/**
 * Return the context name
 *
 * @param  ctx The AVClass context
 *
 * @return The AVClass class_name
 */
//const char* av_default_item_name(void* ctx);
var avDefaultItemName func(ctx ffcommon.FVoidP) ffcommon.FCharP
var avDefaultItemNameOnce sync.Once

func AvDefaultItemName(ctx ffcommon.FVoidP) ffcommon.FCharP {
	avDefaultItemNameOnce.Do(func() {
		purego.RegisterLibFunc(&avDefaultItemName, ffcommon.GetAvutilDll(), "av_default_item_name")
	})
	return avDefaultItemName(ctx)
}

// AVClassCategory av_default_get_category(void *ptr);
var avDefaultGetCategory func(ptr ffcommon.FVoidP) AVClassCategory
var avDefaultGetCategoryOnce sync.Once

func AvDefaultGetCategory(ptr ffcommon.FVoidP) AVClassCategory {
	avDefaultGetCategoryOnce.Do(func() {
		purego.RegisterLibFunc(&avDefaultGetCategory, ffcommon.GetAvutilDll(), "av_default_get_category")
	})
	return avDefaultGetCategory(ptr)
}

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line
 * @param line_size     size of the buffer
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 */
//void av_log_format_line(void *ptr, int level, const char *fmt, va_list vl,
//char *line, int line_size, int *print_prefix);
var avLogFormatLine func(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList,
	line ffcommon.FCharP, line_size ffcommon.FInt, print_prefix *ffcommon.FInt)
var avLogFormatLineOnce sync.Once

func AvLogFormatLine(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList,
	line ffcommon.FCharP, line_size ffcommon.FInt, print_prefix *ffcommon.FInt) {
	avLogFormatLineOnce.Do(func() {
		purego.RegisterLibFunc(&avLogFormatLine, ffcommon.GetAvutilDll(), "av_log_format_line")
	})
	avLogFormatLine(ptr, level, fmt0, vl, line, line_size, print_prefix)
}

/**
 * Format a line of log the same way as the default callback.
 * @param line          buffer to receive the formatted line;
 *                      may be NULL if line_size is 0
 * @param line_size     size of the buffer; at most line_size-1 characters will
 *                      be written to the buffer, plus one null terminator
 * @param print_prefix  used to store whether the prefix must be printed;
 *                      must point to a persistent integer initially set to 1
 * @return Returns a negative value if an error occurred, otherwise returns
 *         the number of characters that would have been written for a
 *         sufficiently large buffer, not including the terminating null
 *         character. If the return value is not less than line_size, it means
 *         that the log message was truncated to fit the buffer.
 */
//int av_log_format_line2(void *ptr, int level, const char *fmt, va_list vl,
//char *line, int line_size, int *print_prefix);
var avLogFormatLine2 func(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList,
	line ffcommon.FCharP, line_size ffcommon.FInt, print_prefix *ffcommon.FInt) ffcommon.FInt
var avLogFormatLine2Once sync.Once

func AvLogFormatLine2(ptr ffcommon.FVoidP, level ffcommon.FInt, fmt0 ffcommon.FConstCharP, vl ffcommon.FVaList,
	line ffcommon.FCharP, line_size ffcommon.FInt, print_prefix *ffcommon.FInt) ffcommon.FInt {
	avLogFormatLine2Once.Do(func() {
		purego.RegisterLibFunc(&avLogFormatLine2, ffcommon.GetAvutilDll(), "av_log_format_line2")
	})
	return avLogFormatLine2(ptr, level, fmt0, vl, line, line_size, print_prefix)
}

/**
 * Skip repeated messages, this requires the user app to use av_log() instead of
 * (f)printf as the 2 would otherwise interfere and lead to
 * "Last message repeated x times" messages below (f)printf messages with some
 * bad luck.
 * Also to receive the last, "last repeated" line if any, the user app must
 * call av_log(NULL, AV_LOG_QUIET, "%s", ""); at the end
 */
const AV_LOG_SKIP_REPEATED = 1

/**
 * Include the log severity in messages originating from codecs.
 *
 * Results in messages such as:
 * [rawvideo @ 0xDEADBEEF] [error] encode did not produce valid pts
 */
const AV_LOG_PRINT_LEVEL = 2

// void av_log_set_flags(int arg);
var avLogSetFlags func(arg ffcommon.FInt) ffcommon.FCharP
var avLogSetFlagsOnce sync.Once

func AvLogSetFlags(arg ffcommon.FInt) ffcommon.FCharP {
	avLogSetFlagsOnce.Do(func() {
		purego.RegisterLibFunc(&avLogSetFlags, ffcommon.GetAvutilDll(), "av_log_set_flags")
	})
	return avLogSetFlags(arg)
}

// int av_log_get_flags(void);
var avLogGetFlags func() ffcommon.FInt
var avLogGetFlagsOnce sync.Once

func AvLogGetFlags() ffcommon.FInt {
	avLogGetFlagsOnce.Do(func() {
		purego.RegisterLibFunc(&avLogGetFlags, ffcommon.GetAvutilDll(), "av_log_get_flags")
	})
	return avLogGetFlags()
}

/**
 * @}
 */

//#endif /* AVUTIL_LOG_H */
