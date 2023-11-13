package main

import (
	cfg "github.com/GGmaz/BookManager/config"
	"github.com/GGmaz/BookManager/internal/startup"
)

func main() {
	config := cfg.NewConfig()
	server := startup.NewServer(config)
	server.Start()
}
