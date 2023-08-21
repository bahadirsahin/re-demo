package environ

import (
	log "github.com/sirupsen/logrus"
)

func Init() {
	// initialize logrus
	initLogrus()
}

func initLogrus() {
	log.SetFormatter(
		&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05 MST",
		},
	)
}
