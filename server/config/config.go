package config

import "os"

var (
	Port        = os.Getenv("PORT")
	KafkaBroker = os.Getenv("KafkaBroker")
)
