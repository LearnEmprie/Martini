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
	"fmt"
	"github.com/LearnEmprie/Martini/sample"
	"net/http"
	C "github.com/LearnEmprie/Martini/http/controller"
	"log"
)



func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := sample.LoadPage(title)
	if err != nil {
		p = &sample.Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	"<form action=\"/save/%s\" method=\"POST\">"+
	"<textarea name=\"body\">%s</textarea><br>"+
	"<input type=\"submit\" value=\"Save\">"+
	"</form>",
		p.Title, p.Title, p.Body)
}


func main() {


	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/view/", C.ViewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":80", nil))

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
