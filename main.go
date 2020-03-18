package main

import (
	"io"
	"os"
	"text/template"
)

func mkDirs(dirName string) {
	err := os.MkdirAll(dirName, 0777)
	if err != nil {
		panic(err)
	}
}

func createFile(fileName string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	return f
}

func copyOverCSS() {
	mkDirs("dest/css/")

	f := createFile("dest/css/site.css")
	defer f.Close()

	srcCSS, err := os.Open("src/assets/css/site.css")
	if err != nil {
		panic(err)
	}
	defer srcCSS.Close()

	io.Copy(f, srcCSS)
}

func createIndexPage() {
	homeTemplate, err := template.ParseFiles("src/layouts/base.gohtml", "src/index.gohtml")
	if err != nil {
		panic(err)
	}

	f := createFile("dest/index.html")
	defer f.Close()

	homeTemplate.Execute(f, nil)
}


func main() {
	mkDirs("dest")
	copyOverCSS()
	createIndexPage()
}
