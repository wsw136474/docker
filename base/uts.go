package main

//UTS Namespace例子
import (
	"log"
	"os"
	"os/exec"
	"syscall"
)
// NEWUTS这个标识符去创建一个UTSNamespace。 Go帮我们封装了对clone() 函数的调用， 这段代码执行后就会进入到一个 sh运行环境中
func main() {
	// exec.Command (“h”)用来指定被 fork 出来的新进程内的初始命令，默 认使用 sh 来执行
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

/*
由于 UTS Namespace对 hostname做了 隔离， 所以在这个环境内修改 hostname应该不影响外部主机， 下面来做一下实验。
在这个 sh 环境内 执行如下代码示例 。
#修改 hostname为 bird然后打印出来 # hostname -b bird
# hostname
bird
另外启动一个 shell，在宿主机上运行 hostname，看一下效果。 root@iZ254rt8xflZ:~# host name
iZ254rt8xflZ
可以看到， 外部的 hostname井没有被内部的修改所影响，由此可了解UTSNamespace的 作用 。
*/

