package main

import (
	"math"
	"unicode"
)

func atoi(str string) int {
	if str == "" {
		return 0
	}

	i := 0
	n := len(str)

	// check for whitespaces preceeding the number
	for i < n && str[i] == ' ' {
		i++
	}

	//check for + or - preceeding the number
	sign := 1
	if i < n {
		switch str[i] {
		case '-':
			sign = -1
			i++
		case '+':
			i++
		}
	}

	//check for unicode number, and convert to int
	num := 0
	for i < n && unicode.IsDigit(rune(str[i])) {
		digit := int(str[i] - '0')

		//check for overflow
		if num > math.MaxInt32/10 || (num == (math.MaxInt32/10) && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		num = num*10 + digit
		i++
	}
	return num * sign
}
