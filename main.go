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

func copyFile(src, dest string) {
	destF := createFile(dest)
	defer destF.Close()

	srcF, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer srcF.Close()

	io.Copy(destF, srcF)
}

func createIndexPage(data PageDetails) {
	t, err := template.ParseFiles("src/layouts/base.gohtml", "src/index.gohtml")
	if err != nil {
		panic(err)
	}

	f := createFile("dest/index.html")
	defer f.Close()

	t.Execute(f, data)
}

func createWishlistPage(data PageDetails) {
	t, err := template.ParseFiles("src/layouts/base.gohtml", "src/wishlist.gohtml")
	if err != nil {
		panic(err)
	}

	f := createFile("dest/wishlist.html")
	defer f.Close()

	t.Execute(f, data)
}

type PageDetails struct {
	PageName string
}
func main() {
	mkDirs("dest/img/")
	mkDirs("dest/css/")

	copyFile("src/assets/css/site.css", "dest/css/site.css")
	copyFile("src/assets/css/home.css", "dest/css/home.css")
	copyFile("src/assets/css/wishlist.css", "dest/css/wishlist.css")
	copyFile("src/assets/images/louis-sayers.png", "dest/img/louis-sayers.png")

	createIndexPage(PageDetails{ PageName: "home" })
	createWishlistPage(PageDetails{ PageName: "wishlist" })
}
