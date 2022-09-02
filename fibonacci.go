package challenge_go

func Fibonacci(index int) int {
	a := 0
	b := 1
	if index < 0 {
		return -1
	} else {
		return GetFibonacci(index, a, b)
	}
}

func GetFibonacci(index int, a int, b int) int {
	if index == 0 {
		return a
	}
	temps := a
	a = b
	b = temps + a
	return GetFibonacci(index-1, a, b)
}
