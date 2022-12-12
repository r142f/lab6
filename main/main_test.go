package main

import (
	"lab6/calc_visitor"
	"lab6/parser_visitor"
	"lab6/print_visitor"
	"testing"
)

func TestCalc(t *testing.T) {
	var tests = []struct {
		input      string
		want       int
		parseError bool
		calcError  bool
	}{
		{"(23 + 10) * 5 - 3 * (32 + 5) * (10 - 4 * 5) + 8 / 2", (23+10)*5 - 3*(32+5)*(10-4*5) + 8/2, false, false},
		{"(2 + 2) * ((2 * 2) + ((2 * 2) * (2 * 2)))", (2 + 2) * ((2 * 2) + ((2 * 2) * (2 * 2))), false, false},
		{"1 + 2", 1 + 2, false, false},
		{"2 - 1 - 2", 2 - 1 - 2, false, false},
		{"2 - (1 - 2)", 2 - (1 - 2), false, false},
		{"2 / 4 / 8", 2 / 4 / 8, false, false},
		{"2 / (16 / 8)", 2 / (16 / 8), false, false},
		{"2 ^ 2", 4, true, false},
		{"2 *(2", 0, true, false},
		{"2 *(2)", 4, false, false},
		{"2 2", 4, false, true},
		{"", 0, false, false},
	}

	parser := parser_visitor.ParserVisitor{}
	calc := calc_visitor.CalcVisitor{}

	for _, test := range tests {
		tokens, err := parser.Parse(test.input)
		if test.parseError && err == nil {
			t.Errorf("Parse(%q), expected error, got %v", test.input, err)
		}
		if !test.parseError && err != nil {
			t.Errorf("Parse(%q), unexpected error, got %v", test.input, err)
		}
		if test.parseError {
			continue
		}

		got, err := calc.Calc(tokens)
		if test.calcError && err == nil {
			t.Errorf("Calc(%q), expected error, got %v", test.input, err)
		}
		if !test.calcError && err != nil {
			t.Errorf("Calc(%q), unexpected error, got %v", test.input, err)
		}
		if test.calcError {
			continue
		}

		if got != test.want {
			t.Errorf("Calc(%q) = %v", test.input, got)
		}
	}
}

func TestPrint(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"3 + 4 * (2 - 1)", "3 4 2 1 - * +"},
		{"1 + 2", "1 2 +"},
	}

	parser := parser_visitor.ParserVisitor{}
	printer := print_visitor.PrintVisitor{}

	for _, test := range tests {
		tokens, _ := parser.Parse(test.input)
		if got := printer.Print(tokens); got != test.want {
			t.Errorf("Print(%q) = %v", test.input, got)
		}
	}
}
