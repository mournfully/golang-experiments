# Exercise #3, choose-your-own-adventure from GopherExercises

*golang project structure?*

**Requirements**
- [ ] Use the html/template package to create your HTML pages. Part of the purpose of this exercise is to get practice using this package.
Create an http.Handler to handle the web requests instead of a handler function.
Use the encoding/json package to decode the JSON file. You are welcome to try out third party packages afterwards, but I recommend starting here.
A few things worth noting:

Stories could be cyclical if a user chooses options that keep leading to the same place. This isn't likely to cause issues, but keep it in mind.
For simplicity, all stories will have a story arc named "intro" that is where the story starts. That is, every JSON file will have a key with the value intro and this is where your story should start.
Matt Holt's JSON-to-Go is a really handy tool when working with JSON in Go! Check it out - https://mholt.github.io/json-to-go/