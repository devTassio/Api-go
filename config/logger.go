package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Definição do tipo Logger
type Logger struct {
	debug   *log.Logger // Logger para mensagens de debug
	info    *log.Logger // Logger para mensagens de informação
	warning *log.Logger // Logger para mensagens de aviso
	err     *log.Logger // Logger para mensagens de erro
	writer  io.Writer   // Escritor de saída para os loggers
}

// Função NewLogger cria uma nova instância de Logger
// Parâmetro p é o prefixo para os loggers
func NewLogger(p string) *Logger {
	// Define o escritor de saída para o padrão (stdout)
	writer := io.Writer(os.Stdout)
	// Cria um novo logger com o escritor de saída, prefixo e opções de formatação
	logger := &Logger{
		debug:   log.New(writer, p+" DEBUG: ", log.Ldate|log.Ltime),   // Logger para mensagens de debug
		info:    log.New(writer, p+" INFO: ", log.Ldate|log.Ltime),    // Logger para mensagens de informação
		warning: log.New(writer, p+" WARNING: ", log.Ldate|log.Ltime), // Logger para mensagens de aviso
		err:     log.New(writer, p+" ERROR: ", log.Ldate|log.Ltime),   // Logger para mensagens de erro
		writer:  writer,                                               // Define o escritor de saída para o logger
	}
	// Retorna a nova instância do logger
	return logger
}

// Create Non-Format Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Create Format Enable Logs

func (l *Logger) Debugf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.debug.Println(msg)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.info.Println(msg)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.warning.Println(msg)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	l.err.Println(msg)
}
