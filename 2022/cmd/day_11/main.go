package main

import (
	"fmt"
	"math"
)

func main() {
	monkeys := loadMonkeys()
	fmt.Printf("Task1: %q\n", Task1(monkeys))
	fmt.Printf("Task2: %q\n", Task2(monkeys))
}

func Task1(monkeys []Monkey) string {
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, item := range monkeys[j].Items {
				worry := monkeys[j].Operation(item, math.MaxInt) / 3
				next := monkeys[j].Test(worry)
				monkeys[next].Items = append(monkeys[next].Items, worry)
				monkeys[j].Inspected++
			}
			monkeys[j].Items = []int{}
		}
	}

	max1 := monkeys[0].Inspected
	max2 := monkeys[1].Inspected
	if max2 > max1 {
		max1, max2 = max2, max1
	}

	for i := 2; i < len(monkeys); i++ {
		if monkeys[i].Inspected > max2 {
			max2 = monkeys[i].Inspected
			if max2 > max1 {
				max1, max2 = max2, max1
			}
		}
	}

	return fmt.Sprint(max1 * max2)
}

func Task2(monkeys []Monkey) string {
	div := 1
	for _, monkey := range monkeys {
		div *= monkey.Eval()
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			for _, item := range monkeys[j].Items {
				worry := monkeys[j].Operation(item, div)
				next := monkeys[j].Test(worry)
				monkeys[next].Items = append(monkeys[next].Items, worry)
				monkeys[j].Inspected++
			}
			monkeys[j].Items = []int{}
		}
	}

	max1 := monkeys[0].Inspected
	max2 := monkeys[1].Inspected
	if max2 > max1 {
		max1, max2 = max2, max1
	}

	fmt.Println(max1)
	fmt.Println(max2)

	for i := 2; i < len(monkeys); i++ {
		fmt.Println(monkeys[i].Inspected)
		if monkeys[i].Inspected > max2 {
			max2 = monkeys[i].Inspected
			if max2 > max1 {
				max1, max2 = max2, max1
			}
		}
	}

	fmt.Println()
	fmt.Println(max1)
	fmt.Println(max2)

	return fmt.Sprint(max1 * max2)
}

type Monkey struct {
	Items     []int
	Operation func(int, int) int
	Test      func(int) int
	Eval      func() int
	Inspected int
}

func loadMonkeys() []Monkey {
	return []Monkey{
		{
			Items: []int{93, 98},
			Operation: func(v, mod int) int {
				return (v * 17) % mod
			},
			Test: func(v int) int {
				if v%19 == 0 {
					return 5
				}
				return 3
			},
			Eval: func() int {
				return 19
			},
		},
		{
			Items: []int{95, 72, 98, 82, 86},
			Operation: func(v int, mod int) int {
				return (v + 5) % mod
			},
			Test: func(v int) int {
				if v%13 == 0 {
					return 7
				}
				return 6
			},
			Eval: func() int {
				return 13
			},
		},
		{
			Items: []int{85, 62, 82, 86, 70, 65, 83, 76},
			Operation: func(v int, mod int) int {
				return (v + 8) % mod
			},
			Test: func(v int) int {
				if v%5 == 0 {
					return 3
				}
				return 0
			},
			Eval: func() int {
				return 5
			},
		},
		{
			Items: []int{86, 70, 71, 56},
			Operation: func(v int, mod int) int {
				return (v + 1) % mod
			},
			Test: func(v int) int {
				if v%7 == 0 {
					return 4
				}
				return 5
			},
			Eval: func() int {
				return 7
			},
		},
		{
			Items: []int{77, 71, 86, 52, 81, 67},
			Operation: func(v int, mod int) int {
				return (v + 4) % mod
			},
			Test: func(v int) int {
				if v%17 == 0 {
					return 1
				}
				return 6
			},
			Eval: func() int {
				return 17
			},
		},
		{
			Items: []int{89, 87, 60, 78, 54, 77, 98},
			Operation: func(v int, mod int) int {
				return (v * 7) % mod
			},
			Test: func(v int) int {
				if v%2 == 0 {
					return 1
				}
				return 4
			},
			Eval: func() int {
				return 2
			},
		},
		{
			Items: []int{69, 65, 63},
			Operation: func(v int, mod int) int {
				return (v + 6) % mod
			},
			Test: func(v int) int {
				if v%3 == 0 {
					return 7
				}
				return 2
			},
			Eval: func() int {
				return 3
			},
		},
		{
			Items: []int{89},
			Operation: func(v int, mod int) int {
				return (v * v) % mod
			},
			Test: func(v int) int {
				if v%11 == 0 {
					return 0
				}
				return 2
			},
			Eval: func() int {
				return 11
			},
		},
	}
}
