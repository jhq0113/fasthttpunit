package fasthttpunit

import (
	"bufio"
	"bytes"
	"net"
	"time"

	"github.com/valyala/fasthttp"
)

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

type ReadWriter struct {
	net.Conn

	server *fasthttp.Server
	r      bytes.Buffer
	w      bytes.Buffer
}

func NewReadWriter(server *fasthttp.Server) *ReadWriter {
	rw := &ReadWriter{
		server: server,
	}

	return rw
}

func (rw *ReadWriter) Close() error {
	return nil
}

func (rw *ReadWriter) Read(b []byte) (int, error) {
	return rw.r.Read(b)
}

func (rw *ReadWriter) Write(b []byte) (int, error) {
	return rw.w.Write(b)
}

func (rw *ReadWriter) Request(req *fasthttp.Request, resp *fasthttp.Response) (err error) {
	rw.r.WriteString(req.String())

	err = rw.server.ServeConn(rw)
	if err != nil {
		return err
	}

	br := bufio.NewReader(&rw.w)
	return resp.Read(br)
}

func (rw *ReadWriter) RemoteAddr() net.Addr {
	return zeroTCPAddr
}

func (rw *ReadWriter) LocalAddr() net.Addr {
	return zeroTCPAddr
}

func (rw *ReadWriter) SetDeadline(t time.Time) error {
	return nil
}

func (rw *ReadWriter) SetReadDeadline(t time.Time) error {
	return nil
}

func (rw *ReadWriter) SetWriteDeadline(t time.Time) error {
	return nil
}
