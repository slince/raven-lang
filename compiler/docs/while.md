

```js
let a = 1

while (a < 10) {
    a ++
    // if (a == 5) {
    //     break
    // }
}
```

```
entry:
  local a 1
  jmp l1
  
l1:
  le t1 a 10
  cjmp t1 l2 leave
  
l2:
  add a a 1
  jmp l1
```