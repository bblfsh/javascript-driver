package main

import (
	_ "github.com/bblfsh/javascript-driver/driver/impl"
	"github.com/bblfsh/javascript-driver/driver/normalizer"

	"github.com/bblfsh/sdk/v3/driver/server"
)

func main() {
	server.Run(normalizer.Transforms)
}
