package ddd

import (
	"fmt"
	"path"

	tmpl "github.com/hixuym/devctl/internal/grpc/ddd/template"
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
		Dir:      "github.com/hixuym/" + dir,
		Files: []util.File{
			{Path: "api/main.go", Tmpl: tmpl.MainSRV},

			{Path: "application/event/pub/main.go", Tmpl: tmpl.MainSRV},
			{Path: "application/event/sub/main.go", Tmpl: tmpl.MainSRV},
			{Path: "application/service/provider/main.go", Tmpl: tmpl.MainSRV},
			{Path: "application/service/consumer/main.go", Tmpl: tmpl.MainSRV},

			{Path: "domain/" + dir + "/entity/main.go", Tmpl: tmpl.MainSRV},
			{Path: "domain/" + dir + "/event/main.go", Tmpl: tmpl.MainSRV},
			{Path: "domain/" + dir + "/repository/main.go", Tmpl: tmpl.MainSRV},
			{Path: "domain/" + dir + "/service/main.go", Tmpl: tmpl.MainSRV},

			{Path: "infrastructure/config/main.go", Tmpl: tmpl.MainSRV},

			{Path: "client/main.go", Tmpl: tmpl.MainSRV},
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
