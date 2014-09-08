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
	"log"
	"reflect"
	"unsafe"
)

const bufsiz = 4096

type IO struct {
	ctx    *C.AVIOContext
	stream interface{}
	buf    *C.uchar
	name   string
}

func (i *IO) Name() string {
	return i.name
}

func (i *IO) Close() {
	if i.ctx != nil {
		C.av_freep(unsafe.Pointer(&i.ctx))
	}
	if i.buf != nil {
		C.av_freep(unsafe.Pointer(&i.buf))
	}
}

// NewIO creates a new IO context from a ReadSeeker.  stream should implement
// some set of {Reader, Writer, Seeker}.
func NewIO(stream interface{}, name string) (*IO, error) {
	ctx := &IO{}

	var readFcn, writeFcn, seekFcn unsafe.Pointer
	var c_writable C.int
	if _, ok := stream.(io.Reader); ok {
		readFcn = C.cgo_read_packet_wrap
	}
	if _, ok := stream.(io.Writer); ok {
		writeFcn = C.cgo_write_packet_wrap
		c_writable = 1
	}
	if _, ok := stream.(io.Seeker); ok {
		seekFcn = C.cgo_seek_wrap
	}

	if readFcn == nil && writeFcn == nil && seekFcn == nil {
		return nil, errors.New("stream does not implement anything useful")
	}

	ctx.buf = (*C.uchar)(C.av_malloc(bufsiz))
	if ctx.buf == nil {
		return nil, errors.New("out of memory")
	}

	ctx.name = name

	ctx.stream = stream
	ctx.ctx = C.avio_alloc_context(ctx.buf,
		bufsiz,
		c_writable,
		unsafe.Pointer(ctx),
		(*[0]byte)(readFcn),
		(*[0]byte)(writeFcn),
		(*[0]byte)(seekFcn))
	if ctx.ctx == nil {
		ctx.Close()
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
	log.Printf("ret = %v, err = %v", ret, err)
	log.Print(p)
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
	log.Printf("ret = %v, err = %v", ret, err)
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
	log.Printf("ret = %v, err = %v", ret, err)
	if err != nil {
		return -1
	}
	return 0
}
