//go:build windows

package client

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/bradewing/gobwapi/internal/shm"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procOpenFileMapping = kernel32.NewProc("OpenFileMappingA")
	procMapViewOfFile   = kernel32.NewProc("MapViewOfFile")
	procUnmapViewOfFile = kernel32.NewProc("UnmapViewOfFile")
	procCloseHandle     = kernel32.NewProc("CloseHandle")
)

const (
	fileMapAllAccess = 0xF001F
)

func mapGameTable() (unsafe.Pointer, func(), error) {
	return mapSharedMemory(
		"Local\\bwapi_shared_memory_game_list",
		shm.GameTableSize,
	)
}

func mapGameData(serverPID uint32) (unsafe.Pointer, func(), error) {
	return mapSharedMemory(
		fmt.Sprintf("Local\\bwapi_shared_memory_%d", serverPID),
		shm.GameDataSize,
	)
}

func mapSharedMemory(name string, size int) (unsafe.Pointer, func(), error) {
	namePtr, err := syscall.BytePtrFromString(name)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid name %s: %w", name, err)
	}

	handle, _, err := procOpenFileMapping.Call(
		uintptr(fileMapAllAccess),
		0,
		uintptr(unsafe.Pointer(namePtr)),
	)
	if handle == 0 {
		return nil, nil, fmt.Errorf("OpenFileMapping %s: %w", name, err)
	}

	ptr, _, err := procMapViewOfFile.Call(
		handle,
		uintptr(fileMapAllAccess),
		0, 0,
		uintptr(size),
	)
	if ptr == 0 {
		procCloseHandle.Call(handle)
		return nil, nil, fmt.Errorf("MapViewOfFile %s: %w", name, err)
	}

	cleanup := func() {
		procUnmapViewOfFile.Call(ptr)
		procCloseHandle.Call(handle)
	}
	return unsafe.Pointer(ptr), cleanup, nil
}
