package spider

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Article struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Created string `json:"created"`
	Author  string `json:"author"`
	Origin  string `json:"origin"`
	Link    string `json:"link"`
}

func WriteDou(i int, data *Article) {
	WriteFile("豆瓣", i, data)
}
func WriteDan(i int, data *Article) {
	WriteFile("煎蛋", i, data)
}

func WriteFile(folderName string, i int, data *Article) {
	folderPath := fmt.Sprintf("./%s_%d", folderName, i)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777)
	}

	fileName := strings.Replace(data.Title, "/", "or", -1)
	filePath := fmt.Sprintf("%s/%s.txt", folderPath, fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(data.Created + "\n")
	write.WriteString(data.Content + "\n")
	write.WriteString(data.Link + "\n")
	//Flush将缓存的文件真正写入到文件中
	write.Flush()
}

func appendFile(i int, data *Article) {
	folderPath := fmt.Sprintf("./故事_%d", i)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777)
	}

	fileName := strings.Replace(data.Title, "/", "or", -1)
	filePath := fmt.Sprintf("%s/%s.txt", folderPath, fileName)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	str := []byte("源链接：" + data.Link + "\n")
	//及时关闭file句柄
	defer file.Close()
	n, err := file.Write(str)
	// 当 n != len(b) 时，返回非零错误
	if err == nil && n != len(str) {
		println(`错误代码：`, n)
		panic(err)
	}
	//写入文件时，使用带缓存的 *Writer
	// write := bufio.NewWriter(file)
	// write.WriteString("源链接：" + data.Link + "\n")
	//Flush将缓存的文件真正写入到文件中
	// write.Flush()
}
