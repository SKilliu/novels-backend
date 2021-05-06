package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	FacebookSocialKey = "facebook"
	GoogleSocialKey   = "google_play"
	AppleSocialKey    = "apple_id"
)

func GenerateName(source string) (string, error) {
	var randomRow string

	content, err := ioutil.ReadFile("./static/counter.txt")
	if err != nil {
		return randomRow, err
	}

	switch source {
	case FacebookSocialKey:
		randomRow = fmt.Sprintf("fbUser%s", string(content))
	case AppleSocialKey:
		randomRow = fmt.Sprintf("appleUser%s", string(content))
	case GoogleSocialKey:
		randomRow = fmt.Sprintf("googleUser%s", string(content))
	default:
		randomRow = fmt.Sprintf("guest%s", string(content))
	}

	contentNumber, err := strconv.Atoi(string(content))
	if err != nil {
		return randomRow, err
	}

	contentNumber++

	newContent := strconv.Itoa(contentNumber)

	destFile, err := os.OpenFile("./static/counter.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return randomRow, err
	}
	defer destFile.Close()

	err = destFile.Truncate(0)
	if err != nil {
		return randomRow, err
	}

	_, err = destFile.Write([]byte(newContent))
	if err != nil {
		return randomRow, err
	}

	return randomRow, err
}
