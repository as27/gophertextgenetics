package main

import (
	"testing"
)

func TestMe(t *testing.T) {
	main()
	t.Error("--")
}

func Test_fitness(t *testing.T) {
	type args struct {
		s      []byte
		target []byte
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"match all",
			args{[]byte("test"), []byte("test")},
			1,
		},
		{
			"no match",
			args{[]byte("test"), []byte("aaaa")},
			0,
		},
		{
			"half match",
			args{[]byte("test"), []byte("teaa")},
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
