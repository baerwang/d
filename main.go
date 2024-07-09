package main

import (
	"net/http"
)

func main() {
	// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o app

	http.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte(`{"hello":"world"}`))
		},
	)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
