package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type SwapStat struct {
	SwapFileName string
	SwapType string
	SwapSizeInKBs int64
	SwapUsedInKBs int64
	SwapPriority int64
}
func GetOSSwap() []*SwapStat {
	swapStats := []*SwapStat{}
	defaultProcFolder :="/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/swaps")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		swapStats = nil

	}

	if len(lines) > 1 {
		for i := 1; i < len(lines); i++ {
			dataLine := lines[i]
			fields := strings.Fields(dataLine)
			swapStat := new(SwapStat)
			swapStat.SwapFileName = fields[0]
			swapStat.SwapType = fields[1]
			swapStat.SwapSizeInKBs, _ = strconv.ParseInt(fields[2], 10, 64)
			swapStat.SwapUsedInKBs, _ = strconv.ParseInt(fields[3], 10, 64)
			swapStat.SwapPriority, _ = strconv.ParseInt(fields[4], 10, 64)
			swapStats = append(swapStats, swapStat)
		}
	}
	return swapStats
}
