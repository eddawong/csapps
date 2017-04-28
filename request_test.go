package csapps

import (
	"testing"
	"net/url"
	"io"
	"io/ioutil"
)

var URL, _ = url.Parse("rtsp://10.47.214.112:554/hahaha")
var req = &Request{
	Method: "DESCRIBE",
	URL:    URL,
	Body:   noBody{},
	Proto: "RTSP",
	ProtoMajor: 1,
	ProtoMinor: 0,
}

func BenchmarkRequest_String(b *testing.B) {
	for i:=0;i<b.N;i++{
		req.String()
	}
}

func BenchmarkRequest_Read(b *testing.B) {
	for i:=0;i<b.N;i++{
		io.Copy(ioutil.Discard,req)
	}
}