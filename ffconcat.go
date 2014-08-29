package main

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
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"unsafe"
)

var (
	infiles stringList
	outfile string
)

type stringList []string

func (s *stringList) String() string {
	return fmt.Sprint(*s)
}

func (s *stringList) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type input struct {
	fctx    *C.AVFormatContext
	c_fname *C.char
}

func newInput() *input {
	i := &input{}
	runtime.SetFinalizer(i, func(i *input) {
		i.Close()
	})
	return i
}

func (i *input) Open(filename string) error {
	if i.fctx != nil {
		return errors.New("input already open")
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

func (i *input) OpenDecoder(idx int) error {
	// bounds check
	nb_streams := int(i.fctx.nb_streams)
	if idx > nb_streams {
		return errors.New(fmt.Sprintf("index out of bounds: %d > %d", idx, nb_streams))
	}

	stream := i.Streams()[idx]

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

func (i *input) CloseDecoder(idx int) {
	// bounds check
	nb_streams := int(i.fctx.nb_streams)
	if idx > nb_streams {
		panic(fmt.Sprintf("index out of bounds: %d > %d", idx, nb_streams))
	}

	stream := i.Streams()[idx]

	if stream != nil && stream.codec != nil {
		C.avcodec_close(stream.codec)
	}
}

func (i *input) Close() {
	if i.fctx != nil {
		for j := range i.Streams() {
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

func (i *input) Dump(stream int) {
	C.av_dump_format(i.fctx, C.int(stream), i.c_fname, 0)
}

func (i *input) Streams() []*C.AVStream {
	nstr := i.fctx.nb_streams
	return (*[1 << 30]*C.AVStream)(unsafe.Pointer(i.fctx.streams))[:nstr:nstr]
}

func averror(ret C.int) error {
	var errbuf [256]C.char
	C.av_make_error_string(&errbuf[0], C.size_t(len(errbuf)), ret)
	return errors.New(C.GoStringN(&errbuf[0], C.int(len(errbuf))))
}

func init() {
	flag.Var(&infiles, "i", "input file, can be given multiple times")
	flag.StringVar(&outfile, "o", "", "output file")
}

func main() {
	flag.Parse()

	if len(infiles) == 0 || outfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	C.av_register_all()

	inputs := make([]*input, len(infiles))
	for i, s := range infiles {
		ipt := newInput()
		err := ipt.Open(s)
		if err != nil {
			log.Fatal(err)
		}
		defer ipt.Close()
		ipt.Dump(i)
		for j := range ipt.Streams() {
			ipt.OpenDecoder(j)
		}
		inputs = append(inputs, ipt)
	}
}
