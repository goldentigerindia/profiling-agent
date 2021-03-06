package os

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/goldentigerindia/profiling-agent/config"
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)
type TcpType string
const(
	TCP4 TcpType = "TCP4"
	TCP6 ="TCP6"
)

type ProcessStat struct {
	ProcessTree []*ProcessInfoStat
	Top         []*TopStat
}

/*
http://man7.org/linux/man-pages/man5/proc.5.html
$ head -1 /proc/24784/net/tcp; grep 15907701 /proc/24784/net/tcp
  sl  local_address rem_address   st  tx_queue  rx_queue tr tm->when  retrnsmt   uid  timeout inode
  46: 010310AC:9C4C 030310AC:1770 01 0100000150:00000000  01:00000019 00000000  1000 0 54165785 4 cd1e6040 25 4 27 3 -1

46: 010310AC:9C4C 030310AC:1770 01
|   |         |   |        |    |--> connection state
|   |         |   |        |------> remote TCP port number
|   |         |   |-------------> remote IPv4 address
|   |         |--------------------> local TCP port number
|   |---------------------------> local IPv4 address
|----------------------------------> number of entry

00000150:00000000 01:00000019 00000000
|        |        |  |        |--> number of unrecovered RTO timeouts
|        |        |  |----------> number of jiffies until timer expires
|        |        |----------------> timer_active (see below)
|        |----------------------> receive-queue
|-------------------------------> transmit-queue

1000 0 54165785 4 cd1e6040 25 4 27 3 -1
|    | |        | |        |  | |  |  |--> slow start size threshold,
|    | |        | |        |  | |  |       or -1 if the treshold
|    | |        | |        |  | |  |       is >= 0xFFFF
|    | |        | |        |  | |  |----> sending congestion window
|    | |        | |        |  | |-------> (ack.quick<<1)|ack.pingpong
|    | |        | |        |  |---------> Predicted tick of soft clock
|    | |        | |        |               (delayed ACK control data)
|    | |        | |        |------------> retransmit timeout
|    | |        | |------------------> location of socket in memory
|    | |        |-----------------------> socket reference count
|    | |-----------------------------> inode
|    |----------------------------------> unanswered 0-window probes
|---------------------------------------------> uid
timer_active:
  0  no timer is pending
  1  retransmit-timer is pending
  2  another timer (e.g. delayed ack or keepalive) is pending
  3  this is a socket in TIME_WAIT state. Not all fields will contain
     data (or even exist)
  4  zero window probe timer is pending
package main

import "fmt"
import "encoding/hex"
func main() {
    z:="0A010100"
    a,_:=hex.DecodeString(z)
    fmt.Printf("%v.%v.%v.%v",a[0],a[1],a[2],a[3])

 }
enum {
    TCP_ESTABLISHED = 1,
    TCP_SYN_SENT,
    TCP_SYN_RECV,
    TCP_FIN_WAIT1,
    TCP_FIN_WAIT2,
    TCP_TIME_WAIT,
    TCP_CLOSE,
    TCP_CLOSE_WAIT,
    TCP_LAST_ACK,
    TCP_LISTEN,
	TCP_MAX_STATES  Leave at the end!
};

sudo cat /proc/1/ns/net
 */
