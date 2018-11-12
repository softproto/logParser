package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

//read and parse the lines from selected file
func readLinesFromFile(outChannel chan<- LogRecord, fileName string, fileReadingTimeout time.Duration) {
	var startLineNumber uint = 0
	var fileLastModifiedTime time.Time = time.Unix(0, 0)

	for {
		var currentLogRecord LogRecord
		var rawLogString string
		var currentLineNumber uint = 0

		fileHandler, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}

		fileStat, err := fileHandler.Stat()
		if err != nil {
			log.Fatalf("stat: %v", err)
		}

		var fileModifiedTime time.Time = fileStat.ModTime()

		if fileModifiedTime.After(fileLastModifiedTime) {
			scanner := bufio.NewScanner(fileHandler)
			for scanner.Scan() {
				currentLineNumber++
				if currentLineNumber >= startLineNumber {
					currentLogRecord.File_name = fileName
					rawLogString = scanner.Text()
					splittedLines := strings.Split(rawLogString, " | ")
					if len(splittedLines) == 2 {
						currentLogRecord.Log_message = splittedLines[1]
						for formatNumber, layoutString := range logTimeLayout {
							logTime, err := time.Parse(layoutString, splittedLines[0])
							if err != nil {
								//log.Fatal(err)
								currentLogRecord.Log_format = "unknown format"
							} else {
								currentLogRecord.Log_time = logTime
								currentLogRecord.Log_format = formatNumber
								outChannel <- currentLogRecord
								//fmt.Println(currentLogRecord)
								break
							}
						}
					}
				}
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			startLineNumber = currentLineNumber + 1
		}

		fileLastModifiedTime = fileModifiedTime
		time.Sleep(fileReadingTimeout * time.Millisecond)
		fileHandler.Close()
	}
}
