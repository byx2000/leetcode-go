package basic_calculator

import "strings"

// https://leetcode.cn/problems/basic-calculator/
// https://leetcode.cn/problems/basic-calculator-ii/
// https://www.lintcode.com/problem/849

var expr string
var index int

func next() uint8 {
	index++
	return expr[index-1]
}

func peek() uint8 {
	return expr[index]
}

func end() bool {
	return index == len(expr)
}

func parseInt() int {
	r := 0
	for !end() && peek() >= '0' && peek() <= '9' {
		d := next()
		r = r*10 + int(d-'0')
	}
	return r
}

// factor = int | (expr)
func parseFactor() int {
	if peek() == '(' {
		next()
		r := parseExpr()
		next()
		return r
	} else {
		return parseInt()
	}
}

// term = factor *|/ factor *|/ ... *|/ factor
func parseTerm() int {
	r := parseFactor()
	for !end() && (peek() == '*' || peek() == '/') {
		op := next()
		if op == '*' {
			r *= parseFactor()
		} else {
			r /= parseFactor()
		}
	}
	return r
}

// expr = term +|- term +|- ... +|- term
func parseExpr() int {
	r := parseTerm()
	for !end() && (peek() == '+' || peek() == '-') {
		op := next()
		if op == '+' {
			r += parseTerm()
		} else {
			r -= parseTerm()
		}
	}
	return r
}

func Calculate(s string) int {
	expr = strings.Replace(s, " ", "", -1)
	index = 0
	return parseExpr()
}
