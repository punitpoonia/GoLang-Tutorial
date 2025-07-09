package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")
	if err != nil {
		// log the error panic mode
		panic(err)
	}

	fileinfo, err := f.Stat() // to check file info
	if err != nil {
		// log the error panic mode
		panic(err)
	}
	fmt.Println("file name :", fileinfo.Name())
	fmt.Println("file directory :", fileinfo.IsDir())
	fmt.Println("file Size :", fileinfo.Size())
	fmt.Println("file Mode :", fileinfo.Mode())
	fmt.Println("file Modified time :", fileinfo.ModTime())

	// read file

	defer f.Close()

	buf := make([]byte, fileinfo.Size())

	d, err := f.Read(buf)
	if err != nil {
		// log the error panic mode
		panic(err)
	}
	for i := 0; i < len(buf); i++ {
		print(string(buf[i]))
	}
	fmt.Print(d)

	// Another way to read file

	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Print(string(data))

	// read folders
	fmt.Println("Folder os")
	dir, err := os.Open("../")

	if err != nil {
		panic(err)
	}

	defer dir.Close()

	fileInfo, err := dir.ReadDir(-1)

	for _, fi := range fileInfo {
		fmt.Println(fi.Name())
	}

	// How to create a file and write into file

	filecreate, err := os.Create("exmaple2.txt")

	if err != nil {
		panic(err)
	}

	defer filecreate.Close()

	filecreate.WriteString("Hello this file is created using go lang")
	filecreate.WriteString("This is another line appendedS")

}
