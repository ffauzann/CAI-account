package main

import (
	"github.com/ffauzann/CAI-account/internal/app"
)

var cfg app.Config

func init() {
	cfg.Setup()
}

func main() {
	cfg.StartServer()
}
