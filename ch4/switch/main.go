package main

import "fmt"

func main() {
	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthoropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "は短い単語")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "はまあまあの長さの単語", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "は長い単語")
		}

	}

	// その他
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "は偶数")
		case i%3 == 0:
			fmt.Println(i, "は3で割り切れる")
		case i%7 == 0:
			fmt.Println("ループを抜けます")
			break loop
		default:
			fmt.Println(i, "は退屈な数字")
		}
	}

	// Boolean Comparison
	words2 := []string{"hi", "salutations", "hello"}
	for _, word := range words2 {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "は短い単語")
		case wordLen > 10:
			fmt.Println(word, "は長い単語")
		default:
			fmt.Println(word, "はちょうどいい長さの単語")
		}
	}

}
