package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestConvToString(t *testing.T) {
	expected := "4242423.14Golangtrue(1+2i)"
	result := convToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)

	if result != expected {
		t.Errorf("convToString() = %v; expected %v", result, expected)
	}
}

func TestSHA256(t *testing.T) {
	str := "4242423.14Golangtrue(1+2i)"
	salt := []rune("go-2024")
	expectedHash := func() string {
		strRune := []rune(str)
		mid := len(strRune) / 2
		runesWithSalt := append(append(strRune[:mid], salt...), strRune[mid:]...)

		hasher := sha256.New()
		hasher.Write([]byte(string(runesWithSalt)))
		return hex.EncodeToString(hasher.Sum(nil))
	}()

	resultHash := SHA256(str)

	if resultHash != expectedHash {
		t.Errorf("SHA256() = %v; expected %v", resultHash, expectedHash)
	}
}
