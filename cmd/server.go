package cmd

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

type Handler struct {
	client    *dns.Client
	config    Config
	tcpServer *dns.Server
}

func NewHandler(client *dns.Client,config Config) *Handler {
	return &Handler{
		client: client,
		config: config,
	}
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

func (h *Handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	rString := ""

	for _, v := range r.Question {
		rString += v.String()
	}
	
	log.Printf("request: '%s'", rString)
	
	a, rtt, err := h.client.Exchange(r, fmt.Sprintf("%s:%s", h.config.UpstreamServer, h.config.UpstreamPort))
	
	if err != nil {
		log.Printf("failed to communicate with upstream: %s", err)
		return
	}
	
	log.Printf("%s:%s", rString, rtt.String())
	w.WriteMsg(a)
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
