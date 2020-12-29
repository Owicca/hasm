# An assembler for a dummy assembly language

## Explore golang low-level functionality by implementing an assembler
for a dummy assembly language

### Trying to see how expressive is golang in binary manipulation
by writing an assembler for a dummy assembly language that runs
on a 16 bit cpu

### Usage:
- `./assembler -i <path/to/input/file> -o <path/to/output/file>`

### Grammar:
- whitespace:
  - comment
  - newline
  - tab
  - space
- instruction:
  - A-instruction: @55
  - C-instruction: destination = computation; jump
- symbol:
  - variable: @variable_that_is_not_predefined_or_label
  - label: (LABEL)
  - pre-defined:
    - R1..R15: 0..15
    - SCREEN: 16384
    - KEYBOARD: 24576
    - SP: 0
    - LCL: 1
    - ARG: 2
    - THIS: 3
    - THAT: 4
