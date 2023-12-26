package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1(actual))
	fmt.Println(part2(actual))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	var score int
	for _, line := range lines {
		numWinners := parseCard(line).countWinners()
		if numWinners > 1 {
			score += int(math.Pow(2, float64(numWinners-1)))
		} else {
			score += numWinners
		}
	}
	return score
}

type card struct {
	winners []int
	board   []int
}

func (c card) countWinners() int {
	var numWinners int
	for _, v := range c.winners {
		if slices.Contains(c.board, v) {
			numWinners++
		}
	}
	return numWinners
}

func part2(input string) int {
	cards := []card{}
	copies := map[int]int{}
	for i, line := range strings.Split(input, "\n") {
		cards = append(cards, parseCard(line))
		copies[i] = 1
	}
	for idx, c := range cards {
		for i := 1; i <= c.countWinners(); i++ {
			copies[idx+i] += copies[idx]
		}
	}
	var sum int
	for _, i := range copies {
		sum += i
	}
	return sum
}

func SS2IS(in []string) []int {
	out := []int{}
	for _, s := range in {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic("unclean")
		}
		out = append(out, i)
	}
	return out
}

func parseCard(c string) card {
	numbers := strings.Split(strings.Split(c, ":")[1], "|")
	return card{
		winners: SS2IS(strings.Fields(numbers[0])),
		board:   SS2IS(strings.Fields(numbers[1])),
	}
}

const example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

