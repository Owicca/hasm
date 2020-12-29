package utils

import "strings"

func FormatInstruction(mnemonic string) string {
	fields := strings.Fields(mnemonic)
	return strings.ToUpper(strings.Join(fields, ""))
}
