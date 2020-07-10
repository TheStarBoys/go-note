package main

import (
	"fmt"
	"time"
	"os"
)

func checkLogFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("checkLogFile open err: %s", err))
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(fmt.Sprintf("check log file err: %s", err))
	}
	for {
		mb := int64(1 * 1024 * 1024) // 1MB
		if stat.Size() > 10 * mb {
			panic(fmt.Sprintf("log file <%s> exceed 100 MB", stat.Name()))
		}

		time.Sleep(10 * time.Microsecond)
	}
}
