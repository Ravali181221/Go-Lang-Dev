package Sum

//import "fmt"

func Sum() {
	// define the sequence of numbers
	numbers := []int{1, 2, 3, 4, 5}

	// initialize the sum to 0
	sum := 0

	// use a for loop to iterate over the numbers and add them to the sum
	for _, num := range numbers {
		sum += num
	}

	// print the sum
	//fmt.Println(sum)
}
