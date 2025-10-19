package post

import (
	"fmt"
	"io"
	"net/http"
)

func post(url string, body io.Reader) string {
	post, err := http.Post(url, "Content-Type", body)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer post.Body.Close()

	bytes, er := io.ReadAll(post.Body)

	if er != nil {
		fmt.Println(er)
		return ""
	}

	return string(bytes)

}

func postRequest() {

	fmt.Println("hello")
}
