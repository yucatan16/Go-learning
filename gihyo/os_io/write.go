package main

import (
	"log"
	"os"
	"fmt"
)

func main() {
	// ファイルを生成
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}

	// プログラムが終わったらファイルを閉じる
	defer file.Close()

	// case1
	// message := []byte("hello world\n")
	// _, err = file.Write(message)

	//case2
	// WriteString()を用いると，毎回[]byteに変換する必要がなくなる
	// _, err = file.WriteString("hello world\n")

	//case3
	// 書き込む対象のio.WriterがWriteString()のようなメソッドを実装していない場合は，fmt.Fprint()を用いると，[]byteを経由せずio.Writerに対して文字列を直接書き込むことができる
	_, err = fmt.Fprint(file, "hello world\n")
	if err != nil {
		log.Fatal(err)
	}
}