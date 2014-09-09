#include <stdint.h>
#include <stdio.h>

#include "callback.h"
#include "_cgo_export.h"


int cgo_avio_read_packet(void *opaque, uint8_t *buf, int buf_size)
{
  int ret = go_avio_read_packet(opaque, buf, buf_size);
  return ret;
}

int cgo_avio_write_packet(void *opaque, uint8_t *buf, int buf_size)
{
  int ret = go_avio_read_packet(opaque, buf, buf_size);
  return ret;
}

int64_t cgo_avio_seek(void *opaque, int64_t offset, int whence)
{
  int ret = go_avio_seek(opaque, offset, whence);
  return ret;
}
