#ifndef _CALLBACK_H_
#define _CALLBACK_H_

#include <stdint.h>

int cgo_read_packet_wrap(void *opaque, uint8_t *buf, int buf_size);
int cgo_write_packet_wrap(void *opaque, uint8_t *buf, int buf_size);
int64_t cgo_seek_wrap(void *opaque, int64_t offset, int whence);

#endif // _CALLBACK_H_
