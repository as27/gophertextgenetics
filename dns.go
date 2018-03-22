package main

import (
	"fmt"
	"math/rand"
)

type DNS struct {
	content []byte
	Fitness int
}

func NewDNS(b []byte) *DNS {
	return &DNS{
		content: b,
	}
}

func NewRandomDNS(length int) *DNS {
	var b []byte
	for i := 0; i < length; i++ {
		b = append(b, randByte())
	}
	return NewDNS(b)
}

func (d *DNS) Mutate(mutationRate float64) {
	for i, _ := range d.content {
		if rand.Float64() <= mutationRate {
			d.content[i] = randByte()
		}
	}
}

func randByte() byte {
	//return byte(rand.Intn(65) + 60)
	return byte(rand.Intn(255))
}

func (d *DNS) Reproduce(father *DNS) *DNS {
	childDNS := make([]byte, len(d.content))
	for i, _ := range childDNS {
		/*if rand.Float64() <= 0.2 {
			childDNS[i] = father.content[i]
		}*/
		if i%2 == 0 {
			childDNS[i] = father.content[i]
		} else {
			childDNS[i] = d.content[i]
		}
	}
	/*fmt.Println("mum", d)
	fmt.Println("dad", father)
	fmt.Println("chi", string(childDNS))*/
	return NewDNS(childDNS)
}

func (d *DNS) String() string {
	return fmt.Sprintf("%s", string(d.content))
}
