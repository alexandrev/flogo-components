package keypair

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLogger is the default logger for the Aggregate Activity
var activityLogger = logger.GetLogger("activity-alexandrev-keypair")

func init() {
	activityLogger.SetLogLevel(logger.InfoLevel)
}

type KeyPair struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &KeyPair{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *KeyPair) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Aggregates the Message
func (a *KeyPair) Eval(context activity.Context) (done bool, err error) {

	activityLogger.Debugf("Input Keys [%s]", context.GetInput("keys").(*data.ComplexObject).Value)
	activityLogger.Debugf("Input Values [%f]", context.GetInput("values").(*data.ComplexObject).Value)

	inputKeys := context.GetInput("keys").(*data.ComplexObject).Value.([]interface{})
	inputValues := context.GetInput("values").(*data.ComplexObject).Value.([]interface{})

	result := make([]map[string]interface{}, len(inputValues))
	for i := 0; i < len(inputValues); i++ {
		result[i] = make(map[string]interface{})
		result[i]["operation"] = inputKeys[i].(string)
		result[i]["value"] = inputValues[i].(float64)
	}

	context.SetOutput("values", &data.ComplexObject{Metadata: "", Value: result})
	activityLogger.Debugf("Output Values [%s]", result)	

	return true, nil
}
