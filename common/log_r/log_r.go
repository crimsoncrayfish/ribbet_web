package logr

import (
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

func New(source string) *log.Logger {

	return log.New(os.Stdout, source+" :", logFlags)
}
