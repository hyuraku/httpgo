package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/hyuraku/httpgo/json"
)

func main() {
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	minetype := resp.Header.Get("Content-Type")
	if strings.Contains(minetype, "application/json") {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		json.Json(body)
	} else if strings.Contains(minetype, "image") {
		f, err := os.Create("image.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, resp.Body)
		fmt.Println("File image.jpg downlaod in current working directory")
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}
