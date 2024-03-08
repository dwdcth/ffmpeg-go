package libavfilter

import (
	"sync"
	"unsafe"

	"github.com/dwdcth/ffmpeg-go/v6/ffcommon"
	"github.com/dwdcth/ffmpeg-go/v6/libavutil"
	"github.com/ebitengine/purego"
)

/*
 * filter layer
 * Copyright (c) 2007 Bobby Bingham
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

//#ifndef AVFILTER_AVFILTER_H
//#define AVFILTER_AVFILTER_H

/**
 * @file
 * @ingroup lavfi
 * Main libavfilter public API header
 */

/**
 * @defgroup lavfi libavfilter
 * Graph-based frame editing library.
 *
 * @{
 */

//#include <stddef.h>
//
//#include "../libavutil/attributes.h"
//#include "../libavutil/avutil.h"
//#include "../libavutil/buffer.h"
//#include "../libavutil/dict.h"
//#include "../libavutil/frame.h"
//#include "../libavutil/log.h"
//#include "../libavutil/samplefmt.h"
//#include "../libavutil/pixfmt.h"
//#include "../libavutil/rational.h"
//
//#include "../libavfilter/version.h"

/**
 * Return the LIBAVFILTER_VERSION_INT constant.
 */
//unsigned avfilter_version(void);
var avfilterVersion func() ffcommon.FUnsigned
var avfilterVersionOnce sync.Once

func AvfilterVersion() ffcommon.FUnsigned {
	avfilterVersionOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterVersion, ffcommon.GetAvfilterDll(), "avfilter_version")
	})
	if avfilterVersion != nil {
		return avfilterVersion()
	}
	return 0
}

/**
 * Return the libavfilter build-time configuration.
 */
//const char *avfilter_configuration(void);
var avfilterConfiguration func() ffcommon.FConstCharP
var avfilterConfigurationOnce sync.Once

func AvfilterConfiguration() ffcommon.FConstCharP {
	avfilterConfigurationOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterConfiguration, ffcommon.GetAvfilterDll(), "avfilter_configuration")
	})
	if avfilterConfiguration != nil {
		return avfilterConfiguration()
	}
	return ""
}

/**
 * Return the libavfilter license.
 */
//const char *avfilter_license(void);
var avfilterLicense func() ffcommon.FConstCharP
var avfilterLicenseOnce sync.Once

func AvfilterLicense() ffcommon.FConstCharP {
	avfilterLicenseOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterLicense, ffcommon.GetAvfilterDll(), "avfilter_license")
	})
	if avfilterLicense != nil {
		return avfilterLicense()
	}
	return ""
}

// typedef struct AVFilterContext AVFilterContext;
// typedef struct AVFilterLink    AVFilterLink;
// typedef struct AVFilterPad     AVFilterPad;
type AVFrame = libavutil.AVFrame
type AVMediaType = libavutil.AVMediaType
type AVFilterPad struct {
}

// typedef struct AVFilterFormats AVFilterFormats;
type AVFilterFormats struct {
}

// typedef struct AVFilterChannelLayouts AVFilterChannelLayouts;
type AVFilterChannelLayouts struct {
}

/**
 * Get the number of elements in a NULL-terminated array of AVFilterPads (e.g.
 * AVFilter.inputs/outputs).
 */
//int avfilter_pad_count(const AVFilterPad *pads);
var avfilterPadCount func(pads *AVFilterPad) ffcommon.FInt
var avfilterPadCountOnce sync.Once

func (pads *AVFilterPad) AvfilterPadCount() ffcommon.FInt {
	avfilterPadCountOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterPadCount, ffcommon.GetAvfilterDll(), "avfilter_pad_count")
	})
	if avfilterPadCount != nil {
		return avfilterPadCount(pads)
	}
	return 0
}

/**
 * Get the name of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return name of the pad_idx'th pad in pads
 */
//const char *avfilter_pad_get_name(const AVFilterPad *pads, int pad_idx);
var avfilterPadGetName func(pads *AVFilterPad, pad_idx ffcommon.FInt) ffcommon.FConstCharP
var avfilterPadGetNameOnce sync.Once

func (pads *AVFilterPad) AvfilterPadGetName(pad_idx ffcommon.FInt) ffcommon.FConstCharP {
	avfilterPadGetNameOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterPadGetName, ffcommon.GetAvfilterDll(), "avfilter_pad_get_name")
	})
	return avfilterPadGetName(pads, pad_idx)
}

/**
 * Get the type of an AVFilterPad.
 *
 * @param pads an array of AVFilterPads
 * @param pad_idx index of the pad in the array; it is the caller's
 *                responsibility to ensure the index is valid
 *
 * @return type of the pad_idx'th pad in pads
 */
//enum AVMediaType avfilter_pad_get_type(const AVFilterPad *pads, int pad_idx);
var avfilterPadGetType func(pads *AVFilterPad, pad_idx ffcommon.FInt) AVMediaType
var avfilterPadGetTypeOnce sync.Once

func (pads *AVFilterPad) AvfilterPadGetType(pad_idx ffcommon.FInt) AVMediaType {
	avfilterPadGetTypeOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterPadGetType, ffcommon.GetAvfilterDll(), "avfilter_pad_get_type")
	})
	return avfilterPadGetType(pads, pad_idx)
}

/**
 * The number of the filter inputs is not determined just by AVFilter.inputs.
 * The filter might add additional inputs during initialization depending on the
 * options supplied to it.
 */
const AVFILTER_FLAG_DYNAMIC_INPUTS = (1 << 0)

/**
 * The number of the filter outputs is not determined just by AVFilter.outputs.
 * The filter might add additional outputs during initialization depending on
 * the options supplied to it.
 */
const AVFILTER_FLAG_DYNAMIC_OUTPUTS = (1 << 1)

/**
 * The filter supports multithreading by splitting frames into multiple parts
 * and processing them concurrently.
 */
const AVFILTER_FLAG_SLICE_THREADS = (1 << 2)

/**
 * Some filters support a generic "enable" expression option that can be used
 * to enable or disable a filter in the timeline. Filters supporting this
 * option have this flag set. When the enable expression is false, the default
 * no-op filter_frame() function is called in place of the filter_frame()
 * callback defined on each input pad, thus the frame is passed unchanged to
 * the next filters.
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC = (1 << 16)

/**
 * Same as AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC, except that the filter will
 * have its filter_frame() callback(s) called as usual even when the enable
 * expression is false. The filter will disable filtering within the
 * filter_frame() callback(s) itself, for example executing code depending on
 * the AVFilterContext->is_disabled value.
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL = (1 << 17)

/**
 * Handy mask to test whether the filter supports or no the timeline feature
 * (internally or generically).
 */
const AVFILTER_FLAG_SUPPORT_TIMELINE = (AVFILTER_FLAG_SUPPORT_TIMELINE_GENERIC | AVFILTER_FLAG_SUPPORT_TIMELINE_INTERNAL)

/**
 * Filter definition. This defines the pads a filter contains, and all the
 * callback functions used to interact with the filter.
 */
