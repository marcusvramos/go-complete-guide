package arrays

import "fmt"

func main() {
	var productNames [4]string = [4]string{"laptop"}
	prices := []float64{33.55, 32.43, 56.77, 98.67}

	productNames[2] = "mobile"

	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(productNames[0])
	fmt.Println(productNames[2])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	anotherPrices := prices[:2]
	fmt.Println(anotherPrices)

	anotherSlicedNotation := prices[2:]
	fmt.Println(anotherSlicedNotation)


	productPrices := [4]int{10, 9, 8, 7}
	slicedPrices := productPrices[1:]
	highlightedPrices := slicedPrices[:1]

	fmt.Println("HP ", highlightedPrices)

	fmt.Println("Length of prices: ", len(productPrices))
	fmt.Println("Capacity of prices: ", cap(productPrices))

	fmt.Println("Length of highlightedPrices: ", len(highlightedPrices))
	fmt.Println("Capacity of highlightedPrices: ", cap(highlightedPrices))


	dynamicArray := []int{}
	fmt.Println("Length of dynamicArray: ", len(dynamicArray))
	fmt.Println("Capacity of dynamicArray: ", cap(dynamicArray))

	dynamicArray = append(dynamicArray, 1)
	fmt.Println("Length of dynamicArray: ", len(dynamicArray))
	fmt.Println("Capacity of dynamicArray: ", cap(dynamicArray))
	fmt.Println("Dynamic Array: ", dynamicArray)

	dynamicArray = append(dynamicArray, 2)
	fmt.Println("Length of dynamicArray: ", len(dynamicArray))
	fmt.Println("Capacity of dynamicArray: ", cap(dynamicArray))
	fmt.Println("Dynamic Array: ", dynamicArray)

	anotherList := []int{3, 4, 5}

	newList := append(dynamicArray, anotherList...)
	fmt.Println("NewList: ", newList)
	fmt.Println("Length of NewList: ", len(newList))
}