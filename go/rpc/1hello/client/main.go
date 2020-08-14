package main

import (
	"fmt"
	hello "go_test/rpc/1hello"
	"log"
)

func main() {
	client, err := hello.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
