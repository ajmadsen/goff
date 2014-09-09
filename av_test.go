package av

import "testing"

func openDemux(t *testing.T, file string) Demuxer {
	ioctx := openFile(t, fname)

	fmt, err := OpenDemuxer(ioctx, fname)
	if err != nil {
		t.Fatalf("could not open reader: %v", err)
	}

	return fmt
}

func TestOpenDemuxer(t *testing.T) {
	fmt := openDemux(t, fname)
	fmt.Close()
}

func TestDump(t *testing.T) {
	fmt := openDemux(t, fname)
	fmt.Dump(0)
	fmt.Close()
}
