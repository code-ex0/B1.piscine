package challenge_go

func Capitalize(s string) string {
	Out := ""
	Cap := true

	for i := 0; i < len(s); i++ {
		if (rune(s[i]) >= 'a' && rune(s[i]) <= 'z') && Cap {
			Out += string(rune(s[i] - 32))
			Cap = false
		} else if rune(s[i]) >= 'A' && rune(s[i]) <= 'Z' {
			if Cap {
				Cap = false
				Out += string(rune(s[i]))
			} else {
				Out += string(rune(s[i] + 32))
			}
		} else if !((rune(s[i]) >= 'a' && rune(s[i]) <= 'z') || (rune(s[i]) >= 'A' && rune(s[i]) <= 'Z') || (rune(s[i]) >= '0' && rune(s[i]) <= '9')) {
			Cap = true
			Out += string(rune(s[i]))
		} else if rune(s[i]) >= '0' && rune(s[i]) <= '9' {
			Out += string(rune(s[i]))
			Cap = false
		} else {
			Out += string(rune(s[i]))
		}

	}
	return Out
}
