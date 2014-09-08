package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <stdlib.h>
*/
import "C"

// IOContext is an opaque representation of the AVIOContext type in the C API.
type IOContext struct {
	io *C.AVIOContext
}
