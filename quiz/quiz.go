package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// Parse commandline
	csvfile := flag.String("csv", "problems.csv", "a csv file with problems")
	timelimit := flag.Int("limit", 30, "time duration with seconds")
	shuffle := flag.Bool("shuffle", false, "a shuffle to ask questions")
	flag.Parse()

	// Check error
	file, err := os.Open(*csvfile)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the file %s\n", *csvfile))
	}
	// read csv

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}
	fmt.Println(lines)

	problems := parseLines(lines)
	if *shuffle == true {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You Scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct")
				correct++
			}
		}
	}

	fmt.Printf("You Scored %d out of %d.\n", correct, len(problems))
}

// fucntion to parse lines
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// type of problem
type problem struct {
	q string
	a string
}

// finish the code
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
