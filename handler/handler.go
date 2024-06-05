package handler

import (
	"file/dir"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func LargestFile(c *gin.Context) {

	name := c.Param("name")
	ls, ln, err := dir.FindLargestFile(name)
	if err != nil {
		slog.Error("largest file failed", slog.String("Error", err.Error()))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		fmt.Println()
		return
	}
	c.JSON(http.StatusFound, gin.H{"Largest File Size in bytes": ls, "Largest File Name": ln})
}

func TotalDirSize(c *gin.Context) {

	name := c.Param("name")
	cds, err := dir.CompleteDirSize(name)
	if err != nil {
		slog.Error("total directory size failed", slog.String("Error", err.Error()))
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		fmt.Println()
		return
	}
	c.JSON(http.StatusOK, gin.H{"Total Directory Size": cds})

}
