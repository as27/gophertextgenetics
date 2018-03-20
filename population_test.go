package main

import (
	"reflect"
	"testing"
)

func TestPopulation_Sort(t *testing.T) {
	dns1 := NewDNS([]byte{200})
	dns2 := NewDNS([]byte{180})
	dns3 := NewDNS([]byte{10})

	tests := []struct {
		name string
		dnss []*DNS
		want []*DNS
	}{
		{
			"sorted",
			[]*DNS{dns1, dns2, dns3},
			[]*DNS{dns1, dns2, dns3},
		},
		{
			"unsorted",
			[]*DNS{dns2, dns1, dns3},
			[]*DNS{dns1, dns2, dns3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pop := &Population{
				dnss:    tt.dnss,
				Fitness: func(d *DNS) float64 { return float64(d.content[0]) },
			}
			pop.Sort()
			if !reflect.DeepEqual(tt.want, pop.dnss) {
				t.Errorf("Sort() = \n%s, want \n%s", pop.dnss, tt.want)
			}

		})
	}
}

func TestPopulation_PickDNS(t *testing.T) {
	dns1 := NewDNS([]byte{200})
	dns2 := NewDNS([]byte{180})
	dns3 := NewDNS([]byte{100})
	dns4 := NewDNS([]byte{80})
	dns5 := NewDNS([]byte{1})
	dns6 := NewDNS([]byte{0})

	tests := []struct {
		name string
		dnss []*DNS
		want *DNS
	}{
		{
			"with 0",
			[]*DNS{dns6, dns5, dns6, dns6},
			nil,
		},
		{
			"t1",
			[]*DNS{dns1, dns2, dns3, dns4, dns5},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pop := &Population{
				dnss:    tt.dnss,
				Fitness: func(d *DNS) float64 { return float64(d.content[0]) },
			}
			pop.Sort()
			if got := pop.PickDNS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Population.PickDNS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewPopulation(t *testing.T) {
	type args struct {
		popSize   int
		dnsLength int
	}
	tests := []struct {
		name string
		args args
		want *Population
	}{
		{
			"10 entries",
			args{10, 5},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateNewPopulation(tt.args.popSize, tt.args.dnsLength); !reflect.DeepEqual(got, tt.want) {
				got.PrintN(5)
				t.Errorf("CreateNewPopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckNextGeneration(t *testing.T) {
	pop := CreateNewPopulation(10, 10)
	pop.Fitness = func(d *DNS) float64 { return float64(d.content[0]) }
	pop.Sort()
	pop.PrintN(10)
	pop = pop.NextGeneration(float64(0.1))
	pop.Sort()
	pop.PrintN(10)
	pop = pop.NextGeneration(float64(0.1))
	pop.Sort()
	pop.PrintN(10)
	pop = pop.NextGeneration(float64(0.1))
	pop.Sort()
	pop.PrintN(10)
	t.Error("---")
}
