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

func copyOverImages() {
	mkDirs("dest/img/")

	f := createFile("dest/img/louis-sayers.jpg")
	defer f.Close()

	img, err := os.Open("src/assets/images/louis-sayers.jpg")
	if err != nil {
		panic(err)
	}
	defer img.Close()

	io.Copy(f, img)
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
	copyOverImages()
	createIndexPage()
}
