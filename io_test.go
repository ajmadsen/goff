package av

import (
	"io/ioutil"
	"os"
	"testing"
)

const fname = "audio.mp3"

func diff(b1 []byte, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := range b1 {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}

func openFile(t *testing.T) *IO {
	file, err := os.Open(fname)
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}

	ioctx, err := NewIO(file, fname, false)
	if err != nil {
		t.Errorf("failed to create IO ctx: %v", err)
	}

	return ioctx
}

func TestNewIO(t *testing.T) {
	ioctx := openFile(t)
	ioctx.Close()
}

func TestIORead(t *testing.T) {
	ioctx := openFile(t)

	b1, err := ioutil.ReadAll(ioctx)
	if err != nil {
		t.Fatal("could not read file via avio")
	}
	b2, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal("could not read file via go")
	}

	if !diff(b1, b2) {
		t.Fatal("read bytes differ")
	} else {
		t.Log("read bytes are the same")
	}

	ioctx.Close()
}
