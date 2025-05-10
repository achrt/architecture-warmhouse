package handlers

import (
	"net/http"
	"smarthome/models"
	"smarthome/services"
	"time"

	"github.com/gin-gonic/gin"
)

type TemperatureHandler struct {
	TemperatureAPI services.TemperatureAPI
}

func NewTemperatureHandler(api services.TemperatureAPI) *TemperatureHandler {
	return &TemperatureHandler{
		TemperatureAPI: api,
	}
}

func (h *TemperatureHandler) RegisterRoutes(router *gin.RouterGroup) {
	{
		router.GET("/temperature", h.GetTemperatureByLocation)
	}
}

func (h *TemperatureHandler) GetTemperatureByLocation(c *gin.Context) {

	location := c.Query("location")
	sensorID := c.Query("sensorID")

	if location == "" && sensorID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "location or sensorID must be given"})
		return
	}

	temperature, err := h.TemperatureAPI.Get(sensorID, location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := models.Sensor{
		Value:       float64(temperature),
		LastUpdated: time.Now(),
		Location:    location,
	}

	c.JSON(http.StatusOK, resp)
}
