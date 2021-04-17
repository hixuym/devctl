package template

var (
	Module = `module {{.Dir}}

go 1.16

require (
	google.golang.org/grpc v1.37.0
)

`
)
