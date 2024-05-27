package main

import (
	"memento/cmd/memento/mock"
	"memento/config"
	"memento/internal/api"
)

func main() {
	cfg := config.MustLoad("config/config.yaml")
	var _ = cfg
	storage := mock.MockStorage{}

	server := api.New(&storage)
	server.MustRun()
}
