#include <string>
#include <sstream>
#include <iostream>
#include "Assembler.h"
#include "CommandType.h"
#include "DestinationTypes.h"
#include "Utils.h"

Assembler::Assembler(std::string inputFile) {
	inStream.open(inputFile);
	lineNum = 0;
}

Assembler::~Assembler() {
	std::cout << "Deleting the assembler.\n";
	inStream.close();
}

bool Assembler::hasMoreCommands() {
	return (inStream.peek() != EOF);
}

void Assembler::advance() {
	if (!hasMoreCommands()) {
		throw std::out_of_range("Tried to advance past end of file.");
	}
	lineNum += 1;
	std::getline(inStream, currentCommand);
	trim(currentCommand);
	if (currentCommand.length() <= 0) {
		advance();
	}
	if (currentCommand.length() >= 2 && currentCommand.substr(0,2) == "//") {
		advance();
	}
	if (currentCommand.front() == '@') {
		commandType = CommandType::A_COMMAND;
	}
	else if (currentCommand.front() == '(') {
		commandType = CommandType::L_COMMAND;
	}
	else {
		commandType = CommandType::C_COMMAND;
		parseCompCommand();
	}
}

std::string Assembler::getCommand() {
	return currentCommand;
}

CommandType Assembler::getCommandType() {
	return commandType;
}

void Assembler::parseCompCommand() {
	// Screw it we are doing it live
	try {
		if (currentCommand.find("=") != std::string::npos) {
			parseAssignment();
		}
		else if (currentCommand.find(";") != std::string::npos) {
			parseJump();
		}
		else {
			std::stringstream ss;
			ss << "Invalid expression (" << currentCommand << ")" << std::endl;
			throw std::invalid_argument(ss.str());
		}
	}
	catch (std::invalid_argument e) {
		std::stringstream ss;
		ss << "Invalid syntax found at line " << lineNum << std::endl << e.what();
		throw std::invalid_argument(ss.str());
	}
}

void Assembler::parseAssignment() { 
	int equalsIndex = currentCommand.find("=");
	std::string lhs = currentCommand.substr(0, equalsIndex);
	dest = parseVariable(lhs);
	std::string rhs = currentCommand.substr(equalsIndex + 1, currentCommand.length());
	std::cout << lhs << std::endl;
	std::cout << rhs << std::endl;
}

Destination Assembler::parseVariable(std::string token) {
	if (token.length() == 1) {
		if (token == "A") {
			return A_DEST;
		}
		if (token == "D") {
			return D_DEST;
		}
		if (token == "M") {
			return M_DEST;
		}
	}
	std::stringstream ss;
	ss << "Invalid argument " << token << std::endl;
	throw std::invalid_argument(ss.str());
}

void Assembler::parseJump() {}