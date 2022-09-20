package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const URL = "http://localhost:18888"

// curl http://localhost:18888
func curl1() {
	resp, err := http.Get(URL)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		os.Exit(1)
	}

	log.Println(string(body))
}

// curl -G --data-urlencode "query=hello world" http://localhost:18888
func curl2() {
	values := url.Values{
		"query": {"hello world"},
	}

	resp, err := http.Get(URL + "?" + values.Encode())
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	log.Println(string(body))

}

// curl --head http://localhost:18888
func curl3() {
	resp, err := http.Head(URL)
	if err != nil {
		os.Exit(1)
	}

	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)

}

// x-www-form-urlencoded形式のPOSTメソッドの送信
// curl -d test=value http://localhost:18888
// クエリー付きのGETアクセスの時に出てきたurl.Valuesが再度登場していますが、url.Values.
// Encode()メソッドは呼ばずに、http.PostForm()関数に渡します
func curl4() {
	values := url.Values{
		"test": {"value"},
	}

	resp, err := http.PostForm(URL, values)
	if err != nil {
		os.Exit(1)
	}
	log.Println("Status", resp.Status)
}

// $ curl -T main.go -H "Content-Type: text/plain" http://localhost:18888
func curl5() {
	file, err := os.Open("main.go")
	if err != nil {
		os.Exit(1)
	}
	resp, err := http.Post(URL, "text/plain", file)
	if err != nil {
		os.Exit(1)
	}

	log.Println("Status:", resp.Status)
}

// 文字列をio.Readerインタフェース化する
func curl6() {
	reader := strings.NewReader("テキスト")
	resp, err := http.Post(URL, "text/plain", reader)
	if err != nil {
		os.Exit(1)
	}
	log.Println("Status:", resp.Status)
}

// curl -F "name=Michael Jackson" -F "thumbnail=@photo.jpg" http://localhost:18888
func curl7() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		os.Exit(1)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		os.Exit(1)
	}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(URL, writer.FormDataContentType(), &buffer)
	if err != nil {
		os.Exit(1)
	}
	log.Println("Status:", resp.Status)
}

func curl8() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "MJ")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		os.Exit(1)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		os.Exit(1)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(URL, writer.FormDataContentType(), &buffer)
	if err != nil {
		os.Exit(1)
	}
	log.Println("status", resp.Status)

}

func curl9() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "MJ")

	fileWriter, err := writer.CreateFormFile("thumbnail", "photo.jpg")
	if err != nil {
		os.Exit(1)
	}

	readFile, err := os.Open("photo.jpg")
	if err != nil {
		os.Exit(1)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post(URL, writer.FormDataContentType(), &buffer)
	if err != nil {
		os.Exit(1)
	}
	log.Println("status", resp.Status)
}

func curl10() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "MJ")
}
