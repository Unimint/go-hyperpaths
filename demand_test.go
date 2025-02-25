package hyperpaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignDemand(t *testing.T) {
	// Test right from the paper
	allNodes := map[string]struct{}{
		"A": {},
		"X": {}, "X2": {},
		"Y": {}, "Y3": {},
		"B": {},
	}
	allLinks := []*Link{
		{"A", "B", "Line 1", 25, 6},
		{"A", "X2", "Line 2", 7, 6},
		{"X2", "X", "Line 2", 0, 0},
		{"X", "X2", "Line 2", 0, 6},
		{"X2", "Y", "Line 2", 6, 0},
		{"Y3", "Y", "Line 3", 0, 15},
		{"Y", "B", "Line 4", 10, 3},
		{"X", "Y3", "Line 3", 4, 15},
		{"Y", "Y3", "Line 3", 0, 15},
		{"Y3", "B", "Line 3", 4, 0},
	}
	destinationNode := "B"
	odMatrix := map[string]map[string]float32{
		"A": {
			"B": 1,
		},
	}
	optimalStrategy := Strategy{
		Labels: map[string]float32{
			"A":  27.75,
			"X":  19.071426,
			"X2": 17.5,
			"Y":  11.500001,
			"Y3": 4,
			"B":  0,
		},
		Freqs: map[string]float32{
			"A":  0.33333334,
			"X":  0.23333335,
			"X2": 1e+11,
			"Y":  0.4,
			"Y3": 1e+11,
			"B":  0,
		},
		ASet: []*Link{
			allLinks[9],
			allLinks[6],
			allLinks[8],
			allLinks[7],
			allLinks[4],
			allLinks[1],
			allLinks[0],
			allLinks[3],
		},
	}
	volumes := AssignDemand(allLinks, allNodes, &optimalStrategy, odMatrix, destinationNode)
	correctVolumes := Volumes{
		Links: map[string]map[string]float32{
			"A": {
				"B":  0.5,
				"X2": 0.5,
			},
			"X2": {
				"X": 0.0,
				"Y": 0.5,
			},
			"X": {
				"X2": 0.0,
				"Y3": 0.0,
			},
			"Y": {
				"Y3": 0.083333336,
				"B":  0.4166667,
			},
			"Y3": {
				"Y": 0.0,
				"B": 0.083333336,
			},
		},
		Nodes: map[string]float32{
			"A":  1.0,
			"X2": 0.5,
			"X":  0.0,
			"Y3": 0.083333336,
			"Y":  0.5,
			"B":  0.0,
		},
	}
	assert.Equal(t, len(volumes.Links), len(correctVolumes.Links), "Incorrect number of links in volumes data")
	assert.Equal(t, len(volumes.Nodes), len(correctVolumes.Nodes), "Incorrect number of nodes in volumes data")
	const eps = 1e-6
	for fromNode := range volumes.Links {
		assert.Contains(t, correctVolumes.Links, fromNode, "No 'FromNode' in correct volumes data")
		for toNode, volume := range volumes.Links[fromNode] {
			assert.Contains(t, correctVolumes.Links[fromNode], toNode, "No 'ToNode' in correct volumes data")
			assert.InDelta(t, volume, correctVolumes.Links[fromNode][toNode], eps, "Incorrect volume in link (%s, %s)", fromNode, toNode)
		}
	}
	for i, nodeVolume := range volumes.Nodes {
		assert.InDelta(t, nodeVolume, correctVolumes.Nodes[i], eps, "Incorrect volume in node %d", i)
	}
}
