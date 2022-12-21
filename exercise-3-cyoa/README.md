# Exercise #3, choose-your-own-adventure from GopherExercises

**Suggestions**
- Stories could become cyclical if a user chooses options that keep leading to the same place.
- For simplicity, all stories will start from a story arc named `intro`.
- [Matt Holt's JSON-to-Go](https://mholt.github.io/json-to-go/) or [Paste JSON as Code](https://marketplace.visualstudio.com/items?itemName=quicktype.quicktype) are really useful when working with JSON in Go. 

**Tasks**
- Part 1 (Basic)
  - [ ] Use the `html/template` package to create your HTML pages.
  - [ ] Create an `http.Handler` to handle the web requests instead of a handler function.
  - [ ] Use the `encoding/json` package to decode the JSON file. 

- Part 2 (Advanced)
  - [ ] Support cli, where stories are printed out to the terminal and options are picked via typing in numbers `Press 1 to venture ...`.

- Part 3 (Bonus)
  - [ ] Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn't start on an arc named intro? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one.