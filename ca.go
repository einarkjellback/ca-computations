package cacomp

/*
  1-dimensional, binary cellular automaton.
*/

type CA struct {
	history []Config
	size    int // Must be <= 64
	rules   Rules
}

// Config is the configuration of a binary cellular automaton. It is a sequence of 0's and 1's and thus representable as an unsigned integer.
type Config uint64

// false is the 'quiescent' state while true is the 'alive' state.
type State bool

// Specifies the rules for the state transitions given a particular neighborhood configuration.
type Rules uint32

// Update applies the rule table to every cell in the current configuration and moves the cellular automaton to the next iteration.
func (ca *CA) Update() *CA {
	return nil
}

// NewCA makes a new binary cellular automata with a given configuration,  neighborhood radius, and rule table. It is required that 3 <= c.length <= 64.
func NewCA(c Config, r Rules) (*CA, error) {
	return nil, nil
}
