package vns

import (
	"container/heap"
	"errors"
	"fmt"
	"github.com/einarkjellback/cacomp/ca"
	"log"
	"math/rand"
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

// Controls the number of different configurations that a rule will be tested against.
const SIMS = 20

// Controls the size of the configurations. Must be an odd number.
const CONFIG_SIZE = 51

// Controls the size of the neighborhood.
const RADIUS = 2

// Controls number of generations that each CA will run.
var ITERS int

// The number of bits required to represent a rule with radius r.
var RULE_WIDTH int

func init() {
	ITERS = 2 * CONFIG_SIZE

	temp, err := pow(2, 2*RADIUS+1)
	if err != nil {
		log.Fatal(err)
	}
	RULE_WIDTH = int(temp)
}

/*
Vns contains various details of a variable neighborhood (VNS) run.
*/
type Vns struct {
	// Rules sorted by fitness, highest to lowest.
	Rules *RuleHeap
	// Maps a rule to its fitness value.
	RuleFits map[uint32]float64
	// Complete history of density classification tasks performed.
	RuleConfigs map[uint32][][][]bool
}

type RuleHeap []struct {
	r   uint32
	fit float64
}

// Sort interface.
func (h RuleHeap) Len() int           { return len(h) }
func (h RuleHeap) Less(i, j int) bool { return h[i].fit > h[j].fit }
func (h RuleHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Heap interface.
func (h *RuleHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *RuleHeap) Push(x interface{}) {
	*h = append(*h, x.(struct {
		r   uint32
		fit float64
	}))
}

/* Run performs a general variable neighborhood search on the density classification problem. It stops after iterating for maxIter. kmax is the maximum neighborhood size for the shake function, while lmax is the maximum neighborhood size for the variable neighborhood descent function.
 */
func Run(kmax int, lmax int, maxIter int) (*Vns, error) {
	vns := &Vns{new(RuleHeap), nil, nil}
	heap.Init(vns.Rules)
	// Generate random rule
	r := rand.Uint32()
	heap.Push(vns.Rules, struct {
		r   uint32
		fit float64
	}{r, vns.fitness(r)})

	// run GVNS (See pseudocode above).
	for i := 0; i < maxIter; i++ {
		for k := 1; k <= kmax; {
			newR, err := vns.shake(r, k)
			if err != nil {
				log.Fatal(err)
			}
			newR = vns.vnd(newR, lmax)
			r, k = vns.neighborhoodChange(r, newR, k)
		}
	}

	return vns, nil
}

// shake picks a random, untried solution from the k-neighborhood.
func (vns *Vns) shake(r uint32, k int) (uint32, error) {
	// Flip k randomly selected bits and returns result.
	if k <= 0 || k > RULE_WIDTH {
		return 0, errors.New(fmt.Sprintf("k = %v not in interval [0, %v]", k, RULE_WIDTH))
	}

	pos := make([]int, k)
	for i := range pos {
		pos[i] = int(rand.Uint32()) // Duplicates are possible, but highly unlikely.
	}
	return flipN(r, pos), nil
}

// vnd finds the fittest candidate from the k=1 to k=kMax neighborhoods.
func (vns *Vns) vnd(r uint32, kMax int) uint32 {
	for k := 1; k <= kMax; {
		next := vns.findFittest(r, k)
		r, k = vns.neighborhoodChange(r, next, k)
	}
	// Return fittest candidate found.
	return r
}

// neighborhoodChange returns the rule with highest fitness and updates k accordingly.
func (vns *Vns) neighborhoodChange(curr uint32, next uint32, k int) (uint32, int) {
	currFit := vns.fitness(curr)
	nextFit := vns.fitness(next)
	if nextFit > currFit {
		k = 1
		curr = next
	} else {
		k += 1
	}
	return curr, k
}

// findFittest finds the fittest candidate from the k-neighborhood.
func (vns *Vns) findFittest(r uint32, k int) uint32 {
	hood, err := getNeighborhood(r, k)
	if err != nil {
		log.Fatal(err)
	}
	for _, n := range hood {
		if vns.fitness(r) < vns.fitness(n) {
			r = n
		}
	}
	return r
}

// fitness fetches the fitness of rule r from Vns.RuleFits if present. If fitness of r not calculated, then calculates it and stores it to the map.
func (vns *Vns) fitness(r uint32) float64 {
	fit, ok := vns.RuleFits[r]
	if !ok {
		fit = vns.calcFitness(r)
		vns.RuleFits[r] = fit
	}
	// Fitness is the fraction of correct states after running x amount of steps.
	return fit
}

// CountAlive counts the number of cells in the alive state in the configuration provided.
func CountAlive(config []bool) int {
	acc := 0
	for _, alive := range config {
		if alive {
			acc += 1
		}
	}
	return acc
}

// calcFitness calculates the fitness of r and stores it in Vns.RuleFits.
func (vns *Vns) calcFitness(r uint32) float64 {
	fit := 0.0
	for i := 0; i < SIMS; i++ {
		config, err := ca.RandConfig(CONFIG_SIZE)
		if err != nil {
			log.Fatal(err)
		}
		moreThanHalfAliveOld := CountAlive(config)*2 > CONFIG_SIZE
		sim, err := ca.UpdateN(config, RADIUS, r, ITERS)
		if err != nil {
			log.Fatal(err)
		}
		vns.RuleConfigs[r] = append(vns.RuleConfigs[r], sim)
		config = sim[len(sim)-1]
		aliveCnt := CountAlive(config)
		currFit := float64(aliveCnt) / float64(CONFIG_SIZE)
		if !moreThanHalfAliveOld {
			currFit = 1.0 - currFit
		}
		fit += currFit
	}
	return fit / SIMS
}

// getNeighborhood returns every rule that can be generated by flipping k bits of r.
func getNeighborhood(r uint32, k int) ([]uint32, error) {
	if k > 13 {
		return nil, errors.New(fmt.Sprintf("want k <= 13 due to too many permutations above 13 flips, but got k = %v", k))
	}

	totPerms, err := pow(32, uint(k))
	if err != nil {
		log.Fatal(err)
	}
	allRules := make([]uint32, totPerms)
	genAllRules(r, k, 0, allRules)
	return allRules, nil
}

// Recursive support function for getNeighborhood function.
func genAllRules(r uint32, k int, start int, acc []uint32) {
	if k == 0 {
		acc = append(acc, r)
		return
	}
	for i := start; i <= 32-k; i++ {
		flip(r, i)
		genAllRules(r, k, i+1, acc)
	}
}

// flipN flips the bits in r at the positions contained in pos.
func flipN(r uint32, pos []int) uint32 {
	var err error
	for _, p := range pos {
		r, err = flip(r, p)
		if err != nil {
			log.Fatal(err)
		}
	}
	return r
}

// flip flips the bit at position pos in r.
func flip(r uint32, pos int) (uint32, error) {
	if pos < 0 || pos >= RULE_WIDTH {
		return 0, errors.New(fmt.Sprintf("flip at position %v outside interval [0, %v]", pos, RULE_WIDTH-1))
	}

	flipper := uint32(1) << pos
	if r&flipper == 0 {
		r |= flipper
	} else {
		flipper = ^flipper
		r &= flipper
	}
	return r, nil
}

// pow calculates n**m. Equivalent to math.Pow, but can handle higher numbers at the cost of only working for positive integers.
func pow(n, m uint) (uint64, error) {
	if n == 0 && m == 0 {
		return 0, errors.New("pow(0, 0) is undefined")
	}
	acc := uint64(1)
	p := uint64(n)
	for i := uint(0); i < m; i++ {
		acc *= p
	}
	return acc, nil
}
