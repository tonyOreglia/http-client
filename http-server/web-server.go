package server

import "fmt"

const AF_INET string = ""
const SOCK_STREAM string = ""

type socket struct{}

func openSocket(a string, b string) socket {
	return socket{}
}

func getFilename(msg string) string {
	return msg
}

func open(filename string) string {
	return filename
}

func main() {
	serverSocket := openSocket(AF_INET, SOCK_STREAM)
	address := ""
	fmt.Print(serverSocket)
	fmt.Print(address)
	for true {
		fmt.Println("Ready to serve...")
		message := ""
		filename := getFilename(message)
		open(filename)
		// get data
		// send data
		// close socket
	}
	// close server socket
	// exit
}
