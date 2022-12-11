function main(){
   let a = 10

   if (a > 12) {
      a = 20
   } else {
      a = 30
   }

   switch (a) {
     case 10:
       a += 1
       break;
     case 20:
       a += 2
     case 30:
       a += 3
       break
     default:
       a += 4
   }
}