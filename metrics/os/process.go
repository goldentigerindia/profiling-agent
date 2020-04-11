package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"os"
	"strconv"
	"strings"
)

type ProcessStat struct {
	Processes []ProcessInfoStat
}
type ProcessInfoStat struct {
	ProcessId                      int64
	Command                        string
	State                          string
	ParentProcessId                int64
	ProcessGroupId                 int64
	SessionId                      int64
	TerminalProcessId              int64
	TerminalProcessGroupId         int64
	KernelFlagWord                 int64
	MinorFaults                    int64
	ChildrenMinorFaultWaited       int64
	MajorFaults                    int64
	ChildrenMajorFaultWaited       int64
	UserModeTime                   int64
	SystemModeTime                 int64
	ChildrenUserModeWaitedTime     int64
	ChildrenKernelModeWaitedTime   int64
	Priority                       int64
	PriorityNice                   int64
	NumberOfThreads                int64
	IntervalTimerValue             int64
	StartTimeAfterSystemBoot       int64
	VirtualMemorySizeInBytes       int64
	ResidentSetSize                int64
	ResidentSetSizeLimit           int64
	CodeStartAddress               int64
	CodeEndAddress                 int64
	StartStackAddress              int64
	KernelStackESP                 int64
	KernelStackEIP                 int64
	PendingSignals                 int64
	BlockedSignals                 int64
	IgnoredSignals                 int64
	CaughtSignals                  int64
	WaitingChannel                 int64
	NumberOfPagesSwapped           int64
	CumilativeNumebrOfSwaps        int64
	ExitSignal                     int64
	LastExecutedCpuCore            int64
	RealTimeSchedulingPriority     int64
	SchedulingPolicy               int64
	BlockIODelaysCentiSecond       int64
	GuestCPUTime                   int64
	ChildrenGuestCPUTime           int64
	ProgramBSSStartDataAddress     int64
	ProgramBSSEndDataAddress       int64
	ProgramHeapStartAddress        int64
	ProgramArgumentStartAddress    int64
	ProgramArgumentEndAddress      int64
	ProgramEnvironmentStartAddress int64
	ProgramEnvironmentEndAddress   int64
	ThreadExitCode                 int64
}

