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

func copyOverCSS() {
	mkDirs("dest/css/")

	destCSS, err := os.Create("dest/css/site.css")
	if err != nil {
		panic(err)
	}
	defer destCSS.Close()

	srcCSS, err := os.Open("src/assets/css/site.css")
	if err != nil {
		panic(err)
	}
	defer srcCSS.Close()

	io.Copy(destCSS, srcCSS)
}

func createIndexPage() {
	homeTemplate, err := template.ParseFiles("src/layouts/base.gohtml", "src/index.gohtml")
	if err != nil {
		panic(err)
	}

	indexHtml, err := os.Create("dest/index.html")
	if err != nil {
		panic(err)
	}
	defer indexHtml.Close()

	homeTemplate.Execute(indexHtml, nil)
}

func main() {
	mkDirs("dest")
	copyOverCSS()
	createIndexPage()
}
