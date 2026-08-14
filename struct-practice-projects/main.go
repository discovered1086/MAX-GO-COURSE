package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct-practice-projects/note"
)

func main() {
	noteTitle := inputNoteDetails("Enter a title for the note: ")
	noteText := inputNoteDetails("Enter text for the note: ")

	userNote, err := note.NewNote(noteTitle, noteText)

	if err != nil {
		fmt.Println("Terminating the program due to invalid input")
		return
	}

	userNote.PrintNote()

	message, err := userNote.WriteNoteToFile()

	if err != nil {
		fmt.Println("Terminating the program due to file saving error")
		return
	}

	fmt.Println(message)

}

func inputNoteDetails(prompt string) string {
	var inputNote string
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	inputNote = strings.TrimSuffix(text, "\n")
	inputNote = strings.TrimSuffix(text, "\r")

	return inputNote
}
