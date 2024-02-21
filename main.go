package main

import (
	"gopsutils/dto"
	"fmt"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	memory := MemoryInfo()
	fmt.Println(memory)
	cpu1 := CpuInfo()
	fmt.Println(cpu1)
	disk1 := DiskInfo()
	fmt.Println(disk1)
	host1 := HostInfo()
	fmt.Println(host1)
	diskIO1 := DiskIO()
	fmt.Println(diskIO1)
	load1 := CpuLoad()
	fmt.Println(load1)
}

func MemoryInfo() dto.MemoryInfoDto {

	v, _ := mem.VirtualMemory()

	memoryState := dto.MemoryInfoDto{
		Tatol:       fmt.Sprintf("%vM", v.Total/1024/1024),
		Available:   fmt.Sprintf("%vM", v.Available/1024/1024),
		UsedPercent: fmt.Sprintf("%.2f%%", v.UsedPercent),
	}
	if v.UsedPercent >= 80 {
		memoryState.Abnormal = true
	}
	return memoryState
}

func CpuInfo() dto.CpuInfoDto {
	var (
		cpuInfo dto.CpuInfoDto
		logical bool
	)
	cpuPercent, _ := cpu.Percent(time.Second*1, false)
	CpuState, _ := cpu.Info()
	//基于性能考虑，只获取第一颗CPU的信息
	physicalCpu := CpuState[0]
	cpuInfo = dto.CpuInfoDto{
		ModelName: physicalCpu.ModelName,
		Mhz:       int(physicalCpu.Mhz),
	}
	if strings.Contains(strings.Join(physicalCpu.Flags, " "), " avx2 ") {
		cpuInfo.EnableAvx2 = true
	}
	count, _ := cpu.Counts(logical)
	cpuInfo.CoreCount = count

	if cpuPercent[0] >= 80 {
		cpuInfo.Abnormal = false
	}
	cpuInfo.UsedPercent = fmt.Sprintf("%.2f%%", cpuPercent[0])
	return cpuInfo
}
func CpuLoad() dto.CpuLoadDto {
	var cpuLoad dto.CpuLoadDto
	loadSate, _ := load.Avg()
	cpuLoad = dto.CpuLoadDto{
		Load1:  fmt.Sprintf("%.2f", loadSate.Load1),
		Load5:  fmt.Sprintf("%.2f", loadSate.Load5),
		Load15: fmt.Sprintf("%.2f", loadSate.Load15),
	}
	return cpuLoad
}

func DiskInfo() dto.DiskInfoDto {
	var diskInfo dto.DiskInfoDto
	var disks []dto.SingleDiskInfoDto
	//先查询分区信息
	diskPartitions, _ := disk.Partitions(false)
	for _, diskPartition := range diskPartitions {
		//不显示k8s内部磁盘信息
		if !strings.Contains(strings.Join(diskPartition.Opts, ""), "bind") {
			singleFs := dto.SingleDiskInfoDto{
				Fstype:     diskPartition.Fstype,
				FsPath:     diskPartition.Mountpoint,
				DeviceName: diskPartition.Device,
			}
			//磁盘用量
			fs, _ := disk.Usage(diskPartition.Mountpoint)
			singleFs.Tatal = fmt.Sprintf("%vG", fs.Total/1024/1024/1024)
			singleFs.Free = fmt.Sprintf("%vG", fs.Free/1024/1024/1024)
			if fs.UsedPercent >= 80 {
				singleFs.Abnormal = true
			}
			singleFs.UsedPercent = fmt.Sprintf("%.2f%%", fs.UsedPercent)
			disks = append(disks, singleFs)
		}
	}
	diskInfo.Disks = disks
	return diskInfo
}

func DiskIO() dto.DiskIODto {
	var diskIO dto.DiskIODto
	var singleDiskIOs []dto.SingleDiskIODto
	IOMapState, _ := disk.IOCounters()
	for _, ioStat := range IOMapState {
		singleDiskIO := dto.SingleDiskIODto{
			DiskName:   ioStat.Name,
			ReadCount:  ioStat.ReadCount,
			WriteCount: ioStat.WriteCount,
			ReadBytes:  fmt.Sprintf("%.2fG", float64(ioStat.ReadBytes)/1024/1024/1024),
			WriteBytes: fmt.Sprintf("%.2fG", float64(ioStat.WriteBytes)/1024/1024/1024),
		}
		singleDiskIOs = append(singleDiskIOs, singleDiskIO)
	}
	diskIO.DiskIO = singleDiskIOs
	return diskIO
}

func HostInfo() dto.HostInfoDto {
	hostState, _ := host.Info()
	t := time.Unix(int64(hostState.BootTime), 0)
	d := time.Duration(hostState.Uptime * 1000 * 1000 * 1000)
	var hostInfo = dto.HostInfoDto{
		Hostname:        hostState.Hostname,
		BootTime:        t.Format(time.DateTime),
		Uptime:          d.Abs().String(),
		Procs:           hostState.Procs,
		OS:              hostState.OS,
		Platform:        hostState.Platform,
		PlatformFamily:  hostState.PlatformFamily,
		PlatformVersion: hostState.PlatformVersion,
		KernelVersion:   hostState.KernelVersion,
		KernelArch:      hostState.KernelArch,
	}
	if d.Hours() <= 24 {
		hostInfo.Abnormal = true
	}
	return hostInfo
}
