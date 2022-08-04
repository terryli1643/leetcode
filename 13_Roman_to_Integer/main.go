package main

func main() {
	input := "III"
	romanToInt(input)
}

func romanToInt(s string) int {
	result := 0
	for i, _ := range s {
		li := string(s[i])
		if i+1 >= len(s) {
			result = result + getValue(li)
		} else {
			li1 := string(s[i+1])
			if li == "I" && li1 == "X" ||
				li == "I" && li1 == "V" ||
				li == "X" && li1 == "L" ||
				li == "X" && li1 == "C" ||
				li == "C" && li1 == "D" ||
				li == "C" && li1 == "M" {
				result = result - getValue(li)
			} else {
				result = result + getValue(li)
			}

		}
	}
	return result
}

func compair(s1, s2 string) bool {
	return getValue(s1) < getValue(s2)
}

func getValue(l string) int {
	switch l {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	}
	return 0
}
