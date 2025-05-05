package handlers

import (
	"net/http"
	"smarthome/services"

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

	c.JSON(http.StatusOK, gin.H{"temperature": temperature})
}
