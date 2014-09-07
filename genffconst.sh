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
#include <libavutil/error.h>
EOF

$CC -dM $CFLAGS -E -o _ff.1 _ff.c
$CC $CFLAGS -E -o _ff.2 _ff.c
cat _ff.1 _ff.2 > _ff.ii

CODECS=$(grep AV_CODEC_ID_ _ff.ii | gsed -E 's/.*(AV_CODEC_ID_\w+).*/\1/' | sort | uniq | gsed -E 's/(.*)/\1 = CodecID(C.\1)/')
ERRORS=$(grep AVERROR_ _ff.ii | gsed -E 's/.*(AVERROR_\w+).*/\1/' | sort | uniq | gsed -E 's/(.*)/\1 = Error{C.\1}/')

gofmt > const.go <<EOF
/* GENERATED BY $(basename "$0"), DO NOT EDIT */

package ff

/*
#cgo pkg-config: libavformat libavcodec libavutil
#include <libavformat/avformat.h>
#include <libavcodec/avcodec.h>
#include <libavutil/avutil.h>
#include <libavutil/error.h>
*/
import "C"

// Codec definitions
var (
$CODECS
)

// Error definitions
var (
$ERRORS
)
EOF

#rm -f _ff.{1,2,ii,c}
