package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net"
	"net/http"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalf("Oops: " + err.Error() + "\n")
	}

	h2s := &http2.Server{}
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					fmt.Fprintf(w, "My IP is  %v", ipnet.IP.String())
				}
			}
		}
	})
	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: h2c.NewHandler(handler, h2s),
	}
	fmt.Printf("Listening [0.0.0.0:80]...\n")
	server.ListenAndServe()
}
