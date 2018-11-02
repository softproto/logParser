package main

import (
	//	"bufio"
	"fmt"
	//	"log"
	"os"
	//	"strings"
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

//main
func main() {
	logfilesList := os.Args[1:]
	fmt.Println("\nStarted")

	if len(logfilesList) == 0 {
		//logfilesList = append(logfilesList, "data3.dat", "data2.dat", "data1.dat")
		logfilesList = append(logfilesList, "data1.dat")
		//fmt.Println("usage: logParser filename.dat ...")
	}

	for filenameIndex := range logfilesList {
		readLinesFromFile(logfilesList[filenameIndex], 1000)
	}

}
