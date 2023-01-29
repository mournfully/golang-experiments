# Exercise #2, url-redirection from 

This time around, instead of banging my head against my keyboard when I get stuck. I decided to look at the solution briefley and then try and understand how they'd come to that conlclusion from the official documentation. I'm not sure if this is a better method...

I think I learnt something from spending all day today on this. Before I actually start coding, I should probably come up with a half-decent plan before I go off and do random crap that isn't even relevant .-.

**Tasks**
- Part 1 (Basic)
  - [x] url redirection, test with `go run main.go` and then on your browser go to `http://localhost:8080/` - was confused for a bit lol
  - [x] parse map pathsToUrls and redirect from entries there

- Part 2 (Advanced)
  - [x] create a cli flag to use yaml input file
  - [x] parse from yaml file
  - [x] convert data to a map
  - [x] reuse maphandler() to parse map
  - [x] parse from json file too

- Part 3 (Bonus)
  - [ ] dockerize environment?
  - [ ] high availability redis database?
  - [ ] read from database instead of map
  - [ ] improve your program so that it's comparable to [ptman/urlredir: Educational URL redirector service in Go](https://github.com/ptman/urlredir)?

**References**
[ptman/urlredir: Educational URL redirector service in Go](https://github.com/ptman/urlredir)
[urlshort/main.go at master · gophercises/urlshort](https://github.com/gophercises/urlshort/blob/master/students/dennisvis/main.go)
[urlshort/handler.go at master · gophercises/urlshort](https://github.com/gophercises/urlshort/blob/master/students/dennisvis/urlshort/handler.go)
