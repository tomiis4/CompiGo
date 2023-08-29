section .bss ; block-starting-symbol 
    ;variables

section .data
    ; constants
    var: db "Hello", 10
    varLen: equ $-var

section .text
    global _start ; entry point for linker

    _start ; start
        mov rax,1 ; sys_write (register 1)
        mov rdi,1 ; stdout (1 - standard output)
        mov rsi, var ; message to print
        mov rdx, varLen ; message len
        syscall ; call kernel

        ; end program
        mov rax,60 ; sys_exit
        mov rdi,0 ; err code 0
        syscall
