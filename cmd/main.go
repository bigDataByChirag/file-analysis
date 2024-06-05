package main

import (
	"file/dir"
	"file/handler"
	"file/mid"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

func main() {

	envPath, err := dir.GetEnvPath()
	if err != nil {
		log.Panic("error getting environment variable dir", err)
	}
	err = godotenv.Load(envPath)
	if err != nil {
		log.Panic("error loading environment variables dir", err)
	}
	setupSlog()
	// Create Gin router
	router := gin.Default()

	router.Use(mid.GinLoggerMiddleware())
	// Register Routes
	router.GET("/file/:name", handler.LargestFile)
	router.GET("/dir/:name", handler.TotalDirSize)

	// Start the server
	router.Run(os.Getenv("PORT"))

}

func setupSlog() {
	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		//AddSource: true: This will cause the source dir and line number of the log message to be included in the output
		AddSource: true,
	})

	logger := slog.New(logHandler)
	//SetDefault makes l the default Logger. in our case we would be doing structured logging
	slog.SetDefault(logger)
}
