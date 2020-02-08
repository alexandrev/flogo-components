package keypair

import (
	"fmt"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

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

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)
	tc := test.NewActivityContext(act.Metadata())

	thisMap := make(map[string]interface{})
	thisMap["operation"] = "avg"
	thisMap["value"] = 2
	//setup attrs

	tc.SetInput("keys", []interface{}{"a", "b", "c"})
	tc.SetInput("values", []interface{}{0.01, 0.02, 0.03})

	act.Eval(tc)

	result := tc.GetOutput("values").([]map[string]interface{})
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
