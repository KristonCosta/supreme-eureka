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

// Put your code here.


	@lc 		// Last char
	M=0
	@diff
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
	D;JEQ
	@sp	
	M=0
	@KBD
	D=M
	@WHITE
	D;JEQ
	@BLACK
	0;JMP
(WHITE)
	@color
	M=0
	@SETLC
	0;JMP
(BLACK)
	@color
	M=-1
(SETLC)
	@KBD
	D=M
	@lc
	M=D
(UPDATE)
	@sp
	D=M
	@8192
	D=D-A
	@LOOP
	D;JGE
	
	@SCREEN
	D=A
	@sp
	D=D+M
	@cur
	M=D
	@color
	D=M
	@cur
	A=M
	M=D
	@sp
	M=M+1
	@LOOP
	0;JMP


	
