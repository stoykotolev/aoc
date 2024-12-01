package utils

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var contents []string

	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	return contents

}
