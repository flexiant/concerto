package types

import (
	"encoding/json"
)

type WizardApp struct {
	Id                   string          `json:"id" header:"ID"`
	Name                 string          `json:"name" header:"NAME"`
	Flavour_requirements json.RawMessage `json:"flavour_requirements" header:"FLAVOUR_REQUIREMENTS"`
	Generic_image_id     string          `json:"generic_image_id" header:"GENERIC_IMAGE_ID"`
}
