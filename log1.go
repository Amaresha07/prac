package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer file.Close()

	// Create a multi-writer that writes to both file and console
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Create loggers that log to both file and console
	infoLogger := log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime)
	warnLogger := log.New(multiWriter, "WARNING: ", log.Ldate|log.Ltime)
	errorLogger := log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	infoLogger.Println("Application started successfully.")
	warnLogger.Println("Memory usage is reaching high levels.")
	errorLogger.Println("Failed to connect to the database.")

	file.Sync()
}
