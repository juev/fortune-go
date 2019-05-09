package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	filename = kingpin.Arg("filename", "Fortune filename").String()
	dat      []byte
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	fi, err := os.Stat(*filename)
	if err != nil {
		fmt.Println(*filename + " is not exist...")
		os.Exit(1)
	}
	if fi.Mode().IsRegular() {
		dat, err = ioutil.ReadFile(*filename)
		if err != nil {
			fmt.Println("Cannot open " + *filename + " filename...")
			os.Exit(1)
		}
	} else if fi.Mode().IsDir() {
		fmt.Println(*filename + " is a directory...")
		os.Exit(0)
	}

	fortunes := strings.Split(string(dat), "\n%\n")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(len(fortunes))

	fmt.Println(fortunes[rnd])
}
