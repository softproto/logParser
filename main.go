package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

type LogRecord struct {
	Log_time    time.Time `bson:"log_time"`
	Log_message string    `bson:"log_message"`
	File_name   string    `bson:"file_name"`
	Log_format  string    `bson:"log_format"`
}

var logTimeLayout = map[string]string{
	"first_format":  "Jan 2, 2006 at 3:04:05pm (MST)",
	"second_format": "2006-01-02T15:04:05Z07:00",
}

const dataChannelBuffer = 100
const fileCheckTimeout = 1000

const databaseURL = "mongodb://localhost"
const databaseName = "logDatabase"
const collectionName = "logRecords"

//main
func main() {
	logfilesList := os.Args[1:]

	if len(logfilesList) == 0 {
		logfilesList = append(logfilesList, "data3.dat", "data2.dat", "data1.dat")
		//fmt.Println("usage: logParser filename.dat [...]")
		//log.Fatal("wrong usage")
	}

	fmt.Println("\nStarted for:", logfilesList)

	dataChannel := make(chan LogRecord, dataChannelBuffer)

	session, err := mgo.Dial(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	logCollection := session.DB(databaseName).C(collectionName)

	for filenameIndex := range logfilesList {
		fmt.Println("run parser for", logfilesList[filenameIndex])
		go readLinesFromFile(dataChannel, logfilesList[filenameIndex], fileCheckTimeout)
	}
	fmt.Println("all files processed")

	for {
		for currentLogRecord := range dataChannel {
			err = logCollection.Insert(currentLogRecord)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\n<-New log record")
			fmt.Println("log_time:", currentLogRecord.Log_time)
			fmt.Println("log_message:", currentLogRecord.Log_message)
			fmt.Println("file_name:", currentLogRecord.File_name)
			fmt.Println("log_format:", currentLogRecord.Log_format)
		}
	}

}
