package util

import (
	"context"
	"io/ioutil"
	"os"
	"os/exec"
)

type Cmd struct {
	*exec.Cmd
	ctx context.Context
}

type Cgroup struct{}

func (c *Cgroup) GetInitStub() string {
	return ""
}

func CommandContext(ctx context.Context, user string, isSeniorKill bool, cg *Cgroup, name string, arg ...string) *Cmd {
	return &Cmd{Cmd: exec.CommandContext(ctx, name, arg...)}
}

func CreateTempScript(content string, prefix string) (path string, err error) {
	f, err := ioutil.TempFile("", prefix)
	if err != nil {
		return "", err
	}

	if _, err = f.WriteString(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}
	if err = os.Chmod(f.Name(), 0500); err != nil {
		f.Close()
		os.Remove(f.Name())
		return "", err
	}

	f.Close()
	return f.Name(), nil
}
