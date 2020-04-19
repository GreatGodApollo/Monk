package repl

import (
	"bufio"
	"fmt"
	"github.com/GreatGodApollo/monk/evaluator"
	"github.com/GreatGodApollo/monk/lexer"
	"github.com/GreatGodApollo/monk/object"
	"github.com/GreatGodApollo/monk/parser"
	"io"
	"io/ioutil"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func RunProgram(in io.Reader, out io.Writer, file string) {
	prog, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("could not open file: %s", err.Error())
	}
	env := object.NewEnvironment()

	l := lexer.New(string(prog))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}


func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
