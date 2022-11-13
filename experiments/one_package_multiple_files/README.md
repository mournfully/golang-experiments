I admit it's badly explained, but it's really quite simple:
​
> Is there a way to split this up these functions into separate files to make it easier to deal with?

Yes. You can split up code from the same package just by putting it in multiple files in the same directory. (With the same package declaration.) When you build, build the directory, not a file.


> I know how to make my own packages, but do I actually have to create every additional file as a package?

No. Each package needs it's own (exclusive) directory, but can have multiple files. If you don't want to deal with GOPATH, your main package can use relative imports to find these libraries.


> Is there not a way to just compile multiple files into one?

Yes. Don't name files to compile, name the directory.

​Using multiple .go files. : golang
https://www.reddit.com/r/golang/comments/au0xcj/using_multiple_go_files/

---

```shell
$ pwd
~/golang/experiments/multiple_go_files

$ ls *.go
file-one.go file-two.go

$ go mod init multiple_go_files

$ go run .
Hi from file two
Hi from file one
```

```go
// ~/golang/experiments/multiple_go_files/file-one.go
package main
import ( "fmt" )

func main() {
    funcFromFileTwo()
    fmt.Println("Hi from file one")
}
```

```go
// ~/golang/experiments/multiple_go_files/file-two.go
package main
import ( "fmt" )

func funcFromFileTwo() {
    fmt.Println("Hi from file two")
}
```
