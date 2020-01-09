package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
	"os"
	"define"
	"strings"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:1056")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	fmt.Print("Please type something:")
	var word string
	fmt.Scanln(&word)

	fmt.Print("Async Call or Sync Call?(y or n)")
	var async string
	fmt.Scanln(&async)
	if strings.Contains(strings.ToLower(async),"y") {
		req := define.Request{word}
		var res define.Response
		//res.LetterMap = make(map[string]int)
		fmt.Println("RPC Calling..")

		call := client.Go("Counter.RunAsync", req, &res, nil)
		if call.Error != nil {
			log.Fatal("Call error:", err)
		}
		fmt.Println("RPC Call Return")

		asyncCall := <- call.Done
		asyncRes := asyncCall.Reply.(*define.Response)

		fmt.Fprintf(os.Stdout, "Here are the letters count of \"%s\":\n", word)
		for k,v := range asyncRes.LetterMap {
			fmt.Fprintf(os.Stdout, "\"%s\":%d times\n", k, v)
		}
	} else {
		req := define.Request{word}
		var res define.Response
		//res.LetterMap = make(map[string]int)
		fmt.Println("RPC Calling..")
		err = client.Call("Counter.Run", req, &res)
		if err != nil {
			log.Fatal("Call error:", err)
		}

		fmt.Println("RPC Call Return")
		fmt.Fprintf(os.Stdout, "Here are the letters count of \"%s\":\n", word)
		for k,v := range res.LetterMap {
			fmt.Fprintf(os.Stdout, "\"%s\":%d times\n", k, v)
		}
	}

	fmt.Print("Press enter to exit")
	var temp string
	fmt.Scanln(&temp)

	return
}
