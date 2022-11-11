package cron

import (
	"os/exec"
	"testing"
)

func TestCron_Initialize(t *testing.T) {
	var res []byte
	var err error
	var cmd *exec.Cmd
	cmd = exec.Command(
		"cmd",
		"/c",
		"echo hello world")
	if res, err = cmd.Output(); err != nil {
		t.Fatalf("错误：%s", err)
	}
	//t.Logf(string(res))
	t.Log(res)
}
