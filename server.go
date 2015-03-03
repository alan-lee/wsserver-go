package websocket

import (
	"fmt"
	"net"
)

const (
	DefaultListenPort = 8087
)

type Server struct {
	address string
}

func (s *Server) Serve() error {
	addr := s.Address
	if s.address == nil || s.address == "" {
		addr = fmt.Sprintf(":%d", DefaultListenPort)
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		conn, err := ln.Accept()
		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				fmt.Printf("http: Accept error: %v; retrying in %v", e, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return e
		}

		tempDelay = 0
		c, err = s.newClient(conn)
		if err != nil {
			continue
		}

		go c.serve()
	}
}

func (s *Server) newClient(conn *net.Conn) (c *client, err error) {
	c = new(client)
}

type client struct {
	remoteAddr string
	s          *Server
}
