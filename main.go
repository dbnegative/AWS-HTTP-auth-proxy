package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"flag"

	"github.com/smartystreets/go-aws-auth"
)

var endpoint = flag.String("endpoint_url", "localhost", "the remote endpoint to proxy to")
var scheme = flag.String("scheme", "https", "the http protocol scheme to use")
var port = flag.String("local_port", "8080", "the localhost port to listen on")

func init() {
	flag.Parse()
}

func main() {
	proxy := NewSingleHostReverseProxy(&url.URL{
		Scheme: *scheme,
		Host:   *endpoint})
	log.Fatal(http.ListenAndServe(":"+*port, proxy))
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		awsauth.Sign(req)
	}
	return &httputil.ReverseProxy{Director: director}
}
