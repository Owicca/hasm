package objects

import (
	"fmt"
	"log"
	"strings"

	"github.com/Owicca/hasm/utils"
)

type C struct {
	mnemonic string
	il       int16
}

func (c *C) GetMnemonic() string {
	return c.mnemonic
}

func (c *C) SetMnemonic(mnemonic string) error {
	if mnemonic[0] == '@' {
		return fmt.Errorf("This mnemonic is not an C instruction")
	}
	c.mnemonic = utils.FormatInstruction(mnemonic)

	instructionInt, err := c.parseInstruction()
	if err != nil {
		return fmt.Errorf("Could not parse instruction %s (%s)", c.mnemonic, err)
	}
	log.Println(instructionInt)
	c.il += instructionInt

	return nil
}

func (c *C) GetBinary() string {
	binaryString := []byte(fmt.Sprintf("%016b", c.il))
	binaryString[0] = '1'
	binaryString[1] = '1'
	binaryString[2] = '1'

	return string(binaryString)
}

func (c *C) parseInstruction() (int16, error) {
	var (
		instruction int16  = 0
		mnemonic    string = c.mnemonic
	)

	if strings.Contains(mnemonic, ";") {
		jccIndex := strings.Index(mnemonic, ";")
		if jccIndex < 1 {
			return 0, fmt.Errorf("Invalid JCC part, JCC could not be the first part")
		} else if jccIndex+1 == len(mnemonic) {
			return 0, fmt.Errorf("Invalid JCC part, JCC could not be empty")
		}
		mnemonic = mnemonic[:jccIndex]
		jccStringVal := c.mnemonic[jccIndex+1:]
		jccVal, ok := JCCTable[jccStringVal]
		if !ok {
			return 0, fmt.Errorf("Invalid JCC part, JCC not recognized")
		}
		instruction += jccVal
	}
	if strings.Contains(c.mnemonic, "=") {
		eqIndex := strings.Index(mnemonic, "=")
		if eqIndex < 1 {
			return 0, fmt.Errorf("Invalid Dest/Comp part, '=' could not be the first character")
		} else if eqIndex+1 == len(mnemonic) {
			return 0, fmt.Errorf("Invalid Dest/Comp part, '=' could not be the last character")
		}

		parts := strings.Split(mnemonic, "=")

		destVal, err := c.parseDest(parts[0])
		if err != nil {
			return 0, fmt.Errorf("Invalid Dest part (%s)", err)
		}
		mnemonic = mnemonic[:eqIndex]

		if strings.Contains(parts[1], "M") {
			instruction += 1 << 12
		}

		compVal, err := c.parseComp(parts[1])
		if err != nil {
			return 0, fmt.Errorf("Invalid Comp part (%s)", err)
		}

		instruction += destVal << 3
		instruction += compVal << 6
	} else {
		plainCompVal, err := c.parseComp(mnemonic)
		if err != nil {
			return 0, fmt.Errorf("Invalid Plain Comp part (%s)", err)
		}

		instruction += plainCompVal << 6
	}

	return instruction, nil
}

func (c *C) parseDest(dest string) (int16, error) {
	var destInt int16 = 0

	if len(dest) == 0 {
		return 0, fmt.Errorf("The Dest is empty (%s)", dest)
	}

	destVal, ok := DestTable[dest]
	if !ok {
		return 0, fmt.Errorf("Invalid Dest value (%s)", dest)
	}
	destInt += destVal

	return destInt, nil
}

func (c *C) parseComp(comp string) (int16, error) {
	var compInt int16 = 0

	if len(comp) == 0 {
		return 0, fmt.Errorf("The Comp is empty (%s)", comp)
	}

	compVal, ok := CompTable[comp]
	if !ok {
		return 0, fmt.Errorf("Invalid Comp value (%s)", comp)
	}
	compInt += compVal

	return compInt, nil
}
