#include <stdint.h>

#include "callback.h"
#include "_cgo_export.h"


int cgo_read_packet_wrap(void *opaque, uint8_t *buf, int buf_size)
{
  return ioReadPacket(opaque, buf, buf_size);
}

int cgo_write_packet_wrap(void *opaque, uint8_t *buf, int buf_size)
{
  return ioWritePacket(opaque, buf, buf_size);
}

int64_t cgo_seek_wrap(void *opaque, int64_t offset, int whence)
{
  return ioSeek(opaque, offset, whence);
}
