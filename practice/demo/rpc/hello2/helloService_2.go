package hello2

import (
	"log"
	"net"
	"net/rpc"
)

const (
	HelloServiceName = "path/to/pkg.HelloService"
)

type HelloServiceInterface = interface{
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServiceInterface) {
	err := rpc.RegisterName(HelloServiceName, srv)
	if err != nil {
		log.Fatal("Register HelloService Error:", err)
	}
}

type HelloService struct {}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello: " + request
	return nil
}

func main() {
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen Error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept Failed conn:", err)
		}

		go rpc.ServeConn(conn)
	}
}