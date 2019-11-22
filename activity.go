package jsonworkflow

import "errors"

// ActivityInterface describes main activity methods
type ActivityInterface interface {
	GetID() string
	GetClass() ActivityClass
	Execute(ctx Context) error
	Configure(metadata Metadata)
}

// AbstractActivity implements default methods for activity
type AbstractActivity struct {
	ID       string        `json:"ID"`
	Class    ActivityClass `json:"class"`
	Metadata Metadata      `json:"metadata"`
}

// Metadata represents general metadata structure
type Metadata struct {
	Bit       DeviceBit   `json:"bit"`
	Event     EventType   `json:"event"`
	Direction Direction   `json:"direction"`
	Key       string      `json:"key"`
	Value     interface{} `json:"value"`
}

// GetID returns activity unique id
func (r *AbstractActivity) GetID() string {
	return r.ID
}

// SetID sets activity unique id
func (r *AbstractActivity) SetID(id string) {
	r.ID = id
}

// GetClass return activity class
func (r *AbstractActivity) GetClass() ActivityClass {
	return r.Class
}

// SetClass set activity class
func (r *AbstractActivity) SetClass(class ActivityClass) {
	r.Class = class
}

// Execute function implements activity run logic
func (r *AbstractActivity) Execute(ctx Context) error {
	return nil
}

// Configure function implements configuring activity when its created
func (r *AbstractActivity) Configure(metadata Metadata) {}

// Assign activity
type Assign struct {
	*AbstractActivity
}

// Configure Assign activity
func (r *Assign) Configure(metadata Metadata) {

}

// Execute Assign activity
func (r *Assign) Execute(ctx Context) error {
	return errors.New("nothing to execute")
}
