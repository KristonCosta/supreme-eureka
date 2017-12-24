package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

func parse_a_command(fileLine FileLine) (instruction Instruction, wasParsed bool) {
	label_or_val := strings.TrimLeft(fileLine.line, "@")
	wasParsed = true
	instruction.instrType = A_COMMAND
	instruction.symbol = label_or_val
	return
}

func parse_l_command(fileLine FileLine) (instruction Instruction, wasParsed bool) {
	label := strings.TrimRight(strings.TrimLeft(fileLine.line, "("), ")")
	if len(label) == 0 {
		panic("Empty label at line " + strconv.Itoa(fileLine.num))
	}
	wasParsed = true
	instruction.instrType = L_COMMAND
	instruction.symbol = label
	return
}


func parse_c_command(fileLine FileLine) (instruction Instruction, wasParsed bool) {
	instruction.instrType = C_COMMAND
	if strings.Contains(fileLine.line, "=") {
		splitCommand := strings.Split(fileLine.line, "=")
		if len((splitCommand)) != 2 {
			panic("Invalid line found at line number " + strconv.Itoa(fileLine.num))
		}
		dest, success := parse_dest(splitCommand[0])
		if !success {
			panic("Invalid destination at line number " + strconv.Itoa(fileLine.num))
		}
		instruction.dest = dest
		comp, success := parse_comp(splitCommand[1])
		if !success {
			panic("Invalid computation at line number " + strconv.Itoa(fileLine.num))
		}
		instruction.comp = comp
		instruction.jump = "000"
		wasParsed = true
		return
	} else if strings.Contains(fileLine.line, ";") {
		splitCommand := strings.Split(fileLine.line, ";")
		if len((splitCommand)) != 2 {
			panic("Invalid line found at line number " + strconv.Itoa(fileLine.num))
		}
		comp, success := parse_comp(splitCommand[0])
		if !success {
			panic("Invalid computation at line number " + strconv.Itoa(fileLine.num))
		}
		instruction.comp = comp
		jmp, success := parse_jump(splitCommand[1])
		if !success {
			panic("Invalid jump at line number " + strconv.Itoa(fileLine.num))
		}
		instruction.jump = jmp
		instruction.dest = "000"
		wasParsed = true
		return
	} else {
		panic("Invalid syntax at line number " + strconv.Itoa(fileLine.num))
	}
}

func parse_line(fileLine *FileLine) (instruction Instruction, wasParsed bool) {
	wasParsed = false
	fileLine.line = strings.TrimSpace(fileLine.line)
	if len(fileLine.line) == 0 {
		return
	}
	commentSplit := strings.Split(fileLine.line, "//")
	if len(commentSplit) > 1 && len(commentSplit[0]) > 0 {
		fileLine.line = strings.TrimSpace(commentSplit[0])
	}

	if strings.HasPrefix(fileLine.line, "//") {
		return
	}
	if strings.HasPrefix(fileLine.line, "@") {
		instruction, wasParsed = parse_a_command(*fileLine)
		return
	}
	if strings.HasPrefix(fileLine.line, "(") &&
		strings.HasSuffix(fileLine.line, ")") {
		instruction, wasParsed = parse_l_command(*fileLine)
		return
	}
	instruction, wasParsed = parse_c_command(*fileLine)
	return
}

func parse_asm(filePath string) (instructions []Instruction) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	num := 1
	for scanner.Scan() {
		fileLine := FileLine{scanner.Text(), num}
		instruction, wasParsed := parse_line(&fileLine)
		instruction.rawInstr = fileLine
		if (wasParsed) {
			instructions = append(instructions, instruction)
		}
		num += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}
