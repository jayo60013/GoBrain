package main

import (
	"fmt"
	"os"
	"strings"

	. "gobrain/lexer"
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

type Op struct {
	OpCode  byte
	Operand int
}

type Program struct {
	instructions []Op
	count        int
	capacity     int
}

type Memory struct {
	data     []byte
	count    int
	capacity int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: No input supplied\n")
		fmt.Printf("Usage: %s <input.bf>\n", os.Args[0])
		os.Exit(1)
	}

	filepath := os.Args[1]
	if strings.Split(filepath, ".")[1] != "bf" {
		fmt.Println("ERROR: Input must be brainfuck file in form input.bf")
		os.Exit(1)
	}

	contents := getContents(filepath)
	program := convertToIR(contents)

	var memory Memory
	memory.data = append(memory.data, 0)
	memory.count++
	head := 0
	ip := 0

	for ip < program.count {

		instruction := program.instructions[ip]
		switch instruction.OpCode {

		case OP_ADD:
			memory.data[head] += byte(instruction.Operand)
			ip++

		case OP_SUB:
			memory.data[head] -= byte(instruction.Operand)
			ip++

		case OP_LEFT:
			if head < instruction.Operand {
				fmt.Printf("RUNTIME ERROR: Memory underflow\n")
				os.Exit(1)
			}
			head -= instruction.Operand
			ip++

		case OP_RIGHT:
			head += instruction.Operand
			for head >= memory.count {
				memory.data = append(memory.data, 0)
				memory.count += 1
			}
			ip++

		case OP_INPUT:
			buffer := make([]byte, 1)
			_, err := os.Stdin.Read(buffer)
			if err != nil {
				fmt.Printf("RUNTIME ERROR: Problem reading input\n")
				os.Exit(1)
			}

			// Extract the first byte from the string
			memory.data[head] = buffer[0]
			ip++

		case OP_OUTPUT:
			for i := 0; i < instruction.Operand; i++ {
				fmt.Printf("%c", memory.data[head])
			}
			ip++

		case OP_JUMP_IF_ZERO:
			if memory.data[head] == 0 {
				ip = instruction.Operand
			} else {
				ip++
			}

		case OP_JUMP_IF_NONZERO:
			if memory.data[head] != 0 {
				ip = instruction.Operand
			} else {
				ip++
			}
		}
	}
}

func getContents(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ERROR: Could not open %s\n", filepath)
		os.Exit(1)
	}

	return content
}

func convertToIR(program []byte) Program {
	var instr []Op
	var addrStack []int
	lexer := Lexer{
		Content: program,
		Ip:      0,
		Count:   len(program),
	}

	ch := lexer.Next()

	for ch != 0 {
		switch ch {
		case OP_ADD, OP_SUB, OP_LEFT, OP_RIGHT, OP_INPUT, OP_OUTPUT:
			count := 1
			next := lexer.Next()
			for next == ch {
				count += 1
				next = lexer.Next()
			}
			instr = append(instr, Op{ch, count})
			ch = next

		case OP_JUMP_IF_ZERO:
			addrStack = append(addrStack, len(instr))
			instr = append(instr, Op{ch, 0})
			ch = lexer.Next()
		case OP_JUMP_IF_NONZERO:
			if len(addrStack) == 0 {
				fmt.Println("ERROR: Stack underflow")
				os.Exit(1)
			}

			addr := addrStack[len(addrStack)-1]
			addrStack = addrStack[:len(addrStack)-1]

			instr = append(instr, Op{ch, addr + 1})
			instr[addr].Operand = len(instr)
			ch = lexer.Next()
		}
	}

	return Program{instructions: instr, count: len(instr), capacity: 0}
}
