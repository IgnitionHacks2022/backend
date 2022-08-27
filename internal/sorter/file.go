package sorter

import (
	"errors"
	"io/ioutil"
	"strings"
)

func getBlueBin() (map[string]bool, error) {

	content, err := ioutil.ReadFile("pkg/data/bluebin.txt")
	if err != nil {
		return nil, errors.New("Error reading blue bin")
	}

	blueLines := strings.Split(string(content), "\n")
	blueMap := make(map[string]bool)

	for _, s := range blueLines {
		blueMap[s] = true
	}
	return blueMap, nil
}

func getRedBin() (map[string]bool, error) {

	content, err := ioutil.ReadFile("pkg/data/redbin.txt")
	if err != nil {
		return nil, errors.New("Error reading red bin")
	}

	redLines := strings.Split(string(content), "\n")
	redMap := make(map[string]bool)
	for _, s := range redLines {
		redMap[s] = true
	}
	return redMap, nil
}
