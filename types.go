package main

type Args struct {
	Path string
}

type Body struct {
	Image []byte
}

type Header struct {
	Size uint32
}

type Config struct {
	Listen  string `json:"listen"`
	Port    string `json:"port"`
	Path    string `json:"path"`
	Device  string `json:"device"`
	Keyword string `json:"keyword"`
	Baud    int    `json:"baud"`
}

type ServerOptions struct {
	Listen string
	Port   string
	Path   string
}

type CameraOptions struct {
	Keyword         string
	OnErrorCallback func(error)
	OnDataCallback  func(*Body)
}
