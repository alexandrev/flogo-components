package keypair

import (
	"fmt"
	"github.com/project-flogo/core/data/coerce"
)

type Settings struct {
}

type Input struct {
	Keys   []interface{} `md:"keys"`
	Values []interface{} `md:"values"`
}

type Output struct {
	Values []map[string]interface{} `md:"values"`
}

// ToMap converts the struct Input into a map
func (i *Input) ToMap() map[string]interface{} {
	fmt.Printf("A02 %s", i.Keys)
	return map[string]interface{}{
		"keys":   i.Keys,
		"values": i.Values,
	}
}

// FromMap converts the values from a map into the struct Input
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	fmt.Printf("A1 %s", values["keys"])
	i.Keys, err = coerce.ToArray(values["keys"])
	fmt.Printf("A2 %s", i.Keys)
	if err != nil {
		return err
	}
	i.Values, err = coerce.ToArray(values["values"])
	if err != nil {
		return err
	}
	return nil
}

// ToMap converts the struct Input into a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"values": o.Values,
	}
}

// FromMap converts the values from a map into the struct Input
func (o *Output) FromMap(values map[string]interface{}) error {

	o.Values = values["values"].([]map[string]interface{})
	return nil
}
