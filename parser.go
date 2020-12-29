package main

import (
	"fmt"

	"github.com/Owicca/hasm/objects"
)

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) GetInstruction(line string) (objects.Instruction, error) {
	var (
		instruction objects.Instruction
		tp          rune = 'A'
	)

	if len(line) < 1 || line[0] == '/' && line[1] == '/' {
		return &objects.A{}, nil
	}
	if line[0] == '@' {
		instruction = &objects.A{}
	} else {
		instruction = &objects.C{}
		tp = 'C'
	}

	if err := instruction.SetMnemonic(line); err != nil {
		return instruction, fmt.Errorf("Could not parse %c instruction %s (%s)", tp, line, err)
	}

	return instruction, nil
}
