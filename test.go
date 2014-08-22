package main

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <stdlib.h>
*/
import "C"
import (
	"flag"
	"log"
	"os"
	"unsafe"
)

var (
	infile  string
	outfile string

	errbuf [256]C.char
)

func averror(ret C.int) string {
	C.av_make_error_string(&errbuf[0], C.size_t(len(errbuf)), ret)
	return C.GoStringN(&errbuf[0], C.int(len(errbuf)))
}

func get_streams(fctx *C.AVFormatContext) []*C.AVStream {
	nstr := fctx.nb_streams
	return (*[1 << 30]*C.AVStream)(unsafe.Pointer(fctx.streams))[:nstr:nstr]
}

func init() {
	flag.StringVar(&infile, "in", "", "input file")
	flag.StringVar(&infile, "i", "", "input file [short]")
	flag.StringVar(&outfile, "out", "", "output file")
	flag.StringVar(&outfile, "o", "", "output file [short]")
}

func main() {
	var (
		ifmt_ctx *C.AVFormatContext
		//ofmt_ctx *C.AVFormatContext
		dec_ctx *C.AVCodecContext
		//enc_ctx  *C.AVCodecContext
		//istream  C.int
		pkt      C.AVPacket
		next_pts C.int64_t
		//shift_pts C.int64_t
		frame     *C.AVFrame
		got_frame C.int
	)

	flag.Parse()

	if infile == "" || outfile == "" {
		flag.Usage()
		os.Exit(1)
	}

	c_infile := C.CString(infile)
	defer C.free(unsafe.Pointer(c_infile))

	C.av_register_all()

	ret := C.avformat_open_input(&ifmt_ctx, c_infile, nil, nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}
	defer C.avformat_close_input(&ifmt_ctx)

	ret = C.avformat_find_stream_info(ifmt_ctx, nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}

	log.Print("Input file information:")
	for i := 0; i < int(ifmt_ctx.nb_streams); i++ {
		C.av_dump_format(ifmt_ctx, C.int(i), c_infile, 0)
	}

	streams := get_streams(ifmt_ctx)

	log.Print("Opening first stream")
	dec_ctx = streams[0].codec
	ret = C.avcodec_open2(dec_ctx, C.avcodec_find_decoder(dec_ctx.codec_id), nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}
	defer C.avcodec_close(dec_ctx)
	frame = C.av_frame_alloc()
	if frame == nil {
		log.Fatal("Could not alloc frame")
	}
	defer C.av_frame_free(&frame)

	log.Print("Initializing packet")
	C.av_init_packet(&pkt)
	pkt.data = nil
	pkt.size = 0

	next_pts = -1
	log.Print("Reading packets")
	for C.av_read_frame(ifmt_ctx, &pkt) >= 0 {
		orig_pkt := pkt
		if pkt.pts < 0 {
			log.Printf("WARN: negative pts value: %d", pkt.pts)
		}
		if next_pts > 0 && next_pts != pkt.pts {
			log.Printf("WARN: mismatch pts. Expected [%d] got [%d].", next_pts, pkt.pts)
			pkt.pts = next_pts
		}
		for pkt.size > 0 {
			read := C.avcodec_decode_audio4(dec_ctx, frame, &got_frame, &pkt)
			if read < 0 {
				log.Fatal(averror(read))
			}
			new_data_ptr := uintptr(unsafe.Pointer(pkt.data))
			new_data_ptr += uintptr(read) * unsafe.Sizeof(*pkt.data)
			pkt.data = (*C.uint8_t)(unsafe.Pointer(new_data_ptr))
			pkt.size -= read
			if got_frame != 0 {
				frame.pts = C.av_frame_get_best_effort_timestamp(frame)
				log.Printf("Got frame with pts = %8d", frame.pts)
			}
		}
		pkt.data = nil
		pkt.size = 0
		for {
			C.avcodec_decode_audio4(dec_ctx, frame, &got_frame, &pkt)
			if got_frame != 0 {
				frame.pts = C.av_frame_get_best_effort_timestamp(frame)
				log.Printf("Got frame with pts = %8d", frame.pts)
			} else {
				break
			}
		}
		next_pts = pkt.pts + C.int64_t(pkt.duration)
		C.av_free_packet(&orig_pkt)
	}
}
