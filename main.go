package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func readFile() []string {
	data, err := ioutil.ReadFile("questions.data")
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(data), " ")
}
func questionsAmount(s []string) int {
	return len(s)/2
}
func parseFile(problemString []string,amount int) [100][2]string {
	var problemMatrix [100][2]string
	k1 := 0
	k2 := 1
	for i := 0; i < amount; i++ {
		problemMatrix[i][0] = problemString[k1]
		k1 = k1 + 2
		problemMatrix[i][1] = problemString[k2]
		k2 = k2 + 2
	}
	return problemMatrix
}
func quizzTime(p [100][2]string) {
	i := 0
	var answer string
	for {
		fmt.Println(p[i][0]) //Print question
		fmt.Fscan(os.Stdin, &answer)
		if answer == p[i][1] {
			i++
		} else {
			fmt.Println("No! Try again!")
		}
	}
}

func main() {
	data := readFile()
	amount:=questionsAmount(data)
	questions := parseFile(data,amount)
	quizzTime(questions)
}