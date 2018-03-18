package main

import (
	"math/rand"
)

type DNS struct {
	content []byte
}

func NewDNS(b []byte) *DNS {
	return &DNS{
		content: b,
	}
}

func NewRandomDNS(length int) *DNS {
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, byte(rand.Intn(255)))
	}
	return NewDNS(b)
}

func (d *DNS) Mutate(mutationRate float64) {
	for i, _ := range d.content {
		if rand.Float64() <= mutationRate {
			d.content[i] = byte(rand.Intn(255))
			//fmt.Println("hier")
		}
	}
}

func (d *DNS) Reproduce(father *DNS) *DNS {
	childDNS := d.content
	for i, _ := range childDNS {
		if rand.Float32() <= 0.5 {
			childDNS[i] = father.content[i]
		}
	}
	return NewDNS(childDNS)
}
