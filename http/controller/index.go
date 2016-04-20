package controller

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"html/template"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	//httpGet()
	t,err := template.ParseGlob("./view/index.html")
	//t.Execute(w,err)

	fmt.Println(t,err)
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
