package ca

import (
	"log"
	"reflect"
	"testing"
)

func TestUpdate(t *testing.T) {
	cases := []struct {
		config []bool
		radius int
		rules  uint32
		w      []bool
	}{
		{
			[]bool{false, true, false, false, true, true},
			1,
			1<<0 + 1<<2 + 1<<5,
			[]bool{true, true, false, false, false, false},
		},
		{
			[]bool{true, true, false, true, false, false, false, true, true},
			2,
			1<<30 + 1<<20 + 1<<3 + 1<<7 + 1<<15 + 1<<28 + 1<<31 + 1<<21 + 1<<19 + 1<<1 + 1<<2 + 1<<4,
			[]bool{true, false, false, true, false, false, true, true, true},
		},
		{
			[]bool{true, false, true},
			1,
			0,
			[]bool{false, false, false},
		},
		{
			[]bool{false, false, true},
			1,
			1<<7 + 1<<6 + 1<<5 + 1<<4 + 1<<3 + 1<<2 + 1<<1 + 1<<0,
			[]bool{true, true, true},
		},
		{
			[]bool{true, true, true, true, false},
			2,
			1<<(1+2+4+8) + 1<<(2+4+8+16) + 1<<(1+2+8+16) + 1<<(2+4+16) + 1<<(1+4+8) + 1<<(2+8+16) + 1<<(1+4+16),
			[]bool{false, true, true, false, true},
		},
	}
	for _, tt := range cases {
		got, err := Update(tt.config, tt.radius, tt.rules)
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
		config []bool
		radius int
		rules  uint32
		want   string
	}{
		{make([]bool, 2), 1, 0, "need len(config) >= 2*radius + 1 = 3, but len(config) was 2"},
		{make([]bool, 4), 2, 0, "need len(config) >= 2*radius + 1 = 5, but len(config) was 4"},
		{make([]bool, 10), 3, 0, "radius must be 1 or 2, but was 3"},
		{make([]bool, 10), 0, 0, "radius must be 1 or 2, but was 0"},
	}
	for _, tt := range cases {
		_, err := Update(tt.config, tt.radius, tt.rules)
		if err == nil {
			log.Fatal("want error, but was nil")
		}
		if err.Error() != tt.want {
			log.Fatalf("want error.Error()=%v, but was %v", tt.want, err.Error())
		}
	}
}

func TestUpdateN(t *testing.T) {
	cases := []struct {
		config []bool
		radius int
		rules  uint32
		n      int
		want   [][]bool
	}{
		{
			[]bool{false, false, false}, 1, 1, 4,
			[][]bool{
				[]bool{false, false, false},
				[]bool{true, true, true},
				[]bool{false, false, false},
				[]bool{true, true, true},
				[]bool{false, false, false},
			},
		},
		{
			[]bool{true, false, false},
			1,
			1<<(1+4) + 1<<(2) + 1<<(1),
			6,
			[][]bool{
				[]bool{true, false, false},
				[]bool{true, false, true},
				[]bool{false, true, false},
				[]bool{true, true, false},
				[]bool{false, false, true},
				[]bool{false, true, true},
				[]bool{true, false, false},
			},
		},
		{
			[]bool{true, true, true, true, false},
			2,
			1<<(1+2+4+8) + 1<<(2+4+8+16) + 1<<(1+2+8+16) + 1<<(2+4+16) + 1<<(1+4+8) + 1<<(2+8+16) + 1<<(1+4+16),
			3,
			[][]bool{
				[]bool{true, true, true, true, false},
				[]bool{false, true, true, false, true},
				[]bool{false, true, true, true, true},
				[]bool{true, false, true, true, false},
			},
		},
	}
	for _, tt := range cases {
		got, err := UpdateN(tt.config, tt.radius, tt.rules, tt.n)
		if err != nil {
			log.Fatal(err)
		}
		if !reflect.DeepEqual(tt.want, got) {
			t.Errorf("want %v, but got %v", tt.want, got)
		}
	}
}

func TestUpdateNError(t *testing.T) {
	cases := []struct {
		config []bool
		radius int
		rules  uint32
		n      int
		want   string
	}{
		{make([]bool, 2), 1, 0, 1, "need len(config) >= 2*radius + 1 = 3, but len(config) was 2"},
		{make([]bool, 4), 2, 0, 1, "need len(config) >= 2*radius + 1 = 5, but len(config) was 4"},
		{make([]bool, 10), 3, 0, 1, "radius must be 1 or 2, but was 3"},
		{make([]bool, 10), 0, 0, 1, "radius must be 1 or 2, but was 0"},
		{make([]bool, 10), 1, 0, 0, "need n > 0, but was 0"},
	}
	for _, tt := range cases {
		_, err := UpdateN(tt.config, tt.radius, tt.rules, tt.n)
		if err == nil {
			log.Fatal("want error, but was nil")
		}
		if err.Error() != tt.want {
			log.Fatalf("want error.Error()=%v, but was %v", tt.want, err.Error())
		}
	}
}
