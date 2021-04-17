// Package grpc generates grpc service templates
package grpc

import (
	"github.com/urfave/cli/v2"

	"github.com/hixuym/devctl/cmd"
	"github.com/hixuym/devctl/internal/grpc/ddd"
	"github.com/hixuym/devctl/internal/grpc/simple"
)

func Run(ctx *cli.Context) error {
	srvType := ctx.String("type")
	if srvType == "simple" {
		return simple.Run(ctx)
	}
	return ddd.Run(ctx)
}

func init() {
	cmd.Register(&cli.Command{
		Name:        "grpc",
		Usage:       "Create a grpc service template",
		Description: `'devctl grpc' scaffolds a new service skeleton. Example: 'devctl grpc --type simple helloworld && cd helloworld'`,
		Action:      Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "type",
				Value: "simple",
			},
		},
	})
}
