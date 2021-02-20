package main

import (
	"fmt"
	// "github.com/einarkjellback/cacomp/ca"
	"github.com/einarkjellback/cacomp/vns"
	"log"
)

func main() {
	// runSim()
	m := make(map[int]int)
	_, ok := m[0]
	fmt.Println(ok)
}

func runSim() {
	// Run VNS
	kmax := 3
	iters := 10
	rec, err := vns.Run(kmax, kmax-1, iters)
	if err != nil {
		log.Fatal(err)
	}

	// Display results:
	// Show fitness convergence graph
	// Show 2-dimensional history
	display(rec)
}

func display(rec *vns.Vns) {}
