package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	fileReader := csv.NewReader(file)
	records, _ := fileReader.ReadAll()
	cmdReader := bufio.NewReader(os.Stdin)
	score := 0
	for _, r := range records {
		fmt.Printf("%s: ", r[0])
		text, _ := cmdReader.ReadString('\n')
		t := strings.Trim(text, "\n")
		if t == r[1] {
			score++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", score, len(records))
}