package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type MemStat struct {
	Stats []MemMetric
}
type MemMetric struct {
	Metric string
	Value  int64
	Unit   string
}

func GetOSMem() *MemStat {
	stat := new(MemStat)

	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/meminfo")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		stat = nil

	} else {

		if len(lines) > 0 {
			stat.Stats = []MemMetric{}
			for i := 0; i < len(lines); i++ {
				dataLine := lines[i]

				fields := strings.Fields(dataLine)
				memStat := new(MemMetric)
				memStat.Metric = strings.ReplaceAll(fields[0],":","")
				memStat.Value, _ = strconv.ParseInt(fields[1], 10, 64)
				if len(fields)==3 {
					memStat.Unit = fields[2]
				}
				stat.Stats = append(stat.Stats, *memStat)

			}
		}
	}
	return stat
}
