package repl

import (
	"Interpreter/lexer"
	"Interpreter/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">>>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParserProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}
		_, _ = io.WriteString(out, program.String())
		_, _ = io.WriteString(out, "\n")
	}

}

func printParserErrors(out io.Writer, err []error) {
	for _, err := range err {
		_, _ = io.WriteString(out, "\t"+err.Error()+"\n")
	}
}
