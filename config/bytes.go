package config

import (
	"bytes"
	"fmt"
	"os"
	"encoding/base64"
	"io"
)

func MyBytes()  {
	var b bytes.Buffer // A Buffer needs no initialization.
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!")
	b.WriteTo(os.Stdout)

	//byte 就是字节 和 string差不多
	dog := []byte("ggg")

	fmt.Println(string(dog))
}


func Reader(){

	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding,buf)
	io.Copy(os.Stdout,dec)
}

