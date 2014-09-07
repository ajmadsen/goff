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
	"fmt"
	"unsafe"
)

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