type AVFilter struct {

	/**
	 * Filter name. Must be non-NULL and unique among filters.
	 */
	Name ffcommon.FCharPStruct

	/**
	 * A description of the filter. May be NULL.
	 *
	 * You should use the NULL_IF_CONFIG_SMALL() macro to define it.
	 */
	Description ffcommon.FCharPStruct

	/**
	 * List of inputs, terminated by a zeroed element.
	 *
	 * NULL if there are no (static) inputs. Instances of filters with
	 * AVFILTER_FLAG_DYNAMIC_INPUTS set may have more inputs than present in
	 * this list.
	 */
	Inputs *AVFilterPad
	/**
	 * List of outputs, terminated by a zeroed element.
	 *
	 * NULL if there are no (static) outputs. Instances of filters with
	 * AVFILTER_FLAG_DYNAMIC_OUTPUTS set may have more outputs than present in
	 * this list.
	 */
	Outputs *AVFilterPad

	/**
	 * A class for the private data, used to declare filter private AVOptions.
	 * This field is NULL for filters that do not declare any options.
	 *
	 * If this field is non-NULL, the first member of the filter private data
	 * must be a pointer to AVClass, which will be set by libavfilter generic
	 * code to this class.
	 */
	PrivClass *AVClass

	/**
	 * A combination of AVFILTER_FLAG_*
	 */
	Flags ffcommon.FInt

	/*****************************************************************
	 * All fields below this line are not part of the public API. They
	 * may not be used outside of libavfilter and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */

	/**
	 * Filter pre-initialization function
	 *
	 * This callback will be called immediately after the filter context is
	 * allocated, to allow allocating and initing sub-objects.
	 *
	 * If this callback is not NULL, the uninit callback will be called on
	 * allocation failure.
	 *
	 * @return 0 on success,
	 *         AVERROR code on failure (but the code will be
	 *           dropped and treated as ENOMEM by the calling code)
	 */
	//int (*preinit)(AVFilterContext *ctx);
	Preinit uintptr
	/**
	 * Filter initialization function.
	 *
	 * This callback will be called only once during the filter lifetime, after
	 * all the options have been set, but before links between filters are
	 * established and format negotiation is done.
	 *
	 * Basic filter initialization should be done here. Filters with dynamic
	 * inputs and/or outputs should create those inputs/outputs here based on
	 * provided options. No more changes to this filter's inputs/outputs can be
	 * done after this callback.
	 *
	 * This callback must not assume that the filter links exist or frame
	 * parameters are known.
	 *
	 * @ref AVFilter.uninit "uninit" is guaranteed to be called even if
	 * initialization fails, so this callback does not have to clean up on
	 * failure.
	 *
	 * @return 0 on success, a negative AVERROR on failure
	 */
	//int (*init)(AVFilterContext *ctx);
	Init uintptr
	/**
	 * Should be set instead of @ref AVFilter.init "init" by the filters that
	 * want to pass a dictionary of AVOptions to nested contexts that are
	 * allocated during init.
	 *
	 * On return, the options dict should be freed and replaced with one that
	 * contains all the options which could not be processed by this filter (or
	 * with NULL if all the options were processed).
	 *
	 * Otherwise the semantics is the same as for @ref AVFilter.init "init".
	 */
	//int (*init_dict)(AVFilterContext *ctx, AVDictionary **options);
	InitDict uintptr
	/**
	 * Filter uninitialization function.
	 *
	 * Called only once right before the filter is freed. Should deallocate any
	 * memory held by the filter, release any buffer references, etc. It does
	 * not need to deallocate the AVFilterContext.priv memory itself.
	 *
	 * This callback may be called even if @ref AVFilter.init "init" was not
	 * called or failed, so it must be prepared to handle such a situation.
	 */
	//void (*uninit)(AVFilterContext *ctx);
	Uninit uintptr
	/**
	 * Query formats supported by the filter on its inputs and outputs.
	 *
	 * This callback is called after the filter is initialized (so the inputs
	 * and outputs are fixed), shortly before the format negotiation. This
	 * callback may be called more than once.
	 *
	 * This callback must set AVFilterLink.outcfg.formats on every input link and
	 * AVFilterLink.incfg.formats on every output link to a list of pixel/sample
	 * formats that the filter supports on that link. For audio links, this
	 * filter must also set @ref AVFilterLink.incfg.samplerates "in_samplerates" /
	 * @ref AVFilterLink.outcfg.samplerates "out_samplerates" and
	 * @ref AVFilterLink.incfg.channel_layouts "in_channel_layouts" /
	 * @ref AVFilterLink.outcfg.channel_layouts "out_channel_layouts" analogously.
	 *
	 * This callback may be NULL for filters with one input, in which case
	 * libavfilter assumes that it supports all input formats and preserves
	 * them on output.
	 *
	 * @return zero on success, a negative value corresponding to an
	 * AVERROR code otherwise
	 */
	//int (*query_formats)(AVFilterContext *);
	QueryFormats uintptr
	PrivSize     ffcommon.FInt ///< size of private data to allocate for the filter

	FlagsInternal ffcommon.FInt ///< Additional flags for avfilter internal use only.

	//#if FF_API_NEXT
	/**
	 * Used by the filter registration system. Must not be touched by any other
	 * code.
	 */
	Next *AVFilter
	//#endif

	/**
	 * Make the filter instance process a command.
	 *
	 * @param cmd    the command to process, for handling simplicity all commands must be alphanumeric only
	 * @param arg    the argument for the command
	 * @param res    a buffer with size res_size where the filter(s) can return a response. This must not change when the command is not supported.
	 * @param flags  if AVFILTER_CMD_FLAG_FAST is set and the command would be
	 *               time consuming then a filter should treat it like an unsupported command
	 *
	 * @returns >=0 on success otherwise an error code.
	 *          AVERROR(ENOSYS) on unsupported commands
	 */
	//int (*process_command)(AVFilterContext *, const char *cmd, const char *arg, char *res, int res_len, int flags);
	ProcessCommand uintptr
	/**
	 * Filter initialization function, alternative to the init()
	 * callback. Args contains the user-supplied parameters, opaque is
	 * used for providing binary data.
	 */
	//int (*init_opaque)(AVFilterContext *ctx, void *opaque);
	InitOpaque uintptr
	/**
	 * Filter activation function.
	 *
	 * Called when any processing is needed from the filter, instead of any
	 * filter_frame and request_frame on pads.
	 *
	 * The function must examine inlinks and outlinks and perform a single
	 * step of processing. If there is nothing to do, the function must do
	 * nothing and not return an error. If more steps are or may be
	 * possible, it must use ff_filter_set_ready() to schedule another
	 * activation.
	 */
	//int (*activate)(AVFilterContext *ctx);
	Activate uintptr
}

/**
 * Process multiple parts of the frame concurrently.
 */
const AVFILTER_THREAD_SLICE = (1 << 0)

type AVFilterInternal struct {
}

