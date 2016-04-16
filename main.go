package main



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
}


