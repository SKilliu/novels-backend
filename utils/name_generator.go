package utils

import (
	"fmt"
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

func GenerateName() string {
	bytes := make([]rune, 8)

	for i := range bytes {
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	randomRow := fmt.Sprintf("guest-%s", string(bytes))

	return randomRow
}
