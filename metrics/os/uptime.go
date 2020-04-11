package os

import (
	"fmt"
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strconv"
	"strings"
)

type UpTimeStat struct {
	UpTime TimeStat
	UpTimeString string
	IdleTime TimeStat
	IdleTimeString string
}
type TimeStat struct{
	Years int64
	Months int64
	Days int64
	Hours int64
	Minutes int64
	Seconds int64
	MilliSeconds int64
}
func timeStatToDisplayString(timeStat TimeStat) string{
	timeStatDisplay:=""
	CommaRequired := false
	if timeStat.Years != 0 {
		timeStatDisplay += fmt.Sprintf("%d years",timeStat.Years)
		CommaRequired = true
	}
	if timeStat.Months != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d months",timeStat.Months)
		CommaRequired = true
	}
	if timeStat.Days != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d days",timeStat.Days)
		CommaRequired = true
	}
	if timeStat.Hours != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d hours",timeStat.Hours)
		CommaRequired = true
	}
	if timeStat.Minutes != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d minutes",timeStat.Minutes)
		CommaRequired = true
	}
	if timeStat.Seconds != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d seconds",timeStat.Seconds)
		CommaRequired = true
	}
	if timeStat.MilliSeconds != 0 {
		if CommaRequired {
			timeStatDisplay += ", "
		}
		timeStatDisplay += fmt.Sprintf("%d milliseconds",timeStat.MilliSeconds)
		CommaRequired = true
	}
	return timeStatDisplay
}

func GetOsUpTime() *UpTimeStat {
	upTimeStat := new(UpTimeStat)
	defaultProcFolder :="/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/uptime")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		upTimeStat = nil
	}

	if len(lines) > 0 {
		dataLine := lines[0]
		fields := strings.Fields(dataLine)
		upTime,_:=strconv.ParseFloat(fields[0], 64)
		upTimeStat.UpTime= secondsToTimeStat(upTime)
		upTimeStat.UpTimeString=timeStatToDisplayString(upTimeStat.UpTime)
		idleTime,_:=strconv.ParseFloat(fields[1], 64)
		upTimeStat.IdleTime= secondsToTimeStat(idleTime)
		upTimeStat.IdleTimeString=timeStatToDisplayString(upTimeStat.IdleTime)
		}
	for i, line := range lines {
		fmt.Println(i, line)
	}
	return upTimeStat
}
func secondsToTimeStat(seconds float64) TimeStat{
	timeStat:=new(TimeStat)
	if seconds<0{
		timeStat=nil
	}
	timeStat.Years = (int64(seconds) / (12*31*24*60*60))
	seconds=seconds-float64((timeStat.Years * (12*31*24*60*60)))
	timeStat.Months = (int64(seconds) / (31*24*60*60))
	seconds=seconds-float64((timeStat.Months * (31*24*60*60)))
	timeStat.Days = (int64(seconds) / (24*60*60))
	seconds=seconds-float64((timeStat.Days * (24*60*60)))
	timeStat.Hours = (int64(seconds) / (60*60))
	seconds=seconds-float64((timeStat.Hours * (60*60)))
	timeStat.Minutes = (int64(seconds) / (60))
	seconds=seconds-float64((timeStat.Minutes * (60)))
	timeStat.Seconds = (int64(seconds))
	seconds=seconds-float64(timeStat.Seconds)
	timeStat.MilliSeconds=int64(seconds*1000)
	return *timeStat
}