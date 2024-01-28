package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	OP_INC             = '>'
	OP_DEC             = '<'
	OP_ADD             = '+'
	OP_SUB             = '-'
	OP_OUTPUT          = '.'
	OP_INPUT           = ','
	OP_JUMP_IF_ZERO    = '['
	OP_JUMP_IF_NONZERO = ']'
)

type IR struct {
	OpCode byte
	OpRand int
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: No input supplied\n")
		log.Fatalf("Usage: gobrain <input.bf>\n")
	}

	filepath := os.Args[1]
	if strings.Split(filepath, ".")[1] != "bf" {
		log.Fatalf("ERROR: Input must be brainfuck file in form input.bf")
	}

	program := getContents(filepath)

	fmt.Println(convertToIR(program))
}

func getContents(filepath string) []byte {
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("ERROR: Could not open %s\n", filepath)
	}

	return content
}

func convertToIR(program []byte) []IR {
	ip := 0
	var char byte
	var instructions []IR

	for ; ; ip++ {
		char, ip = getNextToken(program, ip)
		if ip == -1 {
			break
		}

		instructions = append(instructions, IR{char, 1})
	}
	return instructions
}

func getNextToken(program []byte, ip int) (byte, int) {
	for ; ip < len(program); ip++ {
		switch program[ip] {
		case '<', '>', '+', '-', '.', ',', '[', ']':
			return program[ip], ip
		}
	}

	return 0, -1
}
