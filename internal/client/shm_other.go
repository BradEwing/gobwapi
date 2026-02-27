//go:build !linux && !windows

package client

import (
	"fmt"
	"runtime"
	"unsafe"
)

func mapGameTable() (unsafe.Pointer, func(), error) {
	return nil, nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
}

func mapGameData(serverPID uint32) (unsafe.Pointer, func(), error) {
	return nil, nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
}
