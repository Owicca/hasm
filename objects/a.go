package objects

import (
	"fmt"
	"strconv"
	"strings"
)

type A struct {
	mnemonic string
	il       int16
}

func (a *A) GetMnemonic() string {
	return a.mnemonic
}

func (a *A) SetMnemonic(mnemonic string) error {
	if mnemonic[0] != '@' || strings.Contains(mnemonic, ";") || strings.Contains(mnemonic, "=") {
		return fmt.Errorf("This mnemonic is not an A instruction")
	}
	a.mnemonic = mnemonic
	value, err := strconv.Atoi(a.mnemonic[1:])

	if err != nil {
		return fmt.Errorf("Could not convert %s to an integer (%s)",
			a.mnemonic[1:], err)
	}
	a.il = int16(value)

	return nil
}

func (a *A) GetBinary() string {
	return fmt.Sprintf("%016b", a.il)
}
