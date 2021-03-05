package aggregate

import (
	"strings"
	"sync"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/data/metadata"
)

const (
	ivFunction = "function"
	ivValue    = "value"
	ivDataKey  = "key"

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
	aggregators map[string]*BlockAverage
}

// NewActivity creates a new AppActivity
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	act := &Activity{aggregators: make(map[string]*BlockAverage), mutex: &sync.RWMutex{}}

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
	output := &Output{}
	keyTsTokens := strings.Split(input.Key, "-")
	keyTs := keyTsTokens[len(keyTsTokens)-1]
	keyTsNum, err := coerce.ToInt64(keyTs)
	tsKey, err := coerce.ToString(((keyTsNum + 150) / 300) * 300)

	dataKey := ""
	for i := 0; i < len(keyTsTokens)-1; i++ {
		dataKey += keyTsTokens[i] + "-"
	}
	dataKey = dataKey + tsKey
	operation := input.Function
	println(operation)
	if operation == "list" {
		for k := range a.aggregators {
			println(k)
			output.Result = append(output.Result, k)
		}
		output.Report = true
		err = context.SetOutputObject(output)
		if err != nil {
			return false, err
		}
		return true, nil
	}

	aggregatorKey := context.ActivityHost().Name() + ":" + context.Name() + ":" + dataKey
	println(aggregatorKey)
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

			aggr = NewBlockAverage()
			a.aggregators[aggregatorKey] = aggr
			context.Logger().Debug("Aggregator created for ", aggregatorKey)
		}

		a.mutex.Unlock()

	}

	inputValues := input.Value
	values := make([]float64, len(inputValues))
	items := make([]string, len(inputValues))
	operations := make([]string, len(inputValues))
	for i := 0; i < len(values); i++ {
		inputValuesIndex := inputValues[i].(map[string]interface{})
		operations[i] = inputValuesIndex["operation"].(string)
		values[i], ok = inputValuesIndex["value"].(float64)
		if !ok {
			values[i] = float64(inputValuesIndex["value"].(int))
		}
		items[i], ok = inputValuesIndex["items"].(string)

	}
	var result []float64
	var report bool
	if operation == "set" {
		report, result = aggr.Add(operations, items, values)
		output.Report = report
	} else if operation == "get" {
		report, result, items = aggr.Get()
		delete(a.aggregators, aggregatorKey)
		output.Report = report
	}

	output.Result = append(output.Result, result)
	output.Result = append(output.Result, items)
	err = context.SetOutputObject(output)
	if err != nil {
		return false, err
	}

	return true, nil
}
