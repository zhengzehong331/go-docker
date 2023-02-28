package container

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
)

func RunContainerInitProcess() error {
	cmdArray := readUserCommand()
	if cmdArray == nil || len(cmdArray) == 0 {
		return fmt.Errorf("Run contialner get user command error,cmd Array is nil")
	}

	// ststem加入linux之后，mount namespace 变成了shared bt defalult 所以你必须显示声明你要这个新的mount namespace独立
	syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")

	// MS_NOEXEC: 在文本系统中不允许运行其他程序
	// MS_NOSUID: 在本系统运行程序的时候，不允许其他set-user-id和set- group-id
	// MS_NODEV:  默认参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	// 改动，调用exec.LookPath,可以在系统的PATH里面寻找命令的绝对路径
	path, err := exec.LookPath(cmdArray[0])
	if err != nil {
		logrus.Errorf("Exec loop path error %v", err)
		return err

	}
	logrus.Infof("Find path %s", path)
	// 这里是设置初始化进程为用户进程，如果不对这里进行设置，那么这个进程将会是init进程，和我们在docker启动时看到的不符
	// 这里的syscall.Exec方法最终调用了 int execve，调用这个方法,将用户指定的进程运行起来,把最初的init进程给替换掉,这样当进入到容器内部的时候,就会发现容器内的第一个程序就是我们指定的进程了。
	if err := syscall.Exec(path, cmdArray[0:], os.Environ()); err != nil {
		logrus.Errorf(err.Error())
	}
	return nil
}

func readUserCommand() []string {
	// uintptr(3)就是指index为3个文件描述符，也就是传递进来的管道的一端
	pipe := os.NewFile(uintptr(3), "pipe")
	msg, err := ioutil.ReadAll(pipe)
	if err != nil {
		logrus.Errorf("init read pipe error %v", err)
		return nil
	}
	msgStr := string(msg)
	return strings.Split(msgStr, " ")
}
