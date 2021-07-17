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
	if ! h.config.EnableTCP {
		log.Fatalf("TCP server is not enabled in config")
	}

	
}