package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFile  string
	OutputFile string
}

func New(input, output string) FileManager {
	return FileManager{
		InputFile:  input,
		OutputFile: output,
	}
}

func (fm FileManager) ReadFile() ([]string, error) {
	file, err := os.Open(fm.InputFile)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("failed to scanning file")
	}

	return lines, nil
}

func (fm FileManager) WriteJson(data interface{}) error {
	file, err := os.Create(fm.OutputFile)
	if err != nil {
		return errors.New("failed to created file")
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return errors.New("failed to convert to JSON")
	}
	if _, err := file.Write(jsonData); err != nil {
		return errors.New("failed to write file")
	}

	return nil
}
