package main

import (
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

	c := new(dns.Client)
	c.Net = "tcp-tls"
	c.Dialer = &net.Dialer{
		Timeout: cfg.UpstreamTimeout,
	}

	h := cmd.NewHandler(c,cfg)

	h.StartServer()

	dns.Handle(".",h)

	// handle signals
	sig := <-sigChan

	log.Printf("signal termianted. msg:%v\n", sig.String())
	h.Shutdown()
}
