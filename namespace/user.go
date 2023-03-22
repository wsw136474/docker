package main

//User Namespace 主要是隔离用户 的用户组 ID
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
//一个进程的 User ID 和 Group ID在UserNamespace内外可以是不同的。 比较常用的是，在宿主机上以一个非root用户运行
// 创建一个 User Namespace， 然后在 User Namespace 里面却映射成 root 用户。这意味着 ， 这个 进程在 User Namespace 里面有 root权限
//，但是在 User Namespace 外面却没有 root 的权限。从 Linux Kernel 3.8开始， 非root进程也可以创建UserNamespace， 
//并且此用户在Namespace里 面可以被映射成 root， 且在 Namespace 内有 root权限。 
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(-1)
}

/*

*/

