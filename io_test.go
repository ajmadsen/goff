package av

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

const fname = "audio.mp3"

func diff(b1 []byte, b2 []byte) bool {
	if len(b1) != len(b2) {
		return true
	}

	for i := range b1 {
		if b1[i] != b2[i] {
			return true
		}
	}

	return false
}

func openFile(t *testing.T, name string) IOReader {
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	ioctx, err := NewIOReader(file)
	if err != nil {
		t.Fatalf("failed to create IO ctx: %v", err)
	}

	return ioctx
}

func TestNewIO(t *testing.T) {
	ioctx := openFile(t, fname)
	ioctx.Close()
}

func TestIOReadFile(t *testing.T) {
	ioctx := openFile(t, fname)

	b1, err := ioutil.ReadAll(ioctx)
	if err != nil {
		t.Fatal("could not read file via avio")
	}
	b2, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal("could not read file via go")
	}

	if diff(b1, b2) {
		t.Fatal("read bytes differ")
	}

	ioctx.Close()
}

func TestIOOpenURL(t *testing.T) {
	wait := make(chan int)
	go func(ready chan<- int) {
		ready <- 1
		err := http.ListenAndServe(":32149", http.FileServer(http.Dir(".")))
		if err != nil {
			t.Fatal(err)
		}
	}(wait)

	// Hack to wait for the goroutine to start, because we were getting issues
	// where FFmpeg was attempting to open the URL before Go could spin up the
	// HTTP server. Waiting for the goroutine to start appeared to solve the
	// issue, and is slightly more flexible than sleeping for some amount of
	// time.
	<-wait

	ioctx, err := OpenURLSource("http://localhost:32149/" + fname)
	if err != nil {
		t.Fatal(err)
	}

	b1, err := ioutil.ReadAll(ioctx)
	if err != nil {
		t.Fatal(err)
	}
	b2, err := ioutil.ReadFile(fname)
	if err != nil {
		t.Fatal(err)
	}

	if diff(b1, b2) {
		t.Fatal("file differs from local")
	}

	ioctx.Close()
}
