# how do go_tests work?

### basic testing
```shell
$ go test
$ go test -v
# You can also run a specific test from many tests with -run flag.
$ go test -v -run=Test_SayGoodBye
```

### test coverage
```shell
$ go test -coverprofile=cover_out
$ go tool cover -html=cover_out -o cover_out.html
# commands above will create files 'cover_out' and 'cover_out.html'
$ $BROWSER cover_out.html 
```

### Table-Driven Test
*bruh that's exactly what i needed to know 5 hours ago ;-;*
[How to Write Unit Test in Go. In this article, we are going to take a… | by Mert Kimyonşen | Delivery Hero Tech Hub | Medium](https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510)
