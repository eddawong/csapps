package csapps

import (
	"net/url"
	"net/textproto"
	"io"
	"fmt"
	"bytes"
	"sync"

)

type Request struct {
	Method        string
	URL           *url.URL
	Proto         string
	ProtoMajor    int
	ProtoMinor    int
	Header        textproto.MIMEHeader
	ContentLength int64
	Body          io.ReadCloser
	buffer        *bytes.Buffer
}

func (req *Request)fillbuf(){
	req.buffer = &bytes.Buffer{}
	req.buffer.WriteString(fmt.Sprintf("%s %s %s/%d.%d\r\n", req.Method, req.URL, req.Proto, req.ProtoMajor, req.ProtoMinor))
	for k, v := range req.Header {
		for _, value := range v {
			req.buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, value))
		}
	}
	req.buffer.WriteString("\r\n")
	io.Copy(req.buffer,req.Body)
}

func (req *Request) String() string {
	if req.buffer==nil{
		req.fillbuf()
	}
	return req.buffer.String()
}


func (req *Request)Read(p []byte)(n int, err error){
	if req.buffer==nil{
		req.fillbuf()
	}
	return req.buffer.Read(p)
}

func (req *Request)WriteTo(w io.Writer)(n int64, err error){
	if req.buffer==nil{
		req.fillbuf()
	}
	return req.buffer.WriteTo(w)
}

func (req *Request)Len()(n int){
	return len(req.String())
}


type body struct {
	src    io.Reader
	lk     sync.Mutex
	closed bool
}

type noBody struct{}

func (noBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (noBody) Close() error                     { return nil }
func (noBody) WriteTo(io.Writer) (int64, error) { return 0, nil }
