package objects

var (
	JCCTable map[string]int16 = map[string]int16{
		"JGT": int16(1), //0 0 1
		"JEQ": int16(2), //0 1 0
		"JGE": int16(3), //0 1 1
		"JLT": int16(4), //1 0 0
		"JNE": int16(5), //1 0 1
		"JLE": int16(6), //1 1 0
		"JMP": int16(7), //1 1 1
	}
	CompTable map[string]int16 = map[string]int16{
		"0":   int16(42), //1 0 1 0 1 0
		"1":   int16(63), //1 1 1 1 1 1
		"-1":  int16(58), //1 1 1 0 1 0
		"D":   int16(12), //0 0 1 1 0 0
		"A":   int16(48), //1 1 0 0 0 0
		"M":   int16(48), //1 1 0 0 0 0 when a = 1
		"!D":  int16(13), //0 0 1 1 0 1
		"!A":  int16(49), //1 1 0 0 0 1
		"!M":  int16(49), //1 1 0 0 0 1 when a = 1
		"-D":  int16(15), //0 0 1 1 1 1
		"-A":  int16(51), //1 1 0 0 1 1
		"D+1": int16(31), //0 1 1 1 1 1
		"A+1": int16(55), //1 1 0 1 1 1
		"M+1": int16(55), //1 1 0 1 1 1 when a = 1
		"D-1": int16(14), //0 0 1 1 1 0
		"A-1": int16(50), //1 1 0 0 1 0
		"M-1": int16(50), //1 1 0 0 1 0 when a = 1
		"D+A": int16(2),  //0 0 0 0 1 0
		"D+M": int16(2),  //0 0 0 0 1 0 when a = 1
		"D-A": int16(19), //0 1 0 0 1 1
		"D-M": int16(19), //0 1 0 0 1 1 when a = 1
		"A-D": int16(7),  //0 0 0 1 1 1
		"M-D": int16(7),  //0 0 0 1 1 1
		"D&A": int16(0),  //0 0 0 0 0 0
		"D&M": int16(0),  //0 0 0 0 0 0 when a = 1
		"D|A": int16(21), //0 1 0 1 0 1
		"D|M": int16(21), //0 1 0 1 0 1 when a = 1
	}
	DestTable map[string]int16 = map[string]int16{
		"M":   int16(1), // 0 0 1
		"D":   int16(2), // 0 1 0
		"MD":  int16(3), // 0 1 1
		"A":   int16(4), // 1 0 0
		"AM":  int16(5), // 1 0 1
		"AD":  int16(6), // 1 1 0
		"AMD": int16(7), // 1 1 1
	}
)
