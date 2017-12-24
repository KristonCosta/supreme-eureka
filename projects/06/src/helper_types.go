package main

type InstructionType int

const (
	COMMENT InstructionType = iota
	A_COMMAND
	C_COMMAND
	L_COMMAND
)

type Instruction struct {
	rawInstr  		FileLine
	instrType 		InstructionType
	dest 			string
	symbol			string
	comp 			string
	jump 			string
}

type FileLine struct {
	line 	string
	num 	int
}

func (instructionType InstructionType) String() string {
	switch instructionType {
	case 1:
		return "A_COMMAND"
	case 2:
		return "C_COMMAND"
	case 3:
		return "L_COMMAND"
	default:
		return "Unknown"
	}
}