/** An instance of a filter */
type AVClass = libavutil.AVClass
type AVFilterContext struct {
	AvClass *AVClass ///< needed for av_log() and filters common options

	Filter *AVFilter ///< the AVFilter of which this is an instance

	Name ffcommon.FCharPStruct ///< name of this filter instance

	InputPads *AVFilterPad       ///< array of input pads
	Inputs    **AVFilterLink     ///< array of pointers to input links
	NbInputs  ffcommon.FUnsigned ///< number of input pads

	OutputPads *AVFilterPad       ///< array of output pads
	Outputs    **AVFilterLink     ///< array of pointers to output links
	NbOutputs  ffcommon.FUnsigned ///< number of output pads

	Priv ffcommon.FVoidP ///< private data for use by the filter

	Graph *AVFilterGraph ///< filtergraph this filter belongs to

	/**
	 * Type of multithreading being allowed/used. A combination of
	 * AVFILTER_THREAD_* flags.
	 *
	 * May be set by the caller before initializing the filter to forbid some
	 * or all kinds of multithreading for this filter. The default is allowing
	 * everything.
	 *
	 * When the filter is initialized, this field is combined using bit AND with
	 * AVFilterGraph.thread_type to get the final mask used for determining
	 * allowed threading types. I.e. a threading type needs to be set in both
	 * to be allowed.
	 *
	 * After the filter is initialized, libavfilter sets this field to the
	 * threading type that is actually used (0 for no multithreading).
	 */
	ThreadType ffcommon.FInt

	/**
	 * An opaque struct for libavfilter internal use.
	 */
	Internal *AVFilterInternal

	CommandQueue uintptr //*AVFilterCommand//todo

	EnableStr  ffcommon.FCharPStruct ///< enable expression string
	Enable     ffcommon.FVoidP       ///< parsed expression (AVExpr*)
	VarValues  *ffcommon.FDouble     ///< variable values for the enable expression
	IsDisabled *ffcommon.FInt        ///< the enabled state from the last expression evaluation

	/**
	 * For filters which will create hardware frames, sets the device the
	 * filter should create them in.  All other filters will ignore this field:
	 * in particular, a filter which consumes or processes hardware frames will
	 * instead use the hw_frames_ctx field in AVFilterLink to carry the
	 * hardware context information.
	 */
	HwDeviceCtx *AVBufferRef

	/**
	 * Max number of threads allowed in this filter instance.
	 * If <= 0, its value is ignored.
	 * Overrides global number of threads set per filter graph.
	 */
	NbThreads ffcommon.FInt

	/**
	 * Ready status of the filter.
	 * A non-0 value means that the filter needs activating;
	 * a higher value suggests a more urgent activation.
	 */
	Ready ffcommon.FUnsigned

	/**
	 * Sets the number of extra hardware frames which the filter will
	 * allocate on its output links for use in following filters or by
	 * the caller.
	 *
	 * Some hardware filters require all frames that they will use for
	 * output to be defined in advance before filtering starts.  For such
	 * filters, any hardware frame pools used for output must therefore be
	 * of fixed size.  The extra frames set here are on top of any number
	 * that the filter needs internally in order to operate normally.
	 *
	 * This field must be set before the graph containing this filter is
	 * configured.
	 */
	ExtraHwFrames ffcommon.FInt
}

func (this *AVFilterContext) GetInput(index ffcommon.FUnsignedInt) (res *AVFilterLink) {
	t := uintptr(unsafe.Pointer(this.Inputs)) + 8*uintptr(index)
	t = *(*uintptr)(unsafe.Pointer(t))
	res = (*AVFilterLink)(unsafe.Pointer(t))
	return
}

/*
*

  - Lists of formats / etc. supported by an end of a link.
    *

  - This structure is directly part of AVFilterLink, in two copies:

  - one for the source filter, one for the destination filter.

  - These lists are used for negotiating the format to actually be used,

  - which will be loaded into the format and channel_layout members of

  - AVFilterLink, when chosen.
*/
type AVFilterFormatsConfig struct {

	/**
	 * List of supported formats (pixel or sample).
	 */
	Formats *AVFilterFormats

	/**
	 * Lists of supported sample rates, only for audio.
	 */
	Samplerates *AVFilterFormats

	/**
	 * Lists of supported channel layouts, only for audio.
	 */
	ChannelLayouts *AVFilterChannelLayouts
}

/**
 * A link between two filters. This contains pointers to the source and
 * destination filters between which this link exists, and the indexes of
 * the pads involved. In addition, this link also contains the parameters
 * which have been negotiated and agreed upon between the filter, such as
 * image dimensions, format, etc.
 *
 * Applications must not normally access the link structure directly.
 * Use the buffersrc and buffersink API instead.
 * In the future, access to the header may be reserved for filters
 * implementation.
 */
type AVLINK int32

const (
	AVLINK_UNINIT    = 0    ///< not started
	AVLINK_STARTINIT = iota ///< started, but incomplete
	AVLINK_INIT      = iota ///< complete
)

type AVFilterLink struct {
	Src    *AVFilterContext ///< source filter
	Srcpad *AVFilterPad     ///< output pad on the source filter

	Dst    *AVFilterContext ///< dest filter
	Dstpad *AVFilterPad     ///< input pad on the dest filter

	Type AVMediaType ///< filter media type

	/* These parameters apply only to video */
	W                 ffcommon.FInt ///< agreed upon image width
	H                 ffcommon.FInt ///< agreed upon image height
	SampleAspectRatio AVRational    ///< agreed upon sample aspect ratio
	/* These parameters apply only to audio */
	ChannelLayout ffcommon.FUint64T ///< channel layout of current buffer (see libavutil/channel_layout.h)
	SampleRate    ffcommon.FInt     ///< samples per second

	Format ffcommon.FInt ///< agreed upon media format

	/**
	 * Define the time base used by the PTS of the frames/samples
	 * which will pass through this link.
	 * During the configuration stage, each filter is supposed to
	 * change only the output timebase, while the timebase of the
	 * input link is assumed to be an unchangeable property.
	 */
	TimeBase AVRational

	/*****************************************************************
	 * All fields below this line are not part of the public API. They
	 * may not be used outside of libavfilter and can be changed and
	 * removed at will.
	 * New public fields should be added right above.
	 *****************************************************************
	 */

	/**
	 * Lists of supported formats / etc. supported by the input filter.
	 */
	Incfg AVFilterFormatsConfig

	/**
	 * Lists of supported formats / etc. supported by the output filter.
	 */
	Outcfg AVFilterFormatsConfig

	/** stage of the initialization of the link properties (dimensions, etc) */
	//enum {
	//AVLINK_UNINIT = 0,      ///< not started
	//AVLINK_STARTINIT,       ///< started, but incomplete
	//AVLINK_INIT             ///< complete
	//} init_state;
	InitState AVLINK
	/**
	 * Graph the filter belongs to.
	 */
	Graph *AVFilterGraph

	/**
	 * Current timestamp of the link, as defined by the most recent
	 * frame(s), in link time_base units.
	 */
	CurrentPts ffcommon.FInt64T

	/**
	 * Current timestamp of the link, as defined by the most recent
	 * frame(s), in AV_TIME_BASE units.
	 */
	CurrentPtsUs ffcommon.FInt64T

	/**
	 * Index in the age array.
	 */
	AgeIndex ffcommon.FInt

	/**
	 * Frame rate of the stream on the link, or 1/0 if unknown or variable;
	 * if left to 0/0, will be automatically copied from the first input
	 * of the source filter if it exists.
	 *
	 * Sources should set it to the best estimation of the real frame rate.
	 * If the source frame rate is unknown or variable, set this to 1/0.
	 * Filters should update it if necessary depending on their function.
	 * Sinks can use it to set a default output frame rate.
	 * It is similar to the r_frame_rate field in AVStream.
	 */
	FrameRate AVRational

	/**
	 * Buffer partially filled with samples to achieve a fixed/minimum size.
	 */
	PartialBuf *AVFrame

	/**
	 * Size of the partial buffer to allocate.
	 * Must be between min_samples and max_samples.
	 */
	PartialBufSize ffcommon.FInt

	/**
	 * Minimum number of samples to filter at once. If filter_frame() is
	 * called with fewer samples, it will accumulate them in partial_buf.
	 * This field and the related ones must not be changed after filtering
	 * has started.
	 * If 0, all related fields are ignored.
	 */
	MinSamples ffcommon.FInt

	/**
	 * Maximum number of samples to filter at once. If filter_frame() is
	 * called with more samples, it will split them.
	 */
	MaxSamples ffcommon.FInt

	/**
	 * Number of channels.
	 */
	Channels ffcommon.FInt

	/**
	 * Number of past frames sent through the link.
	 */
	FrameCountIn, FrameCountOut ffcommon.FInt64T

	/**
	 * A pointer to a FFFramePool struct.
	 */
	FramePool ffcommon.FVoidP

	/**
	 * True if a frame is currently wanted on the output of this filter.
	 * Set when ff_request_frame() is called by the output,
	 * cleared when a frame is filtered.
	 */
	FrameWantedOut ffcommon.FInt

	/**
	 * For hwaccel pixel formats, this should be a reference to the
	 * AVHWFramesContext describing the frames.
	 */
	HwFramesCtx *AVBufferRef

	//#ifndef FF_INTERNAL_FIELDS

	/**
	 * Internal structure members.
	 * The fields below this limit are internal for libavfilter's use
	 * and must in no way be accessed by applications.
	 */
	Reserved [0xF000]ffcommon.FChar

	//#else /* FF_INTERNAL_FIELDS */
	//
	///**
	// * Queue of frames waiting to be filtered.
	// */
	//FFFrameQueue fifo;
	//
	///**
	// * If set, the source filter can not generate a frame as is.
	// * The goal is to avoid repeatedly calling the request_frame() method on
	// * the same link.
	// */
	//int frame_blocked_in;
	//
	///**
	// * Link input status.
	// * If not zero, all attempts of filter_frame will fail with the
	// * corresponding code.
	// */
	//int status_in;
	//
	///**
	// * Timestamp of the input status change.
	// */
	//int64_t status_in_pts;
	//
	///**
	// * Link output status.
	// * If not zero, all attempts of request_frame will fail with the
	// * corresponding code.
	// */
	//int status_out;
	//
	//#endif /* FF_INTERNAL_FIELDS */

}

