package main

import (
	"fmt"
	"filehandling"
	"dsa"
)

func main() {
	filename := "numbers.txt"
	err := filehandling.generateNumbers(filename, 1000)

	if err != nil {
		fmt.Println("Error in file creation: ", err)
		return
	}

	numbers, err := filehandling.readNumbersFromFile(filename)

	if err != nil {
		fmt.Println("Error in reading file: ", err)
		return
	}

	sortedNumbers := dsa.mergeSort(numbers)

	err = filehandling.writeNumbersToFile("sorted_numbers.txt", sortedNumbers)

	if err != nil {
		fmt.Println("Error in writing file: ", err)
		return
	}

	err = filehandling.writeNumbersToFile("sorted_numbers.txt", sortedNumbers)

	if err != nil {
		fmt.Println("Error in writing file: ", err)
		return
	}

	fmt.Println("File operations completed successfully")

}
