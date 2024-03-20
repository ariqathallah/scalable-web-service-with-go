package controller

import (
	models "assignment-3/model"
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	// get current directory
	dir, err := os.Getwd()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// read file
	byte, err := os.ReadFile(filepath.Join(dir, "lib", "data.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// decode json
	var data models.Status
	err = json.Unmarshal(byte, &data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var waterResponse string
	var windResponse string

	if data.Status.Water <= 5 {
		waterResponse = "aman"
	} else if data.Status.Water >= 6 && data.Status.Water <= 8 {
		waterResponse = "siaga"
	} else {
		waterResponse = "bahaya"
	}

	if data.Status.Wind <= 6 {
		windResponse = "aman"
	} else if data.Status.Wind >= 7 && data.Status.Wind <= 15 {
		windResponse = "siaga"
	} else {
		windResponse = "bahaya"
	}

	// return html file
	c.HTML(http.StatusOK, "index.html", gin.H{
		"water":        data.Status.Water,
		"water_status": waterResponse,
		"wind":         data.Status.Wind,
		"wind_status":  windResponse,
	})
}
