package main

import "fmt"

func whatNumber(num string) {

	nameOfNumbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	fmt.Println(nameOfNumbers[num])

}

func main() {

	whatNumber("six")
}
