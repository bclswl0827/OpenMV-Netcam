package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

func CameraReader(port io.ReadWriteCloser, options CameraOptions, body *Body) error {
	type Header struct {
		Size uint32
	}
	var header Header

	for {
		frameHeader := make([]byte, 3)
		_, err := port.Read(frameHeader)
		if err != nil {
			options.OnErrorCallback(err)
			return err
		}

		if frameHeader[0] == 0x55 &&
			frameHeader[1] == 0xAA &&
			frameHeader[2] == 0x5A {
			break
		}
	}

	size := make([]byte, 4)
	n, err := io.ReadFull(port, size)
	if err != nil {
		options.OnErrorCallback(err)
		return err
	}

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
		options.OnErrorCallback(
			errors.New(
				"incorrect frame: " +
					fmt.Sprintf("%d", header.Size),
			),
		)
		return errors.New(
			"incorrect frame: " +
				fmt.Sprintf("%d", header.Size),
		)
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
			port = OpenCamera(device, baud)

			continue
		}
	}
}
