package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func displayBanner(title string) {
	fmt.Printf("\n---------- %s ------------------------------\n\n", title)
}

func demoCpuCounts() {
	cpuCount, _ := cpu.Counts(true)
	fmt.Printf("cpu.Counts(true): %d\n", cpuCount)
	cpuCount, _ = cpu.Counts(false)
	fmt.Printf("cpu.Counts(false): %d\n\n", cpuCount)
}

func demoCpuTimes(perCpu bool) {
	formatString := "timeStat[%d]:\n\tCPU: %s \n\tGuest: %f\n\tGuestNice: %f\n\tIdle: %f\n\tIowait: %f\n\tIrq: %f\n\tNice: %f\n\tSoftirq: %f\n\tSteal: %f\n\tStolen: %f\n\tSystem: %f\n\tTotal: %f\n\tUser: %f\n"
	timeStats, _ := cpu.Times(perCpu)
	for i, timeStat := range timeStats {
		if !perCpu {
			fmt.Printf("Total ")
		}
		fmt.Printf(formatString,
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
}

func demoCpuInfo() {
	formatString := "infoStat[%d]: \n\tCPU: %d \n\tCacheSize: %d \n\tCoreID: %s \n\tCores: %d \n\tFamily: %s \n\tFlags: %+v \n\tMhz: %4.1f \n\tMicrocode: %s \n\tModel: %s \n\tModelName: %s \n\tPhysicalID: %s \n\tStepping: %d \n\tVendorID: %s\n"
	infoStats, _ := cpu.Info()
	for i, infoStat := range infoStats {
		fmt.Printf(formatString,
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
}

func demoCpuPercent(perCpu bool) {
	formatString := "percent[%d]: %f\n"
	interval := time.Second
	percents, _ := cpu.Percent(interval, perCpu)
	for i, percent := range percents {
		if !perCpu {
			fmt.Printf("Total ")
		}
		fmt.Printf(formatString,
			i,
			percent,
		)
	}
}

func demoCpu() {
	displayBanner("Cpu")
	demoCpuCounts()
	demoCpuTimes(true)
	demoCpuTimes(false)
	demoCpuInfo()
	demoCpuPercent(true)
	demoCpuPercent(false)
}

func demoDiskPartitions(perCpu bool) {
	formatString := "partition[%d]: \n\tDevice: %s \n\tFstype: %s \n\tMountpoint: %s \n\tOpts: %s\n"
	partitions, _ := disk.Partitions(perCpu)
	for i, partition := range partitions {
		if !perCpu {
			fmt.Printf("Total ")
		}
		fmt.Printf(formatString,
			i,
			partition.Device,
			partition.Fstype,
			partition.Mountpoint,
			partition.Opts,
		)
	}
}

func demoDiskIOCounters() {
	formatString := "counter[%s]: \n\tIoTime: %d\n\tIopsInProgress: %d\n\tMergedReadCount: %d\n\tMergedWriteCount: %d\n\tName: %s\n\tReadBytes: %d\n\tReadCount: %d\n\tReadTime: %d\n\tSerialNumber: %s\n\tWeightedIO: %d\n\tWriteBytes: %d\n\tWriteCount: %d\n\tWriteTime: %d\n"
	counters, _ := disk.IOCounters("sda", "sdb")
	for key, counter := range counters {
		fmt.Printf(formatString,
			key,
			counter.IoTime,
			counter.IopsInProgress,
			counter.MergedReadCount,
			counter.MergedWriteCount,
			counter.Name,
			counter.ReadBytes,
			counter.ReadCount,
			counter.ReadTime,
			counter.SerialNumber,
			counter.WeightedIO,
			counter.WriteBytes,
			counter.WriteCount,
			counter.WriteTime,
		)
	}
}

func demoDiskUsage() {
	formatString := "usage[%s]: \n\tFree: %d\n\tFstype: %s\n\tInodesFree: %d\n\tInodesTotal: %d\n\tInodesUsed: %d\n\tInodesUsedPercent: %.1f\n\tPath: %s\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %.1f\n"
	path := "/"
	usage, _ := disk.Usage(path)
	fmt.Printf(formatString,
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

func demoDisk() {
	displayBanner("Disk")
	demoDiskPartitions(true)
	demoDiskPartitions(false)
	demoDiskIOCounters()
	demoDiskUsage()
}

func demoHostBootTime() {
	bootTime, _ := host.BootTime()
	fmt.Printf("host.BootTime(): %d  (%s)\n", bootTime, time.Unix(int64(bootTime), 0))
}

func demoHostUptime() {
	upTime, _ := host.Uptime()
	fmt.Printf("host.Uptime(): %d\n\n", upTime)
}

func demoHostInfo() {
	formatString := "info[]: \n\tBootTime: %d\n\tHostID: %s\n\tHostname: %s\n\tKernelVersion: %s\n\tOS: %s\n\tPlatform: %s\n\tPlatformFamily: %s\n\tPlatformVersion: %s\n\tProcs: %d\n\tUptime: %d\n\tVirtualizationRole: %s\n\tVirtualizationSystem: %s\n"
	info, _ := host.Info()
	fmt.Printf(formatString,
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
}

func demoHostUsers() {
	formatString := "user[%d]: \n\tHost: %s\n\tStarted: %d (%s)\n\tTerminal: %s\n\tUser: %s\n"
	users, _ := host.Users()
	for i, user := range users {
		fmt.Printf(formatString,
			i,
			user.Host,
			user.Started,
			time.Unix(int64(user.Started), 0),
			user.Terminal,
			user.User,
		)
	}
}

func demoHostPlatformInformation() {
	formatString := "platform[]: \n\tPlatform: %s\n\tFamily: %s\n\tVersion: %s\n"
	platform, family, version, _ := host.PlatformInformation()
	fmt.Printf(formatString,
		platform,
		family,
		version,
	)
}

func demoHostKernelVersion() {
	//  Bug reported: https://github.com/shirou/gopsutil/issues/409
	//	formatString := "kernel[]: \n\tVersion: %s\n"
	//	version, _ = host.KernelVersion()
	//	fmt.Printf(formatString,
	//		version,
	//	)
}

func demoHostVirtualization() {
	//  Bug reported: https://github.com/shirou/gopsutil/issues/411
	//	formatString := "virtualization[]: \n\tSystem: %s\n\tRole: %s\n"
	//	system, role, _ := host.Virtualization()
	//	fmt.Printf(formatString,
	//		system,
	//		role,
	//	)
}

func demoHostSensorsTemperatures() {
	formatString := "temperature[%d]: \n\tSensorKey: %s\n\tTemperature: %f\n"
	temperatures, _ := host.SensorsTemperatures()
	for i, temperature := range temperatures {
		fmt.Printf(formatString,
			i,
			temperature.SensorKey,
			temperature.Temperature,
		)
	}
}

func demoHost() {
	displayBanner("Host")
	demoHostBootTime()
	demoHostUptime()
	demoHostInfo()
	demoHostUsers()
	demoHostPlatformInformation()
	demoHostKernelVersion()
	demoHostVirtualization()
	demoHostSensorsTemperatures()
}

func demoLoadAvg() {
	formatString := "average[]: \n\tLoad1: %f\n\tLoad5: %f\n\tLoad15: %f\n"
	average, _ := load.Avg()
	fmt.Printf(formatString,
		average.Load1,
		average.Load5,
		average.Load15,
	)
}

func demoLoadMisc() {
	formatString := "misc[]: \n\tCtxt: %d\n\tProcsBlocked: %d\n\tProcsRunning: %d\n"
	misc, _ := load.Misc()
	fmt.Printf(formatString,
		misc.Ctxt,
		misc.ProcsBlocked,
		misc.ProcsRunning,
	)
}

func demoLoad() {
	displayBanner("Load")
	demoLoadAvg()
	demoLoadMisc()
}

func demoMemVirtualMemory() {
	formatString := "virtual[]: \n\tActive: %d\n\tAvailable: %d\n\tBuffers: %d\n\tCached: %d\n\tDirty: %d\n\tFree: %d\n\tInactive: %d\n\tPageTables: %d\n\tShared: %d\n\tSlab: %d\n\tSwapCached: %d\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %f\n\tWired: %d\n\tWriteback: %d\n\tWritebackTmp: %d\n"
	virtual, _ := mem.VirtualMemory()
	fmt.Printf(formatString,
		virtual.Active,
		virtual.Available,
		virtual.Buffers,
		virtual.Cached,
		virtual.Dirty,
		virtual.Free,
		virtual.Inactive,
		virtual.PageTables,
		virtual.Shared,
		virtual.Slab,
		virtual.SwapCached,
		virtual.Total,
		virtual.Used,
		virtual.UsedPercent,
		virtual.Wired,
		virtual.Writeback,
		virtual.WritebackTmp,
	)
}

func demoMemSwapMemory() {
	formatString := "swap[]: \n\tFree: %d\n\tSin: %d\n\tSout: %d\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %f\n"
	swap, _ := mem.SwapMemory()
	fmt.Printf(formatString,
		swap.Free,
		swap.Sin,
		swap.Sout,
		swap.Total,
		swap.Used,
		swap.UsedPercent,
	)
}

func demoMem() {
	displayBanner("Mem")
	demoMemVirtualMemory()
	demoMemSwapMemory()
}

func demoNetIOCounters(perCpu bool) {
	formatString := "iocounter[%d]: \n\tBytesRecv: %d\n\tBytesSent: %d\n\tDropin: %d\n\tDropout: %d\n\tErrin: %d\n\tErrout: %d\n\tFifoin: %d\n\tFifoout: %d\n\tName: %s\n\tPacketsRecv: %d\n\tPacketsSent: %d\n"
	iocounters, _ := net.IOCounters(true)
	for i, iocounter := range iocounters {
		if !perCpu {
			fmt.Printf("Total ")
		}
		fmt.Printf(formatString,
			i,
			iocounter.BytesRecv,
			iocounter.BytesSent,
			iocounter.Dropin,
			iocounter.Dropout,
			iocounter.Errin,
			iocounter.Errout,
			iocounter.Fifoin,
			iocounter.Fifoout,
			iocounter.Name,
			iocounter.PacketsRecv,
			iocounter.PacketsSent,
		)
	}
}

func demoNetConnections() {
	formatString := "connection[%d]: \n\tFamily: %d\n\tFd: %d\n\tLaddr: %+v\n\tPid: %d\n\tRaddr: %+v\n\tStatus: %s\n\tType: %d\n\tUids: %+v\n"
	connections, _ := net.Connections("all")
	for i, connection := range connections {
		fmt.Printf(formatString,
			i,
			connection.Family,
			connection.Fd,
			connection.Laddr,
			connection.Pid,
			connection.Raddr,
			connection.Status,
			connection.Type,
			connection.Uids,
		)
	}
}

func demoNetProtoCounters() {
	formatString := "protocounter[%d]: \n\tProtocol: %s\n\tStats:\n"
	protoCounterStatsFormatString := "\t\tStats[\"%s\"]: %d\n"
	protoCounters, _ := net.ProtoCounters([]string{})
	for i, protoCounter := range protoCounters {
		fmt.Printf(formatString,
			i,
			protoCounter.Protocol,
		)
		for key, value := range protoCounter.Stats {
			fmt.Printf(protoCounterStatsFormatString,
				key,
				value,
			)
		}
	}
}

func demoNetFilterCounters() {
	formatString := "filterCounter[%d]: \n\tConTrackCount: %d\n\tConTrackMaxt: %d\n"
	filterCounters, _ := net.FilterCounters()
	for i, filterCounter := range filterCounters {
		fmt.Printf(formatString,
			i,
			filterCounter.ConnTrackCount,
			filterCounter.ConnTrackMax,
		)
	}
}

func demoNetInterfaces() {
	formatString := "interface[%d]: \n\tAddrs: %+v\n\tFlags: %+v\n\tHardwareAddr: %s\n\tMTU: %d\n\tName: %s\n"
	interfaces, _ := net.Interfaces()
	for i, anInterface := range interfaces {
		fmt.Printf(formatString,
			i,
			anInterface.Addrs,
			anInterface.Flags,
			anInterface.HardwareAddr,
			anInterface.MTU,
			anInterface.Name,
		)
	}
}

func demoNetPids() {
	//  Bug reported: https://github.com/shirou/gopsutil/issues/410
	//	formatString := "pid[%d]: %d\n"
	//	pids, _ := net.Pids()
	//	for i, pid := range pids {
	//		fmt.Printf(formatString,
	//			i,
	//			pid,
	//		)
	//	}
}

func demoNet() {
	displayBanner("Net")
	demoNetIOCounters(true)
	demoNetIOCounters(false)
	demoNetConnections()
	demoNetProtoCounters()
	demoNetFilterCounters()
	demoNetInterfaces()
	demoNetPids()
}

func displayProcessRlimit(rLimit []process.RlimitStat) {
	formatString := "\trlimit[%d]: %+v\n"
	for i, value := range rLimit {
		fmt.Printf(formatString,
			i,
			value,
		)
	}
}

func displayProcessIOCounterStat(netIOCounters []net.IOCountersStat) {
	formatString := "\tnetIOCounters[%d]: %+v\n"
	for i, value := range netIOCounters {
		fmt.Printf(formatString,
			i,
			value,
		)
	}
}

func displayProcessMemoryMaps(memoryMaps *[]process.MemoryMapsStat) {
	if memoryMaps != nil {
		formatString := "\tmemoryMap[%d]: %+v\n"
		for i, value := range *memoryMaps {
			fmt.Printf(formatString,
				i,
				value,
			)
		}
	}
}

func demoProcessPids() {
	limit := 10 // Limit the number of processes printed.
	formatString := "pid[%d]: \n\tPid: %d\n\tcpuAffinity: %v\n\tcpuPercent: %f\n\tchildren: %+v\n\tcmdline: %s\n\tcmdlineSlice: %+v\n\tconnections: %+v\n\tcreateTime: %d\n\tcwd: %s\n\texe: %s\n\tgids: %+v\n\tioCounters: %+v\n\tioNice: %d\n\tisRunning: %t\n\tmemoryInfo: %+v\n\tmemoryInfoEx: %+v\n\tmemoryPercent: %f\n\tname: %s\n\tnice: %d\n\tnumCtxSwitches: %+v\n\tnumFds: %d\n\tnumThreads: %d\n\topenFiles: %+v\n\tparent: %+v\n\tpercent: %f\n\tpPid: %d\n\tstatus: %s\n\tterminal: %s \n\tthreads: %+v\n\ttimes: %+v\n\tuids: %+v\n\tusername: %s\n"
	pids, _ := process.Pids()
	for i, pid := range pids {
		if i >= limit {
			break
		}

		// Given a PID, create a process object.

		aProcess, _ := process.NewProcess(pid)

		// Extract data from process object.

		cpuAffinity, _ := aProcess.CPUAffinity()
		cpuPercent, _ := aProcess.CPUPercent()
		children, _ := aProcess.Children()
		cmdline, _ := aProcess.Cmdline()
		cmdlineSlice, _ := aProcess.CmdlineSlice()
		connections, _ := aProcess.Connections()
		createTime, _ := aProcess.CreateTime()
		cwd, _ := aProcess.Cwd()
		exe, _ := aProcess.Exe()
		gids, _ := aProcess.Gids()
		ioCounters, _ := aProcess.IOCounters()
		ioNice, _ := aProcess.IOnice()
		isRunning, _ := aProcess.IsRunning()
		// x := aProcess.Kill()
		memoryInfo, _ := aProcess.MemoryInfo()
		memoryInfoEx, _ := aProcess.MemoryInfoEx()
		memoryMapsTrue, _ := aProcess.MemoryMaps(true)
		memoryMapsFalse, _ := aProcess.MemoryMaps(false)
		memoryPercent, _ := aProcess.MemoryPercent()
		name, _ := aProcess.Name()
		netIoCountersTrue, _ := aProcess.NetIOCounters(true)
		netIoCountersFalse, _ := aProcess.NetIOCounters(false)
		nice, _ := aProcess.Nice()
		numCtxSwitches, _ := aProcess.NumCtxSwitches()
		numFds, _ := aProcess.NumFDs()
		numThreads, _ := aProcess.NumThreads()
		openFiles, _ := aProcess.OpenFiles()
		parent, _ := aProcess.Parent()
		percent, _ := aProcess.Percent(time.Microsecond)
		pPid, _ := aProcess.Ppid()
		// x := aProcess.Resume()
		rLimit, _ := aProcess.Rlimit()
		// x := aProcess.SendSignal(...)
		status, _ := aProcess.Status()
		// x := aProcess.Suspend()
		terminal, _ := aProcess.Terminal()
		// x := aProcess.Terminate()
		threads, _ := aProcess.Threads()
		times, _ := aProcess.Times()
		uids, _ := aProcess.Uids()
		username, _ := aProcess.Username()

		fmt.Printf(formatString,
			i,
			pid,
			cpuAffinity,
			cpuPercent,
			children,
			cmdline,
			cmdlineSlice,
			connections,
			createTime,
			cwd,
			exe,
			gids,
			ioCounters,
			ioNice,
			isRunning,
			memoryInfo,
			memoryInfoEx,
			memoryPercent,
			name,
			nice,
			numCtxSwitches,
			numFds,
			numThreads,
			openFiles,
			parent,
			percent,
			pPid,
			status,
			terminal,
			threads,
			times,
			uids,
			username,
		)
		displayProcessRlimit(rLimit)
		displayProcessIOCounterStat(netIoCountersFalse)
		displayProcessIOCounterStat(netIoCountersTrue)
		displayProcessMemoryMaps(memoryMapsFalse)
		displayProcessMemoryMaps(memoryMapsTrue)
	}
}

func demoProcess() {
	displayBanner("Process")
	demoProcessPids()
}

func main() {
	demoCpu()
	demoDisk()
	demoHost()
	demoLoad()
	demoMem()
	demoNet()
	demoProcess()
}
