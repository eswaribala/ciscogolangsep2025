package models

type KafkaConfig struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
	GroupID string   `json:"group_id"`
	Enabled bool     `json:"enabled"`
	Port    int      `json:"port"`
}
