package main

import (
	"context"

	"github.com/ludusrusso/ticket-image-gen/pkg/server"
)

func main() {
	server.Run(context.Background())
}
