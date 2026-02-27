//go:build windows

package client

import (
	"fmt"
	"io"
	"os"
)

func dialPipe(serverPID uint32) (io.ReadWriteCloser, error) {
	name := fmt.Sprintf(`\\.\pipe\bwapi_pipe_%d`, serverPID)
	f, err := os.OpenFile(name, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("open named pipe %s: %w", name, err)
	}
	return f, nil
}
