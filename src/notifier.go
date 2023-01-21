package main

import (
	"net/url"
	"net/http"
)

func httpSend() *http.Response {
	resp, _ := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
	return resp
}
