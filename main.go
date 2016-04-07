package main


import (
	"fmt"
	"github.com/LearnEmprie/Martini/sample"
	"net/http"
	C "github.com/LearnEmprie/Martini/http/controller"
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
	http.ListenAndServe(":8080", nil)

}


