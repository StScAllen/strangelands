// log.go

package main

import "fmt"
import "os"

type Log struct {
	errorCount, warnCount, infoCount, aiCount int
	errors, warns, infos, ais                 string
}

func openLog() Log {
	var log Log

	log.errorCount = 0
	log.warnCount = 0
	log.infoCount = 0
	log.aiCount = 0

	log.errors = ""
	log.warns = ""
	log.infos = ""
	log.ais = ""

	return log
}

func (log *Log) addError(msg string) {
	log.errorCount += 1
	log.errors += msg + "\n"
}

func (log *Log) addWarn(msg string) {
	log.warnCount += 1
	log.warns += msg + "\n"
}

func (log *Log) addInfo(msg string) {
	log.infoCount += 1
	log.infos += msg + "\n"
}

func (log *Log) addAi(msg string) {
	log.aiCount += 1
	log.ais += msg + "\n"
}

func (log *Log) writeToFile() {
	var filename string

	filename = "log.txt"

	file, err := os.Create(filename)

	var saveString string

	saveString += "--Info--\n"
	saveString += log.infos
	saveString += "--Warnings--\n"
	saveString += log.warns
	saveString += "--AI--\n"
	saveString += log.ais
	saveString += "--Errors--\n"
	saveString += log.errors

	if err == nil {
		defer file.Close()
		file.WriteString(saveString)
	}
}

func (log *Log) displayLog() {
	clearConsole()
	fmt.Println("--Info--")
	fmt.Println(log.infos)
	fmt.Println("--Warnings--")
	fmt.Println(log.warns)
	fmt.Println("--AI--")
	fmt.Println(log.ais)
	fmt.Println("--Errors--")
	fmt.Println(log.errors)
	showPause("")
	
	clearConsole()
}
