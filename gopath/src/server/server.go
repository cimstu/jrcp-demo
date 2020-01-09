package main

import "C"

import (
	"define"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
	"time"
)

type Counter int

func (s *Counter) Run(req *define.Request, res *define.Response) error {
	fmt.Fprintf(os.Stdout, "Call Counter.Run With:%s......\n", req.Word)
	letters := []rune(req.Word)

	res.LetterMap = make(map[string]int)
	for i := 0; i < len(letters); i++ {
		if _, inmap := res.LetterMap[string(letters[i])]; !inmap {
			res.LetterMap[string(letters[i])] = 1
		} else {
			res.LetterMap[string(letters[i])]++
		}
	}

	fmt.Println("Done!")

	return nil
}

func (s *Counter) RunAsync(req *define.Request, res *define.Response) error {
	fmt.Fprintf(os.Stdout, "Call Counter.RunAsync With:%s......\n", req.Word)

	letters := []rune(req.Word)
	res.LetterMap = make(map[string]int)
	for i := 0; i < len(letters); i++ {
		if _, inmap := res.LetterMap[string(letters[i])]; !inmap {
			res.LetterMap[string(letters[i])] = 1
		} else {
			res.LetterMap[string(letters[i])]++
		}
	}
	time.Sleep(time.Second*5)
	fmt.Println("Done!")

	return nil
}

//export StartServer
func StartServer() {
	counter := new(Counter)
	err := rpc.Register(counter)
	if err != nil {
		return
	}

	addr, err := net.ResolveTCPAddr("tcp", ":1056")
	if err != nil {
		return
	}

	listener, err := net.ListenTCP("tcp", addr)

	fmt.Println("RPC Server Started!")
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("a client connected failed:", err)
			continue
		}
		fmt.Println("a client connected")

		jsonrpc.ServeConn(conn)
	}
}

func main()  {
	StartServer()
}
