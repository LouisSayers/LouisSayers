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

func copyFile(srcLocation, destLocation string) {
	destF := createFile(destLocation)
	defer destF.Close()

	srcF, err := os.Open(srcLocation)
	if err != nil {
		panic(err)
	}
	defer srcF.Close()

	io.Copy(destF, srcF)
}

func copyOverCSS() {
	srcLocation := "src/assets/css/site.css"
	destLocation := "dest/css/site.css"

	copyFile(srcLocation, destLocation)
}

func copyOverImages() {
	srcLocation := "src/assets/images/louis-sayers.jpg"
	destLocation := "dest/img/louis-sayers.jpg"

	copyFile(srcLocation, destLocation)
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
	mkDirs("dest/img/")
	mkDirs("dest/css/")

	copyOverCSS()
	copyOverImages()
	createIndexPage()
}
