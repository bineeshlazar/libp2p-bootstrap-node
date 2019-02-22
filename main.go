package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bnsh12/p2pnet"
)

func main() {
	help := flag.Bool("help", false, "Display Help")
	listenHost := flag.String("host", "0.0.0.0", "The bootstrap node host listen address\n")
	port := flag.Int("port", 4001, "The bootstrap node listen port")
	flag.Parse()

	if *help {
		fmt.Printf("This is a simple bootstrap node for kad-dht application using libp2p\n\n")
		fmt.Printf("Usage: \n   Run './bootnode'\nor Run './bootnode -host [host] -port [port]'\n")

		os.Exit(0)
	}

	cfg := &p2pnet.Config{}

	cfg.KeyFile = ".key.dat"
	cfg.ListenHost = *listenHost
	cfg.ListenPort = *port
	cfg.RendezvousString = "bootstrapper"
	cfg.EnableRelay = true

	p2p, err := p2pnet.NewNetwork(cfg)
	if err != nil {
		log.Fatalf("Initialization failed (%s)", err)
	}

	defer p2p.Close()

	for _, addr := range p2p.Addrs() {
		log.Printf("Listening on : %s", addr.String())
	}

	select {}
}
