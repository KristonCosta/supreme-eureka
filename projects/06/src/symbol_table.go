package main

import (
	"fmt"
)

type SymbolTable struct {
	table     map[string]string
	varOffset int
}

func createSymbolTable() (symbolTable SymbolTable) {
	symbolTable.varOffset = 0
	symbolTable.table = make(map[string]string)
	symbolTable.table["SP"] = "0000000000000000"
	symbolTable.table["R0"] = "0000000000000000"
	symbolTable.table["LCL"] = "0000000000000001"
	symbolTable.table["R1"] = "0000000000000001"
	symbolTable.table["ARG"] = "0000000000000010"
	symbolTable.table["R2"] = "0000000000000010"
	symbolTable.table["THIS"] = "0000000000000011"
	symbolTable.table["R3"] = "0000000000000011"
	symbolTable.table["THAT"] = "0000000000000100"
	symbolTable.table["R4"] = "0000000000000100"
	symbolTable.table["R5"] = "0000000000000101"
	symbolTable.table["R6"] = "0000000000000110"
	symbolTable.table["R7"] = "0000000000000111"
	symbolTable.table["R8"] = "0000000000001000"
	symbolTable.table["R9"] = "0000000000001001"
	symbolTable.table["R10"] = "0000000000001010"
	symbolTable.table["R11"] = "0000000000001011"
	symbolTable.table["R12"] = "0000000000001100"
	symbolTable.table["R13"] = "0000000000001101"
	symbolTable.table["R14"] = "0000000000001110"
	symbolTable.table["R15"] = "0000000000001111"
	symbolTable.table["SCREEN"] = fmt.Sprintf("%016b", 16384)
	symbolTable.table["KBD"] = fmt.Sprintf("%016b", 24576)
	return
}

func (symbolTable *SymbolTable) addSymbol(symbol string, instrNum int) {
	if _, ok := symbolTable.table[symbol]; ok {
		panic("Found duplicate symbol: " + symbol)
	}
	buf := fmt.Sprintf("%015b", instrNum)
	symbolTable.table[symbol] = "0" + buf
}

func (symbolTable *SymbolTable) addVariable(variable string) {
	symbolTable.table[variable] = fmt.Sprintf("%016b", 16+symbolTable.varOffset)
	symbolTable.varOffset += 1
}
