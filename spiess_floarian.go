package hyperpaths

// ComputeSF computes the Spiess-Florian algorithm
func ComputeSF(allLinks []*Link, allStops map[string]struct{}, destination string, odMatrix map[string]map[string]float32) {
	// Part 1: Find optimal strategy
	ops := FindOptimalStrategy(allLinks, allStops, destination)
	// Part 2: Assign demand according to optimal strategy
	volumes := AssignDemand(allLinks, allStops, ops, odMatrix, destination)
	_ = volumes
	panic("@todo: return final result")
}
