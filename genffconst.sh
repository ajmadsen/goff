#!/usr/bin/env bash

# Generate necessary FFmpeg constants from the headers

PKGCONFIG=pkg-config
CC=clang
SED=gsed

# Internal, do not change

CFLAGS=$($PKGCONFIG --cflags libavformat libavcodec libavutil)

cat > _ff.c <<EOF
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
EOF

$CC $CFLAGS -E -o _ff.i _ff.c

CODECS=$(grep AV_CODEC_ID_ _ff.i | gsed -E 's/.*(AV_CODEC_ID_\w+).*/\1/' | sort | uniq | gsed -E 's/(.*)/\1 = CodecID(C.\1)/')

gofmt > const.go <<EOF
/* GENERATED BY $(basename "$0"), DO NOT EDIT */

package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
*/
import "C"

type CodecID int

var (
$CODECS
)
EOF

rm -f _ff.c _ff.i