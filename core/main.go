package main

import (
	"fmt"
	"net/http"

	"koukai/routing"
)

type HttpSocket struct {
	addr string
	port string
}

func httpSocketDefault() HttpSocket {
	httpSocket := HttpSocket{
		addr: "0.0.0.0",
		port: "8000",
	}
	return httpSocket
}

func (hs HttpSocket) toString() string {
	return hs.addr + ":" + hs.port
}

func main() {
	httpSocket := httpSocketDefault()

	fmt.Println("Sarting http server in ", httpSocket.toString())
	router := &http.ServeMux{}
	routing.Routing(router)

	err := http.ListenAndServe(httpSocket.toString(), router)
	if err != nil {
		fmt.Printf("Server failed in listenning from %s, err: %s \n", httpSocket.toString(), err)
	}
}
