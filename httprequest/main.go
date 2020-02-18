package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error is ", error)
		os.Exit(1)
	}

	//byte :=[]byte
	bs := make([]byte, 9999999) //number of empty elements in the byte slice
	resp.Body.Read(bs)

	fmt.Println(string(bs))
}
