package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	targetText := "SagMirWoDieBlumenSind"
	mutationRate := float64(0.2)
	population := 600
	fmt.Println(targetText, mutationRate, population)
	pop := CreateNewPopulation(population, len(targetText))
	pop.Fitness = func(d *DNS) float64 {
		match := 0
		for i, b := range []byte(targetText) {
			if b == d.content[i] {
				match++
			}
		}

		return float64(match)
	}
	pop.Sort()

	counter := 0
	for {
		//pop.PrintN(10)
		//wait()
		bestMatch := string(pop.dnss[0].content)
		//fmt.Println(pop.Fitness(pop.dnss[0]), bestMatch, pop.dnss[0].content, pop.dnss[1].content, pop.dnss[2].content)
		if bestMatch == targetText {
			fmt.Println("Counter", counter, bestMatch)
			break
		}
		if counter == 200 {
			pop.PrintN(50)
			break
		}

		counter++
		nextGen := pop.NextGeneration(mutationRate)

		pop = nextGen
		pop.Sort()
	}
}
func wait() {
	var userIn string
	fmt.Println("Press ENTER to continue")
	fmt.Println("Press q to quit")
	fmt.Scanln(&userIn)
	if userIn == "q" {
		os.Exit(4)
	}
}

func fitness(s, target []byte) float64 {
	match := 0
	for i := range target {
		if target[i] == s[i] {
			match++
		}
	}
	return float64(match) / float64(len(target))
}
