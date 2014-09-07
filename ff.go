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
	"errors"
	"fmt"
	"runtime"
	"unsafe"
)

type Input interface {
	OpenDemuxer(string) error
	OpenDecoder(int, Dictionary) error
	CloseDecoder(int)
	Close()
	NStreams() int
	Dump(int)
}

type Output interface {
	OpenMuxer(string) error
	NewStream(CodecID) (int, error)
	OpenEncoder(int, Dictionary) error
	CloseEncoder(int) error
	Close()
	NStreams() int
	Dump(int)
}

type Dictionary struct {
	dict *C.AVDictionary
}

func NewDictionary(m map[string]interface{}) Dictionary {
	var d Dictionary
	d.FromMap(m)
	return d
}

func (d *Dictionary) Map() map[string]string {
	m := make(map[string]string)

	empty := C.CString("")
	defer C.free(unsafe.Pointer(empty))

	ent := C.av_dict_get(d.dict, empty, nil, C.AV_DICT_IGNORE_SUFFIX)
	for ; ent != nil; ent = C.av_dict_get(d.dict, empty, ent, C.AV_DICT_IGNORE_SUFFIX) {
		m[C.GoString(ent.key)] = C.GoString(ent.value)
	}

	return m
}

func (d *Dictionary) FromMap(m map[string]interface{}) {
	if d.dict != nil {
		C.av_dict_free(&d.dict)
		d.dict = nil
	}

	for k, v := range m {
		c_k := C.CString(k)

		switch v := v.(type) {
		case string:
			c_v := C.CString(v)
			C.av_dict_set(&d.dict, c_k, c_v, 0)
			C.free(unsafe.Pointer(c_v))
		case *C.char:
			C.av_dict_set(&d.dict, c_k, v, 0)
		case int:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case uint:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case int32:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case uint32:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case int64:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case uint64:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.int:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.uint:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.int32_t:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.uint32_t:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.int64_t:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.uint64_t:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		case C.size_t:
			C.av_dict_set_int(&d.dict, c_k, C.int64_t(v), 0)
		default:
			panic(fmt.Sprintf("cannot encode type %T in dictionary", v))
		}

		C.free(unsafe.Pointer(c_k))
	}
}

func (d *Dictionary) Free() {
	C.av_dict_free(&d.dict)
	d.dict = nil
}

type fmtctx struct {
	fctx    *C.AVFormatContext
	c_fname *C.char
}

// newInput creates a new fmtctx structure and sets the finalizer
func NewInput() Input {
	i := &fmtctx{}
	runtime.SetFinalizer(i, func(i *fmtctx) {
		i.Close()
	})
	return i
}

// OpenMuxer opens a file in FFmpeg by allocating an AVFormatContext and
// opening an AVIO if necessary
func (i *fmtctx) OpenMuxer(filename string) error {
	return nil
}

// OpenDemuxer opens a file in FFmpeg by allocating an AVFormatContext and finding
// stream information
func (i *fmtctx) OpenDemuxer(filename string) error {
	if i.fctx != nil {
		return errors.New("fmtctx already open")
	}

	i.c_fname = C.CString(filename)

	ret := C.avformat_open_input(&i.fctx, i.c_fname, nil, nil)
	if ret < 0 {
		i.fctx = nil
		return averror(ret)
	}

	ret = C.avformat_find_stream_info(i.fctx, nil)
	if ret < 0 {
		i.Close()
		return averror(ret)
	}

	return nil
}

// NewStream allocates a new stream in a fmtctx
func (i *fmtctx) NewStream(c CodecID) (int, error) {
	return -1, nil
}

// OpenEncoder opens a AVCodecContext for an allocated stream
func (i *fmtctx) OpenEncoder(idx int, opts Dictionary) error {
	return nil
}

// OpenDecoder opens an AVCodecContext to decode the stream specified by idx
func (i *fmtctx) OpenDecoder(idx int, opts Dictionary) error {
	// bounds check
	nb_streams := int(i.fctx.nb_streams)
	if idx > nb_streams {
		return errors.New(fmt.Sprintf("index out of bounds: %d > %d", idx,
			nb_streams))
	}

	stream := i.streams()[idx]

	// find decoder
	decoder := C.avcodec_find_decoder(stream.codec.codec_id)
	if decoder == nil {
		return errors.New(fmt.Sprintf("could not find decoder for stream %d", idx))
	}

	// open decoder
	ret := C.avcodec_open2(stream.codec, decoder, nil)
	if ret < 0 {
		return averror(ret)
	}

	return nil
}

// CloseDecoder closes the AVCodecContext associated with the stream, should
// one exist. It will panic if the idx is out of bounds, and will silently
// ignore you if the stream doesn't have an open codec.
func (i *fmtctx) CloseDecoder(idx int) {
	// bounds check
	nb_streams := int(i.fctx.nb_streams)
	if idx > nb_streams {
		panic(fmt.Sprintf("index out of bounds: %d > %d", idx, nb_streams))
	}

	stream := i.streams()[idx]

	if stream != nil && stream.codec != nil {
		C.avcodec_close(stream.codec)
	}
}

// Close attempts to close the fmtctx by closing all open decoders, closing the
// file, and freeing any other resources associated with the fmtctx.
func (i *fmtctx) Close() {
	if i.fctx != nil {
		for j := range i.streams() {
			i.CloseDecoder(j)
		}
		C.avformat_close_input(&i.fctx)
	}
	if i.c_fname != nil {
		C.free(unsafe.Pointer(i.c_fname))
		i.c_fname = nil
	}
	runtime.SetFinalizer(i, nil)
}

// Dump dumps information about the given stream to the console
func (i *fmtctx) Dump(stream int) {
	C.av_dump_format(i.fctx, C.int(stream), i.c_fname, 0)
}

// streams returns a slice containing the streams in an fmtctx. The slice is
// backed by a C array, and any attempts to modify the slice may have undesired
// consequences.
func (i *fmtctx) streams() []*C.AVStream {
	nstr := i.NStreams()
	return (*[1 << 30]*C.AVStream)(unsafe.Pointer(i.fctx.streams))[:nstr:nstr]
}

// NStreams returns the number of streams in a fmtctx.
func (i *fmtctx) NStreams() int {
	return int(i.fctx.nb_streams)
}

// averror converts a return code to a descriptive string
func averror(ret C.int) error {
	var errbuf [1024]C.char
	C.av_make_error_string(&errbuf[0], C.size_t(len(errbuf)), ret)
	return errors.New(C.GoStringN(&errbuf[0], C.int(len(errbuf))))
}

// Init initializes the FFmpeg library for usage
func Init() {
	C.av_register_all()
}
