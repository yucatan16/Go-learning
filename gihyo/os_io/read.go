package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	// ファイルを開く
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// プログラムが終わったらファイルを閉じる
	defer file.Close()

	// 12byte格納可能なスライスを用意する
	message := make([]byte, 12)

	// ファイル内のデータをスライスに読み出す
	// Read()は，読み出したデータを格納するのに十分な長さを持ったスライスを渡すと，そこにデータが格納される
	_, err = file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", message)
}