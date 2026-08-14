package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewNote(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title or content is empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (n Note) WriteNoteToFile() (string, error) {
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	fmt.Println("Writing note to file: " + fileName)

	jsonContent, err := json.Marshal(n)

	if err != nil {
		return "", err
	}

	err = os.WriteFile(fileName, jsonContent, 0755)

	if err != nil {
		return "", err
	}

	return "The file has been saved successfully.", nil
}

//
//func (n Note) ReadNoteFromFile() Note {
//
//}

func (n Note) PrintNote() {
	fmt.Printf("======= %v ========= \n", n.Title)
	fmt.Printf("Content: %v\n", n.Content)
}
