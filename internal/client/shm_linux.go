//go:build linux

package client

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"

	"github.com/bradewing/gobwapi/internal/shm"
)

func mapGameTable() (unsafe.Pointer, func(), error) {
	return mapSharedMemory(
		"/dev/shm/bwapi_shared_memory_game_list",
		shm.GameTableSize,
	)
}

func mapGameData(serverPID uint32) (unsafe.Pointer, func(), error) {
	return mapSharedMemory(
		fmt.Sprintf("/dev/shm/bwapi_shared_memory_%d", serverPID),
		shm.GameDataSize,
	)
}

func mapSharedMemory(path string, size int) (unsafe.Pointer, func(), error) {
	f, err := os.OpenFile(path, os.O_RDWR, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("open shared memory %s: %w", path, err)
	}
	defer f.Close()

	data, err := syscall.Mmap(int(f.Fd()), 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		return nil, nil, fmt.Errorf("mmap %s: %w", path, err)
	}

	cleanup := func() {
		syscall.Munmap(data)
	}
	return unsafe.Pointer(&data[0]), cleanup, nil
}
