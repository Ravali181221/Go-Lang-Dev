package calculation

//import "fmt"

func Multiplication() {
	// define the sequence of numbers
	numbers := []int{1, 2, 3, 4, 6}

	// initialize the product to 1
	product := 1

	// use a for loop to iterate over the numbers and multiply them to the product
	for _, num := range numbers {
		product *= num
	}

	// print the product
	//fmt.Println(product)
}
