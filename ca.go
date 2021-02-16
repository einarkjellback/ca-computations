package cacomp

/*
  1-dimensional, binary cellular automaton.
*/

const RADIUS = 2

// Config is the configuration of a binary cellular automaton. It is a sequence of 0's and 1's and thus representable as an unsigned integer.
type Config []bool

// Specifies the rules for the state transitions given a particular neighborhood configuration.
type Rules uint32

// Update applies the rule table to every cell in the configuration and returns the new configuration.
func Update(c Config, radius int, r Rules) (Config, error) {
	newConfig := make([]bool, len(c))
	for i := range c {
		lo := i - radius
		if lo < 0 {

		}
		hi := i + radius
		if hi >= len(c) {
			hi -= len(c)
		}

		neighState := 0
		for j := lo; j <= hi; j++ {
			isAlive := neighbors[j]
			if isAlive {
				neighState |= 1
			}
			neighState <<= 1
		}

		newState := 1<<neighState&r > 0
		newConfig[i] = newState
	}
	return newConfig, nil
}

// UpdateN updates the cellular automaton n times. The function returns the sequence of configurations, the earliest being at the beginning of the list and the last at the end.
func UpdateN(c Config, r Rules, n int) ([]Config, error) {
	return nil, nil
}

// RandRule returns random rules with radius r. r must be 1 or 2.
func RandRule(r int) (Rules, error) {
	return 0, nil
}

// RandConfig returns a random configuration consisting of n cells
func RandConfig(n int) (Config, error) {
	return nil, nil
}

// randConfig returns a random configuration consisting of n cells where more than half of cells are alive or dead depending on the provided argument.
func randConfigHalf(n int, moreThanHalfAlive bool) (Config, error) {
	return nil, nil
}
