package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var debuggerInstance *Debugger

type Debugger struct {
	debug       bool
	InitialTime time.Time
}

func NewDebugger() *Debugger {
	if debuggerInstance == nil {
		debuggerInstance = &Debugger{
			debug:       false,
			InitialTime: time.Now(),
		}
	}
	return debuggerInstance
}

func (d *Debugger) SetDebug(value bool) {
	d.debug = value
}

func (d *Debugger) IsDebug() bool {
	return d.debug
}

func (d *Debugger) PrintForDebug(format string, args ...any) {
	if d.IsDebug() {
		println("[DEBUG] %s\n", fmt.Sprintf(format, args...))
		f, err := os.OpenFile("debug.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			println("Error opening debug file: %v\n", err)
			return
		}
		defer f.Close()
		if _, err := f.WriteString(fmt.Sprintf("[DEBUG] %s\n", fmt.Sprintf(format, args...))); err != nil {
			println("Error writing to debug file: %v\n", err)
		}
	}
}

func (d *Debugger) PrintLog(format string, args ...any) {
	// Usar d.InitialTime para la hora
	currentHour := time.Now().Format("15:04:05")
	newString := fmt.Sprintf("[LOG] <"+currentHour+"> "+format, args...)
	fmt.Println(newString)

	// Obtener el directorio de logs y el archivo en formato día-mes-año.log
	basePath, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	logDir := filepath.Join(basePath, "logs")
	// Formato de la hora que es compatible con Windows
	logFilePath := filepath.Join(logDir, d.InitialTime.Format("02-01-2006-15-04-05")+".log")

	// Create directories if they don't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating log directory:", err)
		return
	}

	// Open the log file in append mode, creating it if it doesn't exist
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Write newString to the log file
	_, err = file.WriteString(newString + "\n")
	if err != nil {
		fmt.Println("Error writing to log file:", err)
		return
	}
}

func (d *Debugger) PrintError(format string, args ...any) {
	// Usar d.InitialTime para la hora
	currentHour := time.Now().Format("15:04:05")
	newString := fmt.Sprintf("[ERROR] <"+currentHour+"> "+format, args...)
	fmt.Println(newString)

	// Obtener el directorio de logs y el archivo en formato día-mes-año.log
	basePath, err := os.Getwd()
	if err != nil {
		return
	}
	logDir := basePath + "/logs"
	logFilePath := logDir + "/" + d.InitialTime.Format("02-01-2006 15:04:05") + ".log"

	// Crear directorios si no existen
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	// Abrir el archivo en modo de anexado
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Escribir newString en el archivo de log
	_, err = file.WriteString(newString + "\n")
	if err != nil {
		return
	}
}

func (d *Debugger) PrintWarning(format string, args ...any) {
	// Usar d.InitialTime para la hora
	currentHour := time.Now().Format("15:04:05")
	newString := fmt.Sprintf("[WARNING] <"+currentHour+"> "+format, args...)
	fmt.Println(newString)

	// Obtener el directorio de logs y el archivo en formato día-mes-año.log
	basePath, err := os.Getwd()
	if err != nil {
		return
	}
	logDir := basePath + "/logs"
	logFilePath := logDir + "/" + d.InitialTime.Format("02-01-2006 15:04:05") + ".log"

	// Crear directorios si no existen
	err = os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return
	}

	// Abrir el archivo en modo de anexado
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// Escribir newString en el archivo de log
	_, err = file.WriteString(newString + "\n")
	if err != nil {
		return
	}
}

func PrintLog(format string, args ...any) {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	debuggerInstance.PrintLog(format, args...)
}

func PrintError(format string, args ...any) {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	debuggerInstance.PrintError(format, args...)
}

func PrintWarning(format string, args ...any) {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	debuggerInstance.PrintWarning(format, args...)
}

func SetDebugTest(value bool) {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	debuggerInstance.SetDebug(value)
}

func IsDegubTest() bool {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	return debuggerInstance.IsDebug()
}

func PrintForDebug(format string, args ...any) {
	if debuggerInstance == nil {
		debuggerInstance = NewDebugger()
	}
	debuggerInstance.PrintForDebug(format, args...)
}

func GetTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