/**
 * Link two filters together.
 *
 * @param src    the source filter
 * @param srcpad index of the output pad on the source filter
 * @param dst    the destination filter
 * @param dstpad index of the input pad on the destination filter
 * @return       zero on success
 */
//int avfilter_link(AVFilterContext *src, unsigned srcpad,
//AVFilterContext *dst, unsigned dstpad);
var avfilterLink func(src *AVFilterContext, srcpad ffcommon.FUnsigned, dst *AVFilterContext, dstpad ffcommon.FUnsigned) ffcommon.FInt
var avfilterLinkOnce sync.Once

func (src *AVFilterContext) AvfilterLink(srcpad ffcommon.FUnsigned, dst *AVFilterContext, dstpad ffcommon.FUnsigned) ffcommon.FInt {
	avfilterLinkOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterLink, ffcommon.GetAvfilterDll(), "avfilter_link")
	})
	if avfilterLink != nil {
		return avfilterLink(src, srcpad, dst, dstpad)
	}
	return 0 // Return a default value or handle the error accordingly
}

/**
 * Free the link in *link, and set its pointer to NULL.
 */
//void avfilter_link_free(AVFilterLink **link);
var avfilterLinkFree func(link **AVFilterLink)
var avfilterLinkFreeOnce sync.Once

func AvfilterLinkFree(link **AVFilterLink) {
	avfilterLinkFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterLinkFree, ffcommon.GetAvfilterDll(), "avfilter_link_free")
	})
	if avfilterLinkFree != nil {
		avfilterLinkFree(link)
	} else {
		// Handle the error or return a default value
	}
}

//#if FF_API_FILTER_GET_SET
/**
 * Get the number of channels of a link.
 * @deprecated Use av_buffersink_get_channels()
 */
//attribute_deprecated
//int avfilter_link_get_channels(AVFilterLink *link);
var avfilterLinkGetChannels func(link *AVFilterLink) ffcommon.FInt
var avfilterLinkGetChannelsOnce sync.Once

func (link *AVFilterLink) AvfilterLinkGetChannels() ffcommon.FInt {
	avfilterLinkGetChannelsOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterLinkGetChannels, ffcommon.GetAvfilterDll(), "avfilter_link_get_channels")
	})
	if avfilterLinkGetChannels != nil {
		return avfilterLinkGetChannels(link)
	} else {
		// Handle the error or return a default value
		return 0 // Default value
	}
}

//#endif
//#if FF_API_FILTER_LINK_SET_CLOSED
/**
 * Set the closed field of a link.
 * @deprecated applications are not supposed to mess with links, they should
 * close the sinks.
 */
//attribute_deprecated
//void avfilter_link_set_closed(AVFilterLink *link, int closed);
var avfilterLinkSetClosed func(link *AVFilterLink, closed ffcommon.FInt)
var avfilterLinkSetClosedOnce sync.Once

func (link *AVFilterLink) AvfilterLinkSetClosed(closed ffcommon.FInt) {
	avfilterLinkSetClosedOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterLinkSetClosed, ffcommon.GetAvfilterDll(), "avfilter_link_set_closed")
	})
	if avfilterLinkSetClosed != nil {
		avfilterLinkSetClosed(link, closed)
	} else {
		// Handle the error
	}
}

//#endif
/**
 * Negotiate the media format, dimensions, etc of all inputs to a filter.
 *
 * @param filter the filter to negotiate the properties for its inputs
 * @return       zero on successful negotiation
 */
//int avfilter_config_links(AVFilterContext *filter);
var avfilterConfigLinks func(filter *AVFilterContext) ffcommon.FInt
var avfilterConfigLinksOnce sync.Once

func (filter *AVFilterContext) AvfilterConfigLinks() ffcommon.FInt {
	avfilterConfigLinksOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterConfigLinks, ffcommon.GetAvfilterDll(), "avfilter_config_links")
	})
	if avfilterConfigLinks != nil {
		return avfilterConfigLinks(filter)
	} else {
		// Handle the error
		return ffcommon.FInt(0) // or return an appropriate error value
	}
}

const AVFILTER_CMD_FLAG_ONE = 1  ///< Stop once a filter understood the command (for target=all for example), fast filters are favored automatically
const AVFILTER_CMD_FLAG_FAST = 2 ///< Only execute command when its fast (like a video out that supports contrast adjustment in hw)

/**
 * Make the filter instance process a command.
 * It is recommended to use avfilter_graph_send_command().
 */
//int avfilter_process_command(AVFilterContext *filter, const char *cmd, const char *arg, char *res, int res_len, int flags);
var avfilterProcessCommand func(filter *AVFilterContext, cmd, arg, res0 ffcommon.FConstCharP, res_len, flags ffcommon.FInt) ffcommon.FInt
var avfilterProcessCommandOnce sync.Once

func (filter *AVFilterContext) AvfilterProcessCommand(cmd, arg, res0 ffcommon.FConstCharP, res_len, flags ffcommon.FInt) ffcommon.FInt {
	avfilterProcessCommandOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterProcessCommand, ffcommon.GetAvfilterDll(), "avfilter_process_command")
	})
	if avfilterProcessCommand != nil {
		return avfilterProcessCommand(filter, cmd, arg, res0, res_len, flags)
	} else {
		// Handle the error
		return ffcommon.FInt(0) // or return an appropriate error value
	}
}

