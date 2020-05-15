package fhir_types

import (
	"encoding/json"
	"fmt"
)

func (c *Connection) GetPatient(pid string) (*Paient, error) {
	b, err := c.Get(fmt.Sprintf("Patient/%v", pid))

	if err != nil {
		return nil, err
	}

	data := Patient{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}

	return &data, nil
}

func (c *Connection) SavePatient(pid string, patient Patient) (error) {

}

type Patient struct {
	ResourceType    string `json:"resourceType"`
	BirthDate       string `json:"birthData"`
	Active          bool   `json:"active"`
	Gender          string `json:"gender"`
	DeceasedBoolean bool   `json:"deceasedBoolean"`
}
