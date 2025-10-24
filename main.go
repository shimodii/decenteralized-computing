package main

import "net"

func main(){
    // saving clients ip
    clients := make([]string, 0, 2)
    
    // create server
    server, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    defer server.Close()
    println("server is up on 8080")

    for {
        client, err := server.Accept()
        if err != nil {
            panic(err)
        }
        clients = append(clients, client.RemoteAddr().String())
        if (len(clients) == 2){
            println("yay clients are full")
        } else if (len(clients) < 2) {
            println("need more nodes")
        } else {
            panic("nodes overload, decrease them bro")
        }
        //go handleClient(client)

    }
}

