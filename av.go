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
	"unsafe"
)

type Demuxer interface {
	Close()
}

type fmtctx struct {
	fctx   *C.AVFormatContext
	c_name *C.char
}

func (f *fmtctx) Close() {
	if f.c_name != nil {
		C.free(unsafe.Pointer(f.c_name))
	}
	if f.fctx != nil {
		C.avformat_close_input(&f.fctx)
	}
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

func OpenReader(ioctx *IO) (Demuxer, error) {
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

	return &fmtctx{
		fmt,
		C.CString(ioctx.Name()),
	}, nil
}
