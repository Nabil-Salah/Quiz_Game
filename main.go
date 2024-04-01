package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz in seconds")
	filename := flag.String("fname", "problems.csv", "the time limit for the quiz in seconds")
	flag.Parse()
	contents, err := os.ReadFile(*filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	r := csv.NewReader(strings.NewReader(string(contents)))
	var degree int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
questionsloop:
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("What is", record[0], "?")
		ansch := make(chan string)
		go func() {
			var ans string
			_, err = fmt.Scanf("%s\n", &ans)
			ansch <- ans
		}()
		select {
		case <-timer.C:
			break questionsloop
		case answer := <-ansch:
			num := strings.TrimSpace(record[1])
			answer = strings.TrimSpace(answer)
			if num == answer {
				degree++
			}
		}
	}
	fmt.Printf("your answered %d answers correctly\n", degree)
}
