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

// NewFormatContext creates a new FormatContext structure.
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

// FindStreamInfo probes the FormatContext for stream information
func (i *FormatContext) FindStreamInfo() error {
	ret := C.avformat_find_stream_info(i.fctx, nil)
	if ret < 0 {
		return Error{int(ret)}
	}

	return nil
}

// FindProgram finds the next program in stream. If last != nil, it will return
// the next program after last.
func (i *FormatContext) FindProgram(last *Program, stream int) *Program {
	c_program := C.av_find_program_from_stream(i.fctx, last.p, C.int(stream))
	if c_program == nil {
		return nil
	}

	return &Program{c_program}
}

// FindInputFormat finds an InputFormat based on its short name.
func FindInputFormat(short_name string) *InputFormat {
	c_short_name := C.CString(short_name)
	defer C.free(unsafe.Pointer(c_short_name))

	fmt := C.av_find_input_format(c_short_name)
	if fmt == nil {
		return nil
	}

	return &InputFormat{fmt}
}

// ProbeInputFormat attempts to guess the file format from ProbeData.
func ProbeInputFormat(pd *ProbeData, is_opened bool) *InputFormat {
	var opened C.int
	if is_opened {
		opened = 1
	} else {
		opened = 0
	}

	fmt := C.av_probe_input_format(pd.pd, opened)
	if fmt == nil {
		return nil
	}

	return &InputFormat{fmt}
}

// ProbeInputFormat2 attempts to guess the file format from ProbeData.
// score_min is the minimum score required to succeed.  Returns non-nil
// InputFormat on success, as well as the detection score.
func ProbeInputFormat2(pd *ProbeData, is_opened bool, score_min int) (*InputFormat, int) {
	var opened C.int
	if is_opened {
		opened = 1
	} else {
		opened = 0
	}

	score := C.int(score_min)

	fmt := C.av_probe_input_format2(pd.pd, opened, &score)
	if fmt == nil {
		return nil, int(score)
	}

	return &InputFormat{fmt}, int(score)
}

// ProbeInputFormat3 attempts to guess the file format from ProbeData.  Returns
// non-nil InputFormat on success, as well as the detection score.
func ProbeInputFormat3(pd *ProbeData, is_opened bool) (*InputFormat, int) {
	var opened, score C.int
	if is_opened {
		opened = 1
	} else {
		opened = 0
	}

	fmt := C.av_probe_input_format2(pd.pd, opened, &score)
	if fmt == nil {
		return nil, int(score)
	}

	return &InputFormat{fmt}, int(score)
}

// ProbeInputBuffer probes an IOContext to find the most likely input format.
func ProbeInputBuffer(pb *IOContext, offset uint, max_probe_size uint) (*InputFormat, error) {
	var fmt *C.AVInputFormat

	ret := C.av_probe_input_buffer(pb.io, &fmt, nil, nil, C.uint(offset),
		C.uint(max_probe_size))
	if ret < 0 {
		return nil, Error{int(ret)}
	}

	return &InputFormat{fmt}, nil
}

// OpenInput opens filename using opts if opts != nil. The input format will be
// forced to ifmt if non-nil.
func OpenInput(filename string, ifmt *InputFormat, opts *Dictionary) (*FormatContext, error) {
	var fmt *C.AVFormatContext
	c_filename := C.CString(filename)

	if opts == nil {
		opts = &Dictionary{}
	}
	if ifmt == nil {
		ifmt = &InputFormat{}
	}

	ret := C.avformat_open_input(&fmt, c_filename, ifmt.fmt, &opts.dict)
	if ret < 0 {
		return nil, Error{int(ret)}
	}

	return &FormatContext{
		fctx:    fmt,
		c_fname: c_filename,
	}, nil
}
