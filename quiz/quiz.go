package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct {
	P, A string
}

func RunQuiz() {
	csvFileName := flag.String("csv", "probblems.csv", "Csv file with `problem,answer` format")
	timeLimit := flag.Int("limit", 5, "Every problem must run on this time length")
	flag.Parse()

	file, error := os.Open(*csvFileName)
	if error != nil {
		exit(fmt.Sprintf("File not opened: %s", *csvFileName))
	}

	csvReader := csv.NewReader(file)
	lines, error := csvReader.ReadAll()
	if error != nil {
		exit("csv parse error.")
	}
	problems := parseLines(lines)
	quizeHellper(problems, *timeLimit)

}

func quizeHellper(probs []Problem, limit int) {
	ansChan := make(chan string)
	correct := 0
	for i, p := range probs {
		fmt.Printf("\nProblem #%d %s = ", i+1, p.P)
		time := time.NewTimer(time.Duration(limit) * time.Second)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			ansChan <- strings.TrimSpace(ans)
			time.Stop()
		}()
		select {
		case ans := <-ansChan:
			if ans == p.A {
				correct += 1
			}
		case <-time.C:
			fmt.Print("::timeout::")
		}
	}
	fmt.Printf("\nYou have %d correct answers out of %d", correct, len(probs))
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			P: strings.TrimSpace(line[0]),
			A: strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
