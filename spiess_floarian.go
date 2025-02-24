package hyperpaths

// ComputeSF computes the Spiess-Florian algorithm
func ComputeSF(allLinks []*Link, allStops map[string]struct{}, destination string) {
	// Part 1: Find optimal strategy
	ops := FindOptimalStrategy(allLinks, allStops, destination)
	_ = ops
	// Part 2: Assign demand according to optimal strategy
	panic("@todo: part 2")
}
