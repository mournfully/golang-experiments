package main

import (
	"flag"
	"fmt"

	"example.com/cyoa/decoder"
	"github.com/kr/pretty"
)

const (
// starterChapterKey = "intro"
)

var (
	filePath = flag.String("file", "../../example/story.json", "path to the input json file which would be used as the story")
	// outputType = flag.String("output", "cli", "print output to `cli` or to `web`server")
)

func main() {
	flag.Parse()

	content, err := decoder.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("JsonDecoder: %s", err)
	}

	fmt.Printf("%# v\n", pretty.Formatter(content))

}

// func readStory(content decoder.StoryFormat) {
// 	chapterKey, _ := getChapter(content, starterChapterKey)
// 	for {
// 		printChapter()
// 	}
// }

// func printChapter() {

// }
