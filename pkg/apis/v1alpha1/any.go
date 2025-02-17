package v1alpha1

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

// Any can be any type.
// +k8s:deepcopy-gen=false
type Any struct {
	// Value contains the value of the Any object.
	// +kubebuilder:pruning:PreserveUnknownFields
	// +kubebuilder:validation:Schemaless
	// +optional
	Value interface{} `json:",inline"`
}

func (in *Any) DeepCopyInto(out *Any) {
	if err := copier.Copy(out, in); err != nil {
		panic("deep copy failed")
	}
}

func (a *Any) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.Value)
}

func (a *Any) UnmarshalJSON(data []byte) error {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	a.Value = v
	return nil
}
