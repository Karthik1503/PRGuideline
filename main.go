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

func divide(x, y int) int {

	result := 0
	if x < y {
		result = 0
	} else if x == y {
		result = 1
	} else {
		for x > 0 && x > y {
			result += 1
			subtract(x, y)
		}
	}
	return result
}


func simpleMultiply(x, y int) int {
	return x * y
}

func simpleDivide(x, y int) int {
	return x % y
}

