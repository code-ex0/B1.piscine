package main

import "os"

func main() {
	err, Param := ConvertRune(os.Args[1:])
	if err != false && len(Param[1]) == 1 {
		var Result int
		err1, Number1 := ConvertInt(Param[0])
		err2, Number2 := ConvertInt(Param[2])
		if err1 != false && err2 != false {
			switch os.Args[2] {
			case "/":
				if Number2 == 0 {
					os.Stdout.WriteString("No division by 0\n")
				} else {
					os.Stdout.WriteString(itoa(Number1/Number2) + "\n")
				}
			case "*":
				Result = Number1 * Number2
				if Number1 == 0 || (Result/Number1 == Number2) {
					os.Stdout.WriteString(itoa(Result) + "\n")
				}
			case "+":
				Result = Number1 + Number2
				if (Result > Number1) == (Number2 > 0) {
					os.Stdout.WriteString(itoa(Result) + "\n")
				}
			case "-":
				Result = Number1 - Number2
				if (Result < Number1) == (Number2 > 0) {
					os.Stdout.WriteString(itoa(Result) + "\n")
				}
			case "%":
				if Number2 == 0 {
					os.Stdout.WriteString("No modulo by 0\n")
				} else {
					os.Stdout.WriteString(itoa(Number1%Number2) + "\n")
				}
			}
		}
	}
}

func ConvertRune(String []string) (bool, [][]rune) {
	Temp := make([][]rune, len(String))
	for i := 0; i < len(String); i++ {
		for _, letter := range String[i] {
			if letter >= '0' && letter <= '9' || letter == '/' || letter == '*' || letter == '-' || letter == '+' || letter == '%' {
				Temp[i] = append(Temp[i], letter)
			} else {
				return false, Temp
			}
		}
	}
	return true, Temp
}

func ConvertInt(Temp []rune) (bool, int) {
	if len(Temp) <= 20 {
		var Result int
		Mult := 1
		for i := 0; i < len(Temp); i++ {
			if Temp[i] == '-' {
				Mult = -1
			} else {
				Result *= 10
				Result += int(Temp[i]) - 48
			}
		}
		return true, Result * Mult
	}
	return false, 0
}

func itoa(Number int) (Result string) {
	for i := Number; i != 0; i = i / 10 {
		Temp := i % 10
		if Temp < 0 {
			Temp = -Temp
		}
		Result = string(rune(Temp+48)) + Result
	}
	if Number < 0 {
		Result = "-" + Result
	}
	if Number == 0 {
		Result = "0"
	}
	return
}
