package avio

// #cgo pkg-config: libavformat libavutil
// #include <libavformat/avio.h>
// #include <libavutil/avutil.h>
// #include <stdint.h>
/*
	extern int my_avio_read(void *opaque, uint8_t *buf, int buf_size);
*/
import "C"
import "io"
import "reflect"
import "unsafe"

var AVERROR_EOF = C.AVERROR_EOF

type AVIOContext struct {
	ctx *C.AVIOContext
	r   io.Reader
}

func New(r io.Reader) *AVIOContext {
	const SIZE = 4096
	buf := unsafe.Pointer(C.av_malloc(SIZE))
	if buf == nil {
		return nil
	}

	avioctx := &AVIOContext{nil, r}
	avioctx.ctx = C.avio_alloc_context(
		(*C.uchar)(buf),
		SIZE,
		0,
		unsafe.Pointer(avioctx),
		(*[0]byte)(unsafe.Pointer(C.my_avio_read)),
		(*[0]byte)(unsafe.Pointer(uintptr(0))),
		(*[0]byte)(unsafe.Pointer(uintptr(0))))

	return avioctx
}

func (ctx *AVIOContext) Close() {
	C.av_freep(unsafe.Pointer(&ctx.ctx.buffer))
	C.av_freep(unsafe.Pointer(&ctx.ctx))
}

// for testing
func avio_read(ioctx *C.AVIOContext, buf []byte, siz int) int {
	n := int(C.avio_read(ioctx, (*C.uchar)(&buf[0]), C.int(siz)))
	return n
}

//export my_avio_read
func my_avio_read(opaque unsafe.Pointer, buf *C.uint8_t, buf_size C.int) C.int {
	avioctx := (*AVIOContext)(opaque)
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(buf)),
		Len:  int(buf_size),
		Cap:  int(buf_size),
	}
	gobuf := *(*[]byte)(unsafe.Pointer(&hdr))
	n, err := avioctx.r.Read(gobuf)
	if n > 0 {
		return C.int(n)
	}
	if err != nil && err != io.EOF {
		return C.int(C.AVERROR_EXTERNAL)
	}
	return 0
}
