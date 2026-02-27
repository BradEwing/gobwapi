//go:build !linux && !windows

package client

import (
	"fmt"
	"io"
	"runtime"
)

func dialPipe(serverPID uint32) (io.ReadWriteCloser, error) {
	return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
}
