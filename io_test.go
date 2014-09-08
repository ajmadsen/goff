package av

import (
	"os"
	"testing"
)

const fname = "audio.mp3"

func openFile(t *testing.T) *IO {
	file, err := os.Open(fname)
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}

	ioctx, err := NewIO(file, fname)
	if err != nil {
		t.Errorf("failed to create IO ctx: %v", err)
	}

	return ioctx
}

func TestNewIO(t *testing.T) {
	ioctx := openFile(t)
	ioctx.Close()
}
