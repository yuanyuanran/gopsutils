package dto

type MemoryInfoDto struct {
	Tatol       string `json:"tatol"`        //总内存，单位Mb
	Available   string `json:"free"`         //可用内存，单位Mb
	UsedPercent string `json:"used_percent"` //使用率
	Abnormal    bool   `json:"abnormal"`     //是否异常，为ture则为异常
}

type CpuInfoDto struct {
	CoreCount   int    `json:"core_count"`   //物理核心数量
	UsedPercent string `json:"used_percent"` //使用率
	Abnormal    bool   `json:"abnormal"`     //是否异常，为ture则为异常
	ModelName   string `json:"model_name"`   //CPU名称
	Mhz         int    `json:"mhz"`          //CPU频率
	EnableAvx2  bool   `json:"avx2"`         //是否支持avx2指令集
}

type CpuLoadDto struct {
	Load1  string `json:"load1"`
	Load5  string `json:"load5"`
	Load15 string `json:"load15"`
}

type DiskInfoDto struct {
	Disks []SingleDiskInfoDto `json:"disks"`
	// DiskIO []SingleDiskIODto      `json:"disk_ios"`
}

type SingleDiskInfoDto struct {
	Fstype      string `json:"fs_type"`      //磁盘格式
	FsPath      string `json:"fs_path"`      //文件路径
	DeviceName  string `json:"device_name"`  //设备名称
	Tatal       string `json:"tatol"`        //磁盘容量
	Free        string `json:"free"`         //磁盘可用
	UsedPercent string `json:"used_percent"` //使用率
	Abnormal    bool   `json:"abnormal"`     //是否异常，为ture则为异常
}
type HostInfoDto struct {
	Hostname        string `json:"hostname"`        //主机名称
	Uptime          string `json:"uptime"`          //启动时长
	BootTime        string `json:"bootTime"`        //启动时间
	Procs           uint64 `json:"procs"`           // 进程数
	OS              string `json:"os"`              // 系统: freebsd, linux
	Platform        string `json:"platform"`        // 平台: ubuntu, linuxmint
	PlatformFamily  string `json:"platformFamily"`  // 发行版: debian, rhel
	PlatformVersion string `json:"platformVersion"` // 系统版本
	KernelVersion   string `json:"kernelVersion"`   // 内核版本
	KernelArch      string `json:"kernelArch"`      // 架构
	Abnormal        bool   `json:"abnormal"`        //是否异常，为ture则为异常,判断依据启动时间小于24小时
}

type DiskIODto struct {
	DiskIO []SingleDiskIODto `json:"disk_ios"`
}

type SingleDiskIODto struct {
	DiskName         string `json:"disk_name"`
	ReadCount        uint64 `json:"read_count"`  //读取个数
	WriteCount       uint64 `json:"write_count"` //写入个数
	ReadBytes        string `json:"readBytes"`   //总读取的数据单位:G
	WriteBytes       string `json:"writeBytes"`  //总写入的数据单位:G
}
