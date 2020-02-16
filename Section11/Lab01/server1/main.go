package main

import (
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

const (
	localhost = ":12345"
)

type (
	cmdServer struct {
		conn net.Conn
	}
)

func main() {
	ln, err := net.Listen("tcp", localhost)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Infof("Command server listening %v", localhost)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}
func handleConnection(c net.Conn) {
	w := &cmdServer{conn: c}
	io.Copy(w, c)
	c.Close()
}
func (wc *cmdServer) Write(p []byte) (n int, err error) {
	if wc == nil {
		return len(p), nil
	}
	io.WriteString(wc.conn, "[From server] ")
	return wc.conn.Write(p)
}
