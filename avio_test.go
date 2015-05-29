package avio

import "bytes"
import "testing"

func TestRead(t *testing.T) {
	var buf [512]byte

	cases := [][]byte{
		[]byte("Hello, World!"),
		[]byte("世界"),
		[]byte(""),
	}

	for _, c := range cases {
		r := bytes.NewReader(c)

		ioctx := New(r)
		if ioctx == nil {
			t.Error("failed to alloc avio context")
		}

		n := avio_read(ioctx.ctx, buf[:], len(buf))

		if n == AVERROR_EOF {
			continue
		}

		read := buf[:n]

		if n != len(c) {
			t.Errorf("length read does not match input length: %d != %d", n, len(c))
			ioctx.Close()
			continue
		}

		if string(read[:]) != string(c) {
			t.Errorf("read string does not match input string: \"%x\" != \"%x\"", string(read), string(c))
			ioctx.Close()
			continue
		}
	}
}
