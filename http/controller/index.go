package controller

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"html/template"
	"log"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {

	//httpGet()
	//t,err := template.ParseFiles("./view/index.html")
	//t.Execute(w,err)


	const tpl = `   <!DOCTYPE html>
			<html>
				<head>
					<meta charset="UTF-8">
					<title>{{.Title}}</title>
				</head>
				<body>
					{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
				</body>
			</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "大狗升军",
		Items: []string{
			"吃鸡吧去吧",
			"你妹的",
		},
	}

	t.Execute(w,data)
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
