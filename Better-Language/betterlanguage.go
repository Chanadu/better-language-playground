package main

import (
	"bufio"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"

	"github.com/Chanadu/better-language/globals"
	"github.com/Chanadu/better-language/parser"
	"github.com/Chanadu/better-language/parser/interpreter"
	"github.com/Chanadu/better-language/parser/statements"
	"github.com/Chanadu/better-language/scanner"
	"github.com/Chanadu/better-language/utils"
)

func LineReader() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		line := input.Text()
		run(line)
	}

	if err := input.Err(); err != nil {
		utils.CreateAndReportErrorf("Error reading input: %e", err)
	}
}

func FileReader(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		utils.CreateAndReportErrorf("File Reading Error: %e", err)
		return
	}
	run(string(data))
}

func run(source string) {
	globals.HasErrors = false
	tokens, ok := runScanner(source)
	if !ok {
		utils.ReportDebugf("Errors found in scanner, exiting")
		os.Exit(1)
	}
	// printTokens(tokens)
	// utils.ReportDebugf("Scanner completed successfully")

	statement, ok := runParser(tokens)
	if !ok {
		utils.ReportDebugf("Errors found in parsing, exiting")
		os.Exit(1)
	}
	// utils.ReportDebugf("Parser completed successfully")

	ip := interpreter.NewInterpreter()
	ok = ip.Interpret(statement)
	if !ok {
		utils.ReportDebugf("Errors found in runtime, exiting")
		os.Exit(1)
	}
	// utils.ReportDebugf("Runtime completed successfully")
}

func runScanner(source string) (tokens []scanner.Token, ok bool) {
	sc := scanner.NewScanner(source)
	tokens, err := sc.ScanTokens()

	if err != nil {
		utils.CreateAndReportErrorf("Token Scanning Error: %e", err)
		return nil, false
	}
	if globals.HasErrors {
		utils.ReportDebugf("Errors found in scanning, exiting")
		return nil, false
	}
	return tokens, true
}

func runParser(tokens []scanner.Token) (stmts []statements.Statement, ok bool) {
	p := parser.NewParser(tokens)

	stmts, err := p.Parse()

	if err != nil {
		utils.CreateAndReportParsingErrorf("%s", err.Error())
		utils.ReportDebugf("Exited due to preparsing error")
		return nil, false
	}

	if globals.HasErrors {
		utils.ReportDebugf("Errors found in parsing, exiting")
		utils.ReportDebugf("Exited due to parsing error")
		return nil, false
	}

	return stmts, true
}

// func printExpressions(statements expressions.Expression) {
// 	utils.ReportDebugf("Parsed: %v", statements.ToGrammarString())
// }

//goland:noinspection GoUnusedFunction
func printTokens(tokens []scanner.Token) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	_, _ = fmt.Fprintln(w, color.BlueString("Type\tLexeme\tLiteral\tLine"))
	for _, t := range tokens {
		_, _ = fmt.Fprintln(w, color.BlueString(fmt.Sprintf("%s\t%#v\t%#v\t%d", t.Type.String(), t.Lexeme, t.Literal, t.Line)))
	}
	if err := w.Flush(); err != nil {
		utils.CreateAndReportErrorf("Error printing tokens: %e", err)
	}
}