/**
 * Iterate over all registered filters.
 *
 * @param opaque a pointer where libavfilter will store the iteration state. Must
 *               point to NULL to start the iteration.
 *
 * @return the next registered filter or NULL when the iteration is
 *         finished
 */
//const AVFilter *av_filter_iterate(void **opaque);
var avFilterIterate func(opaque *ffcommon.FVoidP) *AVFilter
var avFilterIterateOnce sync.Once

func AvFilterIterate(opaque *ffcommon.FVoidP) *AVFilter {
	avFilterIterateOnce.Do(func() {
		purego.RegisterLibFunc(&avFilterIterate, ffcommon.GetAvfilterDll(), "av_filter_iterate")
	})
	if avFilterIterate != nil {
		return avFilterIterate(opaque)
	} else {
		// Handle the error
		return nil // or return an appropriate error value
	}
}

//#if FF_API_NEXT
/** Initialize the filter system. Register all builtin filters. */
//attribute_deprecated
//void avfilter_register_all(void);
var avfilterRegisterAll func()
var avfilterRegisterAllOnce sync.Once

func AvfilterRegisterAll() {
	avfilterRegisterAllOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterRegisterAll, ffcommon.GetAvfilterDll(), "avfilter_register_all")
	})
	if avfilterRegisterAll != nil {
		avfilterRegisterAll()
	} else {
		// Handle the error
	}
}

/**
 * Register a filter. This is only needed if you plan to use
 * avfilter_get_by_name later to lookup the AVFilter structure by name. A
 * filter can still by instantiated with avfilter_graph_alloc_filter even if it
 * is not registered.
 *
 * @param filter the filter to register
 * @return 0 if the registration was successful, a negative value
 * otherwise
 */
//attribute_deprecated
//int avfilter_register(AVFilter *filter);
var avfilterRegister func(filter *AVFilter) ffcommon.FInt
var avfilterRegisterOnce sync.Once

func AvfilterRegister(filter *AVFilter) ffcommon.FInt {
	avfilterRegisterOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterRegister, ffcommon.GetAvfilterDll(), "avfilter_register")
	})
	if avfilterRegister != nil {
		return avfilterRegister(filter)
	} else {
		// Handle the error
		return ffcommon.FInt(-1) // or return an appropriate error value
	}
}

/**
 * Iterate over all registered filters.
 * @return If prev is non-NULL, next registered filter after prev or NULL if
 * prev is the last filter. If prev is NULL, return the first registered filter.
 */
//attribute_deprecated
//const AVFilter *avfilter_next(const AVFilter *prev);
var avfilterNext func(*AVFilter) *AVFilter
var avfilterNextOnce sync.Once

func AvfilterNext(prev *AVFilter) *AVFilter {
	avfilterNextOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterNext, ffcommon.GetAvfilterDll(), "avfilter_next")
	})
	if avfilterNext != nil {
		return avfilterNext(prev)
	} else {
		// Handle the error
		return nil
	}
}

//#endif

/**
 * Get a filter definition matching the given name.
 *
 * @param name the filter name to find
 * @return     the filter definition, if any matching one is registered.
 *             NULL if none found.
 */
//const AVFilter *avfilter_get_by_name(const char *name);
var avfilterGetByName func(ffcommon.FConstCharP) *AVFilter
var avfilterGetByNameOnce sync.Once

func AvfilterGetByName(name ffcommon.FConstCharP) *AVFilter {
	avfilterGetByNameOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGetByName, ffcommon.GetAvfilterDll(), "avfilter_get_by_name")
	})
	if avfilterGetByName != nil {
		return avfilterGetByName(name)
	} else {
		// Handle the error
		return nil
	}
}

/**
 * Initialize a filter with the supplied parameters.
 *
 * @param ctx  uninitialized filter context to initialize
 * @param args Options to initialize the filter with. This must be a
 *             ':'-separated list of options in the 'key=value' form.
 *             May be NULL if the options have been set directly using the
 *             AVOptions API or there are no options that need to be set.
 * @return 0 on success, a negative AVERROR on failure
 */
//int avfilter_init_str(AVFilterContext *ctx, const char *args);
var avfilterInitStr func(ctx *AVFilterContext, args ffcommon.FConstCharP) ffcommon.FInt
var avfilterInitStrOnce sync.Once

func (ctx *AVFilterContext) AvfilterInitStr(args ffcommon.FConstCharP) ffcommon.FInt {
	avfilterInitStrOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterInitStr, ffcommon.GetAvfilterDll(), "avfilter_init_str")
	})

	// argsPtr := uintptr(0)
	// if args != "" {
	// 	argsPtr = ffcommon.UintPtrFromString(args)
	// }

	return avfilterInitStr(ctx, args)
}

/**
 * Initialize a filter with the supplied dictionary of options.
 *
 * @param ctx     uninitialized filter context to initialize
 * @param options An AVDictionary filled with options for this filter. On
 *                return this parameter will be destroyed and replaced with
 *                a dict containing options that were not found. This dictionary
 *                must be freed by the caller.
 *                May be NULL, then this function is equivalent to
 *                avfilter_init_str() with the second parameter set to NULL.
 * @return 0 on success, a negative AVERROR on failure
 *
 * @note This function and avfilter_init_str() do essentially the same thing,
 * the difference is in manner in which the options are passed. It is up to the
 * calling code to choose whichever is more preferable. The two functions also
 * behave differently when some of the provided options are not declared as
 * supported by the filter. In such a case, avfilter_init_str() will fail, but
 * this function will leave those extra options in the options AVDictionary and
 * continue as usual.
 */
//int avfilter_init_dict(AVFilterContext *ctx, AVDictionary **options);
type AVDictionary = libavutil.AVDictionary

var avfilterInitDict func(ctx *AVFilterContext, options **AVDictionary) ffcommon.FInt
var avfilterInitDictOnce sync.Once

func (ctx *AVFilterContext) AvfilterInitDict(options **AVDictionary) ffcommon.FInt {
	avfilterInitDictOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterInitDict, ffcommon.GetAvfilterDll(), "avfilter_init_dict")
	})

	return avfilterInitDict(ctx, options)
}

/**
 * Free a filter context. This will also remove the filter from its
 * filtergraph's list of filters.
 *
 * @param filter the filter to free
 */
//void avfilter_free(AVFilterContext *filter);

var avfilterFree func(filter *AVFilterContext)
var avfilterFreeOnce sync.Once

func (filter *AVFilterContext) AvfilterFree() {
	avfilterFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterFree, ffcommon.GetAvfilterDll(), "avfilter_free")
	})

	avfilterFree(filter)
}

/**
 * Insert a filter in the middle of an existing link.
 *
 * @param link the link into which the filter should be inserted
 * @param filt the filter to be inserted
 * @param filt_srcpad_idx the input pad on the filter to connect
 * @param filt_dstpad_idx the output pad on the filter to connect
 * @return     zero on success
 */
//int avfilter_insert_filter(AVFilterLink *link, AVFilterContext *filt,
//unsigned filt_srcpad_idx, unsigned filt_dstpad_idx);

var avfilterInsertFilter func(link *AVFilterLink, filt *AVFilterContext, filt_srcpad_idx, filt_dstpad_idx ffcommon.FUnsigned) ffcommon.FInt
var avfilterInsertFilterOnce sync.Once

