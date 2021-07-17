package cmd

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

type Handler struct {
	client *dns.Client
	config Config
}

func NewHandler(client *dns.Client, config Config) *Handler {
	return &Handler{
		client: client,
		config: config,
	}
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
