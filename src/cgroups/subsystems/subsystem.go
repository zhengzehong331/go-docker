package subsystems

// ResourceConfig 用于传递资源配置限制的结构体，CPU时间片权重、CPU核心数
type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

// Subsystem Subsystem接口每个Subsystem可以实现下面的四个接口
// 这里将cgroup抽象成了path，原因是cgroup在hierarchy的路径，便是虚拟文件系统中的路径。
type Subsystem interface {
	// Name 返回Subsystem的名字，比如cpu memory
	Name() string
	// 设置某个cgroup 在这个Subsystem中的资源限制
	Set(path string, res *ResourceConfig) error
	// 将进程添加到某个cgroup中
	Apply(path string, pid int) error
	// 移除某个cgroup
	Remove(path string) error
}

// 不同的subsytem初始化实例创建资源限制处理链数组
var (
	SubsystemIns = []Subsystem{
		&CpuSubSystem{},
		&MemorySubSystem{},
		&CpusetSubSystem{},
	}
)
