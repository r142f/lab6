package main

import (
	"fmt"
	"lab6/calc_visitor"
	"lab6/parser_visitor"
	"lab6/print_visitor"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Wrong args! Proper usage: go run ./main [expression]")
	}

	expression := os.Args[1]

	parser := parser_visitor.ParserVisitor{}
	tokens, err := parser.Parse(expression)
	if err != nil {
		log.Fatal(err)
	}

	printer := print_visitor.PrintVisitor{}
	fmt.Println(printer.Print(tokens))

	calc := calc_visitor.CalcVisitor{}
	result, err := calc.Calc(tokens)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
