package template

var (
	ProtoSRV = `syntax = "proto3";

package {{dehyphen .Alias}};

option go_package = "./proto;{{dehyphen .Alias}}";

// The greeting service definition.
service {{title .Alias}} {
	// Sends a greeting
  	rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
`
)
