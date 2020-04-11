package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type DiskStat struct{
	Disks []DeviceStat
}
type DeviceStat struct{
	MajorNumber int64
	MinorNumber int64
	DeviceName string
	ReadsCompleted int64
	ReadsMerged int64
	SectorsRead int64
	TimeSpendReadingInMillis int64
	WritesCompleted int64
	WritesMerged int64
	SectorsWritten int64
	TimeSpendWritingInMills int64
	NumberOfIOCurrentInProgress int64
	TimeSpendDoingIOInMills int64
	WeightedTimeSpendingIOInMills int64
	DiscardCompleted int64
	DiscardsMerged int64
	SectorsDiscarded int64
	TimeSpendDiscardingInMills int64
	FlushCompleted int64
	TimeSpendFlushingInMills int64
}
func GetOSDisk() *DiskStat {
	stat := new(DiskStat)

	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/diskstats")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		stat = nil

	} else {

		if len(lines) > 0 {
			stat.Disks = []DeviceStat{}
			for i := 0; i < len(lines); i++ {
				dataLine := lines[i]
				fields := strings.Fields(dataLine)
				deviceStat := new(DeviceStat)
				deviceStat.MajorNumber,_= strconv.ParseInt(fields[0], 10, 64)
				deviceStat.MinorNumber,_= strconv.ParseInt(fields[1], 10, 64)
				deviceStat.DeviceName=fields[2]
				deviceStat.ReadsCompleted,_= strconv.ParseInt(fields[3], 10, 64)
				deviceStat.ReadsMerged,_= strconv.ParseInt(fields[4], 10, 64)
				deviceStat.SectorsRead,_= strconv.ParseInt(fields[5], 10, 64)
				deviceStat.TimeSpendReadingInMillis,_= strconv.ParseInt(fields[6], 10, 64)
				deviceStat.WritesCompleted,_= strconv.ParseInt(fields[7], 10, 64)
				deviceStat.WritesMerged,_= strconv.ParseInt(fields[8], 10, 64)
				deviceStat.SectorsWritten,_= strconv.ParseInt(fields[9], 10, 64)
				deviceStat.TimeSpendWritingInMills,_= strconv.ParseInt(fields[10], 10, 64)
				deviceStat.NumberOfIOCurrentInProgress,_= strconv.ParseInt(fields[11], 10, 64)
				deviceStat.TimeSpendDoingIOInMills,_= strconv.ParseInt(fields[12], 10, 64)
				deviceStat.WeightedTimeSpendingIOInMills,_= strconv.ParseInt(fields[13], 10, 64)
				if len(fields)>17{
					deviceStat.DiscardCompleted,_= strconv.ParseInt(fields[14], 10, 64)
					deviceStat.DiscardsMerged,_= strconv.ParseInt(fields[15], 10, 64)
					deviceStat.SectorsDiscarded,_= strconv.ParseInt(fields[16], 10, 64)
					deviceStat.TimeSpendDiscardingInMills,_= strconv.ParseInt(fields[17], 10, 64)
				}
				if len(fields)>19{
					deviceStat.FlushCompleted,_= strconv.ParseInt(fields[18], 10, 64)
					deviceStat.TimeSpendFlushingInMills,_= strconv.ParseInt(fields[19], 10, 64)
				}
				stat.Disks = append(stat.Disks, *deviceStat)

			}
		}
	}
	return stat
}