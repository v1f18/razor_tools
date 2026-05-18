package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func listenPort(port int, shutdown <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	defer listener.Close()

	go func() {
		<-shutdown
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-shutdown:
				return
			default:
				continue
			}
		}

		remoteAddr := conn.RemoteAddr().String()
		fmt.Printf("[!] Port reachable: %d | Source: %s\n", port, remoteAddr)
		conn.Close()
	}
}

func main() {
	startPort := flag.Int("start", 1000, "start port")
	endPort := flag.Int("end", 2000, "end port")

	flag.Parse()

	if *startPort < 1 || *endPort > 65535 || *startPort > *endPort {
		fmt.Println("[-] Invalid port range. Use a range from 1 to 65535, with start <= end.")
		fmt.Printf("[*] Example: %s -start 1000 -end 2000\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("[*] Starting listeners on ports %d-%d...\n", *startPort, *endPort)

	shutdown := make(chan struct{})
	var wg sync.WaitGroup

	for port := *startPort; port <= *endPort; port++ {
		wg.Add(1)
		go listenPort(port, shutdown, &wg)
	}

	fmt.Printf("[+] Listening in background on ports %d-%d.\n", *startPort, *endPort)
	fmt.Println("[*] Scan or connect to this host from outside.")
	fmt.Println("[*] Press Ctrl+C to stop and release all ports.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\n[-] Shutting down...")
	close(shutdown)
	wg.Wait()
	fmt.Println("[+] Done.")
}
