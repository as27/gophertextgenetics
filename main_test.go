package main

import "testing"

func Test_fitness(t *testing.T) {
	type args struct {
		s      string
		target string
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"match all",
			args{"test", "test"},
			1,
		},
		{
			"no match",
			args{"test", "aaaa"},
			0,
		},
		{
			"half match",
			args{"test", "teaa"},
			0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fitness(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}
