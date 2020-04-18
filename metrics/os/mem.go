package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strings"
)

type MemStat struct {
	Entries map[string]string
}

func GetOSMem() *MemStat {
	stat := new(MemStat)
	stat.Entries = make(map[string]string)
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/meminfo")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		stat = nil

	} else {

		if len(lines) > 0 {
			for i := 0; i < len(lines); i++ {
				dataLine := lines[i]

				fields := strings.SplitN(dataLine, ":", -1)
				if fields != nil && len(fields) > 1 && len(strings.TrimSpace(fields[0])) > 0 && len(strings.TrimSpace(fields[1])) > 0 {
					stat.Entries[strings.TrimSpace(fields[0])] = strings.TrimSpace(fields[1])
				}
			}
		}
	}
	return stat
}
