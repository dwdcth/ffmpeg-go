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
 * modify it under the terms of the GNU Lesser General Public License
 * as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * FFmpeg is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with FFmpeg; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA
 */

//#ifndef AVUTIL_THREADMESSAGE_H
//#define AVUTIL_THREADMESSAGE_H

// typedef struct AVThreadMessageQueue AVThreadMessageQueue;
type AVThreadMessageQueue struct {
}
type AVThreadMessageFlags int32

const (

	/**
	 * Perform non-blocking operation.
	 * If this flag is set, send and recv operations are non-blocking and
	 * return AVERROR(EAGAIN) immediately if they can not proceed.
	 */
	AV_THREAD_MESSAGE_NONBLOCK = 1
)

/**
 * Allocate a new message queue.
 *
 * @param mq      pointer to the message queue
 * @param nelem   maximum number of elements in the queue
 * @param elsize  size of each element in the queue
 * @return  >=0 for success; <0 for error, in particular AVERROR(ENOSYS) if
 *          lavu was built without thread support
 */
//int av_thread_message_queue_alloc(AVThreadMessageQueue **mq,
//unsigned nelem,
//unsigned elsize);
// purego func
var avThreadMessageQueueAlloc func(mq **AVThreadMessageQueue, nelem, elsize ffcommon.FUnsigned) ffcommon.FInt
var avThreadMessageQueueAllocOnce sync.Once

func AvThreadMessageQueueAlloc(mq **AVThreadMessageQueue, nelem, elsize ffcommon.FUnsigned) (res ffcommon.FInt) {
	avThreadMessageQueueAllocOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueAlloc, ffcommon.GetAvutilDll(), "av_thread_message_queue_alloc")
	})
	return avThreadMessageQueueAlloc(mq, nelem, elsize)
}

/**
 * Free a message queue.
 *
 * The message queue must no longer be in use by another thread.
 */
//void av_thread_message_queue_free(AVThreadMessageQueue **mq);
// purego func
var avThreadMessageQueueFree func(mq **AVThreadMessageQueue)
var avThreadMessageQueueFreeOnce sync.Once

func AvThreadMessageQueueFree(mq **AVThreadMessageQueue) {
	avThreadMessageQueueFreeOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueFree, ffcommon.GetAvutilDll(), "av_thread_message_queue_free")
	})
	avThreadMessageQueueFree(mq)
}

/**
 * Send a message on the queue.
 */
//int av_thread_message_queue_send(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
// purego struct method
var avThreadMessageQueueSend func(mq *AVThreadMessageQueue, msg ffcommon.FVoidP, flags ffcommon.FUnsigned) ffcommon.FInt
var avThreadMessageQueueSendOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueSend(msg ffcommon.FVoidP, flags ffcommon.FUnsigned) (res ffcommon.FInt) {
	avThreadMessageQueueSendOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueSend, ffcommon.GetAvutilDll(), "av_thread_message_queue_send")
	})
	return avThreadMessageQueueSend(mq, msg, flags)
}

/**
 * Receive a message from the queue.
 */
//int av_thread_message_queue_recv(AVThreadMessageQueue *mq,
//void *msg,
//unsigned flags);
// purego struct method
var avThreadMessageQueueRecv func(mq *AVThreadMessageQueue, msg ffcommon.FVoidP, flags ffcommon.FUnsigned) ffcommon.FInt
var avThreadMessageQueueRecvOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueRecv(msg ffcommon.FVoidP, flags ffcommon.FUnsigned) (res ffcommon.FInt) {
	avThreadMessageQueueRecvOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueRecv, ffcommon.GetAvutilDll(), "av_thread_message_queue_recv")
	})
	return avThreadMessageQueueRecv(mq, msg, flags)
}

/**
 * Set the sending error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_send() will
 * return it immediately. Conventional values, such as AVERROR_EOF or
 * AVERROR(EAGAIN), can be used to cause the sending thread to stop or
 * suspend its operation.
 */
//void av_thread_message_queue_set_err_send(AVThreadMessageQueue *mq,
//int err);
// purego struct method
var avThreadMessageQueueSetErrSend func(mq *AVThreadMessageQueue, err ffcommon.FInt)
var avThreadMessageQueueSetErrSendOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrSend(err ffcommon.FInt) {
	avThreadMessageQueueSetErrSendOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueSetErrSend, ffcommon.GetAvutilDll(), "av_thread_message_queue_set_err_send")
	})
	avThreadMessageQueueSetErrSend(mq, err)
}

/**
 * Set the receiving error code.
 *
 * If the error code is set to non-zero, av_thread_message_queue_recv() will
 * return it immediately when there are no longer available messages.
 * Conventional values, such as AVERROR_EOF or AVERROR(EAGAIN), can be used
 * to cause the receiving thread to stop or suspend its operation.
 */
//void av_thread_message_queue_set_err_recv(AVThreadMessageQueue *mq,
//int err);
// purego struct method
var avThreadMessageQueueSetErrRecv func(mq *AVThreadMessageQueue, err ffcommon.FInt)
var avThreadMessageQueueSetErrRecvOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetErrRecv(err ffcommon.FInt) {
	avThreadMessageQueueSetErrRecvOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueSetErrRecv, ffcommon.GetAvutilDll(), "av_thread_message_queue_set_err_recv")
	})
	avThreadMessageQueueSetErrRecv(mq, err)
}

/**
 * Set the optional free message callback function which will be called if an
 * operation is removing messages from the queue.
 */
//void av_thread_message_queue_set_free_func(AVThreadMessageQueue *mq,
//void (*free_func)(void *msg));
var avThreadMessageQueueSetFreeFunc func(mq *AVThreadMessageQueue, freeFunc func(msg ffcommon.FVoidP) uintptr)
var avThreadMessageQueueSetFreeFuncOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueSetFreeFunc(freeFunc func(msg ffcommon.FVoidP) uintptr) {
	avThreadMessageQueueSetFreeFuncOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueSetFreeFunc, ffcommon.GetAvutilDll(), "av_thread_message_queue_set_free_func")
	})
	avThreadMessageQueueSetFreeFunc(mq, freeFunc)
}

/**
 * Return the current number of messages in the queue.
 *
 * @return the current number of messages or AVERROR(ENOSYS) if lavu was built
 *         without thread support
 */
//int av_thread_message_queue_nb_elems(AVThreadMessageQueue *mq);
// purego struct method
var avThreadMessageQueueNbElems func(mq *AVThreadMessageQueue) ffcommon.FInt
var avThreadMessageQueueNbElemsOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageQueueNbElems() (res ffcommon.FInt) {
	avThreadMessageQueueNbElemsOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageQueueNbElems, ffcommon.GetAvutilDll(), "av_thread_message_queue_nb_elems")
	})
	return avThreadMessageQueueNbElems(mq)
}

/**
 * Flush the message queue
 *
 * This function is mostly equivalent to reading and free-ing every message
 * except that it will be done in a single operation (no lock/unlock between
 * reads).
 */
//void av_thread_message_flush(AVThreadMessageQueue *mq);
// purego struct method
// purego struct method
var avThreadMessageFlush func(mq *AVThreadMessageQueue)
var avThreadMessageFlushOnce sync.Once

func (mq *AVThreadMessageQueue) AvThreadMessageFlush() {
	avThreadMessageFlushOnce.Do(func() {
		purego.RegisterLibFunc(&avThreadMessageFlush, ffcommon.GetAvutilDll(), "av_thread_message_flush")
	})
	avThreadMessageFlush(mq)
}

//#endif /* AVUTIL_THREADMESSAGE_H */
