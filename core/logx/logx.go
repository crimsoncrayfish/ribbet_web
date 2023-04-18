package logx

import (
	"fmt"
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

func New(logOriginName string, verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stdout, logOriginName+": ", logFlags)
	} else {
		//overwrite old logs
		file, err := os.Create("tmp.logs")
		if err != nil {
			fmt.Println(err)
		}
		return log.New(file, logOriginName+": ", logFlags)
	}
}
