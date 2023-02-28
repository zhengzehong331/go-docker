package subsystems

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

// 通过/proc/self/mountinfo 找出某个subsystem的hierachy cgroup根节点坐在的目录
// 例如找到 memtory 中的根节点就为/sys/fs/cgroup/memmory

func FindCgroupMountpoint(subsystem string) string {
	fs, err := os.Open("/proc/self/mountinifo")
	if err != nil {
		return ""
	}

	defer fs.Close()

	scanner := bufio.NewScanner(fs)
	// 获取subsytem的类型
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	// // 如果没找到，就自己创建memory
	// cgPath := "/sys/fs/cgroup/" + subsystem
	return ""
}

// 得到cgroup在文件系统中的绝对路径
func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountpoint(subsystem)
	if cgroupRoot == "" {
		cgroupRoot = "/sys/fs/cgroup/" + subsystem
		if _, err := os.Stat(cgroupRoot); os.IsNotExist(err) {
			if err := os.Mkdir(cgroupRoot, 0755); err == nil {
			} else {
				return "", fmt.Errorf("error create cgroup subsystem di r%v", err)
			}
		}
	}
	fmt.Println(cgroupRoot)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err == nil {
			} else {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}
