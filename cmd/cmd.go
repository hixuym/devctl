package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/hixuym/devctl/plugin"
	"github.com/urfave/cli/v2"
)

func init() {
}

type Cmd interface {
	// Init initialises options
	// Note: Use Run to parse command line
	Init(opts ...Option) error
	// Options set within this command
	Options() Options
	// The cli app within this cmd
	App() *cli.App
	// Run executes the command
	Run() error
	// Implementation
	String() string
}

type command struct {
	opts Options
	app  *cli.App

	before cli.ActionFunc
}

var (
	DefaultCmd  Cmd = New()
	name            = "devctl"
	description     = "a golang microservice development tool"

	defaultFlags = []cli.Flag{}
)

func action(c *cli.Context) error {
	if c.Args().Len() > 0 {
		// if an executable is available with the name of
		// the command, execute it with the arguments from
		// index 1 on.
		v, err := exec.LookPath("devctl-" + c.Args().First())
		if err == nil {
			ce := exec.Command(v, c.Args().Slice()[1:]...)
			ce.Stdout = os.Stdout
			ce.Stderr = os.Stderr
			return ce.Run()
		}

		return UnexpectedCommand(c)

	}

	return MissingCommand(c)
}

func New(opts ...Option) *command {
	options := Options{}
	for _, o := range opts {
		o(&options)
	}

	cmd := new(command)
	cmd.opts = options
	cmd.app = cli.NewApp()
	cmd.app.Name = name
	cmd.app.Version = "1.0.0"
	cmd.app.Usage = description
	cmd.app.Flags = defaultFlags
	cmd.app.Action = action
	cmd.app.Before = beforeFromContext(options.Context, cmd.Before)

	// if this option has been set, we're running a service
	// and no action needs to be performed. The CMD package
	// is just being used to parse flags and configure micro.
	if setupOnlyFromContext(options.Context) {
		cmd.app.Action = func(ctx *cli.Context) error { return nil }
	}

	//flags to add
	if len(options.Flags) > 0 {
		cmd.app.Flags = append(cmd.app.Flags, options.Flags...)
	}
	//action to replace
	if options.Action != nil {
		cmd.app.Action = options.Action
	}

	// cmd to add to use registry
	return cmd
}

func (c *command) App() *cli.App {
	return c.app
}

func (c *command) Options() Options {
	return c.opts
}

// Before is executed before any subcommand
func (c *command) Before(ctx *cli.Context) error {
	// initialize plugins
	for _, p := range plugin.Plugins() {
		if err := p.Init(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (c *command) Init(opts ...Option) error {
	for _, o := range opts {
		o(&c.opts)
	}
	if len(c.opts.Name) > 0 {
		c.app.Name = c.opts.Name
	}
	if len(c.opts.Version) > 0 {
		c.app.Version = c.opts.Version
	}
	c.app.HideVersion = len(c.opts.Version) == 0
	c.app.Usage = c.opts.Description

	//allow user's flags to add
	if len(c.opts.Flags) > 0 {
		c.app.Flags = append(c.app.Flags, c.opts.Flags...)
	}
	//action to replace
	if c.opts.Action != nil {
		c.app.Action = c.opts.Action
	}

	return nil
}

func (c *command) Run() error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v", r)
			panic(r)
		}
	}()
	return c.app.Run(os.Args)
}

func (c *command) String() string {
	return name
}

// Register CLI commands
func Register(cmds ...*cli.Command) {
	app := DefaultCmd.App()
	app.Commands = append(app.Commands, cmds...)
}

// Run the default command
func Run() {
	if err := DefaultCmd.Run(); err != nil {
		log.Printf("Cmd run error: %v", err)
		os.Exit(1)
	}
}
