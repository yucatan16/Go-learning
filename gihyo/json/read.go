package main

import (
	"fmt"
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
	file, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// データを書き込む関数
	var person Person

	//　デコーダの取得
	decoder := json.NewDecoder(file)
	
	// JSONデコードしたデータの書き込み
	err = decoder.Decode(&person)
	if err != nil {
		log.Fatal(err)
	}
	
	// 読み出した結果の表示
	fmt.Println(person)
}