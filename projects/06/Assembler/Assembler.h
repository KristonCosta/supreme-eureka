#pragma once
#include <fstream>
#include <string>
#include "CommandType.h"

class Assembler
{
private:
	std::ifstream inStream;
	CommandType commandType;
	std::string currentCommand;
	std::string symbol;
	std::string dest;
	std::string comp;
	std::string jump;
	int lineNum;
	void parseCompCommand();
	void parseAssignment();
	void parseJump();
	void parseVariable(std::string token);
public:
	Assembler(std::string inputFile);
	~Assembler();
	bool hasMoreCommands();
	void advance();
	CommandType getCommandType();
	std::string getSymbol();
	std::string getDest();
	std::string getComp();
	std::string getJump();
	std::string getCommand();
};
