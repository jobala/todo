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

  WriteSuccess: db "Written to file", 0xa																
  WriteSuccessLen equ $-WriteSuccess

	DummyTodo: db "A todo item",10																		; Placeholder todo item
	DummyTodoLen equ $-DummyTodo

  FileName: db "database.txt"																				; File name for the file storing todo items

section .bss
	fd_out resb 1
	buffer resb 2048
	len equ 2048
	
section .text
	global _start

_start:
	call welcome_message
	call open_or_create_file
	call write_to_file
	call read_from_file
	call exit



open_or_create_file:
	mov eax, 8																; Specifies sys_creat syscall
	mov ebx, FileName													; Provide the filename
	mov ecx, 0777							                ; Specify the permissions with which the file should be created/opened

	int 0x80																	; make the syscall

  mov [fd_out], eax                         ; Store the file descriptor for in memory

	ret

write_to_file:
	mov eax, 4																; Specify the sys_write syscall
	mov ebx, [fd_out]													; Reference file descriptor from memory
	mov ecx, DummyTodo
	mov edx, DummyTodoLen

	int 0x80

  mov eax, 6                                ; Specify sys_close syscall
  mov ebx, [fd_out]                         ; It should close the file described by the file descriptor

  int 0x80
	ret

read_from_file:
	mov eax, 5																; Specify open file name
	mov ebx, FileName
	mov ecx, 0																; Set access permissions to read only
	int 0x80														

	mov [fd_out], eax													; Set the file descriptor for the opened file

	mov eax, 3																; Read from file
	mov ebx, [fd_out]													; Read from the file that was just opened
	mov ecx, buffer														; Specify buffer size
	mov edx, len															; Specify size of the content read
	int 0x80

	mov eax, 4																; Specify the write syscall
	mov ebx, 1                                ; Terminal file descriptor
	mov ecx, buffer                           ; Write the message contained to the buffer
	mov edx, len															; Set the length of bytes to be written 
	int 0x80

	mov eax, 6 																; Specify the sysclose syscall
	mov ebx, [fd_out]													; The file to close is described by the file descriptor
	int 0x80
	ret

welcome_message:
	mov eax, 4																; Specify the sys_write system call
	mov ebx, 1																; Selects the file descriptor to write to in this case it's stdout
	mov ecx, Msg															; Msg to write
	mov edx, MsgLen														; The computer doesn't know how big the message is let it know
	int 0x80																	; Make the system call
	ret

exit:
	mov eax, 1																; Specify exit system call
	int 0x80
