package main

import "fmt"

func main() {
	s := "()"
	fmt.Println(isValid(s))
	fmt.Println(isValidv2(s))
}

func isValid(s string) bool {
	stack := []string{}
	for x := 0; x < len(s); x++ {
		if string(s[x]) == "{" || string(s[x]) == "[" || string(s[x]) == "(" {
			stack = append(stack, string(s[x]))
		}
		if string(s[x]) == "}" || string(s[x]) == "]" || string(s[x]) == ")" {
			if len(stack) == 0 {
				return false
			}
			popped := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			switch popped {
			case "{":
				if string(s[x]) != "}" {
					return false
				}
			case "[":
				if string(s[x]) != "]" {
					return false
				}
			case "(":
				if string(s[x]) != ")" {
					return false
				}
			}
		}
	}

	if len(stack) != 0 {
		return false
	}
	return true
}

func isValidv2(s string) bool {
	stack := []string{}
	m1 := make(map[string]string)
	m1["("] = ")"
	m1["["] = "]"
	m1["{"] = "}"

	for x := 0; x < len(s); x++ {
		st := string(s[x])
		if _, ok := m1[st]; ok {
			stack = append(stack, st) //push
		} else if len(stack) == 0 || st != m1[stack[len(stack)-1]] {
			return false
		} else {
			stack = stack[0 : len(stack)-1] //pop
		}

	}

	if len(stack) != 0 {
		return false
	}
	return true
}
