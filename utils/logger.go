package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

type logger struct {
	method string
	data   interface{}
}

type LogError struct {
	ErrorMessage string `json:"error_message"`
	ErrorLine    int    `json:"line"`
	ErrorFile    string `json:"file"`
}

func Logger() *logger {
	return &logger{
		method: os.Getenv("LOG_METHOD"),
	}
}

func (l *logger) AddData(data interface{}) *logger {
	l.data = data
	return l
}

func (l *logger) Log(action string, desc string) {
	if l.method == "db" {
		l.LogToDB(action, desc)
	} else if l.method == "file" {
		l.LogToFile(action, desc)
	}
}

func (l *logger) LogToDB(action string, desc string) {
	db := DB()
	jsonData := l.data
	if l.data != nil {
		marshal, err := json.Marshal(l.data)
		if err != nil {
			l.Error(err)
		}
		jsonData = string(marshal)
	}
	_, err := db.Query(`insert into log.app (action, description, date_created, json_data) values ($1, $2, $3, $4)`,
		action,
		desc,
		time.Now(),
		jsonData,
	)
	if err != nil {
		l.Error(err)
		log.Fatal(err)
	}
}

func (l *logger) LogToFile(action string, desc string) {
	file, err := os.Create("example.log")
	if err != nil {
		log.Fatal("Error creating log file:", err)
		return
	}
	defer file.Close()

	// create a writer to append content to the file
	writer := bufio.NewWriter(file)

	// write content to the file
	_, err = writer.WriteString(fmt.Sprintf("[%s][%s] -> %s", time.Now(), action, desc))
	if err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}

	// flush the writer to make sure all content is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}

func (l *logger) Error(err error) {
	_, fn, line, _ := runtime.Caller(1)
	l.LogToFile("ERROR PROGRAM", fmt.Sprintf("%s on line %v file %s", err.Error(), line, fn))
}

func (l *logger) AddErrorData(err error) *logger {
	if err == nil {
		l.data = LogError{}
	}

	_, fn, line, _ := runtime.Caller(1)

	l.data = LogError{
		ErrorMessage: err.Error(),
		ErrorLine:    line,
		ErrorFile:    fn,
	}
	return l
}
