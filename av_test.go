package av

import (
	"io"
	"testing"
)

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

func TestPacket(t *testing.T) {
	fmt := openDemux(t, fname)
	pkt, err := fmt.ReadPacket()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("index      = %v", pkt.Index())
	t.Logf("dts        = %v", pkt.Dts())
	t.Logf("pts        = %v", pkt.Pts())
	t.Logf("duration   = %v", pkt.Duration())
	t.Logf("pos        = %v", pkt.Position())
	t.Logf("iskey      = %v", pkt.IsKey())
	t.Logf("iscorrupt  = %v", pkt.IsCorrupt())
	t.Logf("size       = %v", pkt.Size())

	siz := pkt.Size()
	if siz > 16 {
		siz = 16
	}

	t.Logf("data       = %v...", pkt.Data()[:siz])

	pkt.Free()
	fmt.Close()
}

func TestReadPacket(t *testing.T) {
	fmt := openDemux(t, fname)

	count := int(0)
Loop:
	for {
		pkt, err := fmt.ReadPacket()
		switch {
		case err == io.EOF:
			break Loop
		case err != nil:
			t.Fatal(err)
		default:
			count++
			pkt.Free()
		}
	}

	t.Logf("Read %d packets", count)

	fmt.Close()
}
