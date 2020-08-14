package hello

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

type HelloServiceClient struct {
	*rpc.Client
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	return &HelloServiceClient{Client: client}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}
