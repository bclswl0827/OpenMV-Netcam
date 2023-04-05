package main

import "log"

func main() {
	var (
		conf Config
		body Body
		args Args
	)
	args.ReadFlags()
	conf.ReadConfig(args.Path)

	go ReaderDaemon(
		conf.Device, conf.Baud,
		CameraOptions{
			OnErrorCallback: func(err error) {
				log.Println(err)
			},
			OnDataCallback: func(b *Body) {
				log.Println("received ", len(b.Image))
			},
		}, &body)

	StartServer(ServerOptions{
		Listen: conf.Listen,
		Port:   conf.Port,
		Path:   conf.Path,
	}, &body)

}