func (link *AVFilterLink) AvfilterInsertFilter(filt *AVFilterContext, filt_srcpad_idx, filt_dstpad_idx ffcommon.FUnsigned) ffcommon.FInt {
	avfilterInsertFilterOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterInsertFilter, ffcommon.GetAvfilterDll(), "avfilter_insert_filter")
	})

	return avfilterInsertFilter(link, filt, filt_srcpad_idx, filt_dstpad_idx)
}

/**
 * @return AVClass for AVFilterContext.
 *
 * @see av_opt_find().
 */
//const AVClass *avfilter_get_class(void);
var avfilterGetClass func() *AVClass
var avfilterGetClassOnce sync.Once

func AvfilterGetClass() *AVClass {
	avfilterGetClassOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGetClass, ffcommon.GetAvfilterDll(), "avfilter_get_class")
	})

	return avfilterGetClass()
}

// typedef struct AVFilterGraphInternal AVFilterGraphInternal;
type AVFilterGraphInternal struct {
}

/**
 * A function pointer passed to the @ref AVFilterGraph.execute callback to be
 * executed multiple times, possibly in parallel.
 *
 * @param ctx the filter context the job belongs to
 * @param arg an opaque parameter passed through from @ref
 *            AVFilterGraph.execute
 * @param jobnr the index of the job being executed
 * @param nb_jobs the total number of jobs
 *
 * @return 0 on success, a negative AVERROR on error
 */
//typedef int (avfilter_action_func)(AVFilterContext *ctx, void *arg, int jobnr, int nb_jobs);
type AvfilterActionFunc = func(ctx *AVFilterContext, arg ffcommon.FVoidP, jobnr, nb_jobs ffcommon.FInt) uintptr

/**
 * A function executing multiple jobs, possibly in parallel.
 *
 * @param ctx the filter context to which the jobs belong
 * @param func the function to be called multiple times
 * @param arg the argument to be passed to func
 * @param ret a nb_jobs-sized array to be filled with return values from each
 *            invocation of func
 * @param nb_jobs the number of jobs to execute
 *
 * @return 0 on success, a negative AVERROR on error
 */
//typedef int (avfilter_execute_func)(AVFilterContext *ctx, avfilter_action_func *func,
//void *arg, int *ret, int nb_jobs);
type AvfilterExecuteFunc = func(ctx *AVFilterContext, func0 uintptr, arg ffcommon.FVoidP, ret *ffcommon.FInt, nb_jobs ffcommon.FInt) uintptr
type AVFilterGraph struct {
	AvClass   *AVClass
	Filters   **AVFilterContext
	NbFilters ffcommon.FUnsigned

	ScaleSwsOpts ffcommon.FCharPStruct ///< sws options to use for the auto-inserted scale filters
	//#if FF_API_LAVR_OPTS
	//	attribute_deprecated char *resample_lavr_opts;   ///< libavresample options to use for the auto-inserted resample filters
	ResampleLavrOpts ffcommon.FCharPStruct
	//#endif

	/**
	 * Type of multithreading allowed for filters in this graph. A combination
	 * of AVFILTER_THREAD_* flags.
	 *
	 * May be set by the caller at any point, the setting will apply to all
	 * filters initialized after that. The default is allowing everything.
	 *
	 * When a filter in this graph is initialized, this field is combined using
	 * bit AND with AVFilterContext.thread_type to get the final mask used for
	 * determining allowed threading types. I.e. a threading type needs to be
	 * set in both to be allowed.
	 */
	ThreadType ffcommon.FInt

	/**
	 * Maximum number of threads used by filters in this graph. May be set by
	 * the caller before adding any filters to the filtergraph. Zero (the
	 * default) means that the number of threads is determined automatically.
	 */
	NbThreads ffcommon.FInt

	/**
	 * Opaque object for libavfilter internal use.
	 */
	Internal *AVFilterGraphInternal

	/**
	 * Opaque user data. May be set by the caller to an arbitrary value, e.g. to
	 * be used from callbacks like @ref AVFilterGraph.execute.
	 * Libavfilter will not touch this field in any way.
	 */
	Opaque ffcommon.FVoidP

	/**
	 * This callback may be set by the caller immediately after allocating the
	 * graph and before adding any filters to it, to provide a custom
	 * multithreading implementation.
	 *
	 * If set, filters with slice threading capability will call this callback
	 * to execute multiple jobs in parallel.
	 *
	 * If this field is left unset, libavfilter will use its internal
	 * implementation, which may or may not be multithreaded depending on the
	 * platform and build options.
	 */
	//execute *avfilter_execute_func
	AvfilterExecuteFunc uintptr

	AresampleSwrOpts ffcommon.FCharPStruct ///< swr options to use for the auto-inserted aresample filters, Access ONLY through AVOptions

	/**
	 * Private fields
	 *
	 * The following fields are for internal use only.
	 * Their type, offset, number and semantic can change without notice.
	 */

	SinkLinks      **AVFilterLink
	SinkLinksCount ffcommon.FInt

	DisableAutoConvert ffcommon.FUnsigned
}

/**
 * Allocate a filter graph.
 *
 * @return the allocated filter graph on success or NULL.
 */
//AVFilterGraph *avfilter_graph_alloc(void);
var avfilterGraphAlloc func() *AVFilterGraph
var avfilterGraphAllocOnce sync.Once

func AvfilterGraphAlloc() *AVFilterGraph {
	avfilterGraphAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphAlloc, ffcommon.GetAvfilterDll(), "avfilter_graph_alloc")
	})

	return avfilterGraphAlloc()
}

/**
 * Create a new filter instance in a filter graph.
 *
 * @param graph graph in which the new filter will be used
 * @param filter the filter to create an instance of
 * @param name Name to give to the new instance (will be copied to
 *             AVFilterContext.name). This may be used by the caller to identify
 *             different filters, libavfilter itself assigns no semantics to
 *             this parameter. May be NULL.
 *
 * @return the context of the newly created filter instance (note that it is
 *         also retrievable directly through AVFilterGraph.filters or with
 *         avfilter_graph_get_filter()) on success or NULL on failure.
 */
//AVFilterContext *avfilter_graph_alloc_filter(AVFilterGraph *graph,
//const AVFilter *filter,
//const char *name);
var avfilterGraphAllocFilter func(graph *AVFilterGraph, filter *AVFilter, name ffcommon.FConstCharP) *AVFilterContext
var avfilterGraphAllocFilterOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphAllocFilter(filter *AVFilter, name ffcommon.FConstCharP) *AVFilterContext {
	avfilterGraphAllocFilterOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphAllocFilter, ffcommon.GetAvfilterDll(), "avfilter_graph_alloc_filter")
	})

	return avfilterGraphAllocFilter(graph, filter, name)
}

/**
 * Get a filter instance identified by instance name from graph.
 *
 * @param graph filter graph to search through.
 * @param name filter instance name (should be unique in the graph).
 * @return the pointer to the found filter instance or NULL if it
 * cannot be found.
 */
//AVFilterContext *avfilter_graph_get_filter(AVFilterGraph *graph, const char *name);
var avfilterGraphGetFilter func(graph *AVFilterGraph, name ffcommon.FConstCharP) *AVFilterContext
var avfilterGraphGetFilterOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphGetFilter(name ffcommon.FConstCharP) *AVFilterContext {
	avfilterGraphGetFilterOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphGetFilter, ffcommon.GetAvfilterDll(), "avfilter_graph_get_filter")
	})

	return avfilterGraphGetFilter(graph, name)
}

