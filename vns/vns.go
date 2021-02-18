package vns

import (
	"github.com/einarkjellback/cacomp/ca"
)

/*
  Variable neighborhood search algorithm for optimizing the rule table of a 1-dimensional binary cellular automata computing the density classification task.

  General Variable Neighborhood Search (GVNS) Pseudocode:
    while ![stopping condition] {
      k := 1
      while k <= k_max {
        x' := Shake(x, k)
        x'' := VND(x', l_max)
        x, k := NeighborhoodChange(x, x'', k)
      }
    }
    return x
*/

// Run performs a general variable neighborhood search on the density classification problem. It stops after iterating for maxIter.
func Run(kmax int, lmax int, maxIter int) ([]ca.Rules, map[ca.Rules]float64, map[ca.Rules][]ca.Config, error) {
	// Generate random configuration

	// run GVNS (See pseudocode above)
	return nil, nil, nil, nil
}

// shake picks a random solution from the k-neighborhood.
func shake(r ca.Rules, k int) ca.Rules {
	// Flip k randomly selected bits and return result
	return 0
}

// vnd finds the fittest candidate from the k=1 to k=kMax neighborhoods
func vnd(r ca.Rules, kMax int) ca.Rules {
	for k := 1; k <= kMax; {
		// Find fittest candidate from k-neighborhood
		next := findFittest(r, k)
		r, k = neighborhoodChange(r, next, k)
	}
	// Return fittest candidate found
	return 0
}

func neighborhoodChange(curr ca.Rules, next ca.Rules, k int) (ca.Rules, int) {
	currFit := fitness(curr)
	nextFit := fitness(next)
	if nextFit > currFit {
		k = 1
		curr = next
	} else {
		k += 1
	}
	return curr, k
}

func findFittest(r ca.Rules, k int) ca.Rules {
	return 0
}

func fitness(r ca.Rules) float64 {
	// Fitness is fraction of correct states after running x amount of steps
	return 0
}
