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

const bufsiz = 4 * 1024

// IO is a general purpose communication interface between Go and the
// underlying library. It implements ReaderWriterSeeker.
type avio struct {
	ctx    *C.AVIOContext
	stream interface{}
}

type IOReader interface {
	io.Reader
	io.Seeker
	io.Closer
}

type IOWriter interface {
	io.Writer
	io.Seeker
	io.Closer
}

// Close closes the IO stream and frees any resources associated with it.
func (i *avio) Close() error {
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
	return nil
}

// Read reads 0 < n < len(b) bytes from the underlying stream, returning any
// errors encountered.
func (i *avio) Read(b []byte) (int, error) {
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
func (i *avio) Write(b []byte) (int, error) {
	C.avio_write(i.ctx, (*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
	return len(b), nil
}

// Seek sets the offset in the underlying stream.
func (i *avio) Seek(offset int64, whence int) (int64, error) {
	ret := int64(C.avio_seek(i.ctx, C.int64_t(offset), C.int(whence)))
	if ret < 0 {
		return ret, errors.New(averror(C.int(ret)))
	}
	return ret, nil
}

func newIO(s interface{}, write_flag C.int, read_packet, write_packet, seek unsafe.Pointer) (*avio, error) {
	ctx := &avio{stream: s}

	buffer := (*C.uchar)(C.av_mallocz(bufsiz))
	if buffer == nil {
		return nil, errors.New("out of memory")
	}

	ctx.ctx = C.avio_alloc_context(
		buffer,
		bufsiz,
		write_flag,
		unsafe.Pointer(ctx),
		(*[0]byte)(read_packet),
		(*[0]byte)(write_packet),
		(*[0]byte)(seek),
	)
	if ctx.ctx == nil {
		C.av_freep(unsafe.Pointer(&buffer))
		return nil, errors.New("failed to alloc avio")
	}

	return ctx, nil
}

// NewIOReader wraps a Go Reader for use by the underlying library.
func NewIOReader(r io.ReadSeeker) (IOReader, error) {
	return newIO(r, 0, C.cgo_avio_read_packet, nil, C.cgo_avio_seek)
}

// NewIOWriter wraps a Go Writer for use by the underlying library.
func NewIOWriter(r io.WriteSeeker) (IOWriter, error) {
	return newIO(r, 1, nil, C.cgo_avio_write_packet, C.cgo_avio_seek)
}

func openURL(url string, flags C.int) (*avio, error) {
	ctx := &avio{}
	ctx.stream = nil

	c_url := C.CString(url)
	defer C.free(unsafe.Pointer(c_url))

	ret := C.avio_open(&ctx.ctx, c_url, C.AVIO_FLAG_READ)
	if ret < 0 {
		return nil, errors.New(averror(ret))
	}

	return ctx, nil
}

// OpenURLSource opens a url for use by the underlying library.
func OpenURLSource(url string) (IOReader, error) {
	return openURL(url, C.AVIO_FLAG_READ)
}

// OpenURLSink opens a url for use by the underlying library.
func OpenURLSink(url string) (IOWriter, error) {
	return openURL(url, C.AVIO_FLAG_WRITE)
}

func makeByteSlice(buf *C.uint8_t, buf_siz C.int) []byte {
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(buf)),
		Len:  int(buf_siz),
		Cap:  int(buf_siz),
	}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

//export go_avio_read_packet
func go_avio_read_packet(opaque unsafe.Pointer, buf *C.uint8_t, buf_siz C.int) C.int {
	ctx := (*avio)(opaque)
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

//export go_avio_write_packet
func go_avio_write_packet(opaque unsafe.Pointer, buf *C.uint8_t, buf_siz C.int) C.int {
	ctx := (*avio)(opaque)
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

//export go_avio_seek
func go_avio_seek(opaque unsafe.Pointer, offset C.int64_t, whence C.int) C.int64_t {
	ctx := (*avio)(opaque)
	ret, err := ctx.stream.(io.Seeker).Seek(int64(offset), int(whence))
	if err != nil {
		return -1
	}
	return C.int64_t(ret)
}
