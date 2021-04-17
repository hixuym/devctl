package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/xlab/treeprint"
)

func ProtoComments(alias string) []string {
	return []string{
		"\ndownload protoc zip packages and install:\n",
		"visit https://github.com/protocolbuffers/protobuf/releases",
		"go get -u github.com/golang/protobuf/proto",
		"go get -u github.com/golang/protobuf/protoc-gen-go",
		"\ncompile the proto file " + alias + ".proto:\n",
		"cd " + alias,
		"make proto\n",
	}
}

type Config struct {
	// foo
	Alias string
	// github.com/hixuym/foo
	Dir string
	// Files
	Files []File
	// Comments
	Comments []string
}

type File struct {
	Path string
	Tmpl string
}

func write(c Config, file, tmpl string) error {
	fn := template.FuncMap{
		"title": func(s string) string {
			return strings.ReplaceAll(strings.Title(s), "-", "")
		},
		"dehyphen": func(s string) string {
			return strings.ReplaceAll(s, "-", "")
		},
		"lower": func(s string) string {
			return strings.ToLower(s)
		},
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	t, err := template.New("f").Funcs(fn).Parse(tmpl)
	if err != nil {
		return err
	}

	return t.Execute(f, c)
}

func Create(c Config) error {
	// check if dir exists
	if _, err := os.Stat(c.Alias); !os.IsNotExist(err) {
		return fmt.Errorf("%s already exists", c.Alias)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	fmt.Printf("Creating service %s\n\n", c.Alias)

	t := treeprint.New()

	// write the files
	for _, file := range c.Files {
		f := filepath.Join(c.Alias, file.Path)
		dir := filepath.Dir(f)

		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return err
			}
		}

		addFileToTree(t, file.Path)
		if err := write(c, f, file.Tmpl); err != nil {
			return err
		}
	}

	// print tree
	fmt.Println(t.String())

	for _, comment := range c.Comments {
		fmt.Println(comment)
	}

	// just wait
	<-time.After(time.Millisecond * 250)

	return nil
}

func addFileToTree(root treeprint.Tree, file string) {
	split := strings.Split(file, "/")
	curr := root
	for i := 0; i < len(split)-1; i++ {
		n := curr.FindByValue(split[i])
		if n != nil {
			curr = n
		} else {
			curr = curr.AddBranch(split[i])
		}
	}
	if curr.FindByValue(split[len(split)-1]) == nil {
		curr.AddNode(split[len(split)-1])
	}
}
