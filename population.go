package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Population struct {
	dnss       []*DNS
	fitnessSum float64
	Fitness    func(*DNS) float64
}

func CreateNewPopulation(popSize, dnsLength int) *Population {
	var p Population
	for i := 0; i < popSize; i++ {
		p.dnss = append(p.dnss, NewRandomDNS(dnsLength))
	}
	return &p
}

func (p *Population) NextGeneration(mutationRate float64) *Population {
	ng := Population{
		Fitness: p.Fitness,
	}
	for i := 0; i < p.Size(); i++ {
		dnsMum := p.PickDNS()
		dnsDad := p.PickDad(dnsMum)
		child := dnsMum.Reproduce(dnsDad)
		child.Mutate(mutationRate)
		ng.dnss = append(ng.dnss, child)
	}
	ng.Sort()
	return &ng
}

func (p *Population) PickDad(mum *DNS) *DNS {
	dad := p.PickDNS()
	if dad == mum {
		return p.PickDad(mum)
	}
	return dad
}

func (p *Population) PickDNS() *DNS {
	if p.fitnessSum == 0 {
		for _, d := range p.dnss {
			p.fitnessSum = p.fitnessSum + p.Fitness(d)
		}
	}
	r := rand.Float64() * p.fitnessSum

	fitMin := float64(0)
	fitMax := float64(0)
	for _, d := range p.dnss {
		fitMax = fitMin + p.Fitness(d)
		fmt.Println(r, fitMin, fitMax)
		if fitMin <= r && r <= fitMax {
			return d
		}
		fitMin = fitMax
	}
	return p.dnss[0]
}

func (p *Population) Size() int {
	return len(p.dnss)
}

// Sort the population according to the fitness of the dnss higher
// fitness is sorted first
func (p *Population) Sort() {
	sort.Slice(p.dnss, func(i, j int) bool {
		return p.Fitness(p.dnss[i]) > p.Fitness(p.dnss[j])
	})
}
