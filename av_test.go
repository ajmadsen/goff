package av

import "testing"

func TestOpenFile(t *testing.T) {
	ioctx := openFile(t)

	fmt, err := OpenReader(ioctx)
	if err != nil {
		t.Fatalf("could not open reader: %v", err)
	}
	fmt.Close()

	ioctx.Close()
}
