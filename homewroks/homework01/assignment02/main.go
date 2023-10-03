package main

import (
	"CRT-HAO/CS3004302/homework01/iperfer/iperfer"
	"flag"
	"fmt"
	"os"
)

type Mode int

const (
	ServerMode Mode = 1
	ClientMode Mode = 2
)

var (
	mode Mode
	port uint16

	// client
	host string
	time int

	// temp
	serverMode bool
	clientMode bool
	port_uint  uint
)

func init() {
	flag.BoolVar(&serverMode, "s", false, "server mode")
	flag.BoolVar(&clientMode, "c", false, "client mode")

	flag.UintVar(&port_uint, "p", 0, "port to connect or listen on")

	flag.StringVar(&host, "h", "", "host address to connect")
	flag.IntVar(&time, "t", 0, "duration in seconds for which data should be generated")
}

func parseFlags() {
	flag.Parse()

	if serverMode {
		mode = ServerMode
	} else if clientMode {
		mode = ClientMode
	}

	if port_uint != 0 && (port_uint < 1024 || port_uint > 65535) {
		fmt.Println("Error: port number must be in the range 1024 to 65535")
		os.Exit(1)
	}
	port = uint16(port_uint)
}

func main() {
	parseFlags()

	if mode == ServerMode {
		if port == 0 {
			fmt.Println("Error: missing or additional arguments")
			os.Exit(1)
		}

		server, err := iperfer.NewServer(port)
		if err != nil {
			panic(err)
		}

		n, mbps, err := server.Receive()
		if err != nil {
			panic(err)
		}

		fmt.Printf("received=%d KB rate=%f Mbps", n/1000, mbps)
	} else if mode == ClientMode {
		if host == "" || port == 0 || time == 0 {
			fmt.Println("Error: missing or additional arguments")
			os.Exit(1)
		}

		client, err := iperfer.NewClient(host, port)
		if err != nil {
			panic(err)
		}

		n, mbps, err := client.Send(time)
		if err != nil {
			panic(err)
		}

		fmt.Printf("sent=%d KB rate=%f Mbps", n/1000, mbps)
	} else {
		fmt.Println("Error: unknown mode")
		os.Exit(1)
	}
}
