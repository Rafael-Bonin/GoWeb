package utils

func Sort(numbers []int) []int {
	for i := range numbers {
		for j := 0; j < len(numbers) -1 - i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}