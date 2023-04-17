package generator

import (
	"errors"
	"math/rand"
	"strings"
)

func GeneratePassword(length int, excludeUpper, excludeLower, excludeNumbers, excludeSpecial bool) (string, error) {
	if length < 4 {
		return "", errors.New("password length can't be this short")
	} else if excludeUpper && excludeLower && excludeNumbers && excludeSpecial {
		return "", errors.New("password should include at least one type of characters")
	}
	upperChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerChars := "abcdefghijklmnopqrstuvwxyz"
	numberChars := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{};:,.<>/?|"

	characterSets := make([]string, 0, 4)
	if !excludeUpper {
		characterSets = append(characterSets, upperChars)
	}
	if !excludeLower {
		characterSets = append(characterSets, lowerChars)
	}
	if !excludeNumbers {
		characterSets = append(characterSets, numberChars)
	}
	if !excludeSpecial {
		characterSets = append(characterSets, specialChars)
	}

	password := strings.Builder{}
	password.Grow(length)

	for _, set := range characterSets {
		randomChar := set[rand.Intn(len(set))]
		password.WriteString(string(randomChar))
	}

	for i := len(characterSets); i < length; i++ {
		set := characterSets[rand.Intn(len(characterSets))]
		randomChar := set[rand.Intn(len(set))]
		password.WriteString(string(randomChar))
	}

	shuffledPassword := []rune(password.String())
	rand.Shuffle(len(shuffledPassword), func(i, j int) {
		shuffledPassword[i], shuffledPassword[j] = shuffledPassword[j], shuffledPassword[i]
	})
	return string(shuffledPassword), nil
}
