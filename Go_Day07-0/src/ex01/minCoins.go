package main

import (
	"sort"
)

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	if len(coins) == 0 {
		return []int{}
	}
	res := make([]int, 0, len(coins)/2)
	setCoins := make([]int, 0, len(coins))
	set := make(map[int]struct{})
	for i := 0; i < len(coins); i++ {
		_, ok := set[coins[i]]
		if !ok {
			set[coins[i]] = struct{}{}
			setCoins = append(setCoins, coins[i])
		}
	}
	sort.Slice(setCoins, func(i, j int) bool {
		return setCoins[i] < setCoins[j]
	})

	i := len(setCoins) - 1
	for i >= 0 {
		for val >= setCoins[i] {
			val -= setCoins[i]
			res = append(res, setCoins[i])
		}
		i -= 1
	}
	return res
}
