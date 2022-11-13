# Exercise #1, the quiz from GopherExercises

### quiz-game-v1.go
*Put simply, I decided to complete this exercise by myself without looking at the solutions provided. And to instead use `https://gobyexample.com/` for the absolute basics, and to search around for every problem I found. Admittedly, the complexity did grow more than it had to. Unfortunenatly, I didn't know any better at the time.*

### quiz-game-v2.go
*Afterwards, I looked at the provided solutions and made my own updated version but with a solid reference this time around.*



- Part 1 (Basic)
	- [x] read data from problems.csv by default and add [flag -fp] to change input file
	- [x] manipulate data into seperate questions and associated answers
	- [x] ask user questions and track only if they were correct
		- if answer == correct --> counter++ ...then next question
		- if answer != correct --> ...next question
		- if answer == nil --> ...next question
	- [x] at the end output score of # of correct/# of questions

- Part 2 (Advanced)
	- [x] ask user to press enter to start a 30 second quiz by default and add [flag -t] to change time limit
	- [x] stop quiz immediately even if mid-question as soon as time limit is reached
		how about initializing timer() from main() and having it change a global flag after n seconds that would in turn flip an if-else after n seconds in loop(csv_out)
		although, i don't see how this could kick the user if they're in the middle of ask()
		OMG, GO ROUTINES AND CHANNELS!!! (and apparently an os.Exit() too lol)

- Part 3 (Bonus)
	- [x] sanatize user inputs (whitespace & caps) with 'strings' package
	- [x] add [flag -s] to shuffle questions around every run



gophercises/quiz: Ex 1 - Run timed quizzes via the command line
https://github.com/gophercises/quiz