package leetcode

import "math"

func myAtoi(s string) int {
	auto := automaton{}
	auto.Init()
	for _, c := range s {
		auto.Get(c)
	}
	return auto.sign * auto.sum
}

type automaton struct {
	state string
	table map[string][]string
	sign  int
	sum   int
}

func (a *automaton) Init() {
	a.state = "start"
	a.table = map[string][]string{
		"start":     {"start", "signed", "in_number", "end"},
		"signed":    {"end", "end", "in_number", "end"},
		"in_number": {"end", "end", "in_number", "end"},
		"end":       {"end", "end", "end", "end"},
	}
	a.sign = 1
}

func (a *automaton) GetCol(c rune) int {
	switch c {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	default:
		return 3
	}
}

func (a *automaton) Get(c rune) {
	a.state = a.table[a.state][a.GetCol(c)]
	if a.state == "in_number" {
		a.sum = a.sum*10 + int(c-'0')
		if a.sign == 1 {
			a.sum = min(a.sum, math.MaxInt32)
		} else {
			a.sum = min(a.sum, -math.MinInt32)
		}
	} else if a.state == "signed" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}
