package files

import (
	"fmt"
	"os"
	"syscall"
)

func OpenFile(filepath string, flag int) (*os.File, error) {
	if flag == 0 {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	}

	f, err := os.OpenFile(filepath, flag, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}
