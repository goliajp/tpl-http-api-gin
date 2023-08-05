package main

import (
	"github.com/goliajp/http-api-gin/data"
	"github.com/goliajp/http-api-gin/server"
	"os"
)

var signal = make(chan os.Signal)

func main() {
	server.RunCleaner(signal)    // goroutine
	server.RunHttpServer(signal) // block main thread
}

func init() {
	// data preparation
	data.PgPrepare()
	data.KvPrepare()
}
