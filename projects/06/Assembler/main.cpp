#include <iostream>
#include "Assembler.h"

void cleanup(Assembler* assembler) {
	delete assembler;
	char c;
	std::cin >> c;
}

int main() {
	Assembler* assembler = new Assembler("Text.txt");
	while (assembler->hasMoreCommands()) {
		try {
			assembler->advance();
		}
		catch (std::invalid_argument e) {
			std::cout << ">>>> Error: " << e.what() << std::endl;
			cleanup(assembler);
			exit(-1);
		}
		std::cout << assembler->getCommand() << std::endl;
		switch (assembler->getCommandType()) {
		case A_COMMAND: std::cout << "A_COMMAND\n";
			break;
		case C_COMMAND: std::cout << "C_COMMAND\n";
			break;
		case L_COMMAND: std::cout << "L_COMMAND\n";
			break;
		}
	}
	cleanup(assembler);
}