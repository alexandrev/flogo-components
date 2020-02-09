package aggregate

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"testing"
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

func TestEval(t *testing.T) {
	
	iCtx := test.NewActivityInitContext(nil, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())
	
	thisMap := make([]interface{}, 1)
	tmpThisMap := make(map[string]interface{})
	tmpThisMap["operation"] = "avg"
	tmpThisMap["value"] = 2
	thisMap[0] = tmpThisMap
	//setup attrs
	tc.SetInput(ivFunction, "moving")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, thisMap)
	tc.SetInput(ivDataKey, "1")

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult).([]float64)

	if result[0] != 0.0 {
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
	tmpThisMap2["value"] = 3
	thisMap2[0] = tmpThisMap2
	//setup attrs
	tc2.SetInput(ivFunction, "moving")
	tc2.SetInput(ivWindowSize, 5)
	tc2.SetInput(ivValue, thisMap2)
	tc2.SetInput(ivDataKey, "1")

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult).([]float64)

	if result[0] != 2.5 {
		t.Errorf("Result is %f instead of 2.5", result[0])
	}

	if !report {
		t.Error("Window should report after second value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewActivityContext(act.Metadata())
	thisMap3 := make([]interface{}, 1)
	tmpThisMap3 := make(map[string]interface{})
	tmpThisMap3["operation"] = "avg"
	tmpThisMap3["value"] = 3
	thisMap3[0] = tmpThisMap3
	//setup attrs
	tc3.SetInput(ivFunction, "moving")
	tc3.SetInput(ivWindowSize, 5)
	tc3.SetInput(ivValue, thisMap3)
	tc3.SetInput(ivDataKey, "1")

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult).([]float64)

	if result[0] != 3.0 {
		t.Errorf("Result is %f instead of 3.0", result[0])
	}

	if !report {
		t.Error("Window should report after third value")
	}

	fmt.Printf("test %v, %v", report, result)

}

func TestResetEval(t *testing.T) {

	iCtx := test.NewActivityInitContext(nil, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())
	
	thisMap := make([]interface{}, 1)
	tmpThisMap := make(map[string]interface{})
	tmpThisMap["operation"] = "avg"
	tmpThisMap["value"] = 2
	thisMap[0] = tmpThisMap
	//setup attrs

	tc.SetInput(ivFunction, "block")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, thisMap)

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult).([]float64)

	if result[0] != 0.0 {
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
	tmpThisMap2["value"] = 3
	thisMap2[0] = tmpThisMap2
	//setup attrs
	tc2.SetInput(ivFunction, "block")
	tc2.SetInput(ivWindowSize, 2)
	tc2.SetInput(ivValue, thisMap2)

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult).([]float64)

	if result[0] != 2.5 {
		t.Errorf("Result is %f instead of 2.5", result[0])
	}

	if !report {
		t.Error("Window should report after second value")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewActivityContext(act.Metadata())
	thisMap3 := make([]interface{}, 1)
	tmpThisMap3 := make(map[string]interface{})
	tmpThisMap3["operation"] = "avg"
	tmpThisMap3["value"] = 3
	thisMap3[0] = tmpThisMap3
	//setup attrs
	tc3.SetInput(ivFunction, "block")
	tc3.SetInput(ivWindowSize, 2)
	tc3.SetInput(ivValue, thisMap3)

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult).([]float64)

	if report {
		t.Error("Window should not report after third value")
	}

	if result[0] != 0.0 {
		t.Errorf("Result is %f instead of 0.0", result[0])
	}

	fmt.Printf("test %v, %v", report, result)

}

