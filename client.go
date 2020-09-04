package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	ep := flag.String("ep", "https://h2c-server.default.svc.cluster.local:80", "endpoint url ")

	client := http.Client{
		Transport: &http2.Transport{
			// Pretend we are dialing a TLS endpoint.
			// Note, we ignore the passed tls.Config
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

	times := 100
	for times > 0 {
		fmt.Println("starting send req")

		resp, err := client.Get(*ep)
		if err != nil {
			log.Fatal(err)
		}
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		fmt.Printf("get server resp : %d\n", buf.String())
		time.Sleep(time.Second * 1)
		times--
	}
	fmt.Println("end¬")
}
