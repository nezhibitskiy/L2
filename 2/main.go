package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"2/unpack"
)

// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""

// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// qwe\4\5 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\5 => qwe\\\\\ (*)

// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

func readStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return text, nil
}

func main() {
	text, err := readStdin()
	if err != nil {
		log.Fatal(err)
	}

	un, err := unpack.Unpack(text)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(un)
}
