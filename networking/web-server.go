package server

import "fmt"

const AF_INET string = ""
const SOCK_STREAM string = ""

type socket struct{}

func openSocket(a string, b string) socket {
	return socket{}
}

func main() {
	serverSocket := openSocket(AF_INET, SOCK_STREAM)
	fmt.Print(serverSocket)

	for true {
		fmt.Println("Ready to serve...")
	}
}
