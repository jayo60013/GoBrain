package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	OP_LEFT            = '<'
	OP_RIGHT           = '>'
	OP_ADD             = '+'
	OP_SUB             = '-'
	OP_OUTPUT          = '.'
	OP_INPUT           = ','
	OP_JUMP_IF_ZERO    = '['
	OP_JUMP_IF_NONZERO = ']'
)

const TOKENS = "<>+-.,[]"

type Op struct {
	OpCode  byte
	Operand int
}

type Program struct {
	instructions []Op
	count        int
	capacity     int
}

type Lexer struct {
	content []byte
	pos     int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: No input supplied")
		fmt.Println("Usage: gobrain <input.bf>")
		os.Exit(1)
	}

	filepath := os.Args[1]
	if strings.Split(filepath, ".")[1] != "bf" {
		fmt.Println("ERROR: Input must be brainfuck file in form input.bf")
		os.Exit(1)
	}

	contents := getContents(filepath)
	convertToIR(contents)
}

func getContents(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ERROR: Could not open %s\n", filepath)
		os.Exit(1)
	}

	return content
}

func convertToIR(program []byte) []Op {
	var instr []Op
	lexer := Lexer{
		content: program,
		pos:     0,
	}
	ch := lexer.next()

	for ch != 0 {
		switch ch {
		case OP_ADD, OP_SUB, OP_LEFT, OP_RIGHT, OP_INPUT, OP_OUTPUT:
			count := 1
			next := lexer.next()
			for next == ch {
				count += 1
				next = lexer.next()
			}
			instr = append(instr, Op{ch, count})
			ch = next

		case OP_JUMP_IF_ZERO:
			ch = lexer.next()
		case OP_JUMP_IF_NONZERO:
			ch = lexer.next()
		}
	}

	for _, op := range instr {
		fmt.Printf("%c (%d)\n", op.OpCode, op.Operand)
	}
	return instr
}

func (l *Lexer) next() byte {
	for l.pos < len(l.content) && !isToken(l.content[l.pos]) {
		l.pos += 1
	}

	if l.pos >= len(l.content) {
		return 0
	}

	l.pos += 1
	return l.content[l.pos-1]
}

func isToken(symbol byte) bool {
	for _, token := range []byte(TOKENS) {
		if token == symbol {
			return true
		}
	}
	return false
}
