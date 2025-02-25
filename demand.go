package hyperpaths

import (
	"fmt"
	"sort"
)

// Volumes is assignment demand according to optimal strategy
type Volumes struct {
	// Volumes for links
	Links map[string]map[string]float32
	// Volumes for nodes
	Nodes map[string]float32
}

func AssignDemand(allLinks []*Link, allStops map[string]struct{}, optimalStrategy *Strategy, trips map[string]map[string]float32, destination string) *Volumes {
	// Sort attractive links set by u{j} + c{a} in descending order
	sort.Slice(optimalStrategy.ASet, func(i, j int) bool {
		a := optimalStrategy.ASet[i]
		b := optimalStrategy.ASet[j]
		return optimalStrategy.Labels[a.ToNode]+a.TravelCost >= optimalStrategy.Labels[b.ToNode]+b.TravelCost
	})
	nodeVolumes := make(map[string]float32, len(allStops))
	for i := range allStops {
		nodeVolumes[i] = 0
	}
	for origin := range trips {
		if tripsNum, ok := trips[origin][destination]; ok {
			nodeVolumes[origin] = tripsNum
			nodeVolumes[destination] += tripsNum
		}
	}
	nodeVolumes[destination] *= -1

	v := make(map[string]map[string]float32)
	for _, a := range allLinks {
		if _, ok := v[a.FromNode]; !ok {
			v[a.FromNode] = make(map[string]float32)
		}
		v[a.FromNode][a.ToNode] = 0.0
	}

	for _, a := range optimalStrategy.ASet {
		// Calculate frequency (1/headway)
		freq := infiniteFrequency
		if a.Headway > 0 {
			freq = 1 / a.Headway
		}
		va := (freq / optimalStrategy.Freqs[a.FromNode]) * nodeVolumes[a.FromNode]
		if Verbose {
			fmt.Printf("Assigning demand for link: (%s, %s) \\\\ \n", a.FromNode, a.ToNode)
			fmt.Printf("\\quad $v_{(%s, %s)} = \\frac{%v}{%v}%v = %v$ \\\\ \n", a.FromNode, a.ToNode, freq, optimalStrategy.Freqs[a.FromNode], nodeVolumes[a.FromNode], va)
			fmt.Printf("\\quad $V_{%s} = V_{%s} + v_{(%s, %s) = %v + %v = %v}$ \\\\ \n", a.ToNode, a.ToNode, a.FromNode, a.ToNode, nodeVolumes[a.ToNode], va, nodeVolumes[a.ToNode]+va)
		}
		v[a.FromNode][a.ToNode] = va
		nodeVolumes[a.ToNode] += va
	}
	if Verbose {
		fmt.Println("Final node volumes: \\\\")
		for k := range nodeVolumes {
			fmt.Printf("\\quad $V_{%s} = %v$ \\\\ \n", k, nodeVolumes[k])
		}
	}

	result := &Volumes{
		Links: v,
		Nodes: nodeVolumes,
	}
	return result
}
