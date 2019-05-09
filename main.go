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
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	_, err := os.Stat(*filename)
	if err != nil {
		fmt.Println(*filename + " is not exist...")
		os.Exit(1)
	}
	dat, err := ioutil.ReadFile(*filename)
	if err != nil {
		fmt.Println("Cannot open " + *filename + " filename...")
		os.Exit(1)
	}

	fortunes := strings.Split(string(dat), "\n%\n")
	length := len(fortunes)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(length)

	fmt.Println(fortunes[rnd])
}
