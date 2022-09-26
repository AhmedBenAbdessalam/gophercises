package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	filePtr := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer' (default problems.csv)")
	timerPtr := flag.Int("t", 30, "the time limit for the quiz in seconds (default 30)")
	shufflePtr := flag.Bool("shuffle", false, "shuffle questions or not (default false)")
	flag.Parse()
	file, err := os.Open(*filePtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileReader := csv.NewReader(file)
	records, _ := fileReader.ReadAll()
	// shuffle
	if *shufflePtr {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	cmdReader := bufio.NewReader(os.Stdin)
	answers := make(chan (string))
	score := 0
	timer := time.NewTimer(time.Duration(*timerPtr) * time.Second)
quizloop:
	for _, r := range records {
		fmt.Printf("%s: ", r[0])
		go func() {
			text, _ := cmdReader.ReadString('\n')
			t := strings.Trim(text, "\n")
			answers <- t
		}()

		select {
		case <-timer.C:
			break quizloop
		case c := <-answers:
			c = strings.TrimSpace(c)
			c = strings.ToLower(c)
			r[1] = strings.ToLower(r[1])
			if c == r[1] {
				score++
			}
		}

	}
	fmt.Printf("\nYou scored %d out of %d.\n", score, len(records))
}
