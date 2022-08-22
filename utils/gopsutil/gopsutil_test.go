package gopsutil

import (
	"fmt"
	"github.com/dablelv/go-huge-util"
	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"runtime"
	"testing"
	"time"
)

// 采集CPU相关信息
func TestGetCPUInfo(t *testing.T) {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err: %v", err)
	}
	for _, cpuInfo := range cpuInfos {
		fmt.Println(cpuInfo)
	}
	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, true)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// 获取CPU负载消息
func TestGetCpuLoad(t *testing.T) {
	info, _ := load.Avg()
	fmt.Printf("%v\n", info)
}

// Memory
func TestGetMemInfo(t *testing.T) {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("mem info %v\n", memInfo)
}

// Host
func TestGetHost(t *testing.T) {
	hostInfo, _ := host.Info()
	fmt.Printf("host info : %v hostId: %v uptime: %v boottime: %v\n", hostInfo, hostInfo.HostID, hostInfo.Uptime, hostInfo.BootTime)
}

// Disk
func TestGetDiskInfo(t *testing.T) {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get partitions failed, err: %v", err)
	}
	for _, part := range parts {
		fmt.Printf("part: %v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info: \n used:%v free:%v\n", diskInfo.Used, diskInfo.Free)
	}
	ioStat, _ := disk.IOCounters()
	for k, v := range ioStat {
		fmt.Printf("ioStat %v:%v\n", k, v)
	}
}

// IO
func TestGetNetInfo(t *testing.T) {
	info, _ := net.IOCounters(true)
	for i, stat := range info {
		// 转换为JSON结构体打印
		data, _ := util.ToIndentJSON(stat)
		fmt.Printf("%v:%v send:%v recv:%v\n", i, data, stat.BytesSent, stat.BytesRecv)
	}
}

// 获取CPU逻辑核数
// sysctl hw.logicalcpu
// sysctl hw.physicalcpu
func TestGetLogicalCPUNum(t *testing.T) {
	fmt.Println(runtime.NumCPU())
}

// Don't just check errors, handle them gracefully.
func TestWrapErr(t *testing.T) {
	err := newErr()
	if err != nil {
		wrapErr := errors.Wrap(err, "new err")
		fmt.Printf("test wrap err: %v\n", wrapErr)
		causeErr := errors.Cause(err)
		fmt.Printf("test cause err: %v\n", causeErr)
	}

}

func newErr() error {
	err := errors.New("error")
	err = errors.Wrap(err, "open failed")
	err = errors.Wrap(err, "read config failed")
	return err
}
