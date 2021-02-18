package ca

import (
	"log"
	"testing"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		c      Config
		r      Rules
		radius int
		w      Config
	}{
		{
			Config{false, true, false, false, true, true},
			1<<0 + 1<<2 + 1<<5,
			1,
			Config{true, true, false, false, false, false},
		},
		{
			Config{true, true, false, true, false, false, false, true, true},
			1<<30 + 1<<20 + 1<<3 + 1<<7 + 1<<15 + 1<<28 + 1<<31 + 1<<21 + 1<<19 + 1<<1 + 1<<2 + 1<<4,
			2,
			Config{true, false, false, true, false, false, true, true, true},
		},
		{
			Config{true, false, true},
			0,
			1,
			Config{false, false, false},
		},
		{
			Config{false, false, true},
			1<<7 + 1<<6 + 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1 + 1<<0,
			1,
			Config{true, true, true},
		},
	}
	for _, tt := range cases {
		got, err := Update(tt.c, tt.radius, tt.r)
		if err != nil {
			t.Fatal(err)
		}
		want := tt.w
		isEqual := len(want) == len(got)
		for i := 0; isEqual && i < len(want); i++ {
			isEqual = want[i] == got[i]
		}
		if !isEqual {
			t.Errorf("want %v, but got %v", want, got)
		}
	}
}

func TestUpdateError(t *testing.T) {
	cases := []struct {
		c      Config
		r      Rules
		radius int
		want   string
	}{
		{make([]bool, 2), 0, 1, "need len(c) >= 2*radius + 1 = 3, but len(c) was 2"},
		{make([]bool, 4), 0, 2, "need len(c) >= 2*radius + 1 = 5, but len(c) was 4"},
		{nil, 0, 3, "radius must be 1 or 2, but was 3"},
		{nil, 0, 0, "radius must be 1 or 2, but was 0"},
	}
	for _, tt := range cases {
		_, err := Update(tt.c, tt.radius, tt.r)
		if err == nil {
			log.Fatal("want error, but was nil")
		}
		if err.Error() != tt.want {
			log.Fatalf("want error.Error()=%v, but was %v", tt.want, err.Error())
		}
	}
}
