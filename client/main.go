package main

import (
	"net"
)

func main(){
    //println("hello from client")
    serverAddress := "localhost:8080"
    connection, err := net.Dial("tcp", serverAddress)
    if err != nil {
        panic(err)
    }
    
    msg := "hello from client"
    connection.Write([]byte(msg))
}
