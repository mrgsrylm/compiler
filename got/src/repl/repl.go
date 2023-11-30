package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mrgsrylm/compiler/got/src/evaluator"
	"github.com/mrgsrylm/compiler/got/src/lexer"
	"github.com/mrgsrylm/compiler/got/src/object"
	"github.com/mrgsrylm/compiler/got/src/parser"
)

const prompt = ">> "

// Start starts Guut REPL.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
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

		// Process macros
		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		// Evaluate AST
		evaluated := evaluator.Eval(expanded, env)
		if evaluated == nil {
			continue
		}

		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, msg)
		io.WriteString(out, "\n")
	}
}
