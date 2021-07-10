package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id int
	Content string
	Author string
}

func main() {
	// Create は、指定されたファイルを作成または切り捨てます。
	// ファイルがすでに存在する場合は、切り捨てられます。ファイルが存在しない場合は、モード 0666 (umask前)で作成されます。
	// 成功した場合、返されたFileのメソッドはI/Oに使用でき、関連するファイル記述子のモードはO_RDWRです。エラーが発生した場合、そのエラーは *PathError タイプとなります。
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	// NewWriterは渡したファイルに書き込むwriterを返却する
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts{
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	// Flush は、バッファリングされたデータを基礎となる io.Writer に書き込みます。
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post

	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}