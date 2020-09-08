package main

import (
	"fmt"
	"os"
)

func quizzTime(p [][]string){
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

	problems := [][]string{ //Matrix of problems
		{"10+2=", "12"},
		{"2+2=", "4"},
		{"17-2=", "15"},
	}
	quizzTime(problems)
}