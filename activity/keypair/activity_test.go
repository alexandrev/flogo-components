package keypair

import (
	"fmt"
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

	thisMap := make(map[string]interface{})
	thisMap["operation"] = "avg"
	thisMap["value"] = 2
	//setup attrs

	tc.SetInput("keys", &data.ComplexObject{Metadata: "", Value: []interface{}{"a", "b", "c"}})
	tc.SetInput("values", &data.ComplexObject{Metadata: "", Value: []interface{}{0.01, 0.02, 0.03}})

	act.Eval(tc)

	result := tc.GetOutput("values").(*data.ComplexObject).Value.([]map[string]interface{})
	fmt.Printf("%s", result)
	if result[0]["operation"] != "a" {
		t.Errorf("Result is %s instead of a", result[0]["operacion"])
	}
	if result[0]["value"] != 0.01 {
		t.Errorf("Result is %f instead of 0.01", result[0]["value"])
	}

	if result[1]["operation"] != "b" {
		t.Errorf("Result is %s instead of b", result[1]["operacion"])
	}
	if result[1]["value"] != 0.02 {
		t.Errorf("Result is %f instead of 0.02", result[1]["value"])
	}

	if result[2]["operation"] != "c" {
		t.Errorf("Result is %s instead of b", result[2]["operacion"])
	}
	if result[2]["value"] != 0.03 {
		t.Errorf("Result is %f instead of 0.03", result[2]["value"])
	}

}
