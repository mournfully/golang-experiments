### Golang program to convert a decimal number to a binary number.

This table shows the decimal number `122` broken down into binary. 
![](https://content.instructables.com/FHI/OAFN/IKWXG1GA/FHIOAFNIKWXG1GA.png?auto=webp&frame=1&fit=bounds&md=b0e8280c5ef001c4b99248c6857165a4)

---

### I also wish to try out test-driven-development.

*I ended up just reused my readFile() function all the way back from "exercise 1" lol*
[Working with Files in Go | DevDungeon](https://www.devdungeon.com/content/working-files-go#seek)

*As it turns `Table-Driven Tests` were all I needed to replace from ^ and I just wasted my day... again...*
[How to Write Unit Test in Go. In this article, we are going to take a… | by Mert Kimyonşen | Delivery Hero Tech Hub | Medium](https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510)

---

```go
func decimalToBinary(num int){
   var binary []int

   for num !=0 {
      binary = append(binary, num%2)
      num = num / 2
   }
   if len(binary)==0{
      fmt.Printf("%d\n", 0)
   } else {
      for i:=len(binary)-1; i>=0; i--{
         fmt.Printf("%d", binary[i])
      }
      fmt.Println()
   }
}
```
[Write a Golang program to convert a decimal number to its binary form](https://www.tutorialspoint.com/write-a-golang-program-to-convert-a-decimal-number-to-its-binary-form#:~:text=Step%201%3A%20Define%20a%20function,binary%20array%20in%20reverse%20order.)

---

```go
func IntToString2() string {
    a := []int{1, 2, 3, 4, 5}
    b := make([]string, len(a))
    for i, v := range a {
        b[i] = strconv.Itoa(v)
    }

    return strings.Join(b, ",")
}
```
[go - One-liner to transform []int into string - Stack Overflow](https://stackoverflow.com/questions/37532255/one-liner-to-transform-int-into-string)