func GetOSProcess() *ProcessStat {
	stat := new(ProcessStat)
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	f, err := os.Open(procPath)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if (file.IsDir()) {
			if _, err := strconv.ParseInt(file.Name(), 10, 64); err == nil {

				lines, err := util.ReadLines(procPath + "/" + file.Name() + "/stat")
				if err != nil {
					log.Fatalf("readLines: %s", err)
					stat = nil

				} else {
					processInfoStat := new(ProcessInfoStat)
					processInfoStat.ProcessId, _ = strconv.ParseInt(file.Name(), 10, 64)
					if len(lines) > 0 {
						for i := 0; i < len(lines); i++ {
							dataLine := lines[i]
							fields := strings.Fields(dataLine)
							processInfoStat.ProcessId, _ = strconv.ParseInt(fields[0], 10, 64)
							processInfoStat.Command = fields[1]
							processInfoStat.State = getStateName(fields[2])
							processInfoStat.ParentProcessId, _ = strconv.ParseInt(fields[3], 10, 64)
							processInfoStat.ProcessGroupId, _ = strconv.ParseInt(fields[4], 10, 64)
							processInfoStat.SessionId, _ = strconv.ParseInt(fields[5], 10, 64)
							processInfoStat.TerminalProcessId, _ = strconv.ParseInt(fields[6], 10, 64)
							processInfoStat.TerminalProcessGroupId, _ = strconv.ParseInt(fields[7], 10, 64)
							processInfoStat.KernelFlagWord, _ = strconv.ParseInt(fields[8], 10, 64)
							processInfoStat.MinorFaults, _ = strconv.ParseInt(fields[9], 10, 64)
							processInfoStat.ChildrenMinorFaultWaited, _ = strconv.ParseInt(fields[10], 10, 64)
							processInfoStat.MajorFaults, _ = strconv.ParseInt(fields[11], 10, 64)
							processInfoStat.ChildrenMajorFaultWaited, _ = strconv.ParseInt(fields[12], 10, 64)
							processInfoStat.UserModeTime, _ = strconv.ParseInt(fields[13], 10, 64)
							processInfoStat.SystemModeTime, _ = strconv.ParseInt(fields[14], 10, 64)
							processInfoStat.ChildrenUserModeWaitedTime, _ = strconv.ParseInt(fields[15], 10, 64)
							processInfoStat.ChildrenKernelModeWaitedTime, _ = strconv.ParseInt(fields[16], 10, 64)
							processInfoStat.Priority, _ = strconv.ParseInt(fields[17], 10, 64)
							processInfoStat.PriorityNice, _ = strconv.ParseInt(fields[18], 10, 64)
							processInfoStat.NumberOfThreads, _ = strconv.ParseInt(fields[19], 10, 64)
							processInfoStat.IntervalTimerValue, _ = strconv.ParseInt(fields[20], 10, 64)
							processInfoStat.StartTimeAfterSystemBoot, _ = strconv.ParseInt(fields[21], 10, 64)
							processInfoStat.VirtualMemorySizeInBytes, _ = strconv.ParseInt(fields[22], 10, 64)
							processInfoStat.ResidentSetSize, _ = strconv.ParseInt(fields[23], 10, 64)
							processInfoStat.ResidentSetSizeLimit, _ = strconv.ParseInt(fields[24], 10, 64)
							processInfoStat.CodeStartAddress, _ = strconv.ParseInt(fields[25], 10, 64)
							processInfoStat.CodeEndAddress, _ = strconv.ParseInt(fields[26], 10, 64)
							processInfoStat.StartStackAddress, _ = strconv.ParseInt(fields[27], 10, 64)
							processInfoStat.KernelStackESP, _ = strconv.ParseInt(fields[28], 10, 64)
							processInfoStat.KernelStackEIP, _ = strconv.ParseInt(fields[29], 10, 64)
							processInfoStat.PendingSignals, _ = strconv.ParseInt(fields[30], 10, 64)
							processInfoStat.BlockedSignals, _ = strconv.ParseInt(fields[31], 10, 64)
							processInfoStat.IgnoredSignals, _ = strconv.ParseInt(fields[32], 10, 64)
							processInfoStat.CaughtSignals, _ = strconv.ParseInt(fields[33], 10, 64)
							processInfoStat.WaitingChannel, _ = strconv.ParseInt(fields[34], 10, 64)
							processInfoStat.NumberOfPagesSwapped, _ = strconv.ParseInt(fields[35], 10, 64)
							processInfoStat.CumilativeNumebrOfSwaps, _ = strconv.ParseInt(fields[36], 10, 64)
							processInfoStat.ExitSignal, _ = strconv.ParseInt(fields[37], 10, 64)
							processInfoStat.LastExecutedCpuCore, _ = strconv.ParseInt(fields[38], 10, 64)
							processInfoStat.RealTimeSchedulingPriority, _ = strconv.ParseInt(fields[39], 10, 64)
							processInfoStat.SchedulingPolicy, _ = strconv.ParseInt(fields[40], 10, 64)
							processInfoStat.BlockIODelaysCentiSecond, _ = strconv.ParseInt(fields[41], 10, 64)
							processInfoStat.GuestCPUTime, _ = strconv.ParseInt(fields[42], 10, 64)
							processInfoStat.ChildrenGuestCPUTime, _ = strconv.ParseInt(fields[43], 10, 64)
							processInfoStat.ProgramBSSStartDataAddress, _ = strconv.ParseInt(fields[44], 10, 64)
							processInfoStat.ProgramBSSEndDataAddress, _ = strconv.ParseInt(fields[45], 10, 64)
							processInfoStat.ProgramHeapStartAddress, _ = strconv.ParseInt(fields[46], 10, 64)
							processInfoStat.ProgramArgumentStartAddress, _ = strconv.ParseInt(fields[47], 10, 64)
							processInfoStat.ProgramArgumentEndAddress, _ = strconv.ParseInt(fields[48], 10, 64)
							processInfoStat.ProgramEnvironmentStartAddress, _ = strconv.ParseInt(fields[49], 10, 64)
							processInfoStat.ProgramEnvironmentEndAddress, _ = strconv.ParseInt(fields[50], 10, 64)
							processInfoStat.ThreadExitCode, _ = strconv.ParseInt(fields[51], 10, 64)

						}
					}
					stat.Processes = append(stat.Processes, *processInfoStat)
				}

			}
		}
	}
	return stat
}
func getStateName(stateChar string) string {
	state := ""
	switch stateChar {
	case "R":
		state = "Running"
	case "I":
		state = "Interrupted"
	case "S":
		state = "Sleeping in an interruptible wait"
	case "D":
		state = "Waiting in uninterruptible disk sleep"
	case "Z":
		state = "Zombie"
	case "T":
		state = "Stopped"
	case "t":
		state = "Tracing stop"
	case "X":
		state = "Dead"
	case "x":
		state = "Dead"
	case "K":
		state = "Wakekill"
	case "W":
		state = "Waking"
	case "P":
		state = "Parked"
	}
	return state
}
