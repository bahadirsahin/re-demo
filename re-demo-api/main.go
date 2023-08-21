package main

import (
	"re-demo-api/api"
	"re-demo-api/environ"
)

func init() {
	// initialize environment
	environ.Init()
}

func main() {
	// initialize api
	api.Init()
}
