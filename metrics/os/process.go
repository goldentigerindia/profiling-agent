package os

import (
	"fmt"
	"github.com/goldentigerindia/profiling-agent/config"
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
	ProcessId                      int64   //The process ID
	Command                        string  //The filename of the executable, in parentheses.	This is visible whether or not the executable is swapped out.
	State                          string  //One of the following characters, indicating process	state:	R  Running	S  Sleeping in an interruptible wait	D  Waiting in uninterruptible disk sleep	Z  Zombie	T  Stopped (on a signal) or (before Linux 2.6.33) trace stopped 	t  Tracing stop (Linux 2.6.33 onward) W  Paging (only before Linux 2.6.0)	X  Dead (from Linux 2.6.0 onward)	x  Dead (Linux 2.6.33 to 3.13 only)	K  Wakekill (Linux 2.6.33 to 3.13 only) 	W  Waking (Linux 2.6.33 to 3.13 only)	P  Parked (Linux 3.9 to 3.13 only)
	ParentProcessId                int64   //The PID of the parent of this process.
	ProcessGroupId                 int64   //The process group ID of the process.
	SessionId                      int64   //The session ID of the process.
	TerminalProcessId              int64   //The controlling terminal of the process.  (The minor	device number is contained in the combination of bits 31 to 20 and 7 to 0; the major device number is in bits 15 to 8.)
	TerminalProcessGroupId         int64   //The ID of the foreground process group of the controlling terminal of the process.
	KernelFlagWord                 int64   //he kernel flags word of the process.  For bit meanings, see the PF_* defines in the Linux kernel	source file include/linux/sched.h.  Details depend	on the kernel version.
	MinorFaults                    int64   //The number of minor faults the process has made which have not required loading a memory page from disk.
	ChildrenMinorFaultWaited       int64   //The number of minor faults that the process's waited-for children have made.
	MajorFaults                    int64   //The number of major faults the process has made which have required loading a memory page from disk.
	ChildrenMajorFaultWaited       int64   //The number of major faults that the process's waited-for children have made.
	UserModeTime                   int64   //Amount of time that this process has been scheduled in user mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).  This includes guest time, guest_time (time spent running a virtual CPU, see	below), so that applications that are not aware of the guest time field do not lose that time from	their calculations.
	SystemModeTime                 int64   //Amount of time that this process has been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	ChildrenUserModeWaitedTime     int64   //Amount of time that this process's waited-for children have been scheduled in user mode, measured in	clock ticks (divide by sysconf(_SC_CLK_TCK)).  (See	also times(2).)  This includes guest time, cguest_time (time spent running a virtual CPU, see below).
	ChildrenKernelModeWaitedTime   int64   //Amount of time that this process's waited-for children have been scheduled in kernel mode, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	Priority                       int64   //(Explanation for Linux 2.6) For processes running a real-time scheduling policy (policy below; see sched_setscheduler(2)), this is the negated scheduling priority, minus one; that is, a number in the range -2 to -100, corresponding to real-time priorities 1 to 99.  For processes running under a nonreal-time scheduling policy, this is the raw nice value (setpriority(2)) as represented in the kernel.The kernel stores nice values as numbers in the	range 0 (high) to 39 (low), corresponding to the user-visible nice range of -20 to 19. Before Linux 2.6, this was a scaled value based on the scheduler weighting given to this process.
	PriorityNice                   int64   //The nice value (see setpriority(2)), a value in the range 19 (low priority) to -20 (high priority).
	NumberOfThreads                int64   //Number of threads in this process (since Linux 2.6).	Before kernel 2.6, this field was hard coded to 0 as a placeholder for an earlier removed field.
	IntervalTimerValue             int64   //The time in jiffies before the next SIGALRM is sent to the process due to an interval timer.  Since kernel 2.6.17, this field is no longer maintained, and is hard coded as 0.
	StartTimeAfterSystemBoot       int64   //The time the process started after system boot.  In kernels before Linux 2.6, this value was expressed in jiffies.  Since Linux 2.6, the value is expressed in clock ticks (divide by sysconf(_SC_CLK_TCK)). The format for this field was %lu before Linux 2.6.
	VirtualMemorySizeInBytes       int64   //Virtual memory size in bytes
	ResidentSetSize                int64   //Resident Set Size: number of pages the process has in real memory.  This is just the pages which count toward text, data, or stack space.  This does not 	include pages which have not been demand-loaded in,	or which are swapped out.
	ResidentSetSizeLimit           int64   //Current soft limit in bytes on the rss of the process; see the description of RLIMIT_RSS in	getrlimit(2).
	CodeStartAddress               int64   //The address above which program text can run.
	CodeEndAddress                 int64   //The address below which program text can run.
	StartStackAddress              int64   //The address of the start (i.e., bottom) of the stack.
	KernelStackESP                 int64   //The current value of ESP (stack pointer), as found in the kernel stack page for the process.
	KernelStackEIP                 int64   //The current EIP (instruction pointer).
	PendingSignals                 int64   //The bitmap of pending signals, displayed as a decimal number.  Obsolete, because it does not provide	information on real-time signals; use /proc/[pid]/status instead.
	BlockedSignals                 int64   //The bitmap of blocked signals, displayed as a decimal number.  Obsolete, because it does not provide	information on real-time signals; use /proc/[pid]/status instead.
	IgnoredSignals                 int64   //The bitmap of ignored signals, displayed as a decimal number.  Obsolete, because it does not provide	information on real-time signals; use /proc/[pid]/status instead.
	CaughtSignals                  int64   //The bitmap of caught signals, displayed as a decimal	number.  Obsolete, because it does not provide information on real-time signals; use /proc/[pid]/status instead.
	WaitingChannel                 int64   //This is the "channel" in which the process is waiting.  It is the address of a location in the kernel where the process is sleeping.  The corresponding symbolic name can be found in /proc/[pid]/wchan.
	NumberOfPagesSwapped           int64   //Number of pages swapped (not maintained).
	CumilativeNumebrOfSwaps        int64   //Cumulative nswap for child processes (not maintained).
	ExitSignal                     int64   //Signal to be sent to parent when we die.
	LastExecutedCpuCore            int64   //CPU number last executed on.
	RealTimeSchedulingPriority     int64   //Real-time scheduling priority, a number in the range	1 to 99 for processes scheduled under a real-time policy, or 0, for non-real-time processes (see sched_setscheduler(2)).
	SchedulingPolicy               int64   //Scheduling policy (see sched_setscheduler(2)).Decode using the SCHED_* constants in linux/sched.h.
	BlockIODelaysCentiSecond       int64   //Aggregated block I/O delays, measured in clock ticks(centiseconds).
	GuestCPUTime                   int64   //Guest time of the process (time spent running a virtual CPU for a guest operating system), measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	ChildrenGuestCPUTime           int64   //Guest time of the process's children, measured in clock ticks (divide by sysconf(_SC_CLK_TCK)).
	ProgramBSSStartDataAddress     int64   //Address above which program initialized and uninitialized (BSS) data are placed.
	ProgramBSSEndDataAddress       int64   //Address below which program initialized and uninitialized (BSS) data are placed.
	ProgramHeapStartAddress        int64   //Address above which program heap can be expanded	with brk(2).
	ProgramArgumentStartAddress    int64   //Address above which program command-line arguments (argv) are placed.
	ProgramArgumentEndAddress      int64   //Address below program command-line arguments (argv) are placed.
	ProgramEnvironmentStartAddress int64   //Address above which program environment is placed.
	ProgramEnvironmentEndAddress   int64   //Address below which program environment is placed.
	ThreadExitCode                 int64   //The thread's exit status in the form reported by	waitpid(2).
	TotalTime                      int64   //UserModeTime+SystemModeTime+ChildrenUserModeWaitedTime
	Seconds                        int64   //Up Time - ( Start Time/ Clock Ticks)
	CpuUsage                       float64 //( 100 * ((Total Time / Clock Ticks) / Seconds) )
	ChildProcesses                 []*ProcessInfoStat
	Network                        *NetStat
	Memory	                       *ProcessStatM
}
type ProcessStatM struct {
	TotalProgramSize int64  //total program size (same as VmSize in /proc/[pid]/status)
	ResidentSetSize  int64  //resident set size (same as VmRSS in /proc/[pid]/status)
	SharedPages      int64  //number of resident shared pages (i.e., backed by a file) (same as RssFile+RssShmem in /proc/[pid]/status)
	Text             string //text (code)
	Lib              int64  //library (unused since Linux 2.6; always 0)
	Data             int64  //data + stack
	DirtyPages       int64  //dirty pages (unused since Linux 2.6; always 0)
}

