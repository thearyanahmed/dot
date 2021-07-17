package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	"github.com/miekg/dns"
	"github.com/thearyanahmed/dot/cmd"
)

func main() {
	log.Printf("starting dot")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)

	cfg := cmd.Config{}
	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("failed to parse env. \nmsg:%v\nexiting", err.Error())
	}

	fmt.Printf("cfg:%v\n", cfg)

	c := new(dns.Client)
	c.Net = "tcp-tls"
	c.Dialer = &net.Dialer{
		Timeout: cfg.UpstreamTimeout,
	}

	h := cmd.NewHandler(c, cfg)

	log.Printf("starting servers")
	cmd.StartServers(cfg)

	log.Printf("setting up dns handler")
	dns.Handle(".", h)

	// handle signals
	sig := <-sigChan

	log.Printf("signal termianted. msg:%v\n", sig.String())
	cmd.ShutdownServers()
}
