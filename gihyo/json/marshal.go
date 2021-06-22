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
	person := &Person{
		ID: 1,
		Name: "Yukako",
		Email: "yukako@example.com",
		Age: 27, 
		Address: "",
		memo: "hoge",
	}
	
	b,err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(b))
}