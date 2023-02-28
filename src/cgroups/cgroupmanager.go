package cgroups

import (
	"mydocker/src/cgroups/subsystems"

	"github.com/sirupsen/logrus"
)

type CgroupManager struct {
	// cgroup 在 hierarchy中的相对路径
	// 这里其实就是我们所定义的mydocker-cgroup的hierarchy
	Path string
	//资源配置
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

// 将进程加入PID加入到每个cgroup中
func (c *CgroupManager) Apply(pid int) error {
	for _, subSysInts := range subsystems.SubsystemIns {
		subSysInts.Apply(c.Path, pid)
	}
	return nil
}

// 设置各个subsystem挂载中的cgroup资源限制
func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysInts := range subsystems.SubsystemIns {
		subSysInts.Set(c.Path, res)
	}
	return nil
}

// 释放各个subsystem挂载中的cgroup
func (c *CgroupManager) Destroy() error {
	for _, subSysInts := range subsystems.SubsystemIns {
		if err := subSysInts.Remove(c.Path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
