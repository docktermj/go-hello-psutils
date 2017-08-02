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

func demoCpu() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoCpu()")

	// Number of CPUs.

	cpuCount, _ := cpu.Counts(true)
	fmt.Printf("cpu.Counts(true): %d\n", cpuCount)
	cpuCount, _ = cpu.Counts(false)
	fmt.Printf("cpu.Counts(false): %d\n\n", cpuCount)

	// Per CPU statistics.

	timesFormatString := "timeStat[%d]: \n\tCPU: %s \n\tGuest: %f \n\tGuestNice: %f \n\tIdle: %f \n\tIowait: %f \n\tIrq: %f \n\tNice: %f \n\tSoftirq: %f \n\tSteal: %f \n\tStolen: %f \n\tSystem: %f \n\tTotal: %f \n\tUser: %f\n"

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

	timeStats, _ = cpu.Times(false)
	for i, timeStat := range timeStats {
		fmt.Printf("Total ")
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

	infoFormatString := "infoStat[%d]: \n\tCPU: %d \n\tCacheSize: %d \n\tCoreID: %s \n\tCores: %d \n\tFamily: %s \n\tFlags: %v \n\tMhz: %4.1f \n\tMicrocode: %s \n\tModel: %s \n\tModelName: %s \n\tPhysicalID: %s \n\tStepping: %d \n\tVendorID: %s\n"

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
	interval := time.Microsecond

	percents, _ := cpu.Percent(interval, true)
	for i, percent := range percents {
		fmt.Printf(percentFormatString,
			i,
			percent,
		)
	}

	percents, _ = cpu.Percent(interval, false)
	for i, percent := range percents {
		fmt.Printf("Total ")
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

	counterFormatString := "counter[%s]: \n\tIoTime: %d\n\tIopsInProgress: %d\n\tMergedReadCount: %d\n\tMergedWriteCount: %d\n\tName: %s\n\tReadBytes: %d\n\tReadCount: %d\n\tReadTime: %d\n\tSerialNumber: %s\n\tWeightedIO: %d\n\tWriteBytes: %d\n\tWriteCount: %d\n\tWriteTime: %d\n"

	counters, _ := disk.IOCounters("sda", "sdb")
	for key, counter := range counters {
		fmt.Printf(counterFormatString,
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

	//  Bug reported: https://github.com/shirou/gopsutil/issues/409
	//	kernelFormatString := "kernel[]: \n\tVersion: %s\n"
	//
	//	version, _ = host.KernelVersion()
	//	fmt.Printf(kernelFormatString,
	//		version,
	//	)

	// host.Virtualization()

	//  Bug reported: https://github.com/shirou/gopsutil/issues/411
	//	virtualizationFormatString := "virtualization[]: \n\tSystem: %s\n\tRole: %s\n"
	//
	//	system, role, _ := host.Virtualization()
	//	fmt.Printf(virtualizationFormatString,
	//		system,
	//		role,
	//	)

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

func demoLoad() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoLoad()")

	// load.Avg()

	averageFormatString := "average[]: \n\tLoad1: %f\n\tLoad5: %f\n\tLoad15: %f\n"

	average, _ := load.Avg()
	fmt.Printf(averageFormatString,
		average.Load1,
		average.Load5,
		average.Load15,
	)

	// load.Misc()

	miscFormatString := "misc[]: \n\tCtxt: %d\n\tProcsBlocked: %d\n\tProcsRunning: %d\n"

	misc, _ := load.Misc()
	fmt.Printf(miscFormatString,
		misc.Ctxt,
		misc.ProcsBlocked,
		misc.ProcsRunning,
	)

}

func demoMem() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoMem()")

	// mem.VirtualMemory()

	virtualFormatString := "virtual[]: \n\tActive: %d\n\tAvailable: %d\n\tBuffers: %d\n\tCached: %d\n\tDirty: %d\n\tFree: %d\n\tInactive: %d\n\tPageTables: %d\n\tShared: %d\n\tSlab: %d\n\tSwapCached: %d\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %f\n\tWired: %d\n\tWriteback: %d\n\tWritebackTmp: %d\n"

	virtual, _ := mem.VirtualMemory()
	fmt.Printf(virtualFormatString,
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

	// mem.SwapMemory()

	swapFormatString := "swap[]: \n\tFree: %d\n\tSin: %d\n\tSout: %d\n\tTotal: %d\n\tUsed: %d\n\tUsedPercent: %f\n"

	swap, _ := mem.SwapMemory()
	fmt.Printf(swapFormatString,
		swap.Free,
		swap.Sin,
		swap.Sout,
		swap.Total,
		swap.Used,
		swap.UsedPercent,
	)
}

func demoNet() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoNet()")

	// net.IOCounters()

	iocounterFormatString := "iocounter[%d]: \n\tBytesRecv: %d\n\tBytesSent: %d\n\tDropin: %d\n\tDropout: %d\n\tErrin: %d\n\tErrout: %d\n\tFifoin: %d\n\tFifoout: %d\n\tName: %s\n\tPacketsRecv: %d\n\tPacketsSent: %d\n"

	iocounters, _ := net.IOCounters(true)
	for i, iocounter := range iocounters {
		fmt.Printf(iocounterFormatString,
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

	iocounters, _ = net.IOCounters(false)
	for i, iocounter := range iocounters {
		fmt.Printf("Total ")
		fmt.Printf(iocounterFormatString,
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

	// net.Connections()

	connectionFormatString := "connection[%d]: \n\tFamily: %d\n\tFd: %d\n\tLaddr: %+v\n\tPid: %d\n\tRaddr: %+v\n\tStatus: %s\n\tType: %d\n\tUids: %+v\n"

	connections, _ := net.Connections("all")
	for i, connection := range connections {
		fmt.Printf(connectionFormatString,
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

	// net.ProtoCounters()

	protoCounterFormatString := "protocounter[%d]: \n\tProtocol: %s\n\tStats:\n"
	protoCounterStatsFormatString := "\t\tStats[\"%s\"]: %d\n"

	protoCounters, _ := net.ProtoCounters([]string{})
	for i, protoCounter := range protoCounters {
		fmt.Printf(protoCounterFormatString,
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

	// net.FilterCounters()

	filterCounterFormatString := "filterCounter[%d]: \n\tConTrackCount: %d\n\tConTrackMaxt: %d\n"

	filterCounters, _ := net.FilterCounters()
	for i, filterCounter := range filterCounters {
		fmt.Printf(filterCounterFormatString,
			i,
			filterCounter.ConnTrackCount,
			filterCounter.ConnTrackMax,
		)
	}

	// net.Interfaces()

	interfaceCounterFormatString := "interface[%d]: \n\tAddrs: %+v\n\tFlags: %+v\n\tHardwareAddr: %s\n\tMTU: %d\n\tName: %s\n"

	interfaces, _ := net.Interfaces()
	for i, anInterface := range interfaces {
		fmt.Printf(interfaceCounterFormatString,
			i,
			anInterface.Addrs,
			anInterface.Flags,
			anInterface.HardwareAddr,
			anInterface.MTU,
			anInterface.Name,
		)
	}

	// net.Pids()

	//  Bug reported: https://github.com/shirou/gopsutil/issues/410
	//	pidFormatString := "pid[%d]: %d\n"
	//
	//	pids, _ := net.Pids()
	//	for i, pid := range pids {
	//		fmt.Printf(pidFormatString,
	//			i,
	//			pid,
	//		)
	//	}

}

func demoProcess() {

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoProcess()")

	// process.Pids()

	processFormatString := "pid[%d]: \n\tPid: %d\n\tcpuAffinity: %v\n\tcpuPercent: %f\n\tchildren: %+v\n\tcmdline: %s\n\tcmdlineSlice: %+v\n\tconnections: %+v\n\tcreateTime: %d\n\tcwd: %s\n\texe: %s\n\tgids: %+v\n\tioCounters: %+v\n\tioNice: %d\n\tisRunning: %t\n\tmemoryInfo: %+v\n\tmemoryInfoEx: %+v\n\tmemoryMapsTrue: %+v\n\tmemoryMapsFalse: %+v\n\tmemoryPercent: %f\n\tname: %s\n\tnetIoCountersTrue: %+v\n\tnetIoCountersFalse: %+v\n\tnice: %d\n\tnumCtxSwitches: %+v\n\tnumFds: %d\n\tnumThreads: %d\n\topenFiles: %+v\n\tparent: %+v\n\tpercent: %f\n\tpPid: %d\n\trLimit: %+v\n\tstatus: %s\n\tterminal: %s \n\tthreads: %+v\n\ttimes: %+v\n\tuids: %+v\n\tusername: %s\n"

	pids, _ := process.Pids()
	limit := 10 // Limit the number of processes printed.
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
		//	    x := aProcess.Kill()
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
		//	    x := aProcess.Resume()
		rLimit, _ := aProcess.Rlimit()
		//	    x := aProcess.SendSignal(...)
		status, _ := aProcess.Status()
		//	    x := aProcess.Suspend()
		terminal, _ := aProcess.Terminal()
		//	    x := aProcess.Terminate()
		threads, _ := aProcess.Threads()
		times, _ := aProcess.Times()
		uids, _ := aProcess.Uids()
		username, _ := aProcess.Username()

		fmt.Printf(processFormatString,
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
			memoryMapsTrue,
			memoryMapsFalse,
			memoryPercent,
			name,
			netIoCountersTrue,
			netIoCountersFalse,
			nice,
			numCtxSwitches,
			numFds,
			numThreads,
			openFiles,
			parent,
			percent,
			pPid,
			rLimit,
			status,
			terminal,
			threads,
			times,
			uids,
			username,
		)
	}
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
