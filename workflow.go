package jsonworkflow

// Workflow ...
type Workflow struct {
	Nodes []ActivityInterface
	Links []Link
}

// Link ...
type Link struct {
	From string
	To   string
}
