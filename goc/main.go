package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mrgsrylm/compiler/goc/src/compiler"
	"github.com/mrgsrylm/compiler/goc/src/evaluator"
	"github.com/mrgsrylm/compiler/goc/src/lexer"
	"github.com/mrgsrylm/compiler/goc/src/object"
	"github.com/mrgsrylm/compiler/goc/src/parser"
	"github.com/mrgsrylm/compiler/goc/src/vm"

	"github.com/mrgsrylm/compiler/goc/src/repl"
)

func main() {
	// Start Guuc REPL
	if len(os.Args) == 1 {
		fmt.Println("This is the Guuc programming language!")
		fmt.Println("Feel free to type in commands")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// Run a Guuc script
	if err := runScript(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func runScript(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read %s: %v", filename, err)
	}

	p := parser.New(lexer.New(string(data)))
	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		return errors.New(strings.Join(p.Errors(), "\n"))
	}

	// Process macros
	macroEnv := object.NewEnvironment()
	evaluator.DefineMacros(program, macroEnv)
	expanded := evaluator.ExpandMacros(program, macroEnv)

	// Compile the AST to bytecode
	c := compiler.New()
	if err := c.Compile(expanded); err != nil {
		return fmt.Errorf("Woops! Compilation failed: %s", err)
	}

	// Run bytecode instructions
	machine := vm.New(c.Bytecode())
	if err := machine.Run(); err != nil {
		return fmt.Errorf("Woops! Executing bytecode failed: %s", err)
	}

	return nil
}
