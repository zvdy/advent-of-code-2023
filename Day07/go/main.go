package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func strength(hand string, part2 bool) (int, string) {
	hand = strings.ReplaceAll(hand, "T", string(unicode.MaxRune))
	hand = strings.ReplaceAll(hand, "J", string(unicode.MaxRune-1))
	hand = strings.ReplaceAll(hand, "Q", string(unicode.MaxRune-2))
	hand = strings.ReplaceAll(hand, "K", string(unicode.MaxRune-3))
	hand = strings.ReplaceAll(hand, "A", string(unicode.MaxRune-4))

	counter := make(map[rune]int)
	for _, r := range hand {
		counter[r]++
	}

	if part2 {
		target := '1'
		for k := range counter {
			if k != '1' {
				if counter[k] > counter[target] || target == '1' {
					target = k
				}
			}
		}
		if counter['1'] > 0 && target != '1' {
			counter[target] += counter['1']
			delete(counter, '1')
		}
	}

	values := make([]int, 0, len(counter))
	for _, v := range counter {
		values = append(values, v)
	}
	sort.Ints(values)

	switch {
	case len(values) == 5 && values[0] == 5:
		return 10, hand
	case len(values) == 2 && values[0] == 1 && values[1] == 4:
		return 9, hand
	case len(values) == 2 && values[0] == 2 && values[1] == 3:
		return 8, hand
	case len(values) == 3 && values[0] == 1 && values[1] == 1 && values[2] == 3:
		return 7, hand
	case len(values) == 3 && values[0] == 1 && values[1] == 2 && values[2] == 2:
		return 6, hand
	case len(values) == 4 && values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 2:
		return 5, hand
	case len(values) == 5 && values[0] == 1 && values[1] == 1 && values[2] == 1 && values[3] == 1 && values[4] == 1:
		return 4, hand
	default:
		panic(fmt.Sprintf("invalid hand: %v %v %v", counter, hand, values))
	}
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	lines := strings.Split(string(data), "\n")

	for _, part2 := range []bool{false, true} {
		hands := make([]struct {
			hand string
			bid  int
		}, len(lines))
		for i, line := range lines {
			parts := strings.Split(line, " ")
			if len(parts) < 2 {
				fmt.Println("Invalid line:", line)
				continue
			}
			hands[i].hand = parts[0]
			fmt.Sscanf(parts[1], "%d", &hands[i].bid)
		}
		sort.Slice(hands, func(i, j int) bool {
			si, _ := strength(hands[i].hand, part2)
			sj, _ := strength(hands[j].hand, part2)
			return si < sj
		})
		ans := 0
		for i, h := range hands {
			ans += (i + 1) * h.bid
		}
		fmt.Println(ans)
	}
}
