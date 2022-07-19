package main

// Написать функцию поиска всех множеств анаграмм по словарю.
//
//Например:
//'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
//'листок', 'слиток' и 'столик' - другому.

import "fmt"

func main() {
	words := []string{"тяпка", "пятак", "пятка", "листок", "слиток", "столик", "прикол", "роща", "орща", "хехе"}
	fmt.Println(Find(words))
}

func Find(words []string) map[string][]string {
	if len(words) < 2 {
		return nil
	}
	groups := make(map[string][]string)
	groups[words[0]] = append(groups[words[0]], words[0])
	for word := 1; word < len(words); word++ {
		var cur string
		for k, _ := range groups {
			var count int
			iter := make(map[rune]struct{})
			for _, symb := range []rune(k) {
				iter[symb] = struct{}{}
			}
			for i := 0; i < len([]rune(words[word])); i++ {
				if _, ok := iter[[]rune(words[word])[i]]; ok {
					count++
					continue
				}
				break
			}
			if count == len([]rune(k)) {
				cur = k
				groups[k] = append(groups[k], words[word])
				break
			}
		}
		_, ok := groups[cur]
		if !ok {
			groups[words[word]] = []string{words[word]}
		}
	}
	return groups
}
