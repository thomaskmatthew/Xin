package main

import (
	"fmt"
	"io"
	"net/http"
)

func get(url string) string {

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	bytes, er := io.ReadAll(res.Body)
	if er != nil {
		fmt.Println(er)
		return ""
	}

	return string(bytes)

}

func main() {
	fmt.Println("hello world")
	getRequest := get("https://dummyjson.com/products")
	fmt.Println(getRequest)
}
