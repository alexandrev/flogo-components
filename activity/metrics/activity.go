package metrics

import (
	"fmt"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/data"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

var log = logger.GetLogger("activity-alexandrev-metrics")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	vmStat, err := mem.VirtualMemory()
	cpuStat, err := cpu.Info()
	percentage, err := cpu.Percent(0, true)
	hostStat, err := host.Info()

	output := []map[string]interface{}{}

	var field1value = map[string]interface{}{}
	field1value["key"] = "memTotal"
	field1value["value"] = fmt.Sprintf("%.2d", vmStat.Total/(1024*1024))
	output = append(output, field1value)

	var field2value = map[string]interface{}{}
	field2value["key"] = "memUsed"
	field2value["value"] = fmt.Sprintf("%.2d", vmStat.Used/(1024*1024))
	output = append(output, field2value)

	var field3value = map[string]interface{}{}
	field3value["key"] = "memPercentage"
	field3value["value"] = fmt.Sprintf("%.2f", vmStat.UsedPercent)
	output = append(output, field3value)

	var field4value = map[string]interface{}{}
	field4value["key"] = "cpuNumber"
	field4value["value"] = len(percentage)
	output = append(output, field4value)

	var field5value = map[string]interface{}{}
	field5value["key"] = "cpuSpeed"
	field5value["value"] = cpuStat[0].Mhz
	output = append(output, field5value)

	var field6value = map[string]interface{}{}
	field6value["key"] = "cpuUsed"
	field6value["value"] = fmt.Sprintf("%.2f", percentage[0])
	output = append(output, field6value)

	var field7value = map[string]interface{}{}
	field7value["key"] = "procRunning"
	field7value["value"] = hostStat.Procs
	output = append(output, field7value)

	complexField2Value := &data.ComplexObject{Metadata: "", Value: output}
	context.SetOutput("output", complexField2Value)
	return true, nil
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}
