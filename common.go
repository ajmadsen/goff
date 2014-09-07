package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <stdlib.h>
*/
import "C"
import "syscall"

type CodecID int

type Error struct {
	Errno int
}

func (e Error) Error() string {
	var errbuf [1024]C.char
	C.av_make_error_string(&errbuf[0], C.size_t(len(errbuf)), C.int(e.Errno))
	return C.GoStringN(&errbuf[0], C.int(len(errbuf)))
}

func AVError(err syscall.Errno) Error {
	return Error{-int(err)}
}

// averror converts a return code to a descriptive string
func averror(ret C.int) error {
	return Error{int(ret)}
}

// Init initializes the FFmpeg library for usage
func Init() {
	C.av_register_all()
}
