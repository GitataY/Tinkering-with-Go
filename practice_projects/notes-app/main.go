package main

import (
	"errors"
	"fmt"
	"example.com/note/note"
)

func main() {
	title, content, := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string){
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content

}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	var input string
	fmt.Scanln(&input)

	return input
}