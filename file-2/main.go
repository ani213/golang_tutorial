package main

import "os"

func main() {
	err := os.WriteFile("example.txt", []byte("Hello, World! this is Go world"), 0644)
	if err != nil {
		panic(err)
	}
	fs, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer fs.Close()
	fsInfo, err := fs.Stat()
	if err != nil {
		panic(err)
	}
	println("File name:", fsInfo.Name())
	println("File size:", fsInfo.Size())
	println("Is directory:", fsInfo.IsDir())
	println("File mode:", fsInfo.Mode())
	println("Last modified:", fsInfo.ModTime().String())
	println("Is regular file:", fsInfo.Mode().IsRegular())

	f, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	println("File content:", string(f))

}
