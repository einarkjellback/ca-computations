package cacomp

import "log"

func main() {
	// Run VNS
	kmax := 3
	iters := 10
	rs, fits, cs, err := Run(kmax, kmax-1, iters)
	if err != nil {
		log.Fatal(err)
	}

	// Display results:
	// Show fitness convergence graph
	// Show 2-dimensional history
	display(rs, fits, cs)
}

func display(rs []Rules, fits map[Rules]float64, cs map[Rules][]Config) {}