/**
 * Create and add a filter instance into an existing graph.
 * The filter instance is created from the filter filt and inited
 * with the parameter args. opaque is currently ignored.
 *
 * In case of success put in *filt_ctx the pointer to the created
 * filter instance, otherwise set *filt_ctx to NULL.
 *
 * @param name the instance name to give to the created filter instance
 * @param graph_ctx the filter graph
 * @return a negative AVERROR error code in case of failure, a non
 * negative value otherwise
 */
//int avfilter_graph_create_filter(AVFilterContext **filt_ctx, const AVFilter *filt,
//const char *name, const char *args, void *opaque,
//AVFilterGraph *graph_ctx);
var avfilterGraphCreateFilter func(filt_ctx **AVFilterContext, filt *AVFilter, name, args ffcommon.FConstCharP, opaque ffcommon.FVoidP, graph_ctx *AVFilterGraph) ffcommon.FInt
var avfilterGraphCreateFilterOnce sync.Once

func AvfilterGraphCreateFilter(filt_ctx **AVFilterContext, filt *AVFilter, name, args ffcommon.FConstCharP, opaque ffcommon.FVoidP, graph_ctx *AVFilterGraph) ffcommon.FInt {
	avfilterGraphCreateFilterOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphCreateFilter, ffcommon.GetAvfilterDll(), "avfilter_graph_create_filter")
	})

	return avfilterGraphCreateFilter(filt_ctx, filt, name, args, opaque, graph_ctx)
}

/**
 * Enable or disable automatic format conversion inside the graph.
 *
 * Note that format conversion can still happen inside explicitly inserted
 * scale and aresample filters.
 *
 * @param flags  any of the AVFILTER_AUTO_CONVERT_* constants
 */
//void avfilter_graph_set_auto_convert(AVFilterGraph *graph, unsigned flags);
var avfilterGraphSetAutoConvert func(graph *AVFilterGraph, flags ffcommon.FUnsigned)
var avfilterGraphSetAutoConvertOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphSetAutoConvert(flags ffcommon.FUnsigned) {
	avfilterGraphSetAutoConvertOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphSetAutoConvert, ffcommon.GetAvfilterDll(), "avfilter_graph_set_auto_convert")
	})

	avfilterGraphSetAutoConvert(graph, flags)
}

const (
	AVFILTER_AUTO_CONVERT_ALL  = 0  /**< all automatic conversions enabled */
	AVFILTER_AUTO_CONVERT_NONE = -1 /**< all automatic conversions disabled */
)

/**
 * Check validity and configure all the links and formats in the graph.
 *
 * @param graphctx the filter graph
 * @param log_ctx context used for logging
 * @return >= 0 in case of success, a negative AVERROR code otherwise
 */
//int avfilter_graph_config(AVFilterGraph *graphctx, void *log_ctx);
var avfilterGraphConfig func(graphctx *AVFilterGraph, log_ctx ffcommon.FVoidP) ffcommon.FInt
var avfilterGraphConfigOnce sync.Once

func (graphctx *AVFilterGraph) AvfilterGraphConfig(log_ctx ffcommon.FVoidP) ffcommon.FInt {
	avfilterGraphConfigOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphConfig, ffcommon.GetAvfilterDll(), "avfilter_graph_config")
	})

	return avfilterGraphConfig(graphctx, log_ctx)
}

/**
 * Free a graph, destroy its links, and set *graph to NULL.
 * If *graph is NULL, do nothing.
 */
//void avfilter_graph_free(AVFilterGraph **graph);
var avfilterGraphFree func(graphctx **AVFilterGraph)
var avfilterGraphFreeOnce sync.Once

func AvfilterGraphFree(graphctx **AVFilterGraph) {
	avfilterGraphFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphFree, ffcommon.GetAvfilterDll(), "avfilter_graph_free")
	})

	avfilterGraphFree(graphctx)
}

/**
 * A linked-list of the inputs/outputs of the filter chain.
 *
 * This is mainly useful for avfilter_graph_parse() / avfilter_graph_parse2(),
 * where it is used to communicate open (unlinked) inputs and outputs from and
 * to the caller.
 * This struct specifies, per each not connected pad contained in the graph, the
 * filter context and the pad index required for establishing a link.
 */
type AVFilterInOut struct {

	/** unique name for this input/output in the list */
	Name ffcommon.FCharPStruct

	/** filter context associated to this input/output */
	FilterCtx *AVFilterContext

	/** index of the filt_ctx pad to use for linking */
	PadIdx ffcommon.FInt

	/** next input/input in the list, NULL if this is the last */
	Next *AVFilterInOut
}

/**
 * Allocate a single AVFilterInOut entry.
 * Must be freed with avfilter_inout_free().
 * @return allocated AVFilterInOut on success, NULL on failure.
 */
//AVFilterInOut *avfilter_inout_alloc(void);
var avfilterInoutAlloc func() *AVFilterInOut
var avfilterInoutAllocOnce sync.Once

func AvfilterInoutAlloc() *AVFilterInOut {
	avfilterInoutAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterInoutAlloc, ffcommon.GetAvfilterDll(), "avfilter_inout_alloc")
	})

	return avfilterInoutAlloc()
}

/**
 * Free the supplied list of AVFilterInOut and set *inout to NULL.
 * If *inout is NULL, do nothing.
 */
//void avfilter_inout_free(AVFilterInOut **inout);
var avfilterInoutFree func(inout **AVFilterInOut)
var avfilterInoutFreeOnce sync.Once

func AvfilterInoutFree(inout **AVFilterInOut) {
	avfilterInoutFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterInoutFree, ffcommon.GetAvfilterDll(), "avfilter_inout_free")
	})

	avfilterInoutFree(inout)
}

/**
 * Add a graph described by a string to a graph.
 *
 * @note The caller must provide the lists of inputs and outputs,
 * which therefore must be known before calling the function.
 *
 * @note The inputs parameter describes inputs of the already existing
 * part of the graph; i.e. from the point of view of the newly created
 * part, they are outputs. Similarly the outputs parameter describes
 * outputs of the already existing filters, which are provided as
 * inputs to the parsed filters.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  linked list to the inputs of the graph
 * @param outputs linked list to the outputs of the graph
 * @return zero on success, a negative AVERROR code on error
 */
//int avfilter_graph_parse(AVFilterGraph *graph, const char *filters,
//AVFilterInOut *inputs, AVFilterInOut *outputs,
//void *log_ctx);
var avfilterGraphParse func(graph *AVFilterGraph, filters ffcommon.FConstCharP, inputs, outputs *AVFilterInOut, log_ctx ffcommon.FVoidP) ffcommon.FInt
var avfilterGraphParseOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphParse(filters ffcommon.FConstCharP, inputs, outputs *AVFilterInOut, log_ctx ffcommon.FVoidP) ffcommon.FInt {
	avfilterGraphParseOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphParse, ffcommon.GetAvfilterDll(), "avfilter_graph_parse")
	})

	return avfilterGraphParse(graph, filters, inputs, outputs, log_ctx)
}

