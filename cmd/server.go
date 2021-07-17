package cmd

import (
	"log"

	"github.com/miekg/dns"
)

type Handler struct {
	client    *dns.Client
	config    Config
	tcpServer *dns.Server
}

func (h *Handler) StartServer() {
	if !h.config.EnableTCP {
		log.Fatalf("tcp server is not enabled in config")
	}

	h.tcpServer = &dns.Server{
		Addr: ":853",
	}

	go h.tcpServer.ListenAndServe()
	log.Printf("started TCP server on port :853\n")
}




func (h *Handler) Shutdown() {
	if h.tcpServer == nil {
		log.Printf("attempted to shutdown server while it is not running")
		return
	}

	if err := h.tcpServer.Shutdown(); err != nil {
		log.Printf("could not shut down server. \nnet:%v\nmsg:%v\n", h.tcpServer.Net, err.Error())
	}
}
