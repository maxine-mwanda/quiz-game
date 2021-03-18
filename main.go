package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"quiz-game/entities"
	"strings"
	"time"
)

func main(){
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question and answer'")
	timeLimit := flag.Int("limit", 30, "The quiz should last 30 seconds")
	flag.Parse()
	_ = csvFilename

	file, err := os.Open(*csvFilename)
	if err !=nil{
		exit(fmt.Sprintf("Could not open the CSV file: %s", *csvFilename))
		os.Exit(1)
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil{
		exit("Could not parse csv file")
	}
	//fmt.Println(lines)

	problems := parselines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct:= 0

	for i, p := range problems{
		fmt.Printf("problem %d: %s = \n", i+1, p.Question)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d\n", correct, len(problems) )
			return
		case answer:= <-answerCh:
			if answer == p.Answer {
				correct++
			//fmt.Println("Correct")
		}
		}

	}
	//fmt.Printf("You scored %d out of %d\n", correct, len(problems) )
}

func parselines(lines [][]string) []entities.Problem{
	ret := make([]entities.Problem, len(lines))

	for i, line := range lines{
		ret[i] =entities.Problem{
			Question: line[0],
			Answer: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit (msg string){
	fmt.Println(msg)
	os.Exit(1)
}