package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
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

	timesFormatString := "timeStats[%d]: CPU: %-9s Guest: %3.1f GuestNice: %3.1f Idle: %7.1f Iowait: %6.1f Irq: %3.1f Nice: %4.1f Softirq: %3.1f Steal: %3.1f Stolen: %3.1f System: %5.1f Total: %8.1f User: %6.1f\n"

	timeStats, _ := cpu.Times(true)
	for cpuNum := range timeStats {
		fmt.Printf(timesFormatString,
			cpuNum,
			timeStats[cpuNum].CPU,
			timeStats[cpuNum].Guest,
			timeStats[cpuNum].GuestNice,
			timeStats[cpuNum].Idle,
			timeStats[cpuNum].Iowait,
			timeStats[cpuNum].Irq,
			timeStats[cpuNum].Nice,
			timeStats[cpuNum].Softirq,
			timeStats[cpuNum].Steal,
			timeStats[cpuNum].Stolen,
			timeStats[cpuNum].System,
			timeStats[cpuNum].Total(),
			timeStats[cpuNum].User,
		)
	}

	// Total statistics.  Note "false" in cpu.Times(false)

	fmt.Printf("\nCPU total:\n")

	timeStats, _ = cpu.Times(false)
	for cpuNum := range timeStats {
		fmt.Printf(timesFormatString,
			cpuNum,
			timeStats[cpuNum].CPU,
			timeStats[cpuNum].Guest,
			timeStats[cpuNum].GuestNice,
			timeStats[cpuNum].Idle,
			timeStats[cpuNum].Iowait,
			timeStats[cpuNum].Irq,
			timeStats[cpuNum].Nice,
			timeStats[cpuNum].Softirq,
			timeStats[cpuNum].Steal,
			timeStats[cpuNum].Stolen,
			timeStats[cpuNum].System,
			timeStats[cpuNum].Total(),
			timeStats[cpuNum].User,
		)
	}

	// cpu.Info()

	fmt.Printf("\nCPU Info:\n")

	infoFormatString := "infoStats[%d]: \n\tCPU: %d \n\tCacheSize: %d \n\tCoreID: %s \n\tCores: %d \n\tFamily: %s \n\tFlags: %v \n\tMhz: %4.1f \n\tMicrocode: %s \n\tModel: %s \n\tModelName: %s \n\tPhysicalID: %s \n\tStepping: %d \n\tVendorID: %s\n"

	infoStats, _ := cpu.Info()
	for cpuNum := range timeStats {
		fmt.Printf(infoFormatString,
			cpuNum,
			infoStats[cpuNum].CPU,
			infoStats[cpuNum].CacheSize,
			infoStats[cpuNum].CoreID,
			infoStats[cpuNum].Cores,
			infoStats[cpuNum].Family,
			infoStats[cpuNum].Flags,
			infoStats[cpuNum].Mhz,
			infoStats[cpuNum].Microcode,
			infoStats[cpuNum].Model,
			infoStats[cpuNum].ModelName,
			infoStats[cpuNum].PhysicalID,
			infoStats[cpuNum].Stepping,
			infoStats[cpuNum].VendorID,
		)
	}

	// cpu.Percent()

	percentFormatString := "percent[%d]: %f\n"
	interval := time.Second

	fmt.Printf("\nPercents(%d, true):\n", interval)
	percents, _ := cpu.Percent(interval, true)
	for cpuNum := range timeStats {
		fmt.Printf(percentFormatString,
			cpuNum,
			percents[cpuNum],
		)
	}

	fmt.Printf("\nPercents(%d, false):\n", interval)
	percents, _ = cpu.Percent(interval, false)
	for cpuNum := range timeStats {
		fmt.Printf(percentFormatString,
			cpuNum,
			percents[cpuNum],
		)
	}

}

func main() {
	demoCpu()
}
