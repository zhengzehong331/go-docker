package main

import (
	"fmt"
	"mydocker/src/cgroups/subsystems"
	"mydocker/src/command"
	"mydocker/src/container"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var runCommand = cli.Command{
	Name:  "run",
	Usage: "Create a container with namespace and groups limit mydocker run -it [command]",
	Flags: []cli.Flag{
		// 这里定义了runCommand 的 Flag，作用列斯与允许命令时候使用--来指定参数
		cli.BoolFlag{
			Name:  "it",
			Usage: "enable tty",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	Action: func(context *cli.Context) error {
		// 这里是实际执行的函数
		//1. 判断参数是否包含command
		//2. 获取用户指定的command
		//3. 调用Run func去准备启动容器
		if len(context.Args()) < 1 {
			return fmt.Errorf("Missing container command")
		}
		var cmdArrary []string
		for _, arg := range context.Args() {
			cmdArrary = append(cmdArrary, arg)
		}
		tty := context.Bool("it")
		// 添加资源限制
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: context.String("m"),
			CpuSet:      context.String("cpuset"),
			CpuShare:    context.String("cpushare"),
		}
		command.Run(tty, cmdArrary, resConf)
		return nil
	},
}

var initCommand = cli.Command{
	Name:  "init",
	Usage: "Init container process run user's process ini container.Don't call it outside",
	/*
		1. 获取传递过来的command参数
		2. 执行容器初始化操作
	*/
	Action: func(context *cli.Context) error {
		logrus.Infof("init come on")
		cmd := context.Args().Get(0)
		logrus.Infof("command %s", cmd)
		err := container.RunContainerInitProcess()
		return err
	},
}
