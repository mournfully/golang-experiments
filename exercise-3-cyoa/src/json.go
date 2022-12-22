package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadJsonFile(filePath string) (StoryFormat, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("ReadJsonFile: failed to open file: %w", err)
	}

	defer file.Close()

	content, err := decodeStream(file)
	if err != nil {
		return nil, fmt.Errorf("ReadJsonFile: failed to parse file: %w", err)
	}

	return content, nil
}

func decodeStream(reader io.Reader) (StoryFormat, error) {
	data := json.NewDecoder(reader)
	var content StoryFormat
	for {
		err := data.Decode(&content)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decodeStream: failed to decode json into struct: %w", err)
		}
	}

	return content, nil
}
