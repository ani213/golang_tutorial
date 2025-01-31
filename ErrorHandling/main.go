package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func mains() { // main should be
	if _, err := os.Open("myFile.txt"); err != nil {
		fmt.Println(err, "errror")
		fmt.Println(os.ErrNotExist, "os errNotexist", err == os.ErrNotExist)
		if errors.Is(err, os.ErrNotExist) {
			log.Println("file does not exist")
		} else {
			log.Println(err)
		}
		return
	}

	fmt.Println("file successfully opened")
}
