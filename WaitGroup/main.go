package main

import (
	"os"
	"os/signal"
	"syscall"
)

func waitDown() {
	sig := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
}

func main() {
	//chanPrint()
	waitgroupPrint()
	waitDown()
}
