package main

import "github.com/http-wasm/http-wasm-guest-tinygo/handler" //nolint

func main() {
	handler.HandleFn = rewrite
}

func rewrite() {
	if handler.GetURI() == "/v1.0/hi?name=panda" {
		handler.SetURI("/v1.0/hello?name=teddy")
	}
	handler.Next()
}
