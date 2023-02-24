package container

import (
	"os"
	"syscall"

	"github.com/sirupsen/logrus"
)

func RunContainerInitProcess(command string, args []string) error {
	logrus.Infof("command %s", command)

	// MS_NOEXEC: 在文本系统中不允许运行其他程序
	// MS_NOSUID: 在本系统运行程序的时候，不允许其他set-user-id和set- group-id
	// MS_NODEV:  默认参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	argv := []string{command}

	// 这里是设置初始化进程为用户进程，如果不对这里进行设置，那么这个进程将会是init进程，和我们在docker启动时看到的不符
	// 这里的syscall.Exec方法最终调用了 int execve，调用这个方法,将用户指定的进程运行起来,把最初的init进程给替换掉,这样当进入到容器内部的时候,就会发现容器内的第一个程序就是我们指定的进程了。
	if err := syscall.Exec(command, argv, os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}
