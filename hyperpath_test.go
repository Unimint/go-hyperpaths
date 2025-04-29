package hyperpaths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperPaths(t *testing.T) {
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
	ops := FindOptimalStrategy(allLinks, allNodes, destinationNode)
	correctOps := Strategy{
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
		/* In paper it should be this actually */
		// Y3-B
		// Y-Y3
		// X-Y3
		// Y-B
		// X2-Y
		// X-X2
		// A-X2
		// A-B
		ASet: []*Link{
			allLinks[9],
			allLinks[8],
			allLinks[7], // code gives y-b
			allLinks[6], // code gives x-y3
			allLinks[4],
			allLinks[3],
			allLinks[1],
			allLinks[0],
		},
	}
	assert.Equal(t, len(ops.Labels), len(correctOps.Labels), "Incorrect number of labels")
	assert.Equal(t, len(ops.Freqs), len(correctOps.Freqs), "Incorrect number of frequencies")
	assert.Equal(t, len(ops.ASet), len(correctOps.ASet), "Incorrect number of links in attractive set")
	const eps = 1e-20
	for k, v := range ops.Labels {
		assert.Contains(t, correctOps.Labels, k, "Incorrect label key %s has met", k)
		assert.InDelta(t, v, correctOps.Labels[k], eps, "Incorrect label value for node %s", k)
	}
	for k, v := range ops.Freqs {
		assert.Contains(t, correctOps.Freqs, k, "Incorrect frequency key %s has met", k)
		assert.InDelta(t, v, correctOps.Freqs[k], eps, "Incorrect frequency value for node %s", k)
	}
	for i, v := range ops.ASet {
		assert.Equal(t, v, correctOps.ASet[i], "Incorrect link in attractive set at index %d", i)
	}
}
