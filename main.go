package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func demoCpu() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoCpu()")

	// Number of CPUs.

	cpuCount, _ := cpu.Counts(true)
	fmt.Printf("cpu.Counts(true): %d\n", cpuCount)
	cpuCount, _ = cpu.Counts(false)
	fmt.Printf("cpu.Counts(false): %d\n\n", cpuCount)

	// Per CPU statistics.

	timesFormatString := "timeStats[%d]: \n\tCPU: %s \n\tGuest: %f \n\tGuestNice: %f \n\tIdle: %f \n\tIowait: %f \n\tIrq: %f \n\tNice: %f \n\tSoftirq: %f \n\tSteal: %f \n\tStolen: %f \n\tSystem: %f \n\tTotal: %f \n\tUser: %f\n"

	timeStats, _ := cpu.Times(true)
	for i, timeStat := range timeStats {
		fmt.Printf(timesFormatString,
			i,
			timeStat.CPU,
			timeStat.Guest,
			timeStat.GuestNice,
			timeStat.Idle,
			timeStat.Iowait,
			timeStat.Irq,
			timeStat.Nice,
			timeStat.Softirq,
			timeStat.Steal,
			timeStat.Stolen,
			timeStat.System,
			timeStat.Total(),
			timeStat.User,
		)
	}

	// Total statistics.  Note "false" in cpu.Times(false)

	fmt.Printf("Total ")

	timeStats, _ = cpu.Times(false)
	for i, timeStat := range timeStats {
		fmt.Printf(timesFormatString,
			i,
			timeStat.CPU,
			timeStat.Guest,
			timeStat.GuestNice,
			timeStat.Idle,
			timeStat.Iowait,
			timeStat.Irq,
			timeStat.Nice,
			timeStat.Softirq,
			timeStat.Steal,
			timeStat.Stolen,
			timeStat.System,
			timeStat.Total(),
			timeStat.User,
		)
	}

	// cpu.Info()

	infoFormatString := "infoStats[%d]: \n\tCPU: %d \n\tCacheSize: %d \n\tCoreID: %s \n\tCores: %d \n\tFamily: %s \n\tFlags: %v \n\tMhz: %4.1f \n\tMicrocode: %s \n\tModel: %s \n\tModelName: %s \n\tPhysicalID: %s \n\tStepping: %d \n\tVendorID: %s\n"

	infoStats, _ := cpu.Info()
	for i, infoStat := range infoStats {
		fmt.Printf(infoFormatString,
			i,
			infoStat.CPU,
			infoStat.CacheSize,
			infoStat.CoreID,
			infoStat.Cores,
			infoStat.Family,
			infoStat.Flags,
			infoStat.Mhz,
			infoStat.Microcode,
			infoStat.Model,
			infoStat.ModelName,
			infoStat.PhysicalID,
			infoStat.Stepping,
			infoStat.VendorID,
		)
	}

	// cpu.Percent()

	percentFormatString := "percent[%d]: %f\n"
	interval := time.Second

	percents, _ := cpu.Percent(interval, true)
	for i, percent := range percents {
		fmt.Printf(percentFormatString,
			i,
			percent,
		)
	}

	fmt.Printf("Total ")
	percents, _ = cpu.Percent(interval, false)
	for i, percent := range percents {
		fmt.Printf(percentFormatString,
			i,
			percent,
		)
	}

}

