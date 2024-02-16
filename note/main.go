package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"davidwolff.dk/note_cli/note"
	"davidwolff.dk/note_cli/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	//userNote.Display()
	//err = saveData(userNote)
	err = outputData(userNote)

	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}
	fmt.Println("Saving the note succeeded!")

	err = outputData(todo)

	if err != nil {
		fmt.Println("Saving the todo failed.")
		return
	}

	printSomething(1.3)
	printSomething(3)
	printSomething("das")
	printSomething(todo)
	result := add(1, 3)
	fmt.Println("add() result ", result)
}

// switch on type
// can also use "any" as the type. It is the same as interface{}
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float64: ", value)
	case string:
		fmt.Println("String: ", value)
	}
}

// generics
func add[T int | float64 | string](a, b T) T {
	return a + b
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	fmt.Println("Saved the data succesfully")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
