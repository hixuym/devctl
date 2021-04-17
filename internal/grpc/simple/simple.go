package simple

import (
	"fmt"
	"path"

	tmpl "github.com/hixuym/devctl/internal/grpc/simple/template"
	"github.com/hixuym/devctl/internal/util"
	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	dir := ctx.Args().First()

	if len(dir) == 0 {
		fmt.Println("specify service name")
		return nil
	}

	// check if the path is absolute, we don't want this
	// we want to a relative path so we can install in GOPATH
	if path.IsAbs(dir) {
		fmt.Println("require relative path as service will be installed in GOPATH")
		return nil
	}

	c := util.Config{
		Alias:    dir,
		Comments: util.ProtoComments(dir),
		Dir:      dir,
		Files: []util.File{
			{Path: "server/main.go", Tmpl: tmpl.MainSRV},
			{Path: "client/main.go", Tmpl: tmpl.ClientSRV},
			{Path: "proto/" + dir + ".proto", Tmpl: tmpl.ProtoSRV},
			{Path: "generate.go", Tmpl: tmpl.GenerateFile},
			{Path: "Dockerfile", Tmpl: tmpl.DockerSRV},
			{Path: "Makefile", Tmpl: tmpl.Makefile},
			{Path: "README.md", Tmpl: tmpl.Readme},
			{Path: ".gitignore", Tmpl: tmpl.GitIgnore},
			{Path: "go.mod", Tmpl: tmpl.Module},
		},
	}

	// create the files
	return util.Create(c)
}
