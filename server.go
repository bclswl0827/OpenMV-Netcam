package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func StartServer(options ServerOptions, body *Body) {
	http.HandleFunc(options.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")
		writer := bufio.NewWriterSize(w, len(body.Image))

		for {
			fmt.Fprintf(writer, "--frame\nContent-Type: image/jpeg\nContent-Length: %d\n\n", len(body.Image))
			if body.Ready {
				writer.Write(body.Image)
				writer.Flush()
			}

			time.Sleep(10 * time.Millisecond)
		}
	})

	http.ListenAndServe(
		fmt.Sprintf("%s:%s", options.Listen, options.Port), nil,
	)
}
