package main

import (
	logr "ribbet_web/common/log_r"
)

func main() {
	l := logr.New("My Fancy Logger")

	l.Println("Hello World!!!")
}
