package subsystems

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

// memory subsystem的实现
type MemorySubSystem struct{}

func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	// GetCgroupPath 的作用是获取当前subsystem在虚拟文件系统中的路径，GetCgroupPath这个函数将在下面实现
	if subsystemPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {

		if res.MemoryLimit != "" {
			//设置这个cgroup的内存限制，将限制写入到cgroup对应的memory.limit_in_bytes文件中
			if err := ioutil.WriteFile(path.Join(subsystemPath, "memory.limit_in_bytes"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

// 删除cgroupPath 对应的cgroup
func (s *MemorySubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		// 删除cgroup 对应的cgroupPath目录
		return os.Remove(subsysCgroupPath)
	} else {
		return err
	}
}

// 将这个进程加入到cgroupPath对应的cgroup中
func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, false); err == nil {
		// 把进程的PID写到cgroup的虚拟文件系统对应目录下的“task”文件中
		if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "task"), []byte(strconv.Itoa(pid)), 0664); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return fmt.Errorf("get cgrup %s error : %v", cgroupPath, err)
	}
}

// 返回cgroup的名字
func (s *MemorySubSystem) Name() string {
	return "memmory"
}
