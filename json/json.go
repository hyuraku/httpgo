package json

import (
	"encoding/json"
	"fmt"
)

func Json(body []byte) {
	var jsonres interface{}
	if err := json.Unmarshal(body, &jsonres); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	fmt.Println(PrettyPrint(jsonres))
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
