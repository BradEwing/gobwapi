//go:build linux

package client

import (
	"fmt"
	"io"
	"net"
)

func dialPipe(serverPID uint32) (io.ReadWriteCloser, error) {
	path := fmt.Sprintf("/tmp/bwapi_socket_%d", serverPID)
	conn, err := net.Dial("unix", path)
	if err != nil {
		return nil, fmt.Errorf("dial unix socket %s: %w", path, err)
	}
	return conn, nil
}
