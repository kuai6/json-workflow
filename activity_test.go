package jsonworkflow

import (
	"encoding/json"
	"testing"
)

func TestRoot(t *testing.T) {
	test := `{"ID":"292077b1-8e52-4ac2-a269-b8481daa7701", "class":"assign", "metadata": {"event":"ContactClosure", "direction":0, "bit":1}}`

	r := &Assign{}

	err := json.Unmarshal([]byte(test), r)
	if err != nil {
		t.Errorf("cant unmarshall json: %s", test)
	}

	if r.GetID() != "292077b1-8e52-4ac2-a269-b8481daa7701" {
		t.Errorf("getID wrong result: %s", r.GetID())
	}

	if r.GetClass() != ActivityClassAssign {
		t.Errorf("GetClass wrong result: %s", r.GetClass())
	}

	if r.Metadata.Event != EventContactClosure {
		t.Errorf("Metadata.Event wrong result: %s", r.Metadata.Event)
	}

	if r.Metadata.Direction != DirectionOnFall {
		t.Errorf("Metadata.Direction wrong result: %d", r.Metadata.Direction)
	}

	if r.Metadata.Bit != BIT1 {
		t.Errorf("Metadata.Bit wrong result: %d", r.Metadata.Bit)
	}

}
