package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Jane"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Doe")

	fmt.Println(userNames)

	courses := make(floatMap, 3)

	courses["Python"] = 4.5
	courses["Go"] = 4.0
	courses["JavaScript"] = 4.2

	courses.output()

	for userNamesIndex, userName := range userNames {
		fmt.Println(userNamesIndex, userName)
	}

	for key, value := range courses {
		fmt.Println(key, value)
	}
}