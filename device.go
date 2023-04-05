package main

import (
	"io"
	"log"

	"github.com/tarm/serial"
)

func OpenCamera(device string, baud int) io.ReadWriteCloser {
	options := &serial.Config{
		Name: device,
		Baud: baud,
	}

	port, err := serial.OpenPort(options)
	if err != nil {
		log.Fatalln(err)
	}

	return port
}

func CloseCamera(port io.ReadWriteCloser) error {
	err := port.Close()
	if err != nil {
		return err
	}

	return nil
}
