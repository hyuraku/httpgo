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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	minetype := resp.Header.Get("Content-Type")
	if strings.Contains(minetype, "json") {
		json.Json(body)
	} else {
		fmt.Println(string(body))
	}

}
