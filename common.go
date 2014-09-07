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

// CodecID represents a CodecID exported by the FFmpeg API
type CodecID int

// Error represents an error returned by the FFmpeg API
type Error struct {
	// Errno is the return value of the function called in the C API
	Errno int
}

// Error returns the string representation of the error provided by the C API
func (e Error) Error() string {
	var errbuf [1024]C.char
	C.av_make_error_string(&errbuf[0], C.size_t(len(errbuf)), C.int(e.Errno))
	return C.GoStringN(&errbuf[0], C.int(len(errbuf)))
}

// AVError makes an Error out of a syscall.Errno
//
// Some functions in the C API return a negated syscall.Error, so this method
// is provided as a convenience for comparing returned errors to known POSIX
// errors.
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
