package ca

/*
  1-dimensional, binary cellular automaton.
*/

import (
	"errors"
	"fmt"
)

// Update applies the rule table to every cell in the configuration and returns the new configuration.
func Update(config []bool, radius int, rules uint32) ([]bool, error) {
	c := config
	r := rules
	if radius != 1 && radius != 2 {
		return nil, errors.New(fmt.Sprintf("radius must be 1 or 2, but was %v", radius))
	}
	if len(c) < 2*radius+1 {
		return nil, errors.New(fmt.Sprintf("need len(config) >= 2*radius + 1 = %v, but len(config) was %v", 2*radius+1, len(c)))
	}

	newConfig := make([]bool, len(c))
	for i := range c {
		lo := i - radius
		hi := i + radius
		hood := c[Max(lo, 0):Min(hi+1, len(c))]

		if lo < 0 {
			hood = append(c[lo+len(c):], hood...)
		}
		if hi >= len(c) {
			hood = append(hood, c[:hi-len(c)+1]...)
		}

		hoodState := 0
		for _, isAlive := range hood {
			if isAlive {
				hoodState += 1
			}
			hoodState <<= 1
		}
		hoodState >>= 1 // Undo do last shift

		newState := 1<<hoodState&r > 0
		newConfig[i] = newState
	}
	return newConfig, nil
}

// UpdateN updates the cellular automaton n times. The function returns the sequence of configurations, the earliest being at the beginning of the list and the last at the end.
func UpdateN(config []bool, radius int, r uint32, n int) ([][]bool, error) {
	if n <= 0 {
		return nil, errors.New(fmt.Sprintf("need n > 0, but was %v", n))
	}

	configs := make([][]bool, n+1)
	configs[0] = config
	for i := 0; i < n; i++ {
		c, err := Update(configs[i], radius, r)
		if err != nil {
			return nil, err
		}
		configs[i+1] = c
	}
	return configs, nil
}

// RandRule returns random rules with radius r. r must be 1 or 2.
func RandRule(r int) (uint32, error) {
	return 0, nil
}

// RandConfig returns a random configuration consisting of n cells
func RandConfig(n int) ([]bool, error) {
	return nil, nil
}

// randConfig returns a random configuration consisting of n cells where more than half of cells are alive or dead depending on the provided argument.
func randConfigHalf(n int, moreThanHalfAlive bool) ([]bool, error) {
	return nil, nil
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
