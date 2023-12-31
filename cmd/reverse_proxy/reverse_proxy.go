package proxy

import (
	"context"
	"crypto/tls"
	proxy "dhens/drawbridge/cmd/reverse_proxy/ca"
	"io"
	"log"
	"net"
	"time"
)

func TestSetupTCPListener(ca *proxy.CA) {
	log.Printf("Spinning up TCP Listener on localhost:25565")
	l, err := tls.Listen("tcp", "localhost:25565", ca.ServerTLSConfig)
	if err != nil {
		log.Fatalf("TCP Listen failed: %s", err)
	}

	defer l.Close()
	for {
		// wait for connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("TCP Accept failed: %s", err)
		}
		// Handle new connection in a new go routine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(clientConn net.Conn) {
			var d net.Dialer
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			resourceConn, err := d.DialContext(ctx, "tcp", "localhost:25566")
			if err != nil {
				log.Fatalf("Failed to dial: %v", err)
			}

			log.Printf("TCP Accept from: %s\n", clientConn.RemoteAddr())
			// Copy data back and from client and server.
			go io.Copy(resourceConn, clientConn)
			io.Copy(clientConn, resourceConn)
			// Shut down the connection.
			clientConn.Close()
		}(conn)
	}
}
