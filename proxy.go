package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func startProxy() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
