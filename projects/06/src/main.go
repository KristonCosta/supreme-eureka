package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type AsmGenerator struct {
	filename     string
	instructions []Instruction
	symbols      SymbolTable
	compiled     []string
}

func (asmGenerator *AsmGenerator) write_listing() {
	f, err := os.Create(asmGenerator.filename + ".lst")
	if err != nil {
		panic("Could not create listing file.")
	}
	defer f.Close()
	compiledLine := 0
	w := bufio.NewWriter(f)
	for _, instr := range asmGenerator.instructions {
		if instr.instrType == L_COMMAND {
			formatted := fmt.Sprintf("|%-20s|%-20s|%16s|\n",
				instr.rawInstr.line, instr.instrType, "")
			w.WriteString(formatted)
		} else {
			formatted := fmt.Sprintf("|%-20s|%-20s|%16s|\n", instr.rawInstr.line,
				instr.instrType, asmGenerator.compiled[compiledLine])
			w.WriteString(formatted)
			compiledLine += 1
		}
	}
	w.Flush()
}

func (asmGenerator *AsmGenerator) generate_asm() {
	asmGenerator.instructions = parse_asm(asmGenerator.filename + ".asm")
	asmGenerator.symbols = createSymbolTable()
	instrNum := 0
	for _, instr := range asmGenerator.instructions {
		if instr.instrType == L_COMMAND {
			asmGenerator.symbols.addSymbol(instr.symbol, instrNum)
		} else {
			instrNum += 1
		}
	}
	asmGenerator.compiled = asmGenerator.asm_codegen()

}

func (asmGenerator *AsmGenerator) write_asm() {
	f, err := os.Create(asmGenerator.filename + ".hack")
	if err != nil {
		panic("Could not create output file.")
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	for _, cinstr := range asmGenerator.compiled {
		w.WriteString(cinstr + "\n")
	}
	w.Flush()
}

func main() {
	namePtr := flag.String("name", "", "Name of asm file you want to parse (required)")
	listingEnabled := flag.Bool("listing", false, "Generates a listing file")
	flag.Parse()
	if *namePtr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	var asmGenerator AsmGenerator
	asmGenerator.filename = *namePtr
	asmGenerator.generate_asm()
	asmGenerator.write_asm()
	if *listingEnabled {
		asmGenerator.write_listing()
	}
}
