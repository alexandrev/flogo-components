package aggregate

import (
	"fmt"
	"testing"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/test"
)

func TestCreate(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestResetEval(t *testing.T) {

	id, _ := coerce.ToString(time.Now().Unix())
	iCtx := test.NewActivityInitContext(nil, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())

	thisMap := make([]interface{}, 1)
	tmpThisMap := make(map[string]interface{})
	tmpThisMap["operation"] = "avg"
	tmpThisMap["value"] = 2
	thisMap[0] = tmpThisMap
	//setup attrs

	tc.SetInput(ivFunction, "set")
	tc.SetInput(ivValue, thisMap)
	tc.SetInput("key", "t1-"+id)

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult).([]interface{})

	if result[0].([]float64)[0] != 0.0 {
		t.Errorf("Result is %f instead of 0", result[0])
	}
	if report {
		t.Error("Window should not report after first value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc2 := test.NewActivityContext(act.Metadata())
	thisMap2 := make([]interface{}, 1)
	tmpThisMap2 := make(map[string]interface{})
	tmpThisMap2["operation"] = "avg"
	tmpThisMap2["key"] = "t1-20202000"
	tmpThisMap2["value"] = 2
	thisMap2[0] = tmpThisMap2
	//setup attrs
	tc2.SetInput(ivFunction, "get")
	tc2.SetInput(ivValue, thisMap2)
	tc2.SetInput("key", "t1-"+id)

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult).([]interface{})

	if result[0].([]float64)[0] != 2 {
		t.Errorf("Result is %f instead of 2", result[0])
	}

	if !report {
		t.Error("Window should report after second value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewActivityContext(act.Metadata())
	thisMap3 := make([]interface{}, 1)
	tmpThisMap3 := make(map[string]interface{})
	tmpThisMap3["operation"] = "avg"
	tmpThisMap3["key"] = "t1-20202000"
	tmpThisMap3["value"] = 3
	thisMap3[0] = tmpThisMap3
	//setup attrs
	tc3.SetInput(ivFunction, "list")
	tc3.SetInput(ivValue, thisMap3)
	tc3.SetInput("key", "t1-"+id)

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult).([]interface{})

	if result[0].(string) != "0.0" {
		t.Errorf("Result is %f instead of 0.0", result[0])
	}

	fmt.Printf("test %v, %v", report, result)

}
