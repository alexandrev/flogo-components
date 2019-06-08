package aggregate

import (
	"errors"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"sync"
)

// activityLogger is the default logger for the Aggregate Activity
var activityLogger = logger.GetLogger("activity-alexandrev-aggregate")

const (
	ivFunction   = "function"
	ivWindowSize = "windowSize"
	ivValue      = "value"
	ivDataKey    = "key"

	ovResult = "result"
	ovReport = "report"
)

func init() {
	activityLogger.SetLogLevel(logger.InfoLevel)
}

// AggregationActivity is an Activity that is used to Aggregate a message to the console
// inputs : {function, windowSize, autoRest, value}
// outputs: {result, report}
type AggregationActivity struct {
	metadata *activity.Metadata
	mutex    *sync.RWMutex

	// aggregators stateful map of aggregators
	aggregators map[string]Aggregator
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &AggregationActivity{metadata: metadata, aggregators: make(map[string]Aggregator), mutex: &sync.RWMutex{}}
}

// Metadata returns the activity's metadata
func (a *AggregationActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Aggregates the Message
func (a *AggregationActivity) Eval(context activity.Context) (done bool, err error) {

	dataKey := context.GetInput(ivDataKey).(string)

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
			windowSize, _ := context.GetInput(ivWindowSize).(int)
			aggrName, _ := context.GetInput(ivFunction).(string)

			factory := GetFactory(aggrName)

			if factory == nil {
				return false, errors.New("Unknown aggregator: " + aggrName)
			}

			aggr = factory(windowSize)
			a.aggregators[aggregatorKey] = aggr

			activityLogger.Debug("Aggregator created for ", aggregatorKey)
		}

		a.mutex.Unlock()
	}

	inputValues := context.GetInput(ivValue).(*data.ComplexObject).Value.([]interface{})

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

	context.SetOutput(ovReport, report)
	context.SetOutput(ovResult, &data.ComplexObject{Metadata: "", Value: result})
	activityLogger.Debug("Aggregated values returned: %f", result)
	return true, nil
}
