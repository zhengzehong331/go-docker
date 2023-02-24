package container

import (
	"os"
	"os/exec"
	"syscall"
)

func NewParentProcess(tty bool, command string) *exec.Cmd {
	// 这里就是做了colne一个namespace隔离的进程，然后在这个子进程中，调用/proc/self/exe
	args := []string{"init", command}
	/*
		1. proc/self/exe是一个特殊的文件，包含当前可执行文件的内存映像。
		换句话说，会让进程重新运行自己，但是传递child作为第一个参数。
		2. 后面的args实际上是参数，init就是传递给本进程的第一个参数
		3. 下面的clone参数就是去fork出来一个新进程,并且使用了namespace隔离新创建的进程和外部环境
		4. 如果用户指定了t i参数,就需要把当前进程的输入输出导入到标准输入输出上

	*/
	cmd := exec.Command("/proc/self/exe", args...) //这里会先调用init去初始化进程的一些资源与容器
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET,
	}
	if tty {
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}
