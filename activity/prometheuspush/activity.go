package prometheuspush

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"gopkg.in/resty.v1"
)

// activityLogger is the default logger for the Aggregate Activity
var activityLogger = logger.GetLogger("activity-alexandrev-prometheuspush")

func init() {
	activityLogger.SetLogLevel(logger.InfoLevel)
}

type PrometheusPush struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &PrometheusPush{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *PrometheusPush) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Aggregates the Message
func (a *PrometheusPush) Eval(context activity.Context) (done bool, err error) {
	key := context.GetInput("key").(string)
	jobName := context.GetInput("jobName").(string)
	inputMetrics := context.GetInput("metrics").(*data.ComplexObject).Value.([]interface{})
	activityLogger.Debugf("Input Metrics [%s]", inputMetrics)
	gatewayURL := context.GetInput("gatewayUrl").(string) //"http://my-pushgateway-vm1:9091/"
	activityLogger.Debugf("Gateway URL [%s]", gatewayURL)

	activityLogger.Debugf("Key [%s]", key)

	bodyText := ""
	for i := 0; i < len(inputMetrics); i++ {

		metric := inputMetrics[i].(map[string]interface{})
		metricName := metric["key"].(string)
		metricValue, ok := metric["value"].(float64)
		if !ok {
			metricValue = float64(metric["value"].(int))
		}
		activityLogger.Debugf("Metric [%s : %f] ", metricName, metricValue)

		metricValueStr := fmt.Sprintf("%f", metricValue)
		bodyText += "# TYPE " + metricName + " gauge\n" + "# HELP " + metricName + " " + metricName + ".\n" + metricName + metricValueStr + "\n"

	}

	resp, err := resty.R().SetBody(bodyText).Put(gatewayURL + "/metrics/job/" + jobName)
	fmt.Printf("%s", resp)
	fmt.Printf("%s", err)
	if err == nil {
		activityLogger.Debugf("Pushed metric and response returned [%s] ", resp)
		activityLogger.Debugf("Pushed metric to [%s] ", jobName)
	}

	context.SetOutput("result", true)

	return true, nil
}
