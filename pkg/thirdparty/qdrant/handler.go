package qdrant

import "github.com/qdrant/go-client/qdrant"

var _ IHandler = Handler{}

type Handler struct {
	cli *qdrant.Client
}

// New initialize a new qdrant Handler.
func New(conf *Config) (*Handler, error) {
	cli, err := qdrant.NewClient(&qdrant.Config{
		Host:   conf.Host,
		Port:   conf.Port,
		APIKey: conf.APIKey,
		UseTLS: conf.UseTLS,
	})

	return &Handler{cli: cli}, nil
}
