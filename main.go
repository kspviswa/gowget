package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid number of args")
	}
	serverURL := os.Args[1]
	res, err := http.Get(strings.TrimSpace(serverURL))
	if err != nil {
		log.Fatal(err)
	} else {
		body, err := ioutil.ReadAll(res.Body)
		//fmt.Println(string(body))
		//fmt.Errorf(err1.Error())
		if err == nil {
			f, err := os.Create("index.html")
			defer f.Close()
			if err == nil {
				n, err := fmt.Fprintf(f, string(body))
				if err == nil {
					fmt.Println("Sucessfully writen %v bytes", n)
				} else {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}
	fmt.Printf("Gow Get")
}
