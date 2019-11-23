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
	Msg: db "Welcome to the assembly todo manager", 10, "------------------------------------",10								; Create the message variable
	MsgLen equ $- Msg																									; Calculate the length of the message
	FileName: db "database.txt"																				; File name for the file storing todo items
	AccessPermissions equ 0777																				; Read, write & execute by all

section .bss


section .text
	global _start

_start:
	call welcome_message
	call open_or_create_file


open_or_create_file:
	pushad																		; Push the callers general purpose registers into the stack
	mov eax, 8																; Specifies sys_creat syscall
	mov ebx, FileName													; Provide the filename
	mov ecx, AccessPermissions							  ; Specify the permissions with which the file should be created/opened
	popad																			; Pop the callers general purpose registers from the stack

	int 0x80																	; make the syscall
	ret

read_file:


welcome_message:
	mov eax, 4																; Specify the sys_write system call
	mov ebx, 1																; Selects the file descriptor to write to in this case it's stdout
	mov ecx, Msg															; Msg to write
	mov edx, MsgLen														; The computer doesn't know how big the message is let it know
	int 0x80																	; Make the system call

	mov eax, 1																; Specify the sys_exit system call
	int 0x80																	; Make the system calll
