package main

import (
	// "crypto/sha256"
	// "fmt"
	"net/http"
)


func main(){
	dir := http.Dir("./html")

	fs := http.FileServer(dir)

	
	// teste := sha256.New()
	// teste.Write([]byte("hello"))
	// fmt.Printf("%x", teste.Sum(nil))

	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
