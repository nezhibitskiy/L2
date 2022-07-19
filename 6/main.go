package main

//Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
//
//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type flags struct {
	delim     string
	fields    int
	separated bool
}

func getFlags() *flags {
	var f flags
	flag.IntVar(&f.fields, "f", 0, "выбрать поля (колонки)")
	flag.StringVar(&f.delim, "d", "#", "делиметр")
	flag.BoolVar(&f.separated, "s", false, "только строки с разделителем")
	flag.Parse()
	return &f
}

func main() {
	f := getFlags()
	f.fields = 1
	f.delim = "3"
	if f.fields <= 0 {
		log.Fatal("f must be >0")
	}
	p, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/dev06/test.txt", p))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		text := sc.Text()
		checker := strings.Split(text, " ")
		if len(checker) < f.fields {
			continue
		}
		fmt.Println(cut(text, f))
	}
}

func cut(str string, f *flags) string {
	if f.separated && !strings.Contains(str, f.delim) {
		return ""
	}
	splitted := strings.Split(str, f.delim)
	if len(splitted) < f.fields {
		return str
	}
	return splitted[f.fields-1]
}