type Tcp struct {
	SequenceId                       int64
	LocalIpAddress                   string
	LocalPortNumber                  int64
	RemoteIpAddress                  string
	RemotePortNumber                 int64
	ConnectionState                  string
	TransmitQueue                    string
	ReceiveQueue                     string
	TimerActive                      string
	NumberOfJiffiesUntilTimerExpires int64
	NumberOfUnRecoveredRTOTimeouts   int64
	UID                              string
	UnAnsweredZeroWindowProbes       int64
	INode                            int64
	SocketReferenceCount             int64
	SocketLocationAddress            string
	ReTransmitTimeout                int64
	DelayedAcknowledgeControlData    int64
	AcknowledgePingPong              int64
	SendingCongestionWindow          int64
	SlowStartThreshold               int64
}
type TcpStats struct {
	TcpByStatus map[string][]*Tcp
	Tcp6ByStatus map[string][]*Tcp
}
type TopStat struct {
	ProcessId             int64
	ParentProcessId       int64
	UserId                string
	Priotiry              int64
	PriorityNice          int64
	State                 string
	NumberOfThreads       int64
	MemoryUsagePercentage float64
	CpuUsagePercentage    float64
	Command               string
}
type TotalIncludingChildren struct {
	TotalNumberOfThreads       int64
	TotalMemoryUsagePercentage float64
	TotalCpuUsagePercentage    float64
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
	CpuUsage               float64 //( 100 * ((Total Time / Clock Ticks) / Seconds) )
	MemoryUsage            float64 // ((ResidentSetSize + Data)*100)/ Mem Total
	ChildProcesses         []*ProcessInfoStat
	Network                *NetStat
	Memory                 *ProcessStatM
	ProcessDetails         *ProcessStatus
	CommandString          string
	Top                    *TopStat
	TotalIncludingChildren *TotalIncludingChildren
	TCP                    *TcpStats
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
type ProcessStatus struct {
	Entries map[string]string
}

func GetOSProcess() *ProcessStat {

	upTimeStat := GetOsUpTime()
	memStat := GetOSMem()
	osUsers := GetOSUsers()
	totalMemory := int64(0)
	if memStat != nil && memStat.Entries != nil {
		if _, ok := memStat.Entries["MemTotal"]; ok {
			totalMemory, _ = strconv.ParseInt(memStat.Entries["MemTotal"], 10, 64)
		}
	}
	stat := new(ProcessStat)
	stat.ProcessTree = []*ProcessInfoStat{}
	stat.Top = []*TopStat{}
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
							processInfoStat.State = getProcessStateName(fields[2])
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
							processInfoStat.Memory = GetProcessMemoryStat(processInfoStat.ProcessId)
							processInfoStat.ProcessDetails = GetProcessStatus(processInfoStat.ProcessId)
							processInfoStat.TotalTime = processInfoStat.UserModeTime + processInfoStat.SystemModeTime + processInfoStat.ChildrenUserModeWaitedTime
							processInfoStat.Seconds = int64(upTimeStat.UpTimeFloat) - (processInfoStat.StartTimeAfterSystemBoot / config.ApplicationConfig.CpuTicks)
							processInfoStat.CpuUsage = 100 * ((float64(processInfoStat.TotalTime) / float64(config.ApplicationConfig.CpuTicks)) / float64(processInfoStat.Seconds))
							if processInfoStat.Memory != nil && processInfoStat.Memory.ResidentSetSize > 0 && processInfoStat.Memory.Data > 0 && memStat != nil && totalMemory > 0 {
								processInfoStat.MemoryUsage = ((float64(processInfoStat.Memory.ResidentSetSize) + float64(processInfoStat.Memory.Data)) * 100) / float64(totalMemory)
							}
							processInfoStat.CommandString = GetCommand(processInfoStat.ProcessId)
							//Populate TopStatus
							topStat := new(TopStat)
							topStat.ProcessId = processInfoStat.ProcessId
							topStat.ParentProcessId = processInfoStat.ParentProcessId
							if processInfoStat.ProcessDetails != nil && processInfoStat.ProcessDetails.Entries != nil && len(processInfoStat.ProcessDetails.Entries) > 0 {
								if _, ok := processInfoStat.ProcessDetails.Entries["Uid"]; ok {
									uidRow := processInfoStat.ProcessDetails.Entries["Uid"]
									fields := strings.Fields(uidRow)
									if fields != nil && len(fields) > 0 {
										uid := strings.TrimSpace(fields[0])
										if len(uid) > 0 && osUsers != nil && osUsers.UserMap != nil && len(osUsers.UserMap) > 0 {
											if _, ok := osUsers.UserMap[uid]; ok {
												user := osUsers.UserMap[uid]
												if len(user.UserName) > 0 {
													topStat.UserId = user.UserName
												}
											}
										}
									}
								}
							}

							topStat.Priotiry = processInfoStat.Priority
							topStat.PriorityNice = processInfoStat.PriorityNice
							topStat.State = processInfoStat.State
							topStat.NumberOfThreads = processInfoStat.NumberOfThreads
							topStat.MemoryUsagePercentage = processInfoStat.MemoryUsage
							topStat.CpuUsagePercentage = processInfoStat.CpuUsage
							topStat.Command = processInfoStat.CommandString
							stat.Top = append(stat.Top, topStat)
							processInfoStat.Top = topStat
							processInfoStat.TCP = GetTcpStats(processInfoStat.ProcessId)
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
		rootProcess := processIdMap[0]
		if rootProcess != nil && rootProcess.ChildProcesses != nil && len(rootProcess.ChildProcesses) > 0 {
			for _, childProcess := range rootProcess.ChildProcesses {
				if childProcess != nil {
					stat.ProcessTree = append(stat.ProcessTree, childProcess)
					GetTotalIncludingChildren(childProcess)
				}
			}
		}
	}
	return stat
}
func GetTotalIncludingChildren(processInfoStat *ProcessInfoStat) {
	childrenTotal := new(TotalIncludingChildren)
	if processInfoStat.ChildProcesses == nil || len(processInfoStat.ChildProcesses) == 0 {
		childrenTotal.TotalCpuUsagePercentage = processInfoStat.Top.CpuUsagePercentage
		childrenTotal.TotalMemoryUsagePercentage = processInfoStat.Top.MemoryUsagePercentage
		childrenTotal.TotalNumberOfThreads = processInfoStat.Top.NumberOfThreads
		processInfoStat.TotalIncludingChildren = childrenTotal
	} else {
		childrenTotal.TotalCpuUsagePercentage = processInfoStat.Top.CpuUsagePercentage
		childrenTotal.TotalMemoryUsagePercentage = processInfoStat.Top.MemoryUsagePercentage
		childrenTotal.TotalNumberOfThreads = processInfoStat.Top.NumberOfThreads
		for _, childProcess := range processInfoStat.ChildProcesses {
			if childProcess != nil {
				GetTotalIncludingChildren(childProcess)
				if childProcess.TotalIncludingChildren != nil {
					childrenTotal.TotalNumberOfThreads = childrenTotal.TotalNumberOfThreads + childProcess.TotalIncludingChildren.TotalNumberOfThreads
					childrenTotal.TotalMemoryUsagePercentage = childrenTotal.TotalMemoryUsagePercentage + childProcess.TotalIncludingChildren.TotalMemoryUsagePercentage
					childrenTotal.TotalCpuUsagePercentage = childrenTotal.TotalCpuUsagePercentage + childProcess.TotalIncludingChildren.TotalCpuUsagePercentage
				}
			}
		}
		processInfoStat.TotalIncludingChildren = childrenTotal
	}

}
func GetProcessMemoryStat(processId int64) *ProcessStatM {
	processStatM := new(ProcessStatM)
	lines := []string{}
	err := errors.New("error")
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err = util.ReadLines(procPath + "/" + strconv.Itoa(int(processId)) + "/statm")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		processStatM = nil

	} else {

		if len(lines) > 0 {
			dataLine := lines[0]
			fields := strings.Fields(dataLine)
			processStatM.TotalProgramSize, _ = strconv.ParseInt(fields[0], 10, 64)
			processStatM.ResidentSetSize, _ = strconv.ParseInt(fields[1], 10, 64)
			processStatM.SharedPages, _ = strconv.ParseInt(fields[2], 10, 64)
			processStatM.Text = fields[3]
			processStatM.Lib, _ = strconv.ParseInt(fields[4], 10, 64)
			processStatM.Data, _ = strconv.ParseInt(fields[5], 10, 64)
			processStatM.DirtyPages, _ = strconv.ParseInt(fields[6], 10, 64)
		}
	}

	return processStatM
}
func GetProcessStatus(processId int64) *ProcessStatus {
	stat := new(ProcessStatus)
	stat.Entries = make(map[string]string)
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/" + strconv.Itoa(int(processId)) + "/status")
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
func GetCommand(processId int64) string {
	command := ""
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines, err := util.ReadLines(procPath + "/" + strconv.Itoa(int(processId)) + "/comm")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		command = ""

	} else {

		if len(lines) > 0 {
			command = lines[0]

		}
	}
	return command
}
func GetTcpStats(processId int64) *TcpStats{
tcpStats := new(TcpStats)
tcpStats.TcpByStatus=GetTcpByStatus(processId,TCP4)
tcpStats.Tcp6ByStatus=GetTcpByStatus(processId,TCP6)
return tcpStats
}
func GetTcpByStatus(processId int64,tcpType TcpType) map[string][]*Tcp {
	tcpStatusMap := make(map[string][]*Tcp)
	defaultProcFolder := "/proc"
	procPath := util.GetEnv("HOST_PROC", defaultProcFolder)
	lines:=[]string{}
	err:=errors.New("")
	if tcpType==TCP4 {
		lines, err = util.ReadLines(procPath + "/" + strconv.Itoa(int(processId)) + "/net/tcp")
	}else if tcpType==TCP6{
		lines, err = util.ReadLines(procPath + "/" + strconv.Itoa(int(processId)) + "/net/tcp6")
	}
	if err != nil {
		log.Fatalf("readLines: %s", err)
		tcpStatusMap = nil

	} else {
		if len(lines) > 0 {
			for i := 1; i < len(lines); i++ {
				dataLine := lines[i]
				fields := strings.Fields(dataLine)
				if fields != nil && len(fields) > 0 {
					tcp := new(Tcp)
					if len(fields)>=4 {
						/*
							46: 010310AC:9C4C 030310AC:1770 01
|   |         |   |        |    |--> connection state
|   |         |   |        |------> remote TCP port number
|   |         |   |-------------> remote IPv4 address
|   |         |--------------------> local TCP port number
|   |---------------------------> local IPv4 address
|----------------------------------> number of entry
						 */
						tcp.SequenceId, _ = strconv.ParseInt(strings.ReplaceAll(fields[0], ":", ""), 10, 64)
						localIpAddressPort := strings.Split(fields[1], ":")
						if localIpAddressPort != nil && len(localIpAddressPort) == 2 {
							if tcpType==TCP4 {
								tcp.LocalIpAddress = ConvertHexToIPV4(localIpAddressPort[0])
							}else if tcpType==TCP6{
								tcp.LocalIpAddress = ConvertHexToIPV6(localIpAddressPort[0])
							}
							tcp.LocalPortNumber, _ = strconv.ParseInt(hexaNumberToInteger(localIpAddressPort[1]), 16, 64)
						}
						remoteIpAddressPort := strings.Split(fields[2], ":")
						if remoteIpAddressPort != nil && len(remoteIpAddressPort) == 2 {
							if tcpType==TCP4 {
								tcp.RemoteIpAddress = ConvertHexToIPV4(remoteIpAddressPort[0])
							}else if tcpType==TCP6{
								tcp.RemoteIpAddress = ConvertHexToIPV6(remoteIpAddressPort[0])
							}
							tcp.RemotePortNumber, _ = strconv.ParseInt(hexaNumberToInteger(remoteIpAddressPort[1]), 16, 64)
						}
						stateCode, _ := strconv.ParseInt(hexaNumberToInteger(fields[3]), 16, 64)
						tcp.ConnectionState = getTCPStateName(stateCode)
					}
					if len(fields)>=7 {
						/*
						00000150:00000000 01:00000019 00000000
|        |        |  |        |--> number of unrecovered RTO timeouts
|        |        |  |----------> number of jiffies until timer expires
|        |        |----------------> timer_active (see below)
|        |----------------------> receive-queue
|-------------------------------> transmit-queue
						 */
						transmitReceiveQueue := strings.Split(fields[4], ":")
						if transmitReceiveQueue != nil && len(transmitReceiveQueue) == 2 {
							tcp.TransmitQueue = transmitReceiveQueue[0]
							tcp.ReceiveQueue = transmitReceiveQueue[1]
						}
						timerTimerExpires := strings.Split(fields[5], ":")
						if timerTimerExpires != nil && len(timerTimerExpires) == 2 {
							timerStateCode, _ := strconv.ParseInt(hexaNumberToInteger(timerTimerExpires[0]), 10, 64)
							tcp.TimerActive = getTCPTimerState(timerStateCode)
							tcp.NumberOfJiffiesUntilTimerExpires, _ = strconv.ParseInt(hexaNumberToInteger(timerTimerExpires[0]), 16, 64)
						}
						rtoTimeouts, _ := strconv.ParseInt(hexaNumberToInteger(fields[6]), 16, 64)
						tcp.NumberOfUnRecoveredRTOTimeouts = rtoTimeouts
					}
					if len(fields)>=8 {
						tcp.UID = fields[7]
					}
					if len(fields)>=9 {
						tcp.UnAnsweredZeroWindowProbes, _ = strconv.ParseInt(fields[8], 10, 64)
					}
					if len(fields)>=10 {
						tcp.INode, _ = strconv.ParseInt(fields[9], 10, 64)
					}
					if len(fields)>=11 {
						tcp.SocketReferenceCount, _ = strconv.ParseInt(fields[10], 10, 64)
					}
					if len(fields)>=12 {
						tcp.SocketLocationAddress = fields[11]
					}
					if len(fields)>=13 {
						tcp.ReTransmitTimeout, _ = strconv.ParseInt(fields[12], 16, 64)
					}
					if len(fields)>=14 {
						tcp.DelayedAcknowledgeControlData, _ = strconv.ParseInt(hexaNumberToInteger(fields[13]), 10, 64)
					}
					if len(fields)>=15 {
						tcp.AcknowledgePingPong, _ = strconv.ParseInt(fields[14], 10, 64)
					}
					if len(fields)>=16 {
						tcp.SendingCongestionWindow, _ = strconv.ParseInt(fields[15], 10, 64)
					}
					if len(fields)>=17 {
						tcp.SlowStartThreshold, _ = strconv.ParseInt(fields[16], 10, 64)
					}
					/*

1000 0 54165785 4 cd1e6040 25 4 27 3 -1
|    | |        | |        |  | |  |  |--> slow start size threshold,
|    | |        | |        |  | |  |       or -1 if the treshold
|    | |        | |        |  | |  |       is >= 0xFFFF
|    | |        | |        |  | |  |----> sending congestion window
|    | |        | |        |  | |-------> (ack.quick<<1)|ack.pingpong
|    | |        | |        |  |---------> Predicted tick of soft clock
|    | |        | |        |               (delayed ACK control data)
|    | |        | |        |------------> retransmit timeout
|    | |        | |------------------> location of socket in memory
|    | |        |-----------------------> socket reference count
|    | |-----------------------------> inode
|    |----------------------------------> unanswered 0-window probes
|---------------------------------------------> uid
			 */
					if tcpStatusMap[tcp.ConnectionState] == nil {
						tcpStatusMap[tcp.ConnectionState] = []*Tcp{}
					}
					tcpStatusMap[tcp.ConnectionState] = append(tcpStatusMap[tcp.ConnectionState], tcp)

				}
			}

		}
	}
	return tcpStatusMap
}
func getTCPStateName(stateNumber int64) string {
	state := ""
	switch stateNumber {
	case 1:
		state = "TCP_ESTABLISHED"
	case 2:
		state = "TCP_SYN_SENT"
	case 3:
		state = "TCP_SYN_RECV"
	case 4:
		state = "TCP_FIN_WAIT1"
	case 5:
		state = "TCP_FIN_WAIT2"
	case 6:
		state = "TCP_TIME_WAIT"
	case 7:
		state = "TCP_CLOSE"
	case 8:
		state = "TCP_CLOSE_WAIT"
	case 9:
		state = "	TCP_LAST_ACK"
	case 10:
		state = "TCP_LISTEN"

	}
	return state
}
func getTCPTimerState(stateNumber int64) string {
	state := ""
	switch stateNumber {
	case 0:
		state = "NO_TIMER_PENDING"
	case 1:
		state = "TRANSMIT_TIMER_PENDING"
	case 2:
		state = "DELAYED_ACK_OR_KEEPALIVE_PENDING"
	case 3:
		state = "SOCKET_IN_TIME_WAIT Not all fields will contain data (or even exist)"
	case 4:
		state = "ZERO_WINDOW_PROBE_TIMER_PENDING"
	}
	return state
}
func getProcessStateName(stateChar string) string {
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
func ConvertHexToIPV4(hexString string) string {
	result := ""
	a, _ := hex.DecodeString(hexString)

	result = fmt.Sprintf("%v.%v.%v.%v", a[3], a[2], a[1], a[0])
	return result
}
func ConvertHexToIPV6(hexString string) string {
	//result := ""
	ip:=net.ParseIP(hexString)
	//a, _ := hex.DecodeString(hexString)
	//result = fmt.Sprintf("%v.%v.%v.%v.%v.%v", a[5],a[4],a[3],a[2],a[1],a[0])
	return ip.String()
}
func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
