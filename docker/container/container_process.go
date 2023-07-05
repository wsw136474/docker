package container

import (
	"os"
	"os/exec"
	"syscall"
)

// NewParentProcess  /proc/self/exe 自己调用自己, 传入参数init.go继续进入initCommand逻辑
func NewParentProcess(tty bool, command string) *exec.Cmd {
	args := []string{"init.go", command}
	cmd := exec.Command("/proc/self/exe", args...)
	// 启动进程时设置单独的namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS |
			syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
