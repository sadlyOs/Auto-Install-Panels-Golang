package main

import (
	"fmt"
	"os"
)

func readBash(file string) (string, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка с чтением", err)
		return "", fmt.Errorf("ошибка при чтении файла: %v", err)
	}

	read := string(content)

	return read, nil
}
