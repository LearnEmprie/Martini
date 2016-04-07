package main


import (
	"html/template"
	"fmt"
	"github.com/LearnEmprie/Martini/sample"
	"net/http"
	"io/ioutil"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {

	httpGet()
	t,err := template.ParseFiles("./view/index.html")
	t.Execute(w,err)
}

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
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)

}


func httpGet()  {
	resp,err := http.Get("http://127.0.0.1:3000/string/key/1")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body , err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}