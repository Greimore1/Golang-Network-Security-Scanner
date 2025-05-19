package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func scanPort(ip string, port int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return
	}
	conn.Close()
	results <- fmt.Sprintf("Open: %s", address)
}

func scanHost(ip string, ports []int) {
	var wg sync.WaitGroup
	results := make(chan string, 100)

	for _, port := range ports {
		wg.Add(1)
		go scanPort(ip, port, &wg, results)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}

func parsePorts(portRange string) ([]int, error) {
	var ports []int
	parts := strings.Split(portRange, "-")

	if len(parts) == 1 {
		port, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		return []int{port}, nil
	}

	start, err1 := strconv.Atoi(parts[0])
	end, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil || start > end {
		return nil, fmt.Errorf("invalid port range")
	}

	for i := start; i <= end; i++ {
		ports = append(ports, i)
	}
	return ports, nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <IP> <port or port-range (e.g., 80 or 20-100)>")
		return
	}

	ip := os.Args[1]
	portInput := os.Args[2]

	ports, err := parsePorts(portInput)
	if err != nil {
		fmt.Println("Invalid port input:", err)
		return
	}

	fmt.Printf("Scanning %s on ports %v...\n", ip, ports)
	scanHost(ip, ports)
}
