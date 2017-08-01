package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

func demoCpu() {

	formatString := "infoStats[%d]: CPU: %9s Guest: %3.1f GuestNice: %3.1f Idle: %7.1f Iowait: %6.1f Irq: %3.1f Nice: %4.1f Softirq: %3.1f Steal: %3.1f Stolen: %3.1f System: %5.1f Total: %7.1f User: %6.1f\n"

	fmt.Printf("\n---------- %s ------------------------------\n\n", "demoCpu()")

	// Number of CPUs.

	cpuCount, _ := cpu.Counts(true)
	fmt.Printf("Number of CPUs: %d\n", cpuCount)

	// Per CPU statistics.

	timeStats, _ := cpu.Times(true)
	for cpuNum := range timeStats {
		fmt.Printf(formatString,
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
		fmt.Printf(formatString,
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
}

func main() {
	demoCpu()
}
