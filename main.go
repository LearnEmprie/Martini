package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
)

type FastCGIServer struct{}

func (s FastCGIServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte("<h1>Hello, 世界</h1>\n<p>Behold my Go web app.</p>"))
}

func main() {
	listener, _ := net.Listen("tcp", "127.0.0.1:9001")
	srv := new(FastCGIServer)
	fcgi.Serve(listener, srv)
}

/*
import (
	s "github.com/LearnEmprie/Martini/server"
	"fmt"
	"os"
)



func main() {

	if len(os.Args)!=3  {
		fmt.Println("Wrong pare")
		os.Exit(0)
	}

	if os.Args[1]=="server" && len(os.Args)==3 {

		s.StartServer(os.Args[2])
	}


	if os.Args[1]=="client" && len(os.Args)==3 {

		s.StartClient(os.Args[2])
	}
}*/
