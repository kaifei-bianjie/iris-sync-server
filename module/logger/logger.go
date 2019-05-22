package logger

import (
	"io"
	"log"
	"os"
	"io/ioutil"
)

var (
	Trace   *log.Logger // Trace db log
	Info    *log.Logger // Important information
	Warning *log.Logger // Warning information
	Error   *log.Logger // Critical problem
)

const (
	// traceDbFile = ".sync_server_db.log"
	// debugFile = ".sync_server_debug.log"
	errFile     = ".sync_server_err.log"
	warningFile = ".sync_server_warning.log"
)

func init() {
	// traceDbFile, err := os.OpenFile(os.ExpandEnv("$HOME/"+traceDbFile),
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("Failed to open trace db log file:", err)
	// }
	//
	// debugFile, err := os.OpenFile(os.ExpandEnv("$HOME/"+debugFile),
	// 	os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalln("Failed to open debug log file:", err)
	// }
	
	warningFile, err := os.OpenFile(os.ExpandEnv("$HOME/"+warningFile),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warning log file:", err)
	}
	
	errFile, err := os.OpenFile(os.ExpandEnv("$HOME/"+errFile),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(io.MultiWriter(warningFile, os.Stderr),
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(errFile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func test() {
	Trace.Println("This is trace info...")
	Info.Println("This is info info...")
	Warning.Println("This is warning info...")
	Error.Println("This is err info...")
}
