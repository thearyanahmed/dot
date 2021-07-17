package cmd

import (
	"log"

	"github.com/miekg/dns"
)

var (
	tcpServer *dns.Server
	udpServer *dns.Server
)

func StartServers(config Config) {
	if !config.EnableTCP && !config.EnableUDP {
		log.Fatalf("tcp server is not enabled in config")
	}

	if config.EnableTCP {
		tcpServer = &dns.Server{
			Addr: ":53",
			Net:  "tcp",
		}

		go tcpServer.ListenAndServe()
		log.Printf("listening to tcp/53\n")
	}

	if config.EnableUDP {
		tcpServer = &dns.Server{
			Addr: ":53",
			Net:  "udp",
		}

		go tcpServer.ListenAndServe()
		log.Printf("listening to UDP/53\n")
	}
}

func ShutdownServers() {
	shutdown(tcpServer)
	shutdown(udpServer)
}

func shutdown(server *dns.Server) {
	if server == nil {
		return
	}

	if err := server.Shutdown(); err != nil {
		log.Printf("could not shut down server. \nnet:%v\nmsg:%v\n", server.Net, err.Error())
	}
}
