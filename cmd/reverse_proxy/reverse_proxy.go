package proxy

import (
	"context"
	proxy "dhens/drawbridge/cmd/reverse_proxy/ca"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetUpReverseProxy() {
	ca := &proxy.CA{}
	err := ca.SetupRootCA()
	if err != nil {
		log.Fatalf("Error setting up root CA: %s", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", myHandler)
	server := http.Server{
		TLSConfig: ca.ServerTLSConfig,
		Addr:      "localhost:4443",
		Handler:   r,
	}
	log.Printf("Listening Drawbridge reverse rpoxy at %s", server.Addr)

	go func() {
		log.Fatal(server.ListenAndServeTLS("", ""))
	}()

	ca.MakeClientRequest(fmt.Sprintf("https://%s", server.Addr))
}

func TestSetupTCPListener() {
	log.Printf("Spinning up TCP Listener on localhost:25565")
	l, err := net.Listen("tcp", "localhost:25565")
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
			// clientConn.Close()
		}(conn)
	}
}

func myHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("New request from %s", req.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "success!")
}
