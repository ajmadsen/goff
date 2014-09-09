package av

import "testing"

func asserteq(t *testing.T, what string, expected interface{}, got interface{}) {
	if expected != got {
		t.Fatalf("wrong value for %s, expected %v, got %v", what, expected, got)
	}
}

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

func TestNStreams(t *testing.T) {
	fmt := openDemux(t, fname)
	asserteq(t, "NStreams()", 1, fmt.NStreams())
	fmt.Close()
}

func TestStream(t *testing.T) {
	fmt := openDemux(t, fname)
	str := fmt.Stream(fmt.NStreams() - 1)
	if str == nil {
		t.Fatalf("should have stream %d", fmt.NStreams()-1)
	}
	str = fmt.Stream(-1)
	if str != nil {
		t.Fatalf("should not have stream -1")
	}
	str = fmt.Stream(fmt.NStreams() + 1)
	if str != nil {
		t.Fatalf("should not have stream %d", fmt.NStreams()+1)
	}
	fmt.Close()
}

func TestStreamIdx(t *testing.T) {
	fmt := openDemux(t, fname)
	str := fmt.Stream(fmt.NStreams() - 1)
	asserteq(t, "str.Index()", fmt.NStreams()-1, str.Index())
	fmt.Close()
}

func TestStreamIsOpen(t *testing.T) {
	fmt := openDemux(t, fname)
	str := fmt.Stream(fmt.NStreams() - 1)
	asserteq(t, "str.IsOpen()", false, str.IsOpen())
	fmt.Close()
}
