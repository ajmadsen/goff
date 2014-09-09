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
	"io"
	"reflect"
	"unsafe"
)

const NO_PTS = C.AV_NOPTS_VALUE

type Demuxer interface {
	Close()
	Dump(n int)
	NStreams() int
	Stream(idx int) Stream
	ReadPacket() (Packet, error)
}

type Stream interface {
	IsOpen() bool
	Index() int
	//Open(codec string) error
}

type Packet interface {
	Free()
	IsKey() bool
	IsCorrupt() bool
	Pts() int64
	Dts() int64
	Duration() int
	Index() int
	Data() []byte
	Size() int
	Position() int64
}

type avfmtctx struct {
	fctx   *C.AVFormatContext
	c_name *C.char
	ioctx  *avio
	pkt    C.AVPacket
}

type stream struct {
	s *C.AVStream
}

type packet struct {
	p C.AVPacket
}

// AVFMTCTX

func (f *avfmtctx) Close() {
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

func (f *avfmtctx) Dump(n int) {
	C.av_dump_format(f.fctx, C.int(n), f.c_name, 0)
}

func (f *avfmtctx) streams() []*C.AVStream {
	nstr := f.NStreams()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(f.fctx.streams)),
		Len:  nstr,
		Cap:  nstr,
	}
	return *(*[]*C.AVStream)(unsafe.Pointer(&hdr))
}

func (f *avfmtctx) NStreams() int {
	return int(f.fctx.nb_streams)
}

func (f *avfmtctx) Stream(idx int) Stream {
	if idx > f.NStreams() || idx < 0 {
		return nil
	}

	return &stream{f.streams()[idx]}
}

func (f *avfmtctx) ReadPacket() (Packet, error) {
	ret := C.av_read_frame(f.fctx, &f.pkt)
	if ret < 0 {
		if ret == C.AVERROR_EOF {
			return nil, io.EOF
		}
		return nil, errors.New(averror(ret))
	}

	defer C.av_packet_unref(&f.pkt)

	newpkt := &packet{}
	ret = C.av_packet_ref(&newpkt.p, &f.pkt)
	if ret < 0 {
		return nil, errors.New(averror(ret))
	}

	return newpkt, nil
}

// STREAM

func (s *stream) Index() int {
	return int(s.s.index)
}

func (s *stream) IsOpen() bool {
	return (s.s.codec != nil && s.s.codec.codec != nil)
}

// PACKET

func (p *packet) Free() {
	C.av_packet_unref(&p.p)
}

func (p *packet) IsKey() bool {
	return (p.p.flags&C.AV_PKT_FLAG_KEY != 0)
}

func (p *packet) IsCorrupt() bool {
	return (p.p.flags&C.AV_PKT_FLAG_CORRUPT != 0)
}

func (p *packet) Pts() int64 {
	return int64(p.p.pts)
}

func (p *packet) Dts() int64 {
	return int64(p.p.dts)
}

func (p *packet) Duration() int {
	return int(p.p.duration)
}

func (p *packet) Index() int {
	return int(p.p.stream_index)
}

func (p *packet) Size() int {
	return int(p.p.size)
}

func (p *packet) Data() []byte {
	siz := p.Size()
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(p.p.data)),
		Len:  siz,
		Cap:  siz,
	}
	return *(*[]byte)(unsafe.Pointer(&hdr))
}

func (p *packet) Position() int64 {
	return int64(p.p.pos)
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

	var pkt C.AVPacket
	C.av_init_packet(&pkt)

	return &avfmtctx{
		fmt,
		C.CString(name),
		ioctx,
		pkt,
	}, nil
}
