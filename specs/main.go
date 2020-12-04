package specs

import (
	"fmt"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHardwareData() {
	runtimeOS := runtime.GOOS
	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)

	// disk - start from "/" mount point for Linux
	// might have to change for Windows!!
	// don't have a Window to test this out, if detect OS == windows
	// then use "\" instead of "/"

	diskStat, err := disk.Usage("/")
	dealwithErr(err)

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)

	// get interfaces MAC/hardware address
	interfStat, err := net.Interfaces()
	dealwithErr(err)

	fmt.Println("OS" + runtimeOS)
	fmt.Println("Total memory: " + strconv.FormatUint(vmStat.Total, 10) + " bytes ")
	fmt.Println("Free memory: " + strconv.FormatUint(vmStat.Free, 10) + " bytes<")

	// disk
	fmt.Println("Total disk space: " + strconv.FormatUint(diskStat.Total, 10) + " bytes ")
	fmt.Println("Used disk space: " + strconv.FormatUint(diskStat.Used, 10) + " bytes")
	fmt.Println("Free disk space: " + strconv.FormatUint(diskStat.Free, 10) + " bytes")

	// since my machine has one CPU, I'll use the 0 index
	// if your machine has more than 1 CPU, use the correct index
	// to get the proper data
	fmt.Println("CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10))
	fmt.Println("VendorID: " + cpuStat[0].VendorID)
	fmt.Println("Family: " + cpuStat[0].Family)
	fmt.Println("Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10))
	fmt.Println("Model Name: " + cpuStat[0].ModelName)
	fmt.Println("Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz ")

	for idx, cpupercent := range percentage {
		fmt.Println("Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%")
	}

	fmt.Println("Hostname: " + hostStat.Hostname)
	fmt.Println("Uptime: " + strconv.FormatUint(hostStat.Uptime, 10))
	fmt.Println("Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10))

	// another way to get the operating system name
	// both darwin for Mac OSX, For Linux, can be ubuntu as platform
	// and linux for OS

	fmt.Println("OS: " + hostStat.OS)
	fmt.Println("Platform: " + hostStat.Platform)

	// the unique hardware id for this machine
	fmt.Println("Host ID(uuid): " + hostStat.HostID)

	// Memory
	v, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	fmt.Println("GPU used" + fmt.Sprintf("%.2f", v.UsedPercent))

	for _, interf := range interfStat {
		fmt.Println("------------------------------------------------------")
		fmt.Println("Interface Name: " + interf.Name)

		if interf.HardwareAddr != "" {
			fmt.Println("Hardware(MAC) Address: " + interf.HardwareAddr)
		}

		for _, flag := range interf.Flags {
			fmt.Println("Interface behavior or flags: " + flag)
		}

		for _, addr := range interf.Addrs {
			fmt.Println("IPv6 or IPv4 addresses: " + addr.String())

		}

	}

}

func main() {
	GetHardwareData()
}
