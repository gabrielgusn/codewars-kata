package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	url := os.Args[1]

	client := http.Client{
		Timeout: 2000 * time.Millisecond,
	}
	now := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error while making HTTP request to", url)
		panic(err)
	}
	responseTime := time.Since(now).Milliseconds()

	statusCode := resp.StatusCode

	fmt.Print(resp)

	fmt.Println(url, "status is", statusCode, "[", responseTime, "ms ]")
}
