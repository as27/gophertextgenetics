package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	targetText := "You need to guess this text"
	mutationRate := 0.01
	population := 300
	fmt.Println(targetText, mutationRate, population)
}

func fitness(s, target string) float64 {
	match := 0
	for i := range target {
		if target[i] == s[i] {
			match++
		}
	}
	return float64(match) / float64(len(target))
}
