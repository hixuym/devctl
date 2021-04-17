package main

import (
	"github.com/hixuym/devctl/cmd"

	_ "github.com/hixuym/devctl/internal/grpc"
)

func main() {
	cmd.Run()
}
