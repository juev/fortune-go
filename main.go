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
	if fi.Mode().IsRegular() {
		printFortune(*filename)
	} else if fi.Mode().IsDir() {
		files, err := ioutil.ReadDir(*filename)
		if err != nil {
			log.Fatal(err)
		}

		var forts []string
		for _, file := range files {
			if isFortune(*filename + "/" + file.Name()) {
				forts = append(forts, *filename+"/"+file.Name())
			}
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		rnd := r1.Intn(len(forts))
		printFortune(forts[rnd])
	}
}

func printFortune(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot open " + filename + " filename...")
	}

	fortunes := strings.Split(string(data), "\n%\n")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(len(fortunes))

	fmt.Println(fortunes[rnd])
}

func isFortune(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	if _, err := os.Stat(filename + ".dat"); os.IsNotExist(err) {
		return false
	}
	return true
}
