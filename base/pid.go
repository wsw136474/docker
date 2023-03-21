package main

//PID  Namespace例子
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
// PID Namespace是用来隔离进程 ID的。同样一个进程在不同的 PIDNamespace里可以拥 有不同的 PID。这样就可以理解， 在 dockercontainer里面， 使用 ps-ef经常会发现， 在容器 内， 前台运行的那个进程 PID 是 l， 但是在容器外，使用 ps -ef会发现同样的进程却有不同的
//PID， 这就是 PIDNamespace做的事情。
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
sh-4.2# pstree -pl
  ├─sshd(1150)───sshd(1493)─┬─bash(1495)───go(2239)─┬─ipc(2265)─┬─sh(2268)
           │                         │                       │           ├─{ipc}(2266)
           │                         │                       │           └─{ipc}(2267)
           │                         │                       ├─{go}(2240)
           │                         │                       ├─{go}(2241)
           │                         │                       ├─{go}(2242)
           │                         │                       ├─{go}(2263)
           │                         │                       ├─{go}(2264)
           │                         │                       └─{go}(2269)
           │                         ├─bash(1782)───bash(1792)
           │                         └─bash(2070)───go(2339)─┬─pid(2365)─┬─sh(2368)───pstree(2371)
           │                                                 │           ├─{pid}(2366)
           │                                                 │           └─{pid}(2367)

进程树,进程pid为2365

进入shell,进程pid为1
[root@iZwz9fa4vp2ttnrgz8mlkwZ base]# go run pid.go
sh-4.2# eco $$
sh: eco: 未找到命令
sh-4.2# echo $$
1


这里还不能使用 ps 来查看 ， 因为 ps 和 top 等命令会 使用/proc 内容，具体内容在下面的 MountNamespace 部分会进行讲解。
*/

