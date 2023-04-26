package main

import (
	"log"
	"os"
)

const logFlags = log.Ldate | log.Ltime | log.LUTC | log.Lmsgprefix

func main() {

	l := log.New(os.Stdout, "My Fancy Logger: ", logFlags)

	l.Println("Hello World!!!")
}
