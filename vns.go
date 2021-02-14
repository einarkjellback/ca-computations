package cacomp

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

// Run performs a general variable neighborhood search on the provided cellular automaton. It stops after iterating for maxIter.
func Run(ca *CA, kmax int, maxIter int) (RuleTable, error) {
	// Generate random solution

	// run GVNS (See pseudocode above)
	return nil, nil
}

// shake picks a random solution from the k-neighborhood.
func shake(ca *CA, k int) *CA {
	// Flip k randomly selected bits and return result
	return nil
}

// vnd finds the fittest candidate from the k=1 to k=kMax neighborhoods
func vnd(ca *CA, kMax int) *CA {
  curr := ca.rules
	for k := 1; k <= kMax; {
		// Find fittest candidate from k-neighborhood
    curr,
	}
	// Return fittest candidate found
	return nil
}

func neighborhoodChange(curr Rules, next Rules, k int) {
  if
}
