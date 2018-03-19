package main

import (
	"math/rand"
	"time"
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
	rand.Seed(time.Now().UnixNano())
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, byte(rand.Intn(60)+65))
	}
	return NewDNS(b)
}

func (d *DNS) Mutate(mutationRate float64) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range d.content {
		if rand.Float64() <= mutationRate {
			d.content[i] = byte(rand.Intn(60) + 65)
		}
	}
}

func (d *DNS) Reproduce(father *DNS) *DNS {
	rand.Seed(time.Now().UnixNano())
	childDNS := d.content
	for i, _ := range childDNS {
		if rand.Float32() <= 0.5 {
			childDNS[i] = father.content[i]
		}
	}
	return NewDNS(childDNS)
}
