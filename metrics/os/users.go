package os

import (
	"github.com/goldentigerindia/profiling-agent/util"
	"log"
	"strings"
)
type Users struct{
	UserList []*User
	UserMap map[string]*User
}
type User struct {
	UserName string
	UID string
	GID string
	UserDisplayName string
	HomeDirectory string
	LoginShell string
}
func GetOSUsers() *Users {
	users:=new(Users)
	users.UserList=[]*User{}
	users.UserMap=make(map[string]*User)
	defaultEtcFolder := "/etc"
	procPath := util.GetEnv("HOST_ETC", defaultEtcFolder)
	lines, err := util.ReadLines(procPath + "/passwd")
	if err != nil {
		log.Fatalf("readLines: %s", err)
		users = nil

	} else {

		if len(lines) > 0 {
			for i := 0; i < len(lines); i++ {
				dataLine := lines[i]

				fields := strings.SplitN(dataLine, ":", -1)
				if fields != nil && len(fields) > 6 && len(strings.TrimSpace(fields[0])) > 0  {
					user:=new(User)
					user.UserName=strings.TrimSpace(fields[0])
					user.UID=strings.TrimSpace(fields[2])
					user.GID=strings.TrimSpace(fields[3])
					user.UserDisplayName=strings.TrimSpace(fields[4])
					user.HomeDirectory=strings.TrimSpace(fields[5])
					user.LoginShell=strings.TrimSpace(fields[6])
					users.UserList=append(users.UserList,user)
					users.UserMap[user.UID]=user
				}
			}
		}
	}
	return users
}