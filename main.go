package main

import(
	"crypto/sha256"
	"fmt"
)

func main(){
	 teste := sha256.New()
	 teste.Write([]byte("hello"))
	 fmt.Printf("%x", teste.Sum(nil))
}