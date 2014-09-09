#ifndef _CALLBACK_H_
#define _CALLBACK_H_

#include <stdint.h>

int cgo_avio_read_packet(void *opaque, uint8_t *buf, int buf_size);
int cgo_avio_write_packet(void *opaque, uint8_t *buf, int buf_size);
int64_t cgo_avio_seek(void *opaque, int64_t offset, int whence);

#endif // _CALLBACK_H_
