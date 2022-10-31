
section .data
   hello string "helloword"; 字符串常量初始化
   world long 1234

.code
 a:  global main; 程序入口

main:
   load reg0 10 ;这是嵌入行注释
   load reg1 20
   add reg2 reg0  reg1
.loop:
   add reg2 reg0  reg1
.loop2s:
   add reg2 reg0  reg1
;这是独立行注释