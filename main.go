package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/phantommachine/book-scanner/serial"
)

const (
	readTimeout = 25 * time.Millisecond
	delay       = 100 * time.Millisecond
)

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	s, err := serial.New("COM13")
	if err != nil {
		log.Fatal("error getting serial:", err)
	}

	total := 0
	buff := make([]byte, 512)
	for {
		n, err := s.Read(buff[total:])
		if err != nil {
			log.Fatal("error reading serial:", err)
			break
		}
		if n == 0 { // nothing recieved, delay and try again
			time.Sleep(delay)
			continue
		}
		total += n
		if buff[total-1] == '\r' {
			fmt.Println(string(buff[:total]))
			resp, err := http.Get(fmt.Sprintf("https://openlibrary.org/isbn/%s.json", buff[:total-1]))
			if err != nil {
				log.Println("get by isbn error:", err)
			}

			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			total = 0
		}

		if total >= cap(buff) {
			fmt.Println(string(buff[:total]))
			total = 0
		}
	}
}

func usage() {
	fmt.Println("usage: book-scanner <serial-device>")
}
