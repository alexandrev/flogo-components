package aggregate

import (
	"errors"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"sync"
)

const (
	ivFunction   = "function"
	ivWindowSize = "windowSize"
	ivValue      = "value"
	ivDataKey    = "key"

	ovResult = "result"
	ovReport = "report"
)

func init() {
	_ = activity.Register(&Activity{}, New)

}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// AggregationActivity is an Activity that is used to Aggregate a message to the console
// inputs : {function, windowSize, autoRest, value}
// outputs: {result, report}
type Activity struct {
	metadata *activity.Metadata
	mutex    *sync.RWMutex

	// aggregators stateful map of aggregators
	aggregators map[string]Aggregator
}

// NewActivity creates a new AppActivity
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{aggregators: make(map[string]Aggregator), mutex: &sync.RWMutex{}}

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

	dataKey := input.Key

	aggregatorKey := context.ActivityHost().Name() + ":" + context.Name() + ":" + dataKey


	a.mutex.RLock()
	//get aggregator for activity, assumes flow & task names are unique
	aggr, ok := a.aggregators[aggregatorKey]

	a.mutex.RUnlock()

	//if window not create for this flow, create it

	if !ok {

		//go doesn't have lock upgrades or try, so do same check again

		a.mutex.Lock()
		aggr, ok = a.aggregators[aggregatorKey]

		if !ok {
			windowSize := input.WindowSize

			aggrName := input.Function

			factory := GetFactory(aggrName)

			if factory == nil {
				return false, errors.New("Unknown aggregator: " + aggrName)
			}

			aggr = factory(windowSize)
			a.aggregators[aggregatorKey] = aggr

			context.Logger().Debug("Aggregator created for ", aggregatorKey)
		}

		a.mutex.Unlock()

	}

	inputValues := input.Value


	values := make([]float64, len(inputValues))
	operations := make([]string, len(inputValues))
	for i := 0; i < len(values); i++ {
		inputValuesIndex := inputValues[i].(map[string]interface{})
		operations[i] = inputValuesIndex["operation"].(string)
		values[i], ok = inputValuesIndex["value"].(float64)
		if !ok {
			values[i] = float64(inputValuesIndex["value"].(int))
		}

	}

	report, result := aggr.Add(operations, values)

	output := &Output{}
	output.Result = result
	output.Report = report

	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	context.Logger().Debug("Aggregated values returned: %f", result)
	return true, nil
}
