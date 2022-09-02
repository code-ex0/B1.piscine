package challenge_go

func Split(s, sep string) []string {
	LastIndex := 0
	Resolve := false
	Lens := len(sep)
	var Result string
	var TempString string
	for Resolve == false {
		n := Index(s[LastIndex:], sep)
		if !(n == -1) {
			TempString = s[LastIndex : n+LastIndex]
			LastIndex = n + Lens + LastIndex
			Result += TempString + " "
		} else {
			TempString = s[LastIndex:]
			Resolve = true
			Result += TempString
		}

	}
	return SplitWhiteSpaces(Result)
}
