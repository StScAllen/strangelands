// log.go

package main

import "fmt"
import "os"

type Log struct {
	errorCount, warnCount, infoCount int
	errors, warns, infos string
}

func openLog() (Log){
	var log Log
	
	log.errorCount = 0
	log.warnCount = 0
	log.infoCount = 0
	log.errors = ""
	log.warns = ""
	log.infos = ""
	
	return log
}

func (log * Log) addError(msg String){
	log.errorCount += 1
	log.errors += msg + "\n"
}

func (log * Log) addWarn(msg String){
	log.warnCount += 1
	log.warns += msg + "\n"
}

func (log * Log) addInfo(msg String){
	log.infoCount += 1
	log.infos += msg + "\n"
}

func (log * Log) writeToFile(){
	var filename string
	
	filename = "log.txt"
	
	file, err := os.Create(filename)

	var saveString string 
	
	saveString += log.infos
	saveString += log.warns
	saveString += log.errors
	
	if (err == nil){
		defer file.Close()		
		file.WriteString(saveString)	
	}
}

func (log * Log) displayLog() {
	clearConsole()
	fmt.Println(log.infos)
	fmt.Println(log.warns)
	fmt.Println(log.errors)
}