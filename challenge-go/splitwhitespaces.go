package challenge_go

func SplitWhiteSpaces(s string) []string {
	var temp []string
	var tempString string
	for i, lettre := range s {
		if lettre != ' ' {
			tempString += string(lettre)
		} else if lettre == ' ' && s[i-1] != ' ' {
			temp = append(temp, tempString)
			tempString = ""
		}
	}

	temp = append(temp, tempString)
	return temp
}
