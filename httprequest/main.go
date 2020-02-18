package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote these bytes :", len(bs))
	return len(bs), nil
}

func main() {
	resp, error := http.Get("http://google.com")
	if error != nil {
		fmt.Println("Error is ", error)
		os.Exit(1)
	}

	//byte :=[]byte
	//bs := make([]byte, 9999999) //number of empty elements in the byte slice
	//resp.Body.Read(bs)

	//fmt.Println(string(bs))
	lw := logWriter{}
	io.Copy(lw, resp.Body)
	//io.Copy(os.Stdout, resp.Body)
}
