package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"time"
)

func CameraReader(port io.ReadWriteCloser, options CameraOptions, body *Body) error {
	port.Write([]byte(options.Keyword))

	size := make([]byte, 4)
	n, err := io.ReadFull(port, size)
	if err != nil {
		options.OnErrorCallback(err)
		return err
	}

	var header Header
	err = binary.Read(
		bytes.NewReader(size[:n]),
		binary.LittleEndian,
		&header,
	)
	if err != nil {
		options.OnErrorCallback(err)
		return err
	}

	if header.Size > 1000000 {
		err := fmt.Errorf("incorrect packet length: %d", header.Size)
		options.OnErrorCallback(err)

		return err
	}

	body.Image = make([]byte, header.Size)
	_, err = io.ReadFull(port, body.Image)
	if err != nil {
		options.OnErrorCallback(err)
		return err
	}

	options.OnDataCallback(body)
	return nil
}

func ReaderDaemon(device string, baud int, options CameraOptions, body *Body) {
	port := OpenCamera(device, baud)
	defer CloseCamera(port)

	for {
		err := CameraReader(port, options, body)
		if err != nil {
			CloseCamera(port)
			time.Sleep(10 * time.Millisecond)
			port = OpenCamera(device, baud)

			continue
		}
	}
}
