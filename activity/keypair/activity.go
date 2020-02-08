package keypair

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}, New)

}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

type Settings struct {
}

type Input struct {
	keys   []interface{} `md:"keys"`
	values []interface{} `md:"values"`
}

type Output struct {
	result []interface{}
}

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

	context.Logger().Debugf("Input Keys [%s]", context.GetInput("keys").([]interface{}))
	context.Logger().Debugf("Input Values [%f]", context.GetInput("values").([]interface{}))

	inputKeys := context.GetInput("keys").([]interface{})
	inputValues := context.GetInput("values").([]interface{})

	result := make([]map[string]interface{}, len(inputValues))
	for i := 0; i < len(inputValues); i++ {
		result[i] = make(map[string]interface{})
		result[i]["operation"] = inputKeys[i].(string)
		result[i]["value"] = inputValues[i].(float64)
	}

	context.SetOutput("values", result)
	context.Logger().Debugf("Output Values [%s]", result)

	return true, nil
}
