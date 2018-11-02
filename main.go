package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type logRecord struct {
	log_time    time.Time
	log_message string
	file_name   string
	log_format  string
}

var logTimeLayout = map[string]string{
	"first_format":  "Jan 2, 2006 at 3:04:05pm (MST)",
	"second_format": "2006-01-02T15:04:05Z07:00",
}

const dataChannelBuffer = 100

const fileCheckTimeout = 1000

//main
func main() {
	logfilesList := os.Args[1:]

	if len(logfilesList) == 0 {
		//logfilesList = append(logfilesList, "data3.dat", "data2.dat", "data1.dat")
		fmt.Println("usage: logParser filename.dat ...")
		log.Fatal("wrong usage")
	}

	fmt.Println("\nStarted for:", logfilesList)

	dataChannel := make(chan logRecord, dataChannelBuffer)

	for filenameIndex := range logfilesList {
		fmt.Println("run parser for", logfilesList[filenameIndex])
		go readLinesFromFile(dataChannel, logfilesList[filenameIndex], fileCheckTimeout)
	}
	fmt.Println("all files processed")

	for {
		for currentLogRecord := range dataChannel {
			fmt.Println("\n<-New log record")
			fmt.Println("log_time:", currentLogRecord.log_time)
			fmt.Println("log_message:", currentLogRecord.log_message)
			fmt.Println("file_name:", currentLogRecord.file_name)
			fmt.Println("log_format:", currentLogRecord.log_format)
		}
	}

}
