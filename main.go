package main

import (
	"fmt"
	"log"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func main() {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	for _, s := range stat.CPUStats {
		// fmt.Println(s.User)
		// fmt.Println(s.System)
		fmt.Println(s.Idle)
		// s.IOWait
	}
}
