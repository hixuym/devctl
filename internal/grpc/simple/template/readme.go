package template

var (
	Readme = `# {{title .Alias}} Service

This is the {{title .Alias}} service

Generated with

` + "```" +
		`
devctl grpc --type simple {{.Alias}}
` + "```" + `

## Usage

Generate the proto code

` + "```" +
		`
make proto
` + "```" + `

Run the service

` + "```" +
		`
go run .
` + "```"
)
