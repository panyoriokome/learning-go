package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	//Todo is struct
	type Todo struct {
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
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

	// https://noumenon-th.net/programming/2019/09/06/json-unmarshal/ のコピペ
	// url := "https://jsonplaceholder.typicode.com/todos"
	// request, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// //クエリパラメータ
	// params := request.URL.Query()
	// params.Add("userId", "1")
	// request.URL.RawQuery = params.Encode()

	// client := http.Client{}

	// response, err := client.Do(request)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var todos []Todo
	// if err := json.Unmarshal(body, &todos); err != nil {
	// 	log.Fatal(err)
	// }

	// for _, todo := range todos {
	// 	fmt.Printf("id: %v, title: %v, completed: %v\n", todo.ID, todo.Title, todo.Completed)
	// }

	// JSON, Readers, and Writers
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	// ファイルを定義して、
	// TmpFileは*os.Fileを返す https://golang.org/pkg/io/ioutil/#TempFile
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	// os.Fileに対してjson.Encoderで読み書きを行う？
	// The os.File type implements both the io.Reader and io.Writer interfaces, so we can use it to demonstrate json.Decoder and json.Encoder.
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	// 作成したファイルを読み込む
	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fromFile) // {Name:Fred Age:40}
}
