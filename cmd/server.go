package cmd

import (
	"log"

	"github.com/miekg/dns"
)

var tcpServer *dns.Server

type Handler struct {
	client *dns.Client
	config Config
}

func (h *Handler) StartServer() {
	if !h.config.EnableTCP {
		log.Fatalf("TCP server is not enabled in config")
	}

	tcpServer = &dns.Server{
		Addr: ":853",
	}

	go tcpServer.ListenAndServe()
	log.Printf("Started TCP server on port :853\n")
}


