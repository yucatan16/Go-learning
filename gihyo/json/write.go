package main

import (
	"encoding/json"
	"log"
	"os"
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
	
	// ファイルを開く
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// エンコーダの取得
	encoder := json.NewEncoder(file)

	// JSONエンコードしたデータの書き込み
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
}