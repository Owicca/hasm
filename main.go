package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	parser = NewParser()
)

func main() {
	var (
		input  = flag.String("i", "./program.asm", "Path to input file")
		output = flag.String("o", "./program.binary", "Path to output file")
	)
	flag.Parse()

	f_in, err := os.Open(*input)
	if err != nil {
		log.Fatalf("Could not create/open output file %s (%s)", *input, err)
	}
	defer f_in.Close()

	result := []string{}

	scanner := bufio.NewScanner(f_in)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 || line[0] == '/' && line[1] == '/' {
			continue
		}
		instruction, err := parser.GetInstruction(line)
		if err != nil {
			log.Fatalf("Could not parse instruction (%s)", err)
		}

		result = append(result, instruction.GetBinary())
		log.Println(instruction.GetBinary())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error while scanning (%s)", err)
	}

	f_out, err := os.OpenFile(*output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Could not create/open output file %s (%s)", *output, err)
	}
	defer f_out.Close()

	output_content := strings.Join(result, "\n")
	n_written, err := f_out.WriteString(output_content)
	if err != nil {
		log.Fatalf("Could not write to output file %s (%s)", *output, err)
	}
	log.Println(n_written)
}
