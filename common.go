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
)

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
