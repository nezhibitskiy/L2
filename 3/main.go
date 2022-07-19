package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

// - k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
// - n — сортировать по числовому значению
// - r — сортировать в обратном порядке
// - u — не выводить повторяющиеся строки

type str struct {
	value string
	index int
}

type Flags struct {
	path                                 string
	columnToSort                         int
	numberSort, reverseSort, noStrRepeat bool
}

func getFlags() *Flags {
	var f Flags
	path, _ := os.Getwd()
	flag.BoolVar(&f.numberSort, "numberSort", false, "сортировать по числовому значению")
	flag.BoolVar(&f.reverseSort, "reverseSort", false, "сортировать в обратном порядке")
	flag.BoolVar(&f.noStrRepeat, "noStrRepeat", false, "не выводить повторяющиеся строки")
	flag.IntVar(&f.columnToSort, "columnToSort", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.StringVar(&path, "path", "", "путь к файлу")
	flag.Parse()
	return &f
}

func main() {

	f := getFlags()
	if len(f.path) == 0 {
		dir, _ := os.Getwd()
		f.path = fmt.Sprintf("%s/test.txt", dir)
	}
	if f.path == "" {
		log.Fatal(errors.New("не указано имя файла"))
	}

	file, err := os.Open(f.path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	strs := make([][]string, 0)

	if f.noStrRepeat {
		m := make(map[string]struct{})
		for scanner.Scan() {
			str := scanner.Text()
			if _, ok := m[str]; ok {
				continue
			}
			m[str] = struct{}{}
			strs = append(strs, strings.Split(str, " "))
		}
	} else {
		for scanner.Scan() {
			str := scanner.Text()
			if len(str)-1 < f.columnToSort {
				continue
			}
			strs = append(strs, strings.Split(str, " "))
		}
	}

	if f.numberSort {
		err := sortIndex(strs, f)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = sortAlphabet(strs, f)
	if err != nil {
		log.Fatal(err)
	}
}

func sortIndex(strs [][]string, f *Flags) error {
	indexes := make([]int, 0)
	m := make(map[int]string)
	for _, v := range strs {
		a, _ := strconv.Atoi(v[f.columnToSort])
		m[a] = strings.Join(v, " ")
	}
	for _, v := range strs {
		index, err := strconv.Atoi(v[f.columnToSort])
		if err != nil {
			return err
		}
		indexes = append(indexes, index)
	}

	sort.Ints(indexes)

	if f.reverseSort {
		for i := len(indexes) - 1; i >= 0; i-- {
			if str, ok := m[indexes[i]]; ok {
				fmt.Println(str)
			}
		}
		return nil
	}

	for i := 0; i < len(indexes); i++ {
		if str, ok := m[indexes[i]]; ok {
			fmt.Println(str)
		}
	}
	return nil
}

func sortAlphabet(strs [][]string, f *Flags) error {
	indexes := make([]string, 0)
	m := make(map[string]string)
	for _, v := range strs {
		indexes = append(indexes, v[f.columnToSort])
		m[v[f.columnToSort]] = strings.Join(v, " ")
	}
	sort.Strings(indexes)
	if f.reverseSort {
		for i := len(indexes) - 1; i >= 0; i-- {
			if j, ok := m[indexes[i]]; ok {
				fmt.Println(j)
			}
		}
		return nil
	}

	for _, v := range indexes {
		if j, ok := m[v]; ok {
			fmt.Println(j)
		}
	}
	return nil
}
