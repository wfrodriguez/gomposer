package internal

import (
	"log"
	"os"
	"sync"
)

var once sync.Once
var instance *logger
var Verbose bool = true

// var Verbose bool = ncfalse

// FuncLogger es una interfaz que implementa un logger sencillo para la aplicación
type FuncLogger interface {
	Info(string, ...any)
	Warning(string, ...any)
	Error(string, ...any)
}

// logger es una implementación de la interfaz FuncLogger para esta aplicación
type logger struct {
	warningl *log.Logger
	infol    *log.Logger
	errorl   *log.Logger
}

// Crea una instancia de logger
func SimpleLogger() FuncLogger {

	once.Do(func() {
		instance = &logger{
			infol:    log.New(os.Stdout, "\033[34mINFO: \033[0m", log.Ltime),
			warningl: log.New(os.Stdout, "\033[33mWARNING: \033[0m", log.Ltime),
			errorl:   log.New(os.Stdout, "\033[31mERROR: \033[0m", log.Ltime),
		}
	})

	return instance
}

// Info funciona igual que el método `fmt.Printf()`
func (l *logger) Info(msg string, args ...any) {
	if Verbose {
		l.infol.Printf(msg, args...)
	}
}

// Warning funciona igual que el método `fmt.Printf()`
func (l *logger) Warning(msg string, args ...any) {
	l.warningl.Printf(msg, args...)
}

// Error funciona igual que el método `fmt.Printf()`
func (l *logger) Error(msg string, args ...any) {
	l.errorl.Printf(msg, args...)
}
