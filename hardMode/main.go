package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//func Open(name string) (*File, error)

func main() {

	stdOut := os.Args
	fmt.Println(stdOut)

	contents, err := ioutil.ReadFile(stdOut[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", contents)

}
