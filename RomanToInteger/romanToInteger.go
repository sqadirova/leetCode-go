package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("V"))
	fmt.Println(romanToInt2("V"))

}

func romanToInt(s string) int {
	out := 0
	romanNum := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	cIdx := 0
	for cIdx < len(s) {
		c := string(s[cIdx])
		if c == "I" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "V" || next == "X" {
					out += romanNum[next] - romanNum[c]
					cIdx += 2
					continue
				}
			}
		}
		if c == "X" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "L" || next == "C" {
					out += romanNum[next] - romanNum[c]
					cIdx += 2
					continue
				}
			}
		}
		if c == "C" {
			if cIdx < len(s)-1 {
				next := string(s[cIdx+1])
				if next == "D" || next == "M" {
					out += romanNum[next] - romanNum[c]
					cIdx += 2
					continue
				}
			}
		}

		out += romanNum[c]
		cIdx++
	}

	return out
}

func romanToInt2(s string) int {
	romanNum := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	if len(s) == 1 {
		return romanNum[string(s[0])]
	}

	num := romanNum[string(s[len(s)-1])]

	for i := len(s) - 2; i >= 0; i-- {
		if romanNum[string(s[i])] < romanNum[string(s[i+1])] {
			num = num - romanNum[string(s[i])]
		} else {
			num = num + romanNum[string(s[i])]
		}
	}

	return num
}
