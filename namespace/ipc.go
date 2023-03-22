package main

//IPC Namespace例子
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
// IPC Namespace 用来隔离 System V IPC 和 POSIX message queues。 每一个 IPC Namespace 都有自己的 System V IPC 和 POSIX message queu
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
可以看到，仅仅增加 syscall.CLONE_NEWIPC 代表我们希 望创建 IPC Namespace。下面 ， 需要打开两个 shell 来演示隔离的效果
在主机上查看消息队列
[root@iZwz9fa4vp2ttnrgz8mlkwZ ~]# ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息

创建一个
[root@iZwz9fa4vp2ttnrgz8mlkwZ ~]# ipcmk -Q
消息队列 id：0
[root@iZwz9fa4vp2ttnrgz8mlkwZ ~]# ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息
0x9734c163 0          root       644        0            0

再去sh进程中查看消息队列
[root@iZwz9fa4vp2ttnrgz8mlkwZ base]# go run ipc.go
sh-4.2# ipcs -q

--------- 消息队列 -----------
键        msqid      拥有者  权限     已用字节数 消息
*/
