package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type NetStat struct{
	Interfaces []InterfaceStat
}
type InterfaceStat struct{
	InterfaceName string
	Receive ReceiveStat
	Transmit TransmitStat
}
type ReceiveStat struct{
	Bytes int64
	Packets int64
	Errors int64
	Drops int64
	Fifo int64
	Frame int64
	Compressed int64
	Multicast int64
}
type TransmitStat struct{
	Bytes int64
	Packets int64
	Errors int64
	Drops int64
	Fifo int64
	Collision int64
	Carrier int64
	Compressed int64
}

func GetOSNetwork() *NetStat {
	stat := new(NetStat)

	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/net/dev")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		stat = nil

	} else {

		if len(lines) > 0 {
			stat.Interfaces = []InterfaceStat{}
			for i := 2; i < len(lines); i++ {
				dataLine := lines[i]

				fields := strings.Fields(dataLine)
				interfaceStat := new(InterfaceStat)
				interfaceStat.InterfaceName=strings.ReplaceAll(fields[0],":","")
				receiveStat:=new(ReceiveStat)
				receiveStat.Bytes,_= strconv.ParseInt(fields[1], 10, 64)
				receiveStat.Packets,_= strconv.ParseInt(fields[2], 10, 64)
				receiveStat.Errors,_= strconv.ParseInt(fields[3], 10, 64)
				receiveStat.Drops,_= strconv.ParseInt(fields[4], 10, 64)
				receiveStat.Fifo,_= strconv.ParseInt(fields[5], 10, 64)
				receiveStat.Frame,_= strconv.ParseInt(fields[6], 10, 64)
				receiveStat.Compressed,_= strconv.ParseInt(fields[7], 10, 64)
				receiveStat.Multicast,_= strconv.ParseInt(fields[8], 10, 64)
				interfaceStat.Receive=*receiveStat
				transmitStat:=new(TransmitStat)
				transmitStat.Bytes,_= strconv.ParseInt(fields[9], 10, 64)
				transmitStat.Packets,_= strconv.ParseInt(fields[10], 10, 64)
				transmitStat.Errors,_= strconv.ParseInt(fields[11], 10, 64)
				transmitStat.Drops,_= strconv.ParseInt(fields[12], 10, 64)
				transmitStat.Fifo,_= strconv.ParseInt(fields[13], 10, 64)
				transmitStat.Collision,_= strconv.ParseInt(fields[14], 10, 64)
				transmitStat.Carrier,_= strconv.ParseInt(fields[15], 10, 64)
				transmitStat.Compressed,_= strconv.ParseInt(fields[16], 10, 64)
				interfaceStat.Transmit=*transmitStat
				stat.Interfaces = append(stat.Interfaces, *interfaceStat)

			}
		}
	}
	return stat
}