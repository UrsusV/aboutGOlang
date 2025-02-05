package filehandling

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func generateNumbers(filename string, n int) {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error in file creation: ", err)
		return
	}

	defer file.close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		_, err := file.WriteString(fmt.Sprintf("%d\n", rand.Intn(10000)))

		if err != nil {
			fmt.Println("Error in file writing: ", err)
			return
		}

	}

}

func writeNumbersToFile(filename string, numbers []int) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	for _, num := range numbers {
		_, err := file.WriteString(fmt.Sprintf("%d\n", num))

		if err != nil {
			return err
		}
	}

	return nil
}

func readNumbersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
