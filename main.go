package PRGuideline


func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	index := 0
	result := 0
	for index < y {
		result += add(x, x)
	}
	return result
}

func simpleMultiply(x, y int) int {
	return x * y
}


