package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const usage = "This is mydocker, Please enjoy it, just have fun"

func main() {
	// 初始化命令行
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage
	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(ctx *cli.Context) error {
		// 初始化logrus日志配置
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// func Run() {
// 	// 抓到命令行第二个参数
// 	cmd := exec.Command(os.Args[0], "init", os.Args[2])
// 	cmd.SysProcAttr = &syscall.SysProcAttr{
// 		// 这里调用Linux的内核函数，进行进程隔离,隔离UTS与PID,并且隔离CLONE
// 		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
// 		UidMappings: []syscall.SysProcIDMap{
// 			{
// 				ContainerID: 0,
// 				HostID:      os.Geteuid(),
// 				Size:        1,
// 			},
// 		},
// 		GidMappings: []syscall.SysProcIDMap{
// 			{
// 				ContainerID: 0,
// 				HostID:      os.Getuid(),
// 				Size:        1,
// 			},
// 		},
// 	}
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}
// 	cmd.Wait()
// }

// func Init() {
// 	// 在真实的Docker中，实际上并不是直接复制，而是通过Copy andr right机制实现的（可以去查一下）。
// 	imageFolderPath := "/var/lib/mydocker/images/base"
// 	rootFolderPath := "/var/lib/mydocker/containers/" + GenerateContainerId(64)
// 	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err) {
// 		if err := CopyFileOrDirectory(imageFolderPath, rootFolderPath); err != nil {
// 			panic(err)
// 		}
// 	}
// 	if err := syscall.Sethostname([]byte("container")); err != nil {
// 		panic(err)
// 	}
// 	// MS_NOEXEC: 在文本系统中不允许运行其他程序
// 	// MS_NOSUID: 在本系统运行程序的时候，不允许其他set-user-id和set- group-id
// 	// MS_NODEV:  默认参数
// 	// defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
// 	if err := syscall.Chroot(rootFolderPath); err != nil {
// 		panic(err)
// 	}
// 	if err := syscall.Chdir("/"); err != nil {
// 		panic(err)
// 	}
// 	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
// 		panic(err)
// 	}
// 	path, err := exec.LookPath(os.Args[2])
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(path)
// 	// 解决两个shell进程问题
// 	if err := syscall.Exec(path, os.Args[2:], os.Environ()); err != nil {
// 		panic(err)
// 	}
// 	syscall.Unmount("/proc", 0)
// }

// func CopyFileOrDirectory(src string, dst string) error {
// 	fmt.Printf("Copy %s => %s\n", src, dst)
// 	cmd := exec.Command("cp", "-r", src, dst)
// 	return cmd.Run()
// }

// func GenerateContainerId(n uint) string {
// 	rand.Seed(time.Now().UnixNano())
// 	const letters = "abcdefghijklmnopqrstuvwxy0123456789"
// 	b := make([]byte, n)
// 	length := len(letters)
// 	for i := range b {
// 		b[i] = letters[rand.Intn(length)]
// 	}
// 	return string(b)

// }
