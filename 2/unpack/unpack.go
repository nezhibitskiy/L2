package unpack

import (
	"errors"
	"strconv"
	"unicode"
)

func Unpack(array string) (string, error) {
	if array == "" {
		return array, nil
	} else if unicode.IsDigit(rune(array[0])) {
		return "", errors.New("некорректная строка")
	}

	res := make([]rune, 0)
	for i, v := range []rune(array) {
		if v == '\\' {
			continue
		}
		j, err := strconv.Atoi(string(v))
		if err == nil {
			if array[i-1] != '\\' {
				for k := 1; k < j; k++ {
					res = append(res, rune(array[i-1]))
				}
				continue
			}
			res = append(res, v)
			continue
		}
		res = append(res, v)
	}
	return string(res), nil
}
