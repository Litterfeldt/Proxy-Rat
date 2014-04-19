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
			//r.URL.Host = r.Header.Get("Proxy-Host")
			//r.Host = r.Header.Get("Proxy-Host")
			r.Header.Set("Proxy-Host", r.Host)
			r.URL.Host = "proxy.litterfeldt.com"
			r.Host = "proxy.litterfeldt.com"
			return r, nil
		})

	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}
