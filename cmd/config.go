package cmd

import "time"

type Config struct {
	UpstreamTimeout time.Duration `env:"UPSTREAM_TIMEOUT" envDefault:"2000ms"`
	UpstreamServer  string        `env:"UPSTREAM_SERVER" envDefault:"1.1.1.1"`
	UpstreamPort    string        `env:"UPSTREAM_TIMEOUT" envDefault:"853"`
	EnableTCP       bool          `env:"UPSTREAM_TIMEOUT" envDefault:"true"`
}
