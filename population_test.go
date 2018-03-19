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
	dns4 := NewDNS([]byte{10})

	tests := []struct {
		name string
		dnss []*DNS
		want *DNS
	}{
		{
			"t1",
			[]*DNS{dns1, dns2, dns3, dns4},
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
