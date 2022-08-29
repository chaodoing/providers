package verification

import (
	`regexp`
)

func Password(ps string) bool {
	if len(ps) < 6 || len(ps) > 21 {
		return false
	}
	count := 1
	num := `[0-9]{1}`
	a_z := `[a-z]{1}`
	A_Z := `[A-Z]{1}`
	symbol := `[!@#~$%^&*()+-=|_.]{1}`
	if b, err := regexp.MatchString(num, ps); !b || err == nil {
		count = count + 1
	}
	if b, err := regexp.MatchString(a_z, ps); !b || err == nil {
		
		count = count + 1
	}
	if b, err := regexp.MatchString(A_Z, ps); !b || err == nil {
		count = count + 1
	}
	if b, err := regexp.MatchString(symbol, ps); !b || err == nil {
		count = count + 1
	}
	return count >= 4
}
