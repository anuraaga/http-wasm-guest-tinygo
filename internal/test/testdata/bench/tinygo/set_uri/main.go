package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleFn = setURI
}

func setURI(req api.Request, resp api.Response, next api.Next) {
	req.SetURI("/v1.0/hello")
}