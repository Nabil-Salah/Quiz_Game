package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	contents, err := os.ReadFile("problems.csv")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	r := csv.NewReader(strings.NewReader(string(contents)))
	var degree int
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("What is", record[0], "?")
		var ans int
		_, err = fmt.Scanf("%d", &ans)
		num, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if num == ans {
			degree++
		}
	}
	fmt.Printf("your answered %d answers correctly\n", degree)
}
