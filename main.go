package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	filename = kingpin.Arg("filename", "Fortune filename").Default(".").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	fi, err := os.Stat(*filename)
	if err != nil {
		log.Fatal(*filename + " is not exist...")
	}
	var dat []byte
	if fi.Mode().IsRegular() {
		dat, err = ioutil.ReadFile(*filename)
		if err != nil {
			log.Fatal("Cannot open " + *filename + " filename...")
		}
		printFortune(string(dat))
	} else if fi.Mode().IsDir() {
		files, err := ioutil.ReadDir(*filename)
		if err != nil {
			log.Fatal(err)
		}

		var forts []string
		for _, file := range files {
			if isFortune(file.Name()) {
				forts = append(forts, file.Name())
			}
		}
		println(forts)
	}
}

func printFortune(data string) {
	fortunes := strings.Split(data, "\n%\n")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(len(fortunes))

	fmt.Println(fortunes[rnd])

}

func isFortune(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(filename + " is not exist...")
	}

	if fi.Mode().IsRegular() {
		fidat, err := os.Stat(filename + ".dat")
		if err != nil {
			log.Fatal(filename + ".dat is not exist...")
		}
		if fidat.Mode().IsRegular() {
			return true
		}
	}
	return false
}
