package aggregate

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	Function string        `md:"function"`
	Value    []interface{} `md:"value"`
	Key      string        `md:"key"`
}

type Output struct {
	Result []interface{} `md:"result"`
	Report bool          `md:"report"`
}

// FromMap converts the values from a map into the struct Input
func (i *Input) FromMap(values map[string]interface{}) error {

	params, err := coerce.ToArray(values["value"])
	if err != nil {
		return err
	}
	i.Value = params

	function, err := coerce.ToString(values["function"])
	if err != nil {
		return err
	}
	i.Function = function

	key, err := coerce.ToString(values["key"])
	if err != nil {
		return err
	}
	i.Key = key

	return nil
}

// ToMap converts the struct Input into a map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"value":    i.Value,
		"key":      i.Key,
		"function": i.Function,
	}
}

// ToMap converts the struct Input into a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
		"report": o.Report,
	}
}

// FromMap converts the values from a map into the struct Input
func (o *Output) FromMap(values map[string]interface{}) error {

	o.Result = values["result"].([]interface{})
	o.Report = values["report"].(bool)
	return nil
}
