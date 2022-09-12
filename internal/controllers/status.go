package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
)

type appStatus struct {
	Status      string        `json:"status"`
	Uptime      time.Duration `json:"uptime"`
	Environment string        `json:"environment"`
	Version     string        `json:"version"`
}

func (m *Repository) StatusHandler(c *gin.Context) {
	status := appStatus{
		Status:      "Available",
		Uptime:      time.Duration(time.Since(m.App.ServerStartTime).Minutes()),
		Environment: "dev",
		Version:     "1.0.0",
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, status)
}
