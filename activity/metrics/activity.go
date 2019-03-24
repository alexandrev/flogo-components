package command

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
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

	context.SetOutput("memTotal", vmStat.Total)
	context.SetOutput("memUsed", vmStat.Used)
	context.SetOutput("memPercentage", vmStat.UsedPercent)
	context.SetOutput("cpuNumber", len(percentage))

	context.SetOutput("cpuSpeed", cpuStat[0].Mhz)
	context.SetOutput("cpuUsed", percentage)
	context.SetOutput("procRunning", hostStat.Procs)

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
