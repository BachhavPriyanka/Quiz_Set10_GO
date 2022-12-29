package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Questions struct {
	Question       string
	AnswerOption_A string
	AnswerOption_B string
	AnswerOption_C string
	AnswerOption_D string
	CorrectOption  string
}

var score int

func beginToQuiz(lines [][]string) int {
	var totalQue int
	for _, currentLine := range lines {
		totalQue += 1
		data := Questions{
			Question:       currentLine[0],
			AnswerOption_A: currentLine[1],
			AnswerOption_B: currentLine[2],
			AnswerOption_C: currentLine[3],
			AnswerOption_D: currentLine[4],
			CorrectOption:  currentLine[5],
		}
		fmt.Println(data.Question)

		fmt.Println("A.", data.AnswerOption_A, "| B.", data.AnswerOption_B, "| C.", data.AnswerOption_C, "| D.", data.AnswerOption_D)

		fmt.Printf("Please Enter Your Option (A,B,C or D): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		selectedOption := scanner.Text()

		fmt.Println("Correct Option :", data.CorrectOption, "\n")
		finalScore(selectedOption, data.CorrectOption)
	}
	return totalQue
}

func finalScore(selectedOption string, CorrectOption string) {
	res := selectedOption
	ans := CorrectOption
	if res == ans {
		score++
	}
}
func fileOperation(fileName *string) [][]string {
	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}
	return lines
}
func main() {
	fileName := flag.String("csv", "Questions.csv", "a csv file in the format of 'question & answer'")
	lines := fileOperation(fileName)

	var totalQue int
	totalQue = beginToQuiz(lines)
	fmt.Println("You Scored :", score, "Total question :", totalQue)
	fmt.Println()
}