/**
 * Add a graph described by a string to a graph.
 *
 * In the graph filters description, if the input label of the first
 * filter is not specified, "in" is assumed; if the output label of
 * the last filter is not specified, "out" is assumed.
 *
 * @param graph   the filter graph where to link the parsed graph context
 * @param filters string to be parsed
 * @param inputs  pointer to a linked list to the inputs of the graph, may be NULL.
 *                If non-NULL, *inputs is updated to contain the list of open inputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @param outputs pointer to a linked list to the outputs of the graph, may be NULL.
 *                If non-NULL, *outputs is updated to contain the list of open outputs
 *                after the parsing, should be freed with avfilter_inout_free().
 * @return non negative on success, a negative AVERROR code on error
 */
//int avfilter_graph_parse_ptr(AVFilterGraph *graph, const char *filters,
//AVFilterInOut **inputs, AVFilterInOut **outputs,
//void *log_ctx);
var avfilterGraphParsePtr func(graph *AVFilterGraph, filters ffcommon.FConstCharP, inputs, outputs **AVFilterInOut, log_ctx ffcommon.FVoidP) ffcommon.FInt
var avfilterGraphParsePtrOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphParsePtr(filters ffcommon.FConstCharP, inputs, outputs **AVFilterInOut, log_ctx ffcommon.FVoidP) ffcommon.FInt {
	avfilterGraphParsePtrOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphParsePtr, ffcommon.GetAvfilterDll(), "avfilter_graph_parse_ptr")
	})

	return avfilterGraphParsePtr(graph, filters, inputs, outputs, log_ctx)
}

/**
 * Add a graph described by a string to a graph.
 *
 * @param[in]  graph   the filter graph where to link the parsed graph context
 * @param[in]  filters string to be parsed
 * @param[out] inputs  a linked list of all free (unlinked) inputs of the
 *                     parsed graph will be returned here. It is to be freed
 *                     by the caller using avfilter_inout_free().
 * @param[out] outputs a linked list of all free (unlinked) outputs of the
 *                     parsed graph will be returned here. It is to be freed by the
 *                     caller using avfilter_inout_free().
 * @return zero on success, a negative AVERROR code on error
 *
 * @note This function returns the inputs and outputs that are left
 * unlinked after parsing the graph and the caller then deals with
 * them.
 * @note This function makes no reference whatsoever to already
 * existing parts of the graph and the inputs parameter will on return
 * contain inputs of the newly parsed part of the graph.  Analogously
 * the outputs parameter will contain outputs of the newly created
 * filters.
 */
//int avfilter_graph_parse2(AVFilterGraph *graph, const char *filters,
//AVFilterInOut **inputs,
//AVFilterInOut **outputs);
var avfilterGraphParse2 func(graph *AVFilterGraph, filters ffcommon.FConstCharP, inputs, outputs **AVFilterInOut) ffcommon.FInt
var avfilterGraphParse2Once sync.Once

func (graph *AVFilterGraph) AvfilterGraphParse2(filters ffcommon.FConstCharP, inputs, outputs **AVFilterInOut) ffcommon.FInt {
	avfilterGraphParse2Once.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphParse2, ffcommon.GetAvfilterDll(), "avfilter_graph_parse2")
	})

	return avfilterGraphParse2(graph, filters, inputs, outputs)
}

/**
 * Send a command to one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to send, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param res    a buffer with size res_size where the filter(s) can return a response.
 *
 * @returns >=0 on success otherwise an error code.
 *              AVERROR(ENOSYS) on unsupported commands
 */
//int avfilter_graph_send_command(AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, char *res, int res_len, int flags);
var avfilterGraphSendCommand func(graph *AVFilterGraph, target, cmd, arg, res0 ffcommon.FConstCharP, res_len, flags ffcommon.FInt) ffcommon.FInt
var avfilterGraphSendCommandOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphSendCommand(target, cmd, arg, res0 ffcommon.FConstCharP, res_len, flags ffcommon.FInt) ffcommon.FInt {
	avfilterGraphSendCommandOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphSendCommand, ffcommon.GetAvfilterDll(), "avfilter_graph_send_command")
	})

	return avfilterGraphSendCommand(graph, target, cmd, arg, res0, res_len, flags)
}

/**
 * Queue a command for one or more filter instances.
 *
 * @param graph  the filter graph
 * @param target the filter(s) to which the command should be sent
 *               "all" sends to all filters
 *               otherwise it can be a filter or filter instance name
 *               which will send the command to all matching filters.
 * @param cmd    the command to sent, for handling simplicity all commands must be alphanumeric only
 * @param arg    the argument for the command
 * @param ts     time at which the command should be sent to the filter
 *
 * @note As this executes commands after this function returns, no return code
 *       from the filter is provided, also AVFILTER_CMD_FLAG_ONE is not supported.
 */
//int avfilter_graph_queue_command(AVFilterGraph *graph, const char *target, const char *cmd, const char *arg, int flags, double ts);
var avfilterGraphQueueCommand func(graph *AVFilterGraph, target, cmd, arg ffcommon.FConstCharP, flags ffcommon.FInt, ts ffcommon.FDouble) ffcommon.FInt
var avfilterGraphQueueCommandOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphQueueCommand(target, cmd, arg ffcommon.FConstCharP, flags ffcommon.FInt, ts ffcommon.FDouble) ffcommon.FInt {
	avfilterGraphQueueCommandOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphQueueCommand, ffcommon.GetAvfilterDll(), "avfilter_graph_queue_command")
	})

	return avfilterGraphQueueCommand(graph, target, cmd, arg, flags, ts)
}

/**
 * Dump a graph into a human-readable string representation.
 *
 * @param graph    the graph to dump
 * @param options  formatting options; currently ignored
 * @return  a string, or NULL in case of memory allocation failure;
 *          the string must be freed using av_free
 */
//char *avfilter_graph_dump(AVFilterGraph *graph, const char *options);
var avfilterGraphDump func(graph *AVFilterGraph, options ffcommon.FConstCharP) ffcommon.FCharP
var avfilterGraphDumpOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphDump(options ffcommon.FConstCharP) ffcommon.FCharP {
	avfilterGraphDumpOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphDump, ffcommon.GetAvfilterDll(), "avfilter_graph_dump")
	})

	return avfilterGraphDump(graph, options)
}

/**
 * Request a frame on the oldest sink link.
 *
 * If the request returns AVERROR_EOF, try the next.
 *
 * Note that this function is not meant to be the sole scheduling mechanism
 * of a filtergraph, only a convenience function to help drain a filtergraph
 * in a balanced way under normal circumstances.
 *
 * Also note that AVERROR_EOF does not mean that frames did not arrive on
 * some of the sinks during the process.
 * When there are multiple sink links, in case the requested link
 * returns an EOF, this may cause a filter to flush pending frames
 * which are sent to another sink link, although unrequested.
 *
 * @return  the return value of ff_request_frame(),
 *          or AVERROR_EOF if all links returned AVERROR_EOF
 */
//int avfilter_graph_request_oldest(AVFilterGraph *graph);
var avfilterGraphRequestOldest func(graph *AVFilterGraph) ffcommon.FInt
var avfilterGraphRequestOldestOnce sync.Once

func (graph *AVFilterGraph) AvfilterGraphRequestOldest() ffcommon.FInt {
	avfilterGraphRequestOldestOnce.Do(func() {
		purego.RegisterLibFunc(&avfilterGraphRequestOldest, ffcommon.GetAvfilterDll(), "avfilter_graph_request_oldest")
	})

	return avfilterGraphRequestOldest(graph)
}

/**
 * @}
 */

//#endif /* AVFILTER_AVFILTER_H */
