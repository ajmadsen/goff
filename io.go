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

type IO struct {
	ctx    *C.AVIOContext
	stream interface{}
	name   string
}

func (i *IO) Name() string {
	return i.name
}

func (i *IO) Close() {
	if i.ctx.buffer != nil {
		C.av_freep(unsafe.Pointer(&i.ctx.buffer))
	}
	if i.ctx != nil {
		C.av_freep(unsafe.Pointer(&i.ctx))
	}
}

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

func (i *IO) Write(b []byte) (int, error) {
	C.avio_write(i.ctx, (*C.uchar)(unsafe.Pointer(&b[0])), C.int(len(b)))
	return len(b), nil
}

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
