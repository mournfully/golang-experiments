# Exercise #3, choose-your-own-adventure from GopherExercises

**Suggestions**
- Stories could become cyclical if a user chooses options that keep leading to the same place.
- For simplicity, all stories will start from a story arc named `intro`.
- [Matt Holt's JSON-to-Go](https://mholt.github.io/json-to-go/) or [Paste JSON as Code](https://marketplace.visualstudio.com/items?itemName=quicktype.quicktype) are really useful when working with JSON in Go. 

**Tasks**
- Part 1 (Basic)
  - [ ] Use the `html/template` package to create your HTML pages.
  - [ ] Create an `http.Handler` to handle the web requests instead of a handler function.
  - [x] Use the `encoding/json` package to decode the JSON file. 
    - *even after looking through these, i cheated a bit and looked at the solution - but even now i don't get it .-.*
    - [Parse JSON objects with arbitrary key names in Go using interfaces and type assertions](https://gist.github.com/mjohnsullivan/24647cae50928a34b5cc)
    - [Parsing JSON with Go. How to unmarshal or decode JSON data… | by Sau Sheong | Go Recipes](https://go-recipes.dev/parsing-json-with-go-7268937a5f7b)
    - [Parsing JSON files With Golang | TutorialEdge.net](https://tutorialedge.net/golang/parsing-json-with-golang/)

- Part 2 (Advanced)
  - [ ] Support `cli`, where stories are printed out to the terminal and options are picked via typing in numbers `Press 1 to venture ...`.

- Part 3 (Bonus)
  - [ ] Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn't start on an arc named `intro`? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one.