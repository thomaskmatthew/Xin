package get

import (
	"Xin/httpclient/options"
	"fmt"
	"io"
	"net/http"
)

func get(url string, *opt options.Options) string {
	header := opt.Headers
	if header["Content-Type"] == "" {
		header["Content-Type"] = "appliction/json"
	}

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

func getRequest() {
	fmt.Println("hello world")
	getRequest := get("https://dummyjson.com/products")
	fmt.Println(getRequest)
}
