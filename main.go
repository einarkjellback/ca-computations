package main

import (
	"fmt"
	"github.com/einarkjellback/cacomp/ca"
	"github.com/einarkjellback/cacomp/vns"
	"log"
	"reflect"
)

func main() {
	// runSim()
	fmt.Println(reflect.TypeOf([]int{1, 2, 3, 4}[1:3]))
}

func runSim() {
	// Run VNS
	kmax := 3
	iters := 10
	rs, fits, cs, err := vns.Run(kmax, kmax-1, iters)
	if err != nil {
		log.Fatal(err)
	}

	// Display results:
	// Show fitness convergence graph
	// Show 2-dimensional history
	display(rs, fits, cs)
}

func display(rs []ca.Rules, fits map[ca.Rules]float64, cs map[ca.Rules][]ca.Config) {}
