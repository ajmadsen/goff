package av

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>

#include <stdint.h>
#include <stdlib.h>

#include "callback.h"
*/
import "C"
import (
	"errors"
	"io"
	"reflect"
	"unsafe"
)

const bufsiz = 32 * 1024

// IO flags
const (
	FLAG_READ      = 1
	FLAG_WRITE     = 2
	FLAG_READWRITE = FLAG_READ | FLAG_WRITE
)

// IO is a general purpose communication interface between Go and the
// underlying library. It implements ReaderWriterSeeker.
type IO struct {
	ctx    *C.AVIOContext
	stream interface{}
	name   string
}

// Name returns the name of the IO stream, or the name of the underlying file.
func (i *IO) Name() string {
	return i.name
}

// Close closes the IO stream and frees any resources associated with it.
func (i *IO) Close() {
	if i.ctx.av_class != nil {
		C.avio_close(i.ctx)
	} else {
		if i.ctx.buffer != nil {
			C.av_freep(unsafe.Pointer(&i.ctx.buffer))
		}
		if i.ctx != nil {
			C.av_freep(unsafe.Pointer(&i.ctx))
		}
	}
}

// Read reads 0 < n < len(b) bytes from the underlying stream, returning any
// errors encountered.
func (i *IO) Read(b []byte) (int, error) {
	ret := C.avio_read(i.ctx, (*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
	if ret <= 0 {
		if C.avio_feof(i.ctx) != 0 {
			return 0, io.EOF
		}
		return 0, errors.New(averror(ret))
	}
	return int(ret), nil
}

// Write writes len(b) bytes to the underlying stream. Should not return any
// errors.
func (i *IO) Write(b []byte) (int, error) {
	C.avio_write(i.ctx, (*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
	return len(b), nil
}

// Seek sets the offset in the underlying stream.
func (i *IO) Seek(offset int64, whence int) (int64, error) {
	ret := int64(C.avio_seek(i.ctx, C.int64_t(offset), C.int(whence)))
	if ret < 0 {
		return ret, errors.New(averror(C.int(ret)))
	}
	return ret, nil
}

// NewIO creates a new IO context from a ReadSeeker.  stream should implement
// some set of {Reader, Writer, Seeker}. If writable == true and stream
// implements Writer, sets the writer flag internally.
func NewIO(stream interface{}, name string, writable bool) (*IO, error) {
	ctx := &IO{}

	var readFcn, writeFcn, seekFcn unsafe.Pointer
	var c_writable C.int
	if _, ok := stream.(io.Reader); ok {
		readFcn = C.cgo_read_packet_wrap
	}
	if writable {
		c_writable = 1
		if _, ok := stream.(io.Writer); ok {
			writeFcn = C.cgo_write_packet_wrap
		}
	}
	if _, ok := stream.(io.Seeker); ok {
		seekFcn = C.cgo_seek_wrap
	}

	if readFcn == nil && writeFcn == nil && seekFcn == nil {
		return nil, errors.New("stream does not implement anything useful")
	}

	buf := (*C.uchar)(C.av_malloc(bufsiz))
	if buf == nil {
		return nil, errors.New("out of memory")
	}

	ctx.name = name

	ctx.stream = stream
	ctx.ctx = C.avio_alloc_context(buf,
		bufsiz,
		c_writable,
		unsafe.Pointer(ctx),
		(*[0]byte)(readFcn),
		(*[0]byte)(writeFcn),
		(*[0]byte)(seekFcn))
	if ctx.ctx == nil {
		C.av_freep(unsafe.Pointer(&buf))
		return nil, errors.New("could not alloc AVIOContext")
	}

	return ctx, nil
}

// OpenURL opens an IO with the underlying stream opened by the underlying
// library.
func OpenURL(url string, flags int) (*IO, error) {
	ctx := &IO{}
	ctx.name = url
	ctx.stream = nil

	c_url := C.CString(url)
	defer C.free(unsafe.Pointer(c_url))

	ret := C.avio_open(&ctx.ctx, c_url, C.int(flags))
	if ret < 0 {
		return nil, errors.New(averror(ret))
	}

	return ctx, nil
}

func makeByteSlice(buf *C.uint8_t, buf_siz C.int) []byte {
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(buf)),
		Len:  int(buf_siz),
		Cap:  int(buf_siz),
	}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

//export ioReadPacket
func ioReadPacket(opaque unsafe.Pointer, buf *C.uint8_t, buf_siz C.int) C.int {
	ctx := (*IO)(opaque)
	p := makeByteSlice(buf, buf_siz)
	ret, err := ctx.stream.(io.Reader).Read(p)
	if ret > 0 {
		return C.int(ret)
	}
	if err != io.EOF {
		return -1
	}
	return 0
}

//export ioWritePacket
func ioWritePacket(opaque unsafe.Pointer, buf *C.uint8_t, buf_siz C.int) C.int {
	ctx := (*IO)(opaque)
	p := makeByteSlice(buf, buf_siz)
	ret, err := ctx.stream.(io.Writer).Write(p)
	if ret > 0 {
		return C.int(ret)
	}
	if err != nil {
		return -1
	}
	return 0
}

//export ioSeek
func ioSeek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	ctx := (*IO)(opaque)
	ret, err := ctx.stream.(io.Seeker).Seek(int64(offset), int(whence))
	if err != nil {
		return -1
	}
	return C.int64_t(ret)
}