const actual = `Card   1:  2 15 17 11 64 59 45 41 61 19 |  4 36 62 43 94 41 24 25 13 83 97 86 61 90 67  7 15 58 18 19 38 17 49 52 37
Card   2: 41 62 67 93 88 12 78 51 95 49 | 55 63 89 78 45 11 62 50 81  9 32 82 15 36 74 54  4 58  5 56 44 83 90 49 34
Card   3: 51 22 38 33 85 23 56 76 60 93 | 94 40 61 37 38 82 93 96 13 50 81 65 56 26  4 18 86 30  8 16 60 27 23 48 51
Card   4: 51  6 90 10 97 65 19 17 24  3 | 93 82 10 13 17  3 90 74 14  7 77 38 70 97 72 60  6 79 65 94 24 19 51 45 28
Card   5: 76 61 43 95 41  2 40 58 30 96 | 72 23 69  5 30  3 10 17 78 20 13 86 60 81 90 46 96 39  6 32 31 73 65 95 62
Card   6: 32 70 49 99 27 26 75 96 91  4 | 15 74 28  2 17 82 55 96 88 11 95 77 12 38 91  3 56 44  7 32 19 94 85 27 97
Card   7: 38 21 90 66 46 32 55 94 72 75 |  3 10 90 77 41 99 46 82 65 72  9 21 66 94 80 22 97 56 93 61 18 81 34 76 64
Card   8: 32 36 87 79  4 17  1 77 16 63 |  9 63 23 25 91 82 35 74 28 72  6 52 56 12 39 99 57 46 15 76 85  2 75 77 96
Card   9: 42  5 58 80 96 88 37 30 49 69 | 26  6 17 71 46 51 45 10 62 33 38 23 42 74 50 61 95 11 93 57 14 22 72 85 86
Card  10: 10 74 58 71 57 35 34 96 77 18 | 14 27 22 18 70 42 56 94 76 74 85 73 61 34 88 45 39 64 35 87 90 58 91 75 54
Card  11: 24 26 43 62 66 93 38 51 95 86 | 31 75 23 91 12  3 90  9 36 87 76 58 40 35 56 65 17 74 61 93 20 71 82 92 73
Card  12:  2 25 63 37 64 81 20 15 28 88 | 67 26 12  6 58 37 69 93 83 52 81 25 19 30 80 54 73 21 86 20  8 57 47 32 60
Card  13: 48 37 14 80 51 17 85 26 40 33 | 80  7 79 38  8 74 10 78 89 16 81 84 19 49 86 83 63 82 40 58 25 47 53 34  4
Card  14: 93 27  7 80 63 89 50 96  3 15 |  6 19 40 46 51 29  9 86 39 77  4 54 38 41 67 69  1 56  5 35 30 17 71  2 43
Card  15: 33 72 28 75 56 22 16 38 44 51 | 81 67 79 73 52 40 69 43 12 77 31 41 23 14 64 58  5 88 95 13 99  6 42 83 38
Card  16: 36 82 59 39 57 26 61 45 60 65 | 66 51 24  1  4 78 17 16 74 95 34 47 75  6 67 44 48 14 10 89 46  2 31 18 68
Card  17: 40 50 36  8 54 17 29 13 89 98 |  6 59 60 57 18 30 80 32 46 77 86 35 48 84 91  4 38 99 17 33  3 15 82 83 62
Card  18: 81 11 66 67 19 65 39 64 20  9 | 25 14 60  2 50 97 43 18 36 24 89 74 94 71 51 45  5 34 52 31 90  6 48 40 55
Card  19: 41 81 99 97 18 92 79 82 83 62 | 97 34 41 76 87 60 21 63 93 55 77 92 42 56 80 94 81 36 11  1 83 70 52  5 29
Card  20: 50 94 73 61 56 88 35 48 54 21 | 89 35 83 42 45 12 61 18 68 52 13 25 88 54  6 74 21 57 50 56 34 73 48 94 46
Card  21: 95 69 16 96 19 86 94 74 33  1 | 19 95 86 16 94 76 17 39  5  1 97 27 44 69 74 96 29 48 14 92 33  7 12 91 42
Card  22: 17 78 32 62 96 43 68 23 97 52 | 97 28 87 26 23 14 78 39 92 56 31  3 32 52 13 43 80 62 17 69 68  5 96 25 75
Card  23: 57 93  4  6  2 34 18 80 99  9 |  9 53 58 19 35  6 46 87 86 36 59 17 26 54 39 52 99 20 69 18 25 30 34 41 42
Card  24: 72 43  5 36 81 31 77 76 58 48 | 75 48 31 88 33 38 37 69 40 25  4 93 35 23 76 36 64 44  6 65 59 15 43 17 28
Card  25: 14 52 98 94 85 95 27 54 67 34 | 54 14 19  3 24 36  6 52 16 95 76 61 23  8 98 34 28 99 67 85 27 55 70 59 94
Card  26: 33 15 53 51 14 47 12 83  6 48 | 25 15 51 33 12 53 60 49 98 47 30 89 78 83 88 85  6  7 38 92 48 93 74 14  9
Card  27: 72 33 73 12 18 13 15 88 82 39 | 23 33 72 13 18  1 12 15 64  8 59 73 48 31 40 57 76 39 52 88 89 36 82 41 67
Card  28: 42 32 84 56 53 69  5 15  8 63 | 88 32 91 44 62 37 69 72 98 29  8 82 21  5 56 15  6 42 58 61 89 35 96 84 59
Card  29: 51 41 27 43 38 63 67 26  8 50 | 52  8  1 32 26 59  4 67 27 98 75 11 54 14 50 16 70 66  2 86 62 37 65 38 55
Card  30: 79 88 87 54 17 18 16 58 57  1 | 54 17 68 89 10 29 97 16 72 57  2 38 12 65 98 11  1 33 64 40 79 63 70 88 18
Card  31: 61 75 59 95 28 91 57 54 55 82 | 19 57 30 95 54 81 64 28 94 92  3 80 10 63  6 83 88 69 59 15 27 93 96 40 38
Card  32: 67 44 87 82 49  2 85 79 35 77 | 65 21 63 44  4 52 28 89 11 87 37 97 76 71 64 96 69 79 93 13 33  5  7 58 35
Card  33: 34 69 25 16 38  4 48 28 39 80 | 90 71 48 27 46 54 19 16 69 12  6 59 25 17 10 81 47 30 28 13 66 38 58  4 80
Card  34: 65 48 49 80 28  9 22 57 47 95 | 28 59 79  2  4 89 64 98 14 90 18 22 92 52  3 37 95 87 74 54 43 26 51 42 88
Card  35: 60  8 75 58 47 63 90 71 93 50 | 60  9 24 63 98 66 39 91 46 34 86 67 16 95 53 73 48 87 75 96 44 90 68 26 93
Card  36: 31 61 33 38 70 46  2 81 73 58 | 49 67 79  3 20 64 96 62 81 82 54 57 70 30 43 56 71 86  6 19 48 93 58 15 27
Card  37: 86 62 46 70 33 84 90 66 89 96 | 56 39 63 53 45 80 49 61 52 64 12 69 74 26 44 86 96 62 27 31 78  2 28 21 16
Card  38: 95 40 44 15 99 76 85 20 38 11 | 57 22 23 29 69 90 48 51 66 24 11 32 33 25 92 94 70 88 78 39 28 60 77 55 31
Card  39: 67 40 10 87 93 47 23 98 96 91 | 14 30 78 72 70 99 52 65 84 12 18 75 59 85 83 89 55 11 82 34 63 10 97 54 69
Card  40: 95 51 94  2 36 99 98 46 87 82 |  1 40 18 73 50 82 66 14 65 85 83 45 89 58 79 55 84 80 90 72 22  5 38 44 69
Card  41: 68 75 93 29 66 95 27 46 23 82 |  8 60 39 33 41 36 25 73 94 89 50 71 85 19 51 30 53 45  9 31 83 99  5 32 52
Card  42: 73 48 37 23 32 19 72 52 58 91 | 35 19 55  9 23 79 62 87 72 58 95 68 20 38  5 48 60 37 73 46 91 12 74 52 32
Card  43: 42 39 19 59 11 50 88 31 10  9 | 77 11 39 56 42 67 22 12 88 57 75 79 80  6 92 86 99 33 59 90 89 93 68 70 61
Card  44: 86 12  3 41 71 82  6 48 97 93 | 49 75 31 45 41 83 10 86  8 91  9 70 97 69 52 77 51 36 34 14 81 28  2 40 64
Card  45: 51  9 72 44 18 14 40 97  2 25 | 83 80 62 22  7 34 88 33 17 23 38 81 70 89 26 35 57 78 46 76  4 61 56 98 55
Card  46: 22 59 61 94 67 87 21 18 12 69 | 50 22 60 21 67 26 59 94 28 33 30 12 16 69  4 27 95 18 61 53 56 47 85 88 87
Card  47: 91 58 84 29  4 60 77 41 83 97 | 70 35 40 17 62 41  9 59 26 76 78 16 25 92 91 95 71 77 63 21  6 67 54 46 49
Card  48: 90 56 13 75 73  2  8 62 21 88 | 38 88 76 75  8  2 82 77 56 73 28 18 59 21 65 90 23 94  5 13 15 79 26  1 36
Card  49: 84 14 99 51 58 28 63 36 48 91 | 94 46  6 61  2  9 40 38 64 63 91 16 49 81 70 48 80 36 51 58 84 98 96 34 60
Card  50: 79 98 99 11  7 34 38 52 74 27 | 21 57 11 19 89 94 72 52 81 62 92 79 32 77 28 74 84 86 80 69 83 38 98  3 34
Card  51: 58 60 15 93 24 72 94 29 59 12 | 71 72  7 86 23 39 65 87 76 22  6 74 69 94 46 52 91 60 96 24 93 47  4 15 59
Card  52:  6 92 21  7  9 55 17 33 23 16 | 28 69 86 94 34 21 95 51 22 41 99 87 46  1 49 36 90 76 35 30 64  5 44 93 58
Card  53: 57 10 81  4 93 79 51 87 28 50 | 38 34 97 15 49 89 59 37 25 48 70 41 98 33 78 87 95 85 35 71 32 30 96 67 36
Card  54: 25 41 26 14 99 98 71 83 31 54 | 68 17 28 80 79 10 31 99 54 83 27 14 25 32 26 85 42 76 33 49  8 24 67 71 41
Card  55: 38 26 25 70 78 37 13 76 31 73 | 82 10  2 51 13  8 98 12 19 77 84 89 93 22 54 35 23 31 69 42 43 25 39 57 94
Card  56: 31 25 88 98 23  3  2 58 99 67 | 46 27 75 97 84 53  9 47 64 49 23 92 15 74 44 78 58 56 83 30 31  7  5  2 94
Card  57: 50 61 25 29 59  1 47 72 65 84 | 72 74 97 28 30 75 56 61 44 40 62 95  4  8 64 22 34 25 14 16 50 92 65 39 13
Card  58: 73 65 94 47 90 35 95 12 81 78 | 86 87 24 11  1 81 96 33 12 18 74 41  9 22  5 44 48 82 68 67 88 94 20 59 90
Card  59: 41 27  3 55 16 83 49 75 51 88 | 33 66 74 58 60 72 75 67 78 55 68 95 81  3 39 86 69 94 92  4  5 44 57 26 56
Card  60: 33 93 56 90 16 46 65 78 13  6 | 55 41 70  6 63 69 90 82 22 17 66  2 53 79 49 25 10 34 47 54 43 58 59 81 68
Card  61: 72 45 87 76 64 80 96 22 67 98 | 69 25 90 48 67 88 30 34 38 98 19 73 86  8 99 35 44 95 91 36 63 37  2 13 66
Card  62: 96 68 19 52 97 83 60 67 20 35 | 21 96 30 58 10  7 40 25 92 33 15 69 61 34  6 81 87 75 41 16 51 27  8  4 86
Card  63: 17  1 45 26 82 78 85 15 89  8 | 49 38 16 22 71 50 81 74 87 61 14 19 64 93 96 23 59 56 76 86 53  6  7 80 41
Card  64:  3 17 22 94 37 39 67 73 27  4 |  8 27 22  5 28 53 47 52 94 25 29 54  1 59 98 16 78  9 65 91 61 74 73 60 37
Card  65: 80 48 70 23 99 11 20 31 54 94 | 73 69 21 10 66 90 80 29 55 67 92 20 96 94 79 93 31 30 17 14 27 75 13 95 71
Card  66: 14 79 31 27 89 95 69 63 12 67 | 36 85 67 24 87 72 89 26 17 31 27 55 10 32 14 63 71 79 13 12 81 69 28 95 30
Card  67: 87 44 51 53 37 69 39 96 90 16 | 22 63 83 31 32 76 66 68 29 30 53 48  4 72 82 46 95 80 17 21 56 51 89 12 77
Card  68: 39 27 17 15 50 65 38 94 84 42 | 42 84 17 50 52 47 67 93 13 35  7 26 23 95 39 81 94 15 78 27 43 21 38 55 86
Card  69: 13 22  8 54 61 77 75 31  1 67 | 68 54 15 23 40 58 67  3 71 78 90 33 44 22 57 19  8 30 74 38 48 41  9 25 65
Card  70: 10 43 33 21 31  6 94 46 82 83 | 48 62  6 19 67 72 46 81 75  5 54  8 56 92 37 76 96 71 32 36 26 14 30 79  9
Card  71: 91  9 15 33 37 22 61 74 14 70 | 79 31 18 42 43 40 26 56 95 45 86 76 65 23  3 94 24 35  4  2 68 51 25 12 80
Card  72: 30 79 24 55 42  8 13 90 68 73 | 37 61 34 94 51 30  3 10 26 19 50 42 40 31 80 86 66 27 83 90 73 79 99 14 76
Card  73: 25 88 63 85 56 49 30 46 10 99 | 46 68 67 60 11 78  6 85 12 97 91 21 27 81 39 48 69 44 25 63 22 10  3 61 79
Card  74: 21  3 32 71 98 69 44 78 34 11 | 83 56 29 18 34  9 63 92 53 22 61 45 41 38 82 27 49  5 32 99 69 17 55 66 96
Card  75: 68 92 32 65 49 38  3 56 17 44 | 47 68  6  1 10 24 66 84 60  2 48 56 57 61 51 35 28 17 82 52 98 81 23 75 38
Card  76: 30 57  5  8 10 95 45 85 54  3 | 53  2 54 34 23 12 48 16 37 92 85 27 83  8 21 44 81 26 49 67 36 25 52 47 68
Card  77: 31 84 72 26 83 55 20 42 65 61 | 53 96 17 18 59 21 23 77 78 15 43 58 66  3 60 55 38 45  6 72 44 94 11 79 19
Card  78: 62 77 52 72 13 24 31 60 29 26 | 74 51 46 67 97 14 71 70 55 94  6 16 65 85 66 80 60 27 32 39 73 61 34 91 69
Card  79: 69 74 82 58 44 37 86 51 75 20 | 93 70  2 35 30 79 45 56 65 59 22 64 89 36 66  5 34 13 49 43 97 74  7 26 83
Card  80: 84 16 89 49 81  1 79 44 93 25 |  2 35 29  8 87 21 22 57  9 73 67 82 54 45 71 72 59  5  4 97 70 60 91  7 92
Card  81: 25 44 10 12 23 79 60 45 96 90 | 62 23 25 40 12 90 42 10 79 45 69 46  9 27 73 59 34 44 29 94 95 86 66  1 58
Card  82: 42 43 37 18 40 64 75 76 99  3 | 67 42 11 18 91 99 14 90  3 23 73 82 76 40 37 44 28 64 96 16 74 75 55 61 15
Card  83: 38 89 77 47 66 22 87 59  7 43 | 59 46  7 73 79 23 96 22 28 87 66 32 77 89 98 19 38 18 43 56 15 47 74 95 99
Card  84: 86 13 97 81 46 12 88 34 92 52 | 88 69  8  2 76 63 95 27 46 31 81 48 75 92 60 33 36 22 34 13 57 30 21 42 98
Card  85: 57 71 21 43 41 23 18 15 59 93 | 40  8 21 25 19 57 52 71 75 45 16 18 89 99 43 20 27  4 53 98 80 22  7 29 41
Card  86: 84 43 50 94 92 31 48  8  6 91 | 92 53  6 94 25 41 74 54 62 43 95 79 16 31  8 89 84 30 50 91 58 47 17 96 46
Card  87: 45 96 17 51 25 40 39 13 78 82 | 66 61 23 13 59 10  1 86 30 62 74 41 34 45  8 65 36 16 78 93 20 92 35 82 90
Card  88: 29 89 14 85 93 19 56 36 99 15 | 91 14 94 89 96 19 85 71 23 46 21 75 92  7  6 33 60  8 72 81 99 36 29 56 22
Card  89: 77 79 82 96 75 63 72 59 37 11 | 43 97 11 61 80 68 22 19 71 63 16 18 56 49  9 79 45 36 37 89  1  8 48  4 40
Card  90:  2 81 33 10 60 57 38 99 95 65 |  3 98 33 80 32 65 45 96 99 61 25 73 74 37 91  6  2 79 38 34 97 89 36 11  5
Card  91: 31 62 73 12 53 51 64 55  9 20 | 45 76 95 42 81 56 32 62 11 48 93  9 69 79 84 28 46 82 88 96 31 91 13 98 92
Card  92: 64  5 81  2 59 24 40 74 84 58 | 93 43 67 70 53 20 91 18 49  4 14 52 74 19 79 85 61 86 82 57 38 44 34  9 83
Card  93:  6 94 85 43 15 67 68 79 81 65 | 15 35 14 78 21 68 86  9 97 99 75 87 23 46 70 30 41 27 18 84  3 50 53 63 79
Card  94: 61 34 62 94 50 23 69 98 78 60 | 43  9 90 28 49 36 19 97 56 75 62 15 79 70 40 78 24 94 48 88 45 91 66  4 34
Card  95: 73 26 99 83 45 46  9 50 38 14 | 98 81 63 27 55 28 32 82 92 41 87 60 49 50 99 95 11 89  4 78 17 47 20 15 96
Card  96: 52 80  3 67 69 57 16 34 40 77 | 70 35 97  2 71 25  1 15 54 26 19  5 83 20  4 99 56 38 51 87 53 47 30 13 96
Card  97: 54 41 65 23 52 82 81 22 28 47 | 27 90  3 24 15 85 36 88 97 37  4 18 42 50 99 35 83 78 79 46 40 63 92 34 67
Card  98: 47 56 85 60 12  1 35 88 30 86 | 75 83 15 97 66 74 55 62  9 58 19 82 51 54 61 68 47  7 17 10 39 44 20 96 34
Card  99: 52 25 39 18 92 17 93 29  6  9 | 22 83 86 51  5 43 20 73 13 66  6 82 24  8  2 11 87 79 57 50 19 35 45 62 49
Card 100: 27 91 93 24 17 47 80 13 51  8 | 59 78 63 74 81 26 98 64 97 79  6 77 54 83  9 86 69 31 12 10  5 56 34 33 60
Card 101:  8 45 53 15 57 36 69 47 13  1 | 15 13 79 47 69 36 99 57 25 91  1  8 39 73 90 24 84 93 60 40 53 55  3 45 87
Card 102: 45 76 11 54  1 15 24 60 20 29 | 25 32 88  5 66 40 16 58 24 70 79 36 29 53 54 86 60 15 78 57 27 92 99 33 30
Card 103: 46 15 97 62 94 19 99 76 42 93 | 65 19  5 58  3 10 12 24 50 64 54 22 85 31 44 93 30 47 14 21 97 78  2 15 81
Card 104: 61 83 16 66 81 55 52 36 76 95 | 28 51 36 20 11  6 27 33 92 57 94 96 45  9 15  2 63 72 58 44  5 81 43  1 66
Card 105: 44 12 34 11 26 72 57 55 38 69 | 41 57 62 26 14 72 89 99 34 69 28 10 68 70  2 50 98 37 12 38 44 71 11 80 55
Card 106: 43 33 65 67 24 17 58 16 94  9 | 51 54 71 14 95 18 42 45 73 97 24 25 92 49 29 83 38 10 30 61 22 79 43 60 65
Card 107: 74  9 87 65  8 71  6 51 47 79 | 74  5  6 71 65 82 57 88 66 76 70  8 51 87 54 81 96 79  7  9 47 12 24 52 50
Card 108: 32 87 49 10 41 34 68 79 33 23 | 16 67 63 69 60 37  8  2 47 78  5  9 34 61 10 42 36 95 68 24 75 77 85 56 18
Card 109: 98 95 89 44 76 63 83 99 71 79 | 49 64 88 18 67 48 23 37 32 99 65 14  1 63 80 84 60  5 87 94 70  8 96 16 35
Card 110: 73 30 38 99 23 57 68 39 20 45 | 74  9  1  8 35 42 23 24 83 66 76 90 30 56 82 15  6 49 88 45 96 62 43 27 16
Card 111: 48 62 64 91 57 12 68 30 25 18 | 78  3 21 29 20 77 91 74 72  1 18 28 19 25 44 52 34 12 64 17  6  7 57  5 14
Card 112: 62 10 63 18 34 55 72 86 12 45 | 82  7 27 95 76 64 11 31 42 38 75 15 69 29 79 20 17  8 16 39 60 24 87 81  9
Card 113: 91 58 97 55 46 75 74 53 72  2 | 82  7 17 38 89 81 14 24 33 13 11 91 79 93 96 18  2 35 80 15 45 63 41 46 37
Card 114: 76 22 44 83 99 59 21 67 12 17 | 78 68 63 53 16 29 31 90 70 18 69  9 40 38 32  8 27 85 67 26 62 56 97 35 94
Card 115: 11 49 95 65 18 67 29 91 45 25 | 37 44 78 20 88 35 32 60 90 40 26  5 68 53 70 31 74 63 21  4 87 46 12 92 91
Card 116: 50 25 19 90 91  8 80 44  9 38 | 40 47 45 53 99  1 74 65 49 67 16 29 81 70 78 89 42 80 97 76 66 21 24 95 83
Card 117: 38  7 34  3 57 30 29 44 52 21 |  5 48 70 24 26 50 32 27  1  4 25 96 54 16 90 61 33 67 89 78 83 22 72  6 47
Card 118: 35 25 23 73 66 57 86 21 27  8 |  9 98 50 78 75 99 83 10 70 97 49 85 14 73 46 29 16  5 60 82 15 45 62 37  3
Card 119: 51 84  3 55 90 61 93 18 38 35 | 63 50 60 78 87 74 95 58 26 27 34 11 81 30 28 99 94 89 64  1 96 65 72 56 76
Card 120: 17 74 98 56 54 31 10 26 60 45 | 20 90 98 46 31 15 89 10 94 26 35 69 79 60 58 11 74 77 50 30 64 56  9 45 12
Card 121: 27 43 12 84 96 70 61 19 92 46 | 59 18 91 25 23 97 32 72  4 21 95 93 85 60 81  9 58 30 53 94 90 69 48 89  2
Card 122: 88 85  3 82 19 98 72 38 87 93 | 60  3 31 93 51 81 82 23 69 14 72 21 95 45 76 59 86 53 85 73 19 56 61 39 35
Card 123: 22  6 50 96 91 15 33  7 29 57 | 72 67 62 78 91 96 23  7 17 69 26 29 75  2 57 20 43 74 82 68 56 92 54 66 48
Card 124: 73 15 66 38 69  4 83 45 70 95 | 19 71  6  5 15 17 95  1 34  3 68 70 42 75 12 46 38  8 83 39  2 16 74 79 36
Card 125: 32 86  2 71 72 53 73 67 12 85 | 60 31 70 99 33 45 95 22 79 41 25  3 46  5 20 75 37 55 35 18 36 59  8 63 80
Card 126: 44 96 40 68 25 34 13 72  2 80 | 47 18 77 23 63 80 17 73  3 79 21 15 11 58 54 96 95 62 51 41 75 44 40 24 97
Card 127:  4 34 21 87 85  1 44 72 71 24 | 72 46 83 10 50 77 87 49 91 85  1 42 66 38 54 89 35 86 59 65 71 15 34  4 99
Card 128: 24 86 90 74 48 83 88 13  2 38 |  8 16 10 92 76 44  1 59 34 37 80 52 27 70 66 26 22 43  6 84 30 21 60 77 74
Card 129: 28 71 81 50 31 13  6 63 58 51 |  5 35 62 15 72 47 76 44 32 34  2 14 50 45 77 18 24 81 78 59 48 40  9  7 26
Card 130: 19 68 83 40 64 99 55 45 17  9 | 22 77 99 80 90 13 88 44 57 26 85 81 15 56 30 54 18 95  6 83 51 55 94 93 73
Card 131: 47 44 33  5 59 74  1 21 46  3 |  4 14 87 41 82 63 31 45 47 23 55 94 89 39 86 20 52 66 10 59 48 34 68 21 12
Card 132: 78 15 99  9 40  5 34 75  3 79 | 70 30 80 23 95 66 85  7 43 27 12 83 22 44 21 10 31 14 26 45 68 57 49 94 53
Card 133: 54 14 80 44 33 11 53 69  1 71 | 64  2 56 28 84 76 26 32 92  3 39 47 69  9  6 75 50 99 95 30 35 73 98 13 59
Card 134: 58 85 26  9  2 94 59 23 57 52 | 21 39 80 34 50 27 45 76 16 75 54 92  8 66 53 32 17 74 83 13  1 93 99 88 46
Card 135: 43 81 39 26 61 75 13 40 72 48 |  9 87 61 99 51 81 59 55 19 13 35 74 60 45 39 40 65 33 78 46 26  8 71  7 44
Card 136: 76 86 15 32  6 81 29 16 22 92 | 88 73 92 50 75 79 82 61 86 41 30 69 18 76 45 15 67 16 29 81 93 65 32 14  6
Card 137: 37 65 99  6 61 45 27 86 68 84 |  6  7 61 67 65 33 86 64 59 90 96 54 99 23 85 68 18 55 40  4 10 37 97 77 73
Card 138: 39 80 50  4 53 73 48 54 12 60 | 63 50 72 33 17 11 24 18 54 58 68  7 30 75 22 85 87 40 19 44 69 90 97 43 36
Card 139:  8 86 87 39 58 56 80 34 85 81 | 25  9  2 26 19 61 96 72 36 45  6 21 91  7 34 89 59 56 87  8 49 98 85 74 97
Card 140: 19 21 14 84 81  4 62 24 71 77 | 96 14 77 71 38 24 45 51 40 81  5 19 99 57 48 65 72  4 60 62 41 79 87 49  2
Card 141: 87 32 58  9 88 63 22 71 83 18 | 37 26 51 34 44 22 84 48 58 65  3 62 45 90 60 70 74 81 57 68 18 47 85 73 40
Card 142: 17 56 69 87 23 26 25 32 49 47 | 65 83 15 63 73 30 93 85 45 27 84 14 58 76 38 33  8 94 61 74 37 36 52 41 78
Card 143: 52 86 94 23 28 58 81 71 30 21 | 23 96 75 87 55 50 21 76  9 80 97 69 35 33 27 71 86 94 30 42 84 10  3 14 62
Card 144: 80  8 67 32 42 49 76 93 41 90 | 21 59 34 80 71 37 28 33  9 68 95 31 26 47 73 94 41 13 48 27 96  7 85 17 62
Card 145: 22 13 74 92 75 17 19 48 58 55 | 36 31 20 92 11 76 23 51 71 84 59 88 89 38  1 90 97 46 30 41 77  6 17 64  2
Card 146: 51 81 54 87 95 28 30 44 88 85 | 96  1 35 64 45 97 54 84  4 86 40 49 15 28 19 75 10  2 68 60 66 59 12 18 63
Card 147: 44 45 70 23 71 37 17 59 97 53 | 88  9 26 22 62 68  6 96 77 31 95 50 54 42 14  5 30 24 48  3 67 75 56 49  7
Card 148: 38 54  1 75  6 24 70 82 74 94 | 23 34 92  5 14 83 45 88 81 21 43 78 87 56 63 36 48  4 51 60 42 25 18  2 35
Card 149: 11 87 55 57  8 27 48 67 12 45 |  3 37 54 81  2 15 92 30  5 10 38 98 64 93 99 68 36 50 88 97 35 29 79 46 58
Card 150: 65 89 66 91 37  3 49 19 29 17 | 24 34 69 99 15 58 56 28 90  7 73 75 72 66 77 49  8 17 21 74 84 67 19 29 54
Card 151: 86  4 93 67 52 14 43 99  9 38 | 34 60 72 47 24 76 38 67 48 85 14 43 93 92 49 45 18 25 86  9 31 99 63 41  4
Card 152: 79 39 53 12 11 17 27 51 92  5 | 12 79 51 75 92  2 56 39 81 67 17  5 53 55 72 48 60 11 18 27 98 77  6 66 86
Card 153: 22 95 26  5 32 14  7 66  8 35 | 60 32 96 25  8 57 73 95 74 27 66 26  7  1 14  5 12 34 21 35 22 87 77 24 85
Card 154: 79  4 94 63  9 96 21 86 59 46 | 52 76 55 94 57 74 46 15 38 83 32  4  9 73 62 43 36 98 89 18 96 45 79 25 35
Card 155: 53 58 97 98 67  5  8 46 31 54 | 76  5 46 80 30 54 31 33 51 97 47 67  8 83 10 58 60 73 98 71 62 82 48 95 77
Card 156:  3  1 23 18 71 21  8 79 84  2 | 59  8 18 87 29 14 46 99 61 90 66 21 68 62  2 86 31 88 23 74 38 10 73  7 84
Card 157: 24  3 22 90 56 15 78 61  1 81 | 33 70 73  5 71 93 30 67 53 18 62  7 19 85 47 65 43 82 69 75 57 46 44 99 50
Card 158:  2  8 21 71 24 82 47 53 11 25 | 27 25 35 20  3 53 73 96 38  7 65 89 16 78 97 85 74 18 56 49 32 88 59 42 33
Card 159: 69 81 74 51 35 90 33 58 97 22 | 89 45 91  1 77 10 72 35 31 42 28 71 97 80 87 22 64 33 95  3 68 56 69 82 67
Card 160: 99 34 53  3  6 48 93  1 42 13 | 50 35 79 89 26  5 91 10 57 63 30 72 64 17 37 40 90 88 78 18 25 95 67 58 14
Card 161: 41 39 34 75 50 52 94 71 82 13 | 26  8 87 95 20 91 49 96 65 36 51 24 46 69 83 52 47 41 79 11 50 73 89 70 44
Card 162: 48 87 68 57 81 52 49 60 28 15 | 92 11 94 54 45 39  7 84 93 67 42 64 26 90  6 79  4  3 23 96 41 37 97 24 50
Card 163: 67 41 11 38 81 56 45  4 76 70 | 32 87 62 51 58 27 91 37 29 90 22 43 77 97 35 46 93 88 23 31 99 21 79 18 47
Card 164:  2 79 35 21 96 29 22 10 69 99 | 47 91 17 60 20 31 37 81 53 84 16 77  4 46 75 26 78 93 32 76 52  8 27 59 10
Card 165: 32 37 54 95 66 64 59 19 28 44 | 96 49  3 75 94 63 98 65  4  7  2 73 50 24 78 31 21 42 82 25 70 79 35 38  1
Card 166: 24 61 72 44 27 38 81 59 69 89 | 85 29 98  3 73 63 25  1 16 96 57  8 78 92 82 20 79 94 71 62  2 49 97 93 99
Card 167: 55 22 60 95 80 25 28 56 69 18 | 22 38 29 46 90 41 59 93 94 57 74 43 97 49 27 34 19 64 55 81 33 37 13 15 89
Card 168: 90 51 39 37 67 12 65 14  6  4 | 64 70  6 37 25 26  4 48 43 91 28 90 27 41 87 74 14 39 56 51  9 67 12 65 34
Card 169: 57 11 25 88 28 82 60 95 27 26 | 42 82 35 47 87 14 11 88 63 96 65 43 95 25 99  5 26 20 28 57 60 40 76 27  7
Card 170: 31 34 19 29  1 55 20 61 10 94 | 48 77  3 87 91 37 38 57 75 53 12 60 76 45  4 68 13 73 35  7 30 20  6 14 69
Card 171: 67 10 59 89 52 53 40 17 64 29 | 88 27 59  1 75 10 95 67 34 40 83  4 82 64  2 52 89 29 53 33  8 71 91 85 17
Card 172: 87 63 66 32 37 96  3 29 88 90 | 96 18 87 17 14 69 58 65 52 31  1 44 37 60  5 32 41 90  7 49 29 66 35 77 82
Card 173: 70 84 52 81 63 32 61 94 38 97 | 24 34 63 52 81 99  8 57 66 50 94 84 70 38 61 86 15 32 47 78 45 77 95 76 97
Card 174:  4 74 29 30 24 87 91 92  5  2 | 55 86  5 70 12 94 96 89 58 29 14 62  4 99 48 81 91 45 27 54 31 22 68 43 76
Card 175: 91 78  3 20 86 98 89 82 13 57 | 24 76 59 66 73  1  5 82 45  6 92  3 33 43 17 83 12 14 91 71 19 46 54 96 25
Card 176: 12 55 93  3 43 81 28 59 72 76 |  4 90 55 84 69 74 46 45 65 17 26 71 70 21 29 23 61 99 92 44 40 41 51  8 57
Card 177: 56  5 89 93 62 17  2  6 50 40 | 35 60  5 47 76 92 15  2 34 28 69 29 19  9 27 49  1 88 57 48 84  8 26 59 91
Card 178:  1 93 34 57 67 58 84 37 42 91 | 42 69 41 34 91 28 75 70  1 93  4 49 66  6 37 38 32 89 46 44 62 81  9 22 73
Card 179: 90 95  7 29 42 51 22 39 84 75 | 75 25 80 96 55 41 73 43 26 42 46  7 86 22 84  1 39 50 90 12  3 29 51 95 81
Card 180: 72 76  4 68  9 41 16 44 61 13 | 63 92 39 68  5 13 11 65 50 25 76 55 71 29 82 31 53 59  2 99 74 16  1 90 41
Card 181: 29 63 61 82 33 94 51 32 50 84 | 80 85  7  8 54 30 87 56 90 75 69 60 25 43 53 47 72 81  1 59 97 28 52 46 10
Card 182: 98 45 44 32 19  6 51 49 64 84 | 64 20 90  6  4 54 19 80 51 72 69 99 29 44 45 46 84 65 43 53 74 23 42 66 49
Card 183: 23 35  4 50 88 87 38  7 19  2 | 76 79 80  5 66 30 74 29 37 94 11 50  1 67 13 27 90 40 70 91 71 12 81  9 33
Card 184: 55 74  2 73 75 22 32 71 67 27 | 92 19  9 24 74 73  8 67 30 18 14 71 77 45 85 72 15 69 13 27 94 63 98 40 93
Card 185: 14 73 72 38 16 40 83 28 37 96 | 19 74 49 96 90 48 26  1 95 54  3 82 50 41 43 64 69 80 77 97 52 16 23 25 31
Card 186: 24 91 35 21 16 71 94 95 25 53 | 60  5 17 58 82 35 45 80 28 16 13 53 68 97 86  8 52 61 65  1 27 67 91 43  3
Card 187: 96 39 77 52 97 33 80 99 18 15 | 48  8 85 79 81 33 43 90 62 14 36 65  2 32 82 15 91 23 34 68 56 87 11 57 73
Card 188:  8 35 86 95 94 65 26 11 96 31 |  2 98 50 33 59 93 28 49 87 29 79  8 23  6 54 16 82 96 83 61 27 60 53 62 30
Card 189: 60 27 78 55 84  1 36 28 20 77 | 45 85 75  8 54 74 58 73 17 68 13 53 47 79  7 65 35 40 51 87 18 37 71 72 21
Card 190: 41 31 39 33 54 42 71 47 59 24 | 21 96 85 12 81 83 64 87 93 77 92 38 25 52 20 88 65 10 29 16 95 98 22 37 15`
