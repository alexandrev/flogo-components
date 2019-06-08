package metrics

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
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

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs

	act.Eval(tc)
	//check result attr
	output := tc.GetOutput("output").(*data.ComplexObject).Value.([]map[string]interface{})
	fmt.Println("memTotal: ", output)

	fmt.Println("memTotal: ", output[0]["value"])
	fmt.Println("memUsed: ", output[1]["value"])
	fmt.Println("memPercentage: ", output[2]["value"])
	fmt.Println("cpuNumber: ", output[3]["value"])
	fmt.Println("cpuSpeed: ", output[4]["value"])
	fmt.Println("cpuUsed: ", output[5]["value"])
	fmt.Println("procRunning: ", output[6]["value"])

}
