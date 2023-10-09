package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func HandleFile(fileName string, ch chan []int) {

	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file")
		panic(err)
	}

	var elf []int
	var text string
	var food int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = scanner.Text()

		if text == "" {
			ch <- elf
			elf = []int{}
			continue
		}

		food, _ = strconv.Atoi(text)
		elf = append(elf, food)
	}

	ch <- elf
	close(ch)

}

func Sum(slice []int) int {
	var total int
	for _, value := range slice {
		total += value
	}

	return total
}

func main() {
	//elves := HandleFile(os.Args[1])

	inch := make(chan []int)
	//go HandleFile("../inputs/day1_test_input.txt", inch)
	go HandleFile("../inputs/day1_input.txt", inch)

	outch := make(chan int)
	var calories []int

	for elf := range inch {
		go func() {
			outch <- Sum(elf)
		}()

		calories = append(calories, <-outch)

	}

	slices.Sort(calories)

	fmt.Println(calories[len(calories)-1])

}
