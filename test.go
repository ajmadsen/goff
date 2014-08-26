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
		ofmt_ctx *C.AVFormatContext
		dec_ctx  *C.AVCodecContext
		//enc_ctx   *C.AVCodecContext
		istream   C.int
		ostream   *C.AVStream
		pkt       C.AVPacket
		next_pts  C.int64_t
		pts_skew  C.int64_t
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

	log.Print("Finding best stream")
	istream = C.av_find_best_stream(ifmt_ctx, C.AVMEDIA_TYPE_AUDIO, -1, -1, nil, 0)
	if istream < 0 {
		log.Fatal(averror(istream))
	}
	log.Printf("Using stream [%d]", istream)

	log.Print("Opening output file")
	c_outfile := C.CString(outfile)
	defer C.free(unsafe.Pointer(c_outfile))
	ret = C.avformat_alloc_output_context2(&ofmt_ctx, nil, nil, c_outfile)
	if ret < 0 || ofmt_ctx == nil {
		log.Fatal(averror(ret))
	}
	defer C.avformat_free_context(ofmt_ctx)

	log.Print("Creating output stream")
	ostream = C.avformat_new_stream(ofmt_ctx, streams[istream].codec.codec)
	if ostream == nil {
		log.Fatal("Unable to create new stream")
	}
	defer C.avcodec_close(ostream.codec)

	// copy stream information
	ret = C.avcodec_copy_context(ostream.codec, streams[istream].codec)
	if ret < 0 {
		log.Fatal(averror(ret))
	}
	ostream.codec.codec_tag = 0
	if ofmt_ctx.oformat.flags&C.AVFMT_GLOBALHEADER != 0 {
		ostream.codec.flags |= C.CODEC_FLAG_GLOBAL_HEADER
	}

	C.av_dump_format(ofmt_ctx, 0, c_outfile, 1)

	if ofmt_ctx.flags&C.AVFMT_NOFILE == 0 {
		ret = C.avio_open(&ofmt_ctx.pb, c_outfile, C.AVIO_FLAG_WRITE)
		if ret < 0 {
			log.Fatal(averror(ret))
		}
		defer C.avio_close(ofmt_ctx.pb)
	}

	ostream.time_base = ostream.codec.time_base

	ret = C.avformat_write_header(ofmt_ctx, nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}

	dec_ctx = streams[istream].codec
	ret = C.avcodec_open2(dec_ctx, C.avcodec_find_decoder(dec_ctx.codec_id), nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}

	// initialize packet
	C.av_init_packet(&pkt)
	pkt.data = nil
	pkt.size = 0
	frame = C.av_frame_alloc()
	if frame == nil {
		log.Fatal("Failed to alloc frame")
	}

	for C.av_read_frame(ifmt_ctx, &pkt) >= 0 {
		if pkt.stream_index != istream {
			C.av_free_packet(&pkt)
			continue
		}
		if pkt.pts < 0 || pkt.pts == C.AV_NOPTS_VALUE {
			log.Printf("WARN: pts < 0 [%d]", pkt.pts)
		}
		if pkt.pts != next_pts {
			if pkt.pts+pts_skew != next_pts {
				log.Printf("WARN: pts != next_pts [%v,%v]", pkt.pts, next_pts)
				pts_skew = next_pts - pkt.pts
			}
			pkt.pts = next_pts
			pkt.dts = next_pts
		}
		//log.Printf("pts = %10d, dts = %10d", pkt.pts, pkt.dts)
		next_pts = pkt.pts + C.int64_t(pkt.duration)
		//pretty.Log(pkt)

		ret = C.avcodec_decode_audio4(dec_ctx, frame, &got_frame, &pkt)
		if ret < 0 {
			log.Fatal(averror(ret))
		}
		//pretty.Log(frame)

		//C.av_packet_rescale_ts(&pkt, streams[istream].time_base, ostream.time_base)

		// write packet to file
		pkt.stream_index = ostream.index
		pkt.pos = -1
		ret = C.av_write_frame(ofmt_ctx, &pkt)
		if ret < 0 {
			log.Fatal(averror(ret))
		}

		C.av_free_packet(&pkt)
		C.av_frame_unref(frame)
	}

	ret = C.av_write_frame(ofmt_ctx, nil)
	if ret < 0 {
		log.Fatal(averror(ret))
	}

	C.av_frame_free(&frame)

	ret = C.av_write_trailer(ofmt_ctx)
	if ret < 0 {
		log.Fatal(averror(ret))
	}
	C.av_dump_format(ofmt_ctx, 0, c_outfile, 1)
	log.Print("Completed!")
}
