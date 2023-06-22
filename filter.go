package main

import (
	"github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/api"
)

type filter struct {
	callbacks api.FilterCallbackHandler
	path      string
	config    *config
}

func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {
	path, _ := header.Get(":path")
	f.path = path

	if path == "/string" || path == "/byte" {
		f.callbacks.SendLocalReply(301, "", map[string]string{"local-reply": "Set"}, -1, "test-from-go")
		return api.LocalReply
	}
	return api.Continue
}

func (f *filter) DecodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) DecodeTrailers(trailers api.RequestTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) EncodeHeaders(header api.ResponseHeaderMap, endStream bool) api.StatusType {
	if f.path == "/string" {
		location := "/another/string/castle"
		header.Set("location", location)
	}

	if f.path == "/byte" {
		whereTo := []byte("/another/byte/castle")
		location := string(whereTo)
		header.Set("location", location)
	}

	// status, _ := header.Get(":status")
    header.Set("Content-Type", "text/html")
		f.callbacks.SendLocalReply(503, 
			"<b>Bold textM</b>", 
			map[string]string{"Content-Type": "text/html"}, -1, "custom-500-message")
		return api.LocalReply

	return api.Continue
}

func (f *filter) EncodeData(buffer api.BufferInstance, endStream bool) api.StatusType {
	return api.Continue
}

func (f *filter) EncodeTrailers(trailers api.ResponseTrailerMap) api.StatusType {
	return api.Continue
}

func (f *filter) OnDestroy(reason api.DestroyReason) {
}
