package main

import (
	"errors"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Args struct {
	A, B int
}

type Math int

func (m *Math) Mutiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

type Quotient struct {
	Quo, Rem int
}

func (m *Math) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Dicide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

// func main() {
// 	math := new(Math)
// 	rpc.Register(math)
// 	// rpc.HandleHTTP()
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
// 	if err != nil {
// 		fmt.Println("Fatal Error : ", err)
// 		os.Exit(2)
// 	}

// 	// err := http.ListenAndServe(":1234", nil)
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	if err != nil {
// 		fmt.Println("Fatal Error : ", err)
// 		os.Exit(3)
// 	}

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			fmt.Println("conn error : ", err)
// 			continue
// 		}
// 		// rpc.ServeConn(conn)
// 		jsonrpc.ServeConn(conn)
// 	}
// }

func main() {
	math := new(Math)
	err := rpc.Register(math)
	if err != nil {
		panic(err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	listenner, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listenner.Accept()
		if err != nil {
			panic(err)
		}
		// rpc.ServeConn(conn)
		jsonrpc.ServeConn(conn)
	}
}
