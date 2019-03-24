package command

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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
	memTotal := tc.GetOutput("memTotal")
	memUsed := tc.GetOutput("memUsed")
	memPercentage := tc.GetOutput("memPercentage")
	cpuNumber := tc.GetOutput("cpuNumber")
	cpuSpeed := tc.GetOutput("cpuSpeed")
	cpuUsed := tc.GetOutput("cpuUsed")
	procRunning := tc.GetOutput("procRunning")

	fmt.Println("memTotal: ", memTotal)
	fmt.Println("memUsed: ", memUsed)
	fmt.Println("memPercentage: ", memPercentage)
	fmt.Println("cpuNumber: ", cpuNumber)
	fmt.Println("cpuSpeed: ", cpuSpeed)
	fmt.Println("cpuUsed: ", cpuUsed)
	fmt.Println("procRunning: ", procRunning)

}
