package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	// bs := make([]byte, 99999)

	// _, err = resp.Body.Read(bs)

	// // if err != nil {
	// // 	panic(err)
	// // }

	// fmt.Println(string(bs))

	io.Copy(os.Stdout, resp.Body)
}
