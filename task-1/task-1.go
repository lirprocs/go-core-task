package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"        // Тип string
var isActive bool = true          // Тип bool
var complexNum complex64 = 1 + 2i // Тип complex64

func getType(a ...interface{}) {
	for _, val := range a {
		fmt.Printf("%T \n", val)
	}
}

func convToString(a ...interface{}) string {
	str := ""
	for _, val := range a {
		switch val.(type) {
		case int:
			str += fmt.Sprintf("%d", val)
		case float64:
			str += fmt.Sprintf("%.2f", val)
		case string:
			str += val.(string)
		case bool:
			str += fmt.Sprintf("%t", val)
		case complex64:
			str += fmt.Sprintf("%v", val)
		default:
			fmt.Println("Неверный тип данных")
		}
	}
	return str
}

func SHA256(str string) string {
	strRune := []rune(str)
	salt := []rune("go-2024")
	mid := len(strRune) / 2
	runesWithSalt := append(append(strRune[:mid], salt...), strRune[mid:]...)

	hasher := sha256.New()
	hasher.Write([]byte(string(runesWithSalt))) //TODO Не уыерен, что нужно было именно так, делал как в курсе
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	getType(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	str := convToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Println(str)
	hashStr := SHA256(str)
	fmt.Println(hashStr)
}
