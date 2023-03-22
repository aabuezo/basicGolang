package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	s := "names.txt"
	//createFile(s)

	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	readFile(s)
}

func createFile(s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("File created, read and printed :)")

	io.Copy(f, r)
}

func readFile(s string) {
	f, err := os.Open(s)
	if err != nil {
		//fmt.Println("err happened: ", err)
		log.Println("err happened: ", err)
		fmt.Println("Check log.txt in the current directory")
		//log.Fatalln(err)
		//panic(err)
		return
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
