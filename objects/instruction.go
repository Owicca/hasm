package objects

type Instruction interface {
	GetMnemonic() string
	GetBinary() string
	SetMnemonic(string) error
}
