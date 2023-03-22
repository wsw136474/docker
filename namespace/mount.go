package main

//mount Namespace例子
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
// MountNamespace用来隔离各个进程看到的挂载点视图,
//在不同Namespace的进程中， 看 到的文件系统层次是不一样的。在 Mount Namespace 中调用 mount()和 umount()仅仅只会影响 当前 Namespace 内的文件系统 ，而对全局的文件系统是没有影响的。
func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
[root@senwei base]# go run mount.go
sh-4.2# ls /proc/
1     15950  18   31   509	  consoles     key-users     self
10    16     19   32   521	  cpuinfo      kmsg	     slabinfo
1049  17     2	  33   523	  crypto       kpagecount    softirqs
1081  17173  20   350  524	  devices      kpageflags    stat
1083  17246  21   371  59	  diskstats    loadavg	     swaps
1092  17248  22   4    6	  dma	       locks	     sys
11    17250  226  41   7	  driver       mdstat	     sysrq-trigger
1147  17260  23   413  793	  execdomains  meminfo	     sysvipc
1150  17301  234  417  8	  fb	       misc	     timer_list
1151  17313  235  418  859	  filesystems  modules	     timer_stats
1152  17350  236  42   872	  fs	       mounts	     tty
1157  17382  237  43   9	  interrupts   mtrr	     uptime
1158  17393  24   44   95	  iomem        net	     version
1295  17446  240  45   acpi	  ioports      pagetypeinfo  vmallocinfo
13    17499  246  46   buddyinfo  irq	       partitions    vmstat
14    17525  255  494  bus	  kallsyms     sched_debug   zoneinfo
1413  17528  256  5    cgroups	  kcore        schedstat
15    17535  30   508  cmdline	  keys	       scsi

重新挂载proc到新的namespace之下.proc是一个伪文件系统
sh-4.2#  mount -t proc proc /proc/
sh-4.2# ps -aux
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.1 115544  1940 pts/0    S    21:50   0:00 sh
root         5  0.0  0.0 155452  1868 pts/0    R+   22:16   0:00 ps -aux
可 以 看到，在当前 的 Namespace 中， sh 进程是 PID 为 l 的进程。这就说明 ， 当前的 Mount Namespace 中的 mount 和外部空间是隔离的， mount 操作并没有影响到外部。 Docker volume 也 是利用了这个特性。
*/

