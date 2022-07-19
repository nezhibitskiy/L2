package main

//Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
//Реализовать поддержку утилитой следующих ключей:
//-A - "after" печатать +N строк после совпадения
//-B - "before" печатать +N строк до совпадения
//-C - "context" (after+before) печатать ±N строк вокруг совпадения
//-c - "count" (количество строк)
//-i - "ignore-case" (игнорировать регистр)
//-v - "invert" (вместо совпадения, исключать)
//-F - "fixed", точное совпадение со строкой, не паттерн
//-n - "line num", напечатать номер строки

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type flags struct {
	after, before, context             int
	ignore, ver, match, lineNum, count bool
}

func getFlags() *flags {
	var f flags
	flag.IntVar(&f.after, "After", 0, "'after' печатать +N строк после совпадения")
	flag.IntVar(&f.before, "Before", 0, "'before' печатать +N строк до совпадения")
	flag.IntVar(&f.context, "Context", 0, "'context' печатать ±N строк вокруг совпадения")
	flag.BoolVar(&f.count, "count", false, "'count' (количество строк)")
	flag.BoolVar(&f.lineNum, "num", false, "line num, напечатать номер строки")
	flag.BoolVar(&f.ignore, "ignore", false, "игнорировать регистр")
	flag.BoolVar(&f.ver, "ver", false, "вместо совпадения, исключать")
	flag.BoolVar(&f.match, "F", false, "точное совпадение со строкой, не паттерн")
	flag.Parse()
	return &f
}

func main() {
	if len(flag.Args()) == 0 {
		os.Exit(1)
	}
	f := getFlags()
	pattern := flag.Args()[0]
	filepath := flag.Args()[1]

	dir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/%s", dir, filepath))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	strs, matchedLinesNum := make([][]string, 0), make([]int, 0)
	var out string

	scanner := bufio.NewScanner(file)
	var lineCount int
	for scanner.Scan() {
		line := scanner.Text()
		if f.ignore {
			line = strings.ToLower(line)
		}
		if strings.Contains(line, pattern) {
			matchedLinesNum = append(matchedLinesNum, lineCount)
			if f.lineNum {
				out = fmt.Sprintf(out + fmt.Sprintf("Matched row number: %d\n", lineCount+1))
			}
			if f.ver {
				continue
			}
			strs = append(strs, strings.Split(line, " "))
			lineCount++
		} else {
			strs = append(strs, strings.Split(line, " "))
			lineCount++
		}
	}
	if len(matchedLinesNum) == 0 {
		return
	}
	grep(strs, matchedLinesNum, out, f)
}

func grep(strs [][]string, matchesIdx []int, out string, f *flags) {
	var grepABCm map[string][][]string
	if f.after != 0 || f.before != 0 || f.context != 0 {
		grepABCm = grepABC(strs, matchesIdx, f)
		for k, v := range grepABCm {
			fmt.Println(k)
			for _, raw := range v {
				fmt.Println(raw)
			}
		}
	}
	if f.count {
		out = fmt.Sprintf(out + fmt.Sprintf("Number of matched rows: %d", len(matchesIdx)))
	}
	if f.ver {
		for _, v := range strs {
			fmt.Println(v)
		}
	} else {
		for _, v := range matchesIdx {
			fmt.Println(strs[v])
		}
	}
	fmt.Println(out)
}

func grepABC(strs [][]string, matchesIdx []int, f *flags) map[string][][]string {
	res := make(map[string][][]string)
	if f.after != 0 || f.before != 0 {
		for _, v := range matchesIdx {
			if f.before != 0 {
				for i, j := v-1, 0; i >= 0; i, j = i-1, j+1 {
					if j > f.before-1 {
						break
					}
					res["GREP+BEFORE:"+strings.Join(strs[v], " ")] = append(res["GREP+BEFORE:"+strings.Join(strs[v], " ")], strs[i])
				}
			}
			if f.after != 0 {
				for i, j := v+1, 0; i <= len(strs)-1; i, j = i+1, j+1 {
					if j > f.after-1 {
						break
					}
					res["GREP+AFTER:"+strings.Join(strs[v], " ")] = append(res["GREP+AFTER:"+strings.Join(strs[v], " ")], strs[i])
				}
			}
		}
	}
	if f.context != 0 {
		f.after, f.before = f.context, f.context
		f.context = 0
		recMap := grepABC(strs, matchesIdx, f)
		return recMap
	}
	return res
}
