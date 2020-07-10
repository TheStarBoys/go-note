package hello2

import (
	"fmt"
	"log"
	"net/rpc"
)

type HelloServiceClient struct {
	*rpc.Client
}

// 约束客户端跟服务器端的RPC函数方法一致
var _ HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	conn, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}

	return &HelloServiceClient{conn}, nil
}

func (h *HelloServiceClient) Hello(request string, reply *string) error {
	return h.Client.Call(HelloServiceName + ".Hello", request, reply)
}

func main() {
	c, err := DialHelloService("tcp", ":1234")
	if err != nil {
		log.Fatal("Dial Error:", err)
	}

	var reply string
	c.Hello("chenhao", &reply)

	fmt.Println(reply)
}