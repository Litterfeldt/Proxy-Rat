package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
)

func startProxy() {
	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true

	proxy.OnRequest().DoFunc(
		func(r *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			r.Header.Set("Host", r.URL.Host)
			return r, nil
		})

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
