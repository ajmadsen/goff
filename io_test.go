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

func openFile(t *testing.T, name string) *IO {
	file, err := os.Open(name)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	ioctx, err := NewIO(file, fname, false)
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

	<-wait

	ioctx, err := OpenURL("http://localhost:32149/"+fname, FLAG_READ)
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
