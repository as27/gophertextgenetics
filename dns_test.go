package main

import (
	"reflect"
	"testing"
)

func TestDNS_Mutate(t *testing.T) {
	type args struct {
		mutationRate float64
	}
	tests := []struct {
		name string
		d    *DNS
		args args
	}{
		{
			"t1",
			NewDNS([]byte("This is a test")),
			args{0.5},
		},
		{
			"t1",
			NewDNS([]byte("This is a test")),
			args{0.1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.Mutate(tt.args.mutationRate)
		})
	}
	t.Error("")
}

func TestNewRandomDNS(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want *DNS
	}{
		{
			"t1",
			args{10},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomDNS(tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomDNS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNS_Reproduce(t *testing.T) {
	type args struct {
		father *DNS
	}
	tests := []struct {
		name string
		d    *DNS
		args args
		want *DNS
	}{
		{
			"test",
			NewDNS([]byte{100, 100, 100, 100, 100}),
			args{NewDNS([]byte{200, 200, 200, 200, 200})},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Reproduce(tt.args.father); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DNS.Reproduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
