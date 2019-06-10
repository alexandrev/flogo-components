package prometheuspush

import (
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"io/ioutil"
	"testing"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	thisMapArray := make([]interface{}, 1)
	thisMap := make(map[string]interface{})
	thisMap["key"] = "avg_1"
	thisMap["value"] = 3.0
	thisMapArray[0] = thisMap
	//setup attrs

	tc.SetInput("metrics", &data.ComplexObject{Metadata: "", Value: thisMapArray})
	tc.SetInput("gatewayUrl", "http://localhost:9091")
	tc.SetInput("key", "t1")
	tc.SetInput("jobName", "t1")

	act.Eval(tc)

	result := tc.GetOutput("result").(bool)
	if result != true {
		t.Errorf("Result is false")
	}

}
