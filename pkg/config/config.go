package config

import (
	"time"

	cli "github.com/urfave/cli/v2"
)

type Config struct {
	IP            string
	UDP           int
	TCP           int
	LogLvl        string
	DBPath        string
	CrawlDuration time.Duration
}

var DefaultConfig Config = Config{
	IP:            "0.0.0.0",
	UDP:           9001,
	TCP:           9001,
	LogLvl:        "info",
	DBPath:        "eth_nodes.peerstore",
	CrawlDuration: 1 * time.Hour,
}

func (c *Config) Apply(ctx *cli.Context) {
	// read log-level from the ctx
	if ctx.IsSet("log-level") {
		c.LogLvl = ctx.String("log-level")
	}
	// read the port from the ctx
	if ctx.IsSet("port") {
		port := ctx.Int("port")
		c.UDP = port
		c.TCP = port
	}
	// more args?
}
