package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// UnexpectedSubcommand checks for erroneous subcommands and prints help and returns error
func UnexpectedSubcommand(ctx *cli.Context) error {
	if first := Subcommand(ctx); first != "" {
		// received something that isn't a subcommand
		return cli.Exit(fmt.Sprintf("Unrecognized subcommand for %s: %s. Please refer to '%s --help'", ctx.App.Name, first, ctx.App.Name), 1)
	}
	return cli.ShowSubcommandHelp(ctx)
}

func UnexpectedCommand(ctx *cli.Context) error {
	commandName := ctx.Args().First()
	return cli.Exit(fmt.Sprintf("Unrecognized devctl command: %s. Please refer to 'devctl --help'", commandName), 1)
}

func MissingCommand(ctx *cli.Context) error {
	return cli.Exit(fmt.Sprintf("No command provided to devctl. Please refer to 'devctl --help'"), 1)
}

// MicroSubcommand returns the subcommand name
func Subcommand(ctx *cli.Context) string {
	return ctx.Args().First()
}
