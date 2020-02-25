package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EventMessage struct {
	Homepage string `json:"homepage"`
	Branch   string `json:"branch"`
}

func Client(w http.ResponseWriter, r *http.Request) {
	var data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Failed to read request: %s\n", err)
		return
	}

	var msg EventMessage
	err = json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Printf("Failed to parse request: %s\n", err)
		return
	}

	fmt.Println(msg)

}
