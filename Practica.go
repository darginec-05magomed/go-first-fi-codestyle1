package main

import "fmt"

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, fmt.Errorf("нельзя делить на ноль")
	}
	return x / y, nil
}

func main() {
	var a, b int
	res, err := div(a, b)
	fmt.Println(res, err)

	a, b = 27, 3
	res, err = div(a, b)
	fmt.Println(res, err)
}
