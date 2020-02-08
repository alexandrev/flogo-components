package keypair

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)

}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})


type Activity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{}

	return act, nil
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Aggregates the Message
func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	input := &Input{}
	err = context.GetInputObject(input)


	if err != nil {
		return false, err
	}

	context.Logger().Debugf("Input Keys [%s]", input.Keys)
	context.Logger().Debugf("Input Values [%f]", input.Values)


	result := make([]map[string]interface{}, len(input.Values))
	for i := 0; i < len(input.Values); i++ {
		result[i] = make(map[string]interface{})
		result[i]["operation"] = input.Keys[i].(string)
		result[i]["value"] = input.Values[i].(float64)
	}

	context.Logger().Debugf("Output Values [%s]", result)

	output := &Output{}
	output.Values  =  result;

	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
