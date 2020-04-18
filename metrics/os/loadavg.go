package os

import (
	"fmt"
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type LoadAvgStat struct {
	LoadAvg1Minute float64
	LoadAvg5Minute float64
	LoadAvg15Minute float64
	ProcessAverage ProcessAvg
	LastProcessId	int64
}
type ProcessAvg struct {
	CurrentlyRunningProcessCount int64
	TotalProcessCount int64
}
func GetOSLoadAvg() *LoadAvgStat {
	loadAvgStat := new(LoadAvgStat)
	defaultProcFolder :="/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/loadavg")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		loadAvgStat = nil

	}

	if len(lines) > 0 {
		dataLine := lines[0]
		fields := strings.Fields(dataLine)
		loadAvgStat.LoadAvg1Minute, _ = strconv.ParseFloat(fields[0], 64)
		loadAvgStat.LoadAvg5Minute, _ = strconv.ParseFloat(fields[1], 64)
		loadAvgStat.LoadAvg15Minute, _ = strconv.ParseFloat(fields[2], 64)
		processAvg := new(ProcessAvg)
		processFields := strings.Split(fields[3], "/")
		if len(processFields) == 2 {
			processAvg.CurrentlyRunningProcessCount, _ = strconv.ParseInt(processFields[0], 10, 64)
			processAvg.TotalProcessCount, _ = strconv.ParseInt(processFields[1], 10, 64)
		}
		loadAvgStat.ProcessAverage = *processAvg
		loadAvgStat.LastProcessId, _ = strconv.ParseInt(fields[4], 10, 64)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
	return loadAvgStat
}
