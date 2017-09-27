// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

	@lc 		// Last char
	M=0
	@color		// Will just contain which value to fill the screen with
	M=0
	@sp			// Screen pointer
	M=0
(LOOP)
	@lc			
	D=M
	@KBD
	D=D-M 
	@UPDATE
	D;JEQ		// If the characters aren't different we can continue drawing
	@SCREEN
	D=A
	@sp	
	M=D
	@KBD
	D=M
	@SETLC
	D;JEQ		// If KBD is 0 then we can just just ignore setting black
(BLACK)
	D=-1
(SETLC)
	@color
	M=D
	@KBD
	D=M
	@lc
	M=D
(UPDATE)
	@sp
	D=M
	@24576		// Hardcoded...gotta fix
	D=D-A
	@LOOP
	D;JGE		// We are at the end, just keep checking for new keys
	
	@color
	D=M
	@sp			// Get the address at sp location
	A=M
	M=D         // Then set the color for that screen location
	@sp
	M=M+1
	@LOOP
	0;JMP


	
