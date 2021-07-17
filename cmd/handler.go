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

	a, rtt, err := h.client.Exchange(r, fmt.Sprintf("%s:%s", h.config.UpstreamServer, h.config.UpstreamPort))

	if err != nil {
		log.Printf("failed to communicate with upstream: %s", err)
		return
	}
	
	log.Printf("%s:%s\n", rString, rtt.String())
	log.Printf("%v\nID:%v\n", a.Answer,a.Id)
	w.WriteMsg(a)
}
