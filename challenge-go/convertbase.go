package challenge_go

func ConvertBase(nbr string, baseFrom string, baseTo string) string {

	if len(baseFrom) == 2 && len(baseTo) == 10 {

	} else if len(baseTo) == 10 && len(baseFrom) == 2 {

	}

	return "fdfd"
}

func BinaryConvert(BinaryNumber string) int {
	mult := 0
	var result int
	for i := len(BinaryNumber) - 1; i >= 0; i-- {
		result += int(BinaryNumber[i]-48) * IterativePower(2, mult)
		mult++
	}
	return result
}

func DecimalCovert(DecimalNumber string) {
	DecimalNumberInt := Atoi(DecimalNumber)

	for DecimalNumberInt != 0 {

		DecimalNumberInt /= 2
	}
}
