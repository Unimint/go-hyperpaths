package hyperpaths

// Link is just an edge in graph
type Link struct {
	// Source node of the link
	FromNode string
	// Target node of the link
	ToNode string
	// Corresponding route
	RouteID string
	// In most cases this should trave time along the link
	TravelCost float32
	// Interval of public transport
	// Headway could have only dwell (on-board) links
	Headway float32
}
