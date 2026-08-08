package qdrant

// Config define the config of qdrant client.
type Config struct {
	Host   string
	Port   int64
	APIKey string
	UseTLS bool
}
