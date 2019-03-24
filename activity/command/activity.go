package command

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-alexandrev-command")

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

	// do eval
	command := context.GetInput("command").(string)
	os := context.GetInput("os").(string)
	log.Debug("command: " + command)
	fmt.Println("command: ", command)
	cmd := []byte{}
	if os == "windows" {
		cmd, err = exec.Command("cmd", "/C", command).CombinedOutput()
	} else {
		cmd, err = exec.Command(command).CombinedOutput()
	}
	fmt.Println("output: ", string(cmd))
	log.Debugf("output: " + string(cmd))

	if err == nil {
		context.SetOutput("data", string(cmd))
	} else {
		context.SetOutput("data", "-1")
	}

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