func demoDisk() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoDisk()")

	// disk.Partitions()

	partitionFormatString := "partition[%d]: \n\tDevice: %s \n\tFstype: %s \n\tMountpoint: %s \n\tOpts: %s\n"

	partitions, _ := disk.Partitions(true)
	for i, partition := range partitions {
		fmt.Printf(partitionFormatString,
			i,
			partition.Device,
			partition.Fstype,
			partition.Mountpoint,
			partition.Opts,
		)
	}

	// disk.IOCounters()

	countersFormatString := "counters[%s]: \n\tIoTime: %d\n\tIopsInProgress: %d\n\tMergedReadCount: %d\n\tMergedWriteCount: %d\n\tName: %s\n\tReadBytes: %d\n\tReadCount: %d\n\tReadTime: %d\n\tSerialNumber: %s\n\tWeightedIO: %d\n\tWriteBytes: %d\n\tWriteCount: %d\n\tWriteTime: %d\n"

	counters, _ := disk.IOCounters("sda", "sdb")
	for key, value := range counters {
		fmt.Printf(countersFormatString,
			key,
			value.IoTime,
			value.IopsInProgress,
			value.MergedReadCount,
			value.MergedWriteCount,
			value.Name,
			value.ReadBytes,
			value.ReadCount,
			value.ReadTime,
			value.SerialNumber,
			value.WeightedIO,
			value.WriteBytes,
			value.WriteCount,
			value.WriteTime,
		)
	}

	// disk.Usage()

	usageFormatString := "usage[%s]: \n\tFree: %d\n\tFstype: %s\n\tInodesFree: %d\n\tInodesTotal: %d\n\tInodesUsed: %d\n\tInodesUsedPercent: %.1f\n\tPath: %s\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %.1f\n"

	path := "/"
	usage, _ := disk.Usage(path)
	fmt.Printf(usageFormatString,
		path,
		usage.Free,
		usage.Fstype,
		usage.InodesFree,
		usage.InodesTotal,
		usage.InodesUsed,
		usage.InodesUsedPercent,
		usage.Path,
		usage.Total,
		usage.Used,
		usage.UsedPercent,
	)

}

func demoHost() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoHost()")

	// host.BootTime()

	bootTime, _ := host.BootTime()
	fmt.Printf("host.BootTime(): %d  (%s)\n", bootTime, time.Unix(int64(bootTime), 0))

	// host.Uptime()

	upTime, _ := host.Uptime()
	fmt.Printf("host.Uptime(): %d\n\n", upTime)

	// host.Info()

	infoFormatString := "info[]: \n\tBootTime: %d\n\tHostID: %s\n\tHostname: %s\n\tKernelVersion: %s\n\tOS: %s\n\tPlatform: %s\n\tPlatformFamily: %s\n\tPlatformVersion: %s\n\tProcs: %d\n\tUptime: %d\n\tVirtualizationRole: %s\n\tVirtualizationSystem: %s\n"

	info, _ := host.Info()
	fmt.Printf(infoFormatString,
		info.BootTime,
		info.HostID,
		info.Hostname,
		info.KernelVersion,
		info.OS,
		info.Platform,
		info.PlatformFamily,
		info.PlatformVersion,
		info.Procs,
		info.Uptime,
		info.VirtualizationRole,
		info.VirtualizationSystem,
	)

	// host.Users()

	userFormatString := "user[%d]: \n\tHost: %s\n\tStarted: %d (%s)\n\tTerminal: %s\n\tUser: %s\n"

	users, _ := host.Users()
	for i, user := range users {
		fmt.Printf(userFormatString,
			i,
			user.Host,
			user.Started,
			time.Unix(int64(user.Started), 0),
			user.Terminal,
			user.User,
		)
	}

	// host.PlatformInformation()

	platformFormatString := "platform[]: \n\tPlatform: %s\n\tFamily: %s\n\tVersion: %s\n"

	platform, family, version, _ := host.PlatformInformation()
	fmt.Printf(platformFormatString,
		platform,
		family,
		version,
	)

	// host.KernelVersion()

	kernelFormatString := "kernel[]: \n\tVersion: %s\n"

	version, _ = host.KernelVersion()
	fmt.Printf(kernelFormatString,
		version,
	)

	// host.Virtualization()

	virtualizationFormatString := "virtualization[]: \n\tSystem: %s\n\tRole: %s\n"

	system, role, _ := host.Virtualization()
	fmt.Printf(virtualizationFormatString,
		system,
		role,
	)

	// host.SensorsTemperatures()

	temperatureFormatString := "temperature[%d]: \n\tSensorKey: %s\n\tTemperature: %f\n"

	temperatures, _ := host.SensorsTemperatures()
	for i, temperature := range temperatures {
		fmt.Printf(temperatureFormatString,
			i,
			temperature.SensorKey,
			temperature.Temperature,
		)
	}

}

func main() {
		demoCpu()
		demoDisk()
	demoHost()
}
