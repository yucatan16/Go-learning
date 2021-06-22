package main

import (
	"fmt"
	"encoding/json"
	"log"
	)

type Person struct {
  	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"` // jsonに格納しない
	Age int `json:"age"`
	Address string
	memo string // プライベートなフィールド名（小文字）は出力されない
}

func main() {
	var person Person
	// なんでbyteを利用する？
	b := []byte(`{"id": 1, "name": "Yukako", "age": 27}`)
	
	err := json.Unmarshal(b, &person)
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(person)
}