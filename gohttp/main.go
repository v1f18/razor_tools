package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	// 1. Define command-line flags. Default port is set to 8000.
	port := flag.String("p", "8000", "Specify the port to listen on")
	flag.Parse()

	// 2. Get the current working directory.
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	// 3. Create a file server handler that serves the current directory.
	fileServer := http.FileServer(http.Dir(currentDir))

	// 4. Get the local IP address for convenience.
	localIP := getLocalIP()

	fmt.Printf("Dir: %s\n", currentDir)
	fmt.Printf("Listen: 0.0.0.0:%s\n", *port)
	fmt.Printf("Local: http://localhost:%s\n", *port)
	if localIP != "" {
		fmt.Printf("LAN: http://%s:%s\n", localIP, *port)
	}
	fmt.Println("Stop: Ctrl+C")

	// 5. Start the HTTP server.
	address := "0.0.0.0:" + *port
	err = http.ListenAndServe(address, fileServer)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// getLocalIP is a helper function to retrieve the local IPv4 address.
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// Check the address type and exclude loopback addresses.
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
