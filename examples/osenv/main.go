package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	resp, _ := http.Get("https://go.dev/dl/?mode=json")
	defer resp.Body.Close()

	file, _ := os.Create("hoge.json")
	io.Copy(file, resp.Body)

	fmt.Println(filepath.Abs(file.Name()))
}
