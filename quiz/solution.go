package solution

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Parse commandline
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of equations")
	timeLimit := flag.Int("limit", 5, "the time limit for the quiz in seconds")
	flag.Parse()

	// Check error
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv FILE: %s\n", *csvFilename))
	}

	// read csv
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	fmt.Println(lines)

	// solve
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// a new channel
		answerCh := make(chan string)
		// get the answer from scan
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		// select
		select {
		case <-timer.C:
			fmt.Println()
			return
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct")
				correct++
			}
		}
	}
	fmt.Printf("\nYou SCORED %d out of %d.\n", correct, len(problems))
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
