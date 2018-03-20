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
	dnss := make([]*DNS, popSize)
	for i := 0; i < popSize; i++ {
		dnss[i] = NewRandomDNS(dnsLength)
	}
	return &Population{
		dnss: dnss,
	}
}

func (p *Population) NextGeneration(mutationRate float64) *Population {
	ng := Population{}
	ng.Fitness = p.Fitness
	dnss := make([]*DNS, p.Size())
	for i := 0; i < p.Size(); i++ {
		dnsMum := p.PickDNS()
		dnsDad := p.PickDNS()
		child := dnsMum.Reproduce(dnsDad)
		child.Mutate(mutationRate)
		/*if child.String() == dnsMum.String() {
			child = NewRandomDNS(len(child.content))
		}*/
		//fmt.Println(child)
		dnss[i] = child
	}
	ng.dnss = dnss
	return &ng
}

func (p *Population) PickDNS() *DNS {
	if p.fitnessSum == 0 {
		var fitnessSum float64
		for _, d := range p.dnss {
			fitnessSum = fitnessSum + p.Fitness(d)
		}
		p.fitnessSum = fitnessSum
	}

	r := rand.Float64() * p.fitnessSum
	fitMin := float64(0)
	fitMax := float64(0)
	for _, d := range p.dnss {
		fitMax = fitMin + p.Fitness(d)
		if fitMin <= r && r <= fitMax {
			return d
		}
		fitMin = fitMax
	}
	fmt.Println("----->", r, p.fitnessSum)
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

func (p *Population) PrintN(n int) {
	for i, dns := range p.dnss {
		if i == n {
			break
		}
		if dns == nil {
			continue
		}
		fmt.Println(
			dns,
			p.Fitness(dns),
			dns.content)
	}

	fmt.Println("----------------------")
}
