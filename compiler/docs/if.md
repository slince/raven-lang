
```js
let a = 10

if (a > 30) {
    console.log(30)
} else if(a > 20 ) {
    console.log(20)
} else {
    console.log(10)
}
```

```assembly
entry:
  local a 10
  jmp L1
  
L1:
  gt tmp1 a 30 
  jmp tmp1 if.body L2

L2:
  gt tmp2 a 20 
  jmp tmp1 if.body2 if.else

if.body:
   arg 30
   call log
   jmp leave
   
if.body2:
   arg 20
   call log
   jmp leave
   
if.else:
   arg 10
   call log
   jmp leave
```