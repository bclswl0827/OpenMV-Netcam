package main

import "log"

func main() {
	var (
		conf Config
		body Body
		args Args
	)
	args.Read()
	conf.Read(args.Path)

	go ReaderDaemon(
		conf.Device, conf.Baud,
		CameraOptions{
			Keyword: conf.Keyword,
			OnErrorCallback: func(err error) {
				log.Println(err)
			},
			OnDataCallback: func(b *Body) {
				log.Printf("received packet length: %d\n", len(b.Image))
			},
		}, &body)

	StartServer(ServerOptions{
		Listen: conf.Listen,
		Port:   conf.Port,
		Path:   conf.Path,
	}, &body)

}
