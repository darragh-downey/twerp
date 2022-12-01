package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/darragh-downey/twerp/pkg/lexer"
	"github.com/darragh-downey/twerp/pkg/parser"
)

const PROMPT = ">> "

const TWERP = ` 
TWERP
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, TWERP)
	io.WriteString(out, "Whoopsie Daisy!! Some muppet (probably the idiot behind the keyboard) fat fingered their way to Valhalla!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
