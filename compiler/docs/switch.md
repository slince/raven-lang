

```js


let a = 10

switch (a) {
   case 1:
     console.log(1)
     break
   case 2:
     console.log(2)
     break
   default:
    console.log(23)

}

a = 20

```


```code

entry:
  local a 10
  jmp L1

L1:
  equals tmp1 a 1 
  jmp tmp1 case1 l2

L2:
  equals tmp2 a 2 
  jmp tmp2  default

case1:
   arg 1
   call log
   jmp leave

case2:
   arg 2
   call log
   jmp leave

default:
   arg 3
   call log
   jmp leave

leave:
   assign a 20


```