func GetOSProcess() *ProcessStat {
	upTimeStat := GetOsUpTime()
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
	processIdMap := make(map[int64]*ProcessInfoStat)
	processIdMap[0] = new(ProcessInfoStat)
	processIdMap[0].ProcessId = 0
	processIdMap[0].Command = "ROOT"
	parentProcessIdMap := make(map[int64][]*ProcessInfoStat)
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
							processInfoStat.Network = GetOSNetwork(processInfoStat.ProcessId)
							processInfoStat.TotalTime = processInfoStat.UserModeTime + processInfoStat.SystemModeTime + processInfoStat.ChildrenUserModeWaitedTime
							processInfoStat.Seconds = int64(upTimeStat.UpTimeFloat) - (processInfoStat.StartTimeAfterSystemBoot / config.ApplicationConfig.CpuTicks)
							processInfoStat.CpuUsage = 100 * ((float64(processInfoStat.TotalTime) / float64(config.ApplicationConfig.CpuTicks)) / float64(processInfoStat.Seconds))
							processIdMap[processInfoStat.ProcessId] = processInfoStat
							childProcessInfo := []*ProcessInfoStat{}
							if parentProcessIdMap[processInfoStat.ParentProcessId] != nil {
								childProcessInfo = parentProcessIdMap[processInfoStat.ParentProcessId]
							}
							childProcessInfo = append(childProcessInfo, processInfoStat)
							parentProcessIdMap[processInfoStat.ParentProcessId] = childProcessInfo
						}
					}
					//stat.Processes = append(stat.Processes, *processInfoStat)

				}

			}
		}
	}
	for key, element := range parentProcessIdMap {
		if processIdMap[key] != nil {
			processIdMap[key].ChildProcesses = element
		} else {
			fmt.Println("processId : " + string(key) + "not found")
		}
	}
	if processIdMap[0] != nil {
		stat.Processes = append(stat.Processes, *processIdMap[0])
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