func TestVaryingData(t *testing.T) {

	iCtx := test.NewActivityInitContext(nil, nil)
	act, _ := New(iCtx)
	tc := test.NewActivityContext(act.Metadata())

	//setup attrs

	thisMap := make([]interface{}, 1)
	tmpThisMap := make(map[string]interface{})
	tmpThisMap["operation"] = "avg"
	tmpThisMap["value"] = 2
	thisMap[0] = tmpThisMap
	tc.SetInput(ivFunction, "block")
	tc.SetInput(ivWindowSize, 2)
	tc.SetInput(ivValue, thisMap)
	tc.SetInput(ivDataKey, "1")

	act.Eval(tc)

	report := tc.GetOutput(ovReport).(bool)
	result := tc.GetOutput(ovResult).([]float64)

	if result[0] != 0.0 {
		t.Errorf("Result is %f instead of 0", result[0])
	}
	if report {
		t.Error("Window should not report after first value for key 1")
	}

	fmt.Printf("test %v, %v", report, result)

	tca := test.NewActivityContext(act.Metadata())
	thisMapa := make([]interface{}, 1)
	tmpThisMapa := make(map[string]interface{})
	tmpThisMapa["operation"] = "avg"
	tmpThisMapa["value"] = 2
	thisMapa[0] = tmpThisMapa
	//setup attrs
	tca.SetInput(ivFunction, "block")
	tca.SetInput(ivWindowSize, 2)
	tca.SetInput(ivValue, thisMapa)
	tca.SetInput(ivDataKey, "2")

	act.Eval(tca)

	report = tca.GetOutput(ovReport).(bool)
	result = tca.GetOutput(ovResult).([]float64)

	if result[0] != 0.0 {
		t.Errorf("Result is %f instead of 0", result[0])
	}
	if report {
		t.Error("Window should not report after first value for key 2")
	}

	fmt.Printf("test %v, %v", report, result)

	tc2 := test.NewActivityContext(act.Metadata())
	thisMap2 := make([]interface{}, 1)
	tmpThisMap2 := make(map[string]interface{})
	tmpThisMap2["operation"] = "avg"
	tmpThisMap2["value"] = 3
	thisMap2[0] = tmpThisMap2
	//setup attrs
	tc2.SetInput(ivFunction, "block")
	tc2.SetInput(ivWindowSize, 2)
	tc2.SetInput(ivValue, thisMap2)
	tc2.SetInput(ivDataKey, "1")

	act.Eval(tc2)

	report = tc2.GetOutput(ovReport).(bool)
	result = tc2.GetOutput(ovResult).([]float64)

	if result[0] != 2.5 {
		t.Errorf("Result is %f instead of 2.5", result[0])
	}

	if !report {
		t.Error("Window should report after second value for key 1")
	}

	fmt.Printf("test %v, %v", report, result)

	tc3 := test.NewActivityContext(act.Metadata())
	thisMap3 := make([]interface{}, 1)
	tmpThisMap3 := make(map[string]interface{})
	tmpThisMap3["operation"] = "avg"
	tmpThisMap3["value"] = 3
	thisMap3[0] = tmpThisMap3
	//setup attrs
	tc3.SetInput(ivFunction, "block")
	tc3.SetInput(ivWindowSize, 2)
	tc3.SetInput(ivValue, thisMap3)
	tc3.SetInput(ivDataKey, "1")

	act.Eval(tc3)

	report = tc3.GetOutput(ovReport).(bool)
	result = tc3.GetOutput(ovResult).([]float64)

	if report {
		t.Error("Window should not report after third value for key 1")
	}

	if result[0] != 0.0 {
		t.Errorf("Result is %f instead of 0.0", result[0])
	}

	fmt.Printf("test %v, %v", report, result)

	tca2 := test.NewActivityContext(act.Metadata())
	thisMapa2 := make([]interface{}, 1)
	tmpThisMapa2 := make(map[string]interface{})
	tmpThisMapa2["operation"] = "avg"
	tmpThisMapa2["value"] = 3
	thisMapa2[0] = tmpThisMapa2
	//setup attrs
	tca2.SetInput(ivFunction, "block")
	tca2.SetInput(ivWindowSize, 2)
	tca2.SetInput(ivValue, thisMapa2)
	tca2.SetInput(ivDataKey, "2")

	act.Eval(tca2)

	report = tca2.GetOutput(ovReport).(bool)
	result = tca2.GetOutput(ovResult).([]float64)

	if result[0] != 2.5 {
		t.Errorf("Result is %f instead of 2.5", result[0])
	}

	if !report {
		t.Error("Window should report after second value for key 2")
	}

	fmt.Printf("test %v, %v", report, result)

}
