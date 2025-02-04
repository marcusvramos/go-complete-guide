package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structs-practice/note"
	"example.com/structs-practice/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething("Hello, World!")
	printSomething(42)

	title, content := getNoteData()
	todoContent := getTodoData()

	userNote, err := note.New(title, content)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	todo, err := todo.New(todoContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func printSomething(value interface{}) {
	typedVal, ok := value.(int)

	if ok {
		fmt.Println("Int:", typedVal)
		return
	}

	// switch value.(type) {
	// case string:
	// 	fmt.Println("String:", value)
	// case int:
	// 	fmt.Println("Int:", value)
	// default:
	// 	fmt.Println("Unknown type")
	// }
	// fmt.Println(value)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) (error) {
	err := data.Save()

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("saved successfully!")
	return nil
}

func getTodoData() (string) {
	content := getUserInput("Enter todo: ")

	return content
}

func getNoteData() (string, string) {
	title := getUserInput("Enter title: ")
	body := getUserInput("Enter body: ")

	return title, body
}

func getUserInput(promptText string) (string) {
	fmt.Print(promptText)
	
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}