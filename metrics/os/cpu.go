package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type CpuStat struct {
	Cores []CoreStat
}
type CoreStat struct {
	Core        string
	User        int64
	Nice        int64
	System      int64
	Idle        int64
	Iowait      int64
	HardIrq     int64
	Softirq     int64
	Steal       int64
	Guest       int64
	GuestNice   int64
	UsedTotal   int64
	Total       int64
	Percentages *CoreStatPercentage
}
type CoreStatPercentage struct {
	User                float64
	Nice                float64
	System              float64
	Idle                float64
	Iowait              float64
	HardIrq             float64
	Softirq             float64
	Steal               float64
	Guest               float64
	GuestNice           float64
	UsedTotal           float64
	Total               float64
}

func GetOSCpuStat() *CpuStat {
	stat := new(CpuStat)

	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/stat")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		stat = nil

	} else {

		if len(lines) > 0 {
			stat.Cores = []CoreStat{}
			for i := 0; i < len(lines); i++ {
				dataLine := lines[i]
				if strings.HasPrefix(dataLine, "cpu") {
					fields := strings.Fields(dataLine)
					cpuStat := new(CoreStat)
					cpuStat.Core = fields[0]
					cpuStat.User, _ = strconv.ParseInt(fields[1], 10, 64)
					cpuStat.Nice, _ = strconv.ParseInt(fields[2], 10, 64)
					cpuStat.System, _ = strconv.ParseInt(fields[3], 10, 64)
					cpuStat.Idle, _ = strconv.ParseInt(fields[4], 10, 64)
					cpuStat.Iowait, _ = strconv.ParseInt(fields[5], 10, 64)
					cpuStat.HardIrq, _ = strconv.ParseInt(fields[6], 10, 64)
					cpuStat.Softirq, _ = strconv.ParseInt(fields[7], 10, 64)
					cpuStat.Steal, _ = strconv.ParseInt(fields[8], 10, 64)
					cpuStat.Guest, _ = strconv.ParseInt(fields[9], 10, 64)
					cpuStat.GuestNice, _ = strconv.ParseInt(fields[10], 10, 64)
					cpuStat.UsedTotal = cpuStat.User + cpuStat.Nice + cpuStat.System + cpuStat.Iowait + cpuStat.HardIrq + cpuStat.Softirq + cpuStat.Steal + cpuStat.Guest + cpuStat.GuestNice
					cpuStat.Total = cpuStat.UsedTotal + cpuStat.Idle
					cpuStat.Percentages = calculateCpuPercentageStats(cpuStat)
					stat.Cores = append(stat.Cores, *cpuStat)
				}
			}
		}
	}
	return stat
}
func calculateCpuPercentageStats(cpuStat *CoreStat) *CoreStatPercentage {
	cpuStatPercentage := new(CoreStatPercentage)
	if cpuStat.User >0 {
		cpuStatPercentage.User = (float64(cpuStat.User)/float64(cpuStat.Total) ) * float64(100)
	}else{
		cpuStatPercentage.User =0
	}
	if cpuStat.Nice >0 {
		cpuStatPercentage.Nice = (float64(cpuStat.Nice)/float64(cpuStat.Total) ) * float64(100)
	}else{
		cpuStatPercentage.Nice =0
	}
	if cpuStat.System >0 {
		cpuStatPercentage.System = (float64(cpuStat.System)/float64(cpuStat.Total)) * float64(100)
	}else{
		cpuStatPercentage.System =0
	}
	if cpuStat.Idle >0{
		cpuStatPercentage.Idle = (float64(cpuStat.Idle)/float64(cpuStat.Total) ) * float64(100)
	}else{
		cpuStatPercentage.Idle =0
	}
	if cpuStat.Iowait >0{
		cpuStatPercentage.Iowait = (float64(cpuStat.Iowait)/float64(cpuStat.Total)  ) * float64(100)
	}else{
		cpuStatPercentage.Iowait =0
	}
	if cpuStat.HardIrq >0 {
		cpuStatPercentage.HardIrq = (float64(cpuStat.HardIrq)/float64(cpuStat.Total) ) * float64(100)
	}else{
		cpuStatPercentage.HardIrq =0
	}
	if cpuStat.Softirq >0 {
		cpuStatPercentage.Softirq = (float64(cpuStat.Softirq)/float64(cpuStat.Total) ) * float64(100)
	}else{
		cpuStatPercentage.Softirq =0
	}
	if cpuStat.Steal >0 {
		cpuStatPercentage.Steal = (float64(cpuStat.Steal)/float64(cpuStat.Total)) * float64(100)
	}else{
		cpuStatPercentage.Steal =0
	}
	if cpuStat.Guest >0 {
		cpuStatPercentage.Guest = (float64(cpuStat.Guest)/float64(cpuStat.Total) ) * 100
	}else{
		cpuStatPercentage.Guest =0
	}
	if cpuStat.GuestNice >0 {
		cpuStatPercentage.GuestNice = (float64(cpuStat.GuestNice)/float64(cpuStat.Total) ) * 100
	}else{
		cpuStatPercentage.GuestNice =0
	}
	if cpuStat.UsedTotal >0 {
		cpuStatPercentage.UsedTotal = (float64(cpuStat.UsedTotal) / float64(cpuStat.Total)) * 100
	}else{
		cpuStatPercentage.UsedTotal =0
	}
	if cpuStat.Total >0 {
		cpuStatPercentage.Total = (float64(cpuStat.Total) / float64(cpuStat.Total)) * 100
	}else{
		cpuStatPercentage.Total=0
	}
	return cpuStatPercentage
}
