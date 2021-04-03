// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Source: greeter.proto, user.proto

package client

import (
	fmt "fmt"
	reflect "reflect"
	strings "strings"

	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	pb "github.com/erda-project/erda-infra/examples/protocol/pb"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	grpc1 "google.golang.org/grpc"
)

var dependencies = []string{
	"grpc-client@erda.infra.example",
	"grpc-client",
}

type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType              = reflect.TypeOf((*Client)(nil)).Elem()
	greeterServiceClientType = reflect.TypeOf((*pb.GreeterServiceClient)(nil)).Elem()
	greeterServiceServerType = reflect.TypeOf((*pb.GreeterServiceServer)(nil)).Elem()
	userServiceClientType    = reflect.TypeOf((*pb.UserServiceClient)(nil)).Elem()
	userServiceServerType    = reflect.TypeOf((*pb.UserServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.infra.example-client":
		return p.client
	case "erda.infra.example.GreeterService":
		return &greeterServiceWrapper{client: p.client.GreeterService(), opts: opts}
	case "erda.infra.example.GreeterService.client":
		return p.client.GreeterService()
	case "erda.infra.example.UserService":
		return &userServiceWrapper{client: p.client.UserService(), opts: opts}
	case "erda.infra.example.UserService.client":
		return p.client.UserService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case greeterServiceClientType:
		return p.client.GreeterService()
	case greeterServiceServerType:
		return &greeterServiceWrapper{client: p.client.GreeterService(), opts: opts}
	case userServiceClientType:
		return p.client.UserService()
	case userServiceServerType:
		return &userServiceWrapper{client: p.client.UserService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.infra.example-client", &servicehub.Spec{
		Services: []string{
			"erda.infra.example.GreeterService",
			"erda.infra.example.UserService",
			"erda.infra.example-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			greeterServiceClientType,
			userServiceClientType,
			// server types
			greeterServiceServerType,
			userServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
