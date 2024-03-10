package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ra2a(a string) (int, string) {
	switch a {
	case "I":
		return 1, "r"
	case "II":
		return 2, "r"
	case "III":
		return 3, "r"
	case "IV":
		return 4, "r"
	case "V":
		return 5, "r"
	case "VI":
		return 6, "r"
	case "VII":
		return 7, "r"
	case "VIII":
		return 8, "r"
	case "IX":
		return 9, "r"
	case "X":
		return 10, "r"
	default:
		tmp := 0
		tmp, err := strconv.Atoi(a)

		if err != nil || tmp > 10 {
			panic("Invalid number: " + a)
		}
		return tmp, "a"
	}
}
func a2r(a int) string {
	if a > 100 {
		panic("How did we get here? a is " + string(rune(a)))
	}

	cvs := []struct {
		a int
		r string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var r strings.Builder
	for _, cv := range cvs {
		for a >= cv.a {
			r.WriteString(cv.r)
			a -= cv.a
		}
	}

	return r.String()
}
func main() {
	rr := bufio.NewReader(os.Stdin)
	fmt.Println("Kata Academy test case #1, by tigrmango@gmail.com")
	for {
		fmt.Print(">")
		in, _ := rr.ReadString('\n')
		ina := strings.Split(in, " ")
		ina[2] = strings.TrimSpace(ina[2])

		if len(ina) != 3 {
			panic("Wrong input expression format!")
		}

		n1s := ina[0]
		op := ina[1]
		n2s := ina[2]

		n1, t1 := ra2a(n1s)
		n2, t2 := ra2a(n2s)
		if t1 != t2 {
			panic("Mismatching numbers: " + n1s + " and " + n2s + "!")
		}

		out := 0

		switch op {
		case "+":
			out = n1 + n2
		case "-":
			out = n1 - n2
		case "*":
			out = n1 * n2
		case "/":
			out = n1 / n2
		default:
			panic("Wrong operator type!")
		}
		if t1 == "r" {
			if out < 1 {
				panic("Roman numbers must be >0!")
			}
			fmt.Println(a2r(out))
		} else {
			fmt.Println(out)
		}

	}
}
