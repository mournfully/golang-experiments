package decoder

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(filePath string) (StoryFormat, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: failed to open file: %w", err)
	}

	defer file.Close()

	content, err := decodeStream(file)
	if err != nil {
		return nil, fmt.Errorf("ReadFile: failed to parse file: %w", err)
	}

	return content, nil
}

func decodeStream(reader io.Reader) (StoryFormat, error) {
	data := json.NewDecoder(reader)
	var content StoryFormat
	if err := data.Decode(&content); err != nil {
		return nil, fmt.Errorf("decodeStream: failed to decode json into struct: %w", err)
	}

	return content, nil
}

type StoryFormat map[string]Chapter

// Generated by https://quicktype.io
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
