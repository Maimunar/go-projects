package files

import (
	"os"
	"testing"
)

type args struct {
	filepath string
	flag     int
}

func TestOpenCloseFile(t *testing.T) {
	arg := args{
		filepath: "test.csv",
		flag:     0,
	}

	got, err := OpenFile(arg.filepath, arg.flag)
	if err != nil {
		t.Errorf("OpenFile() error = %v", err)
		return
	}
	err = CloseFile(got)
	if err != nil {
		t.Errorf("CloseFile() error = %v", err)
		return
	}
	// Cleanup
	err = os.Remove(arg.filepath)
	if err != nil {
		t.Errorf("Error removing file while cleaning up: %v", err)
	}
}
