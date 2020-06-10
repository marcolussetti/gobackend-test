package rolemodels

// Event is well an event duh
type Event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// AllEvents is an array of events
type AllEvents []Event
