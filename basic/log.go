package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	TRACE   *log.Logger
	INFO    *log.Logger
	WARNING *log.Logger
	ERROR   *log.Logger
	MYFILE  *log.Logger
)

func Init(
	traceHandler io.Writer,
	infoHandler io.Writer,
	warningHandler io.Writer,
	errorHandler io.Writer) {

	TRACE = log.New(traceHandler, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)

	INFO = log.New(infoHandler, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	WARNING = log.New(warningHandler, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	ERROR = log.New(errorHandler, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	// Use file logging
	output := "logging.txt"
	file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", output, ":", err)
	}

	MYFILE = log.New(file, "PREFIX: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	TRACE.Println("trace message")
	INFO.Println("info message")
	WARNING.Println("warning message")
	ERROR.Println("error message")

	MYFILE.Println("Where did it go?")
}
