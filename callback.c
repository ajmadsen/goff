#include <stdint.h>
#include <stdio.h>

#include "callback.h"
#include "_cgo_export.h"


int cgo_read_packet_wrap(void *opaque, uint8_t *buf, int buf_size)
{
  int ret = ioReadPacket(opaque, buf, buf_size);
  return ret;
}

int cgo_write_packet_wrap(void *opaque, uint8_t *buf, int buf_size)
{
  int ret = ioWritePacket(opaque, buf, buf_size);
  return ret;
}

int64_t cgo_seek_wrap(void *opaque, int64_t offset, int whence)
{
  int ret = ioSeek(opaque, offset, whence);
  return ret;
}
