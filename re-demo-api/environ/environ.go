package environ

import (
	"os"
	"re-demo-api/structs"
	"re-demo-api/vars"

	log "github.com/sirupsen/logrus"
)

func Init() {
	// init env vars
	initEnv()

	// initialize logrus
	initLogrus()
}

// initialize environment variables
func initEnv() {
	vars.Env = structs.Env{
		ApiToken: os.Getenv("API_TOKEN"),
	}
}

// initialize logger
func initLogrus() {
	log.SetFormatter(
		&log.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05 MST",
		},
	)
}
