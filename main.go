package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	flagMutationRate   = flag.Float64("mr", float64(0.01), "mutation rate")
	flagPopulation     = flag.Int("p", 800, "population")
	flagTargetText     = flag.String("t", "This is the secret text", "target text")
	flagMaxGenerations = flag.Int("max", 5000, "max generations")
)

func main() {
	flag.Parse()
	start := time.Now()
	targetText := *flagTargetText
	mutationRate := *flagMutationRate
	population := *flagPopulation
	fmt.Println(targetText, mutationRate, population)
	pop := CreateNewPopulation(population, len(targetText))
	pop.Fitness = func(d *DNS) int {
		match := 0
		for i, b := range []byte(targetText) {
			if b == d.content[i] {
				match++
			}
		}

		return match * match * match
	}
	pop.CalcFitness()
	pop.Sort()

	counter := 0
	for {
		if counter%20 == 0 {
			pop.PrintN(5)
		}

		//wait()
		bestMatch := string(pop.dnss[0].content)
		//fmt.Println(pop.Fitness(pop.dnss[0]), bestMatch, pop.dnss[0].content, pop.dnss[1].content, pop.dnss[2].content)
		if bestMatch == targetText {
			fmt.Println(targetText, mutationRate, population)
			fmt.Println("Counter", counter, bestMatch)
			end := time.Now()
			fmt.Println("Time:", end.Sub(start))
			break
		}
		if counter == *flagMaxGenerations {
			pop.PrintN(10)
			break
		}

		counter++
		nextGen := pop.NextGeneration(mutationRate)

		pop = nextGen
		pop.CalcFitness()
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
