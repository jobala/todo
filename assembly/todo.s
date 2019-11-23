; Executable name					:	todo
; Version 								: 1.0
; Description							: A simple todo application that helps users keep track of what they should do
;
;	Run it this way:
; ./todo
;
; Build using this command
; make todo

section .data
	Msg: db "Hello, World", 10 								; Create the message variable
	MsgLen equ $- Msg													; Calculate the length of the message

section .bss


section .text
	global _start

_start:
	call say_hello


say_hello:
	mov eax, 4																; Specify the sys_write system call
	mov ebx, 1																; Selects the file descriptor to write to in this case it's stdout
	mov ecx, Msg															; Msg to write
	mov edx, MsgLen														; The computer doesn't know how big the message is let it know
	int 0x80																	; Make the system call

	mov eax, 1																; Specify the sys_exit system call
	int 0x80																	; Make the system calll
