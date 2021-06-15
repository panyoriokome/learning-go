package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	// {
	// 	"id":"12345",
	// 	"date_ordered":"2020-05-01T13:01:02Z",
	// 	"customer_id":"3",
	// 	"items":[{"id":"xyz123","name":"Thing 1"},{"id":"abc789","name":"Thing 2"}]
	// }
	type Item struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	type Order struct {
		ID          string    `json:"id"`
		DateOrdered time.Time `json:"date_ordered"`
		CustomerID  string    `json:"customer_id"`
		Items       []Item    `json:"items"`
	}

	// jsonからstructへの変換
	data := `{"CustomerID":"22222","ID": "3","DataOrdered":"Mon Jan 30 15:04:05 MST 2006"}`
	var o Order
	if err := json.Unmarshal([]byte(data), &o); err != nil {
		fmt.Println(err)
	}
	fmt.Println(o)

	// Structからbytes of slicesへの変換
	out, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(out)
}
