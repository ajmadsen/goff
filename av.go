package av

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"
)

type Demuxer interface {
	Close()
	Dump(n int)
	NStreams() int
	Stream(idx int) Stream
}

type Stream interface {
	IsOpen() bool
	Index() int
}

type fmtctx struct {
	fctx   *C.AVFormatContext
	c_name *C.char
	ioctx  *avio
}

type stream struct {
	s *C.AVStream
}

func (f *fmtctx) Close() {
	if f.c_name != nil {
		C.free(unsafe.Pointer(f.c_name))
		f.c_name = nil
	}
	if f.fctx != nil {
		C.avformat_close_input(&f.fctx)
	}
	if f.ioctx != nil {
		f.ioctx.Close()
		f.ioctx = nil
	}
}

func (f *fmtctx) Dump(n int) {
	C.av_dump_format(f.fctx, C.int(n), f.c_name, 0)
}

func (f *fmtctx) streams() []*C.AVStream {
	nstr := f.NStreams()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.fctx.streams)),
		Len:  nstr,
		Cap:  nstr,
	}
	return *(*[]*C.AVStream)(unsafe.Pointer(&hdr))
}

func (f *fmtctx) NStreams() int {
	return int(f.fctx.nb_streams)
}

func (f *fmtctx) Stream(idx int) Stream {
	if idx > f.NStreams() || idx < 0 {
		return nil
	}

	return &stream{f.streams()[idx]}
}

func (s *stream) Index() int {
	return int(s.s.index)
}

func (s *stream) IsOpen() bool {
	return (s.s.codec != nil && s.s.codec.codec != nil)
}

func init() {
	C.av_register_all()
	C.avformat_network_init()
}

func averror(ret C.int) string {
	var str [256]C.char

	rv := C.av_strerror(ret, &str[0], C.size_t(len(str)))
	if rv != 0 {
		return "could not get error string"
	}

	return C.GoStringN(&str[0], C.int(len(str)))
}

func OpenDemuxer(r IOReader, name string) (Demuxer, error) {
	ioctx, ok := r.(*avio)
	if !ok {
		tmpctx, err := NewIOReader(r)
		if err != nil {
			return nil, err
		}
		ioctx = tmpctx.(*avio)
	}

	fmt := C.avformat_alloc_context()
	if fmt == nil {
		return nil, errors.New("could not alloc format context")
	}

	fmt.pb = ioctx.ctx

	ret := C.avformat_open_input(&fmt, nil, nil, nil)
	if ret < 0 {
		C.avformat_free_context(fmt)
		return nil, errors.New(averror(ret))
	}

	ret = C.avformat_find_stream_info(fmt, nil)
	if ret < 0 {
		C.avformat_free_context(fmt)
		return nil, errors.New(averror(ret))
	}

	return &fmtctx{
		fmt,
		C.CString(name),
		ioctx,
	}, nil
}
