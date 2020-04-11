package util

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
)

func GetClockTick() int64{
	cmd := exec.Command("getconf", "CLK_TCK")
	out := new(bytes.Buffer)
	cmd.Stdout = out
	_ = cmd.Run()
	hertz:=out.String()
	hertz=strings.TrimSuffix(hertz, "\n")
	hertzInt64,_:=strconv.ParseInt(hertz,10,64)
	return hertzInt64
}
