package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

// FormatContext represents an AVFormatContext struct from the C API. Used for
// Muxing/Demuxing.
type FormatContext struct {
	fctx    *C.AVFormatContext
	c_fname *C.char
}

// InputFormat is an opaque representation of an AVInputFormat struct from the
// C API.
type InputFormat struct {
	fmt *C.AVInputFormat
}

// ProbeData is an opaque representation of an AVProbeData struct from the C
// API.
type ProbeData struct {
	pd *C.AVProbeData
}

// Program is an opaque representation of an AVProgram struct from the C API.
type Program struct {
	p *C.AVProgram
}

// NewFormatContext creates a new FormatContext structure
func NewFormatContext() *FormatContext {
	i := &FormatContext{}
	runtime.SetFinalizer(i, func(i *FormatContext) {
		i.Close()
	})
	return i
}

// Close closes the FormatContext and frees any data associated with it.
func (f *FormatContext) Close() {
}

// streams returns a slice containing the streams in an FormatContext. The
// slice is backed by a C array, and any attempts to modify the slice may have
// undesired consequences.
func (i *FormatContext) streams() []*C.AVStream {
	nstr := i.NStreams()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(i.fctx.streams)),
		Len:  nstr,
		Cap:  nstr,
	}
	return *(*[]*C.AVStream)(unsafe.Pointer(&hdr))
}

// Stream returns the idxth stream of the FormatContext. Panics if idx >
// NStreams().
func (i *FormatContext) Stream(idx int) Stream {
	if idx > i.NStreams() {
		panic(fmt.Sprintf("stream %d is out of bounds"))
	}
	return Stream{i.streams()[idx]}
}

// NStreams returns the number of streams in a FormatContext.
func (i *FormatContext) NStreams() int {
	return int(i.fctx.nb_streams)
}
