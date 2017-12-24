package main

import (
	"strconv"
	"fmt"
)

func (asmGenerator *AsmGenerator) a_codegen(value string) (conv string) {
	if  int_val, err := strconv.Atoi(value); err == nil {
		buf := fmt.Sprintf("%015b", int_val)
		conv = "0" + buf
	} else {
		if bin_val, ok := asmGenerator.symbols.table[value]; ok {
			conv = bin_val
		} else {
			asmGenerator.symbols.addVariable(value)
			conv = asmGenerator.symbols.table[value]
		}
	}
	return
}

func (asmGenerator *AsmGenerator) asm_codegen() (compiledInstrs []string) {
	for _, instr := range asmGenerator.instructions {
		switch instr.instrType {
		case C_COMMAND:
			buf := "111" + instr.comp + instr.dest + instr.jump
			compiledInstrs = append(compiledInstrs, buf)
		case A_COMMAND:
			buf := asmGenerator.a_codegen(instr.symbol)
			compiledInstrs = append(compiledInstrs, buf)
			}
	}
	return
